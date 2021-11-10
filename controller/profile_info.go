package controller

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	path "path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/Clash-Mini/Clash.Mini/common"
	"github.com/Clash-Mini/Clash.Mini/constant"
	cI18n "github.com/Clash-Mini/Clash.Mini/constant/i18n"
	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Clash-Mini/Clash.Mini/notify"
	p "github.com/Clash-Mini/Clash.Mini/profile"
	"github.com/Clash-Mini/Clash.Mini/static"
	commonUtils "github.com/Clash-Mini/Clash.Mini/util/common"
	fileUtils "github.com/Clash-Mini/Clash.Mini/util/file"
	httpUtils "github.com/Clash-Mini/Clash.Mini/util/http"
	stringUtils "github.com/Clash-Mini/Clash.Mini/util/string"

	"github.com/JyCyunMe/go-i18n/i18n"
	"github.com/lxn/walk"
)

const (
	profileInfoLogHeader = logHeader + ".ProfileInfo"
)

type ConfigInfo struct {
	Index   int
	Name    string
	Size    string
	Time    time.Time
	Url     string

	checked 		bool
}

type ConfigInfoModel struct {
	walk.TableModelBase
	items []*ConfigInfo
}

var (
	Profiles       []string
	CurrentProfile string
)

func (r ConfigInfo) displayName() string {
	return stringUtils.TrinocularString(r.checked, "√", "") + r.Name
}

func (m *ConfigInfoModel) ResetRows() {
	fileInfoArr, err := ioutil.ReadDir(constant.ProfileDir)
	if err != nil {
		errMsg := fmt.Sprintf("ResetRows ReadDir error: %v", err)
		log.Errorln(errMsg)
		notify.PushError("", errMsg)
		return
	}
	// TODO: load from /profile
	var match string
	m.items = make([]*ConfigInfo, 0)
	Profiles = []string{}
	for _, f := range fileInfoArr {
		if path.Ext(f.Name()) == constant.ConfigSuffix {
			profileName := strings.TrimSuffix(f.Name(), path.Ext(f.Name()))
			content, err := os.OpenFile(path.Join(constant.ProfileDir, f.Name()), os.O_RDWR, 0666)
			if err != nil {
				errMsg := fmt.Sprintf("ResetRows OpenFile error: %v", err)
				log.Errorln(errMsg)
				notify.PushError("", errMsg)
				return
			}

			reader := bufio.NewReader(content)
			lineData, _, err := reader.ReadLine()
			if err != nil {
				log.Errorln("[profile] updateSubscriptionUserInfo ReadLine error: %v", err)
				return
			}
			match = p.GetTagLineUrl(string(lineData))
			if err = content.Close(); err != nil {
				log.Errorln("[profile] RefreshProfiles CloseFile error: %v", err)
				return
			}

			m.items = append(m.items, &ConfigInfo{
				Name: profileName,
				Size: fileUtils.FormatHumanizedFileSize(f.Size()),
				Time: f.ModTime(),
				Url:  match,
			})
			Profiles = append(Profiles, profileName)
		}
	}
	m.PublishRowsReset()
}

func NewConfigInfoModel() *ConfigInfoModel {
	m := new(ConfigInfoModel)
	m.ResetRows()
	return m
}

func (m *ConfigInfoModel) Checked(row int) bool {
	return m.items[row].checked
}

func (m *ConfigInfoModel) RowCount() int {
	return len(m.items)
}

func (m *ConfigInfoModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	//case 0:
	//	return item.displayName
	case 0:
		return stringUtils.TrinocularString(item.checked, "√", "")
	case 1:
		return item.Name
	case 2:
		return item.Size
	case 3:
		return item.Time
	case 4:
		return item.Url
	}
	panic("unexpected col")
}

func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func copyCacheFile(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer out.Close()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

func copyFileContents(src, dst, name string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	out.WriteString(fmt.Sprintf("# Yaml : %s%s\n", name, constant.ConfigSuffix))
	defer out.Close()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

func PutConfig(name string) {
	cacheName, controllerPort := CheckConfig()
	err := copyCacheFile(constant.CacheFile, path.Join(constant.CacheDir, cacheName+constant.CacheFile))
	if err != nil {
		log.Errorln("[%s] PutConfig copyCacheFile1 error: %v", profileInfoLogHeader, err)
	}
	err = copyFileContents(path.Join(constant.ProfileDir, name+constant.ConfigSuffix), constant.ConfigFile, name)
	if err != nil {
		panic(err)
	}
	err = copyCacheFile(path.Join(constant.CacheDir, name+constant.ConfigSuffix+constant.CacheFile), constant.CacheFile)
	if err != nil {
		log.Errorln("[%s] PutConfig copyCacheFile2 error: %v", profileInfoLogHeader, err)
	}
	time.Sleep(1 * time.Second)
	str := path.Join(constant.Pwd, constant.ConfigFile)
	url := fmt.Sprintf("http://%s:%s/configs", constant.Localhost, controllerPort)
	body := make(map[string]interface{})
	body["path"] = str
	bytesData, err := json.Marshal(body)
	if err != nil {
		log.Errorln("[%s] PutConfig Marshal error: %v", profileInfoLogHeader, err)
		return
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest(http.MethodPut, url, reader)
	if err != nil {
		log.Errorln("[%s] PutConfig NewRequest error: %v", profileInfoLogHeader, err)
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	rsp, err := client.Do(request)
	defer httpUtils.DeferSafeCloseResponseBody(rsp)
	if err != nil {
		log.Errorln("[%s] PutConfig Do error: %v", profileInfoLogHeader, err)
		return
	}
}

func CheckConfig() (config, controllerPort string) {
	controllerPort = constant.ControllerPort
	config = commonUtils.GetExecutablePath(constant.ConfigFile)

	var err error
	exists, err := fileUtils.IsExists(constant.ConfigFile)
	if err != nil {
		err = fmt.Errorf("check config file error: %s", err.Error())
	} else {
		if !exists {
			log.Warnln("cannot find core config file, it will write default core config file")
			err = ioutil.WriteFile(constant.ConfigFile, static.ExampleConfig, 0644)
			if err != nil {
				err = fmt.Errorf("write default core config file error: %s", err.Error())
			}
		}
	}
	if err != nil {
		log.Errorln(err.Error())
		common.CoreRunningStatus = false
		return
	}

	content, err := os.OpenFile(config, os.O_RDWR, 0666)
	if err != nil {
		errMsg := fmt.Sprintf("CheckConfig error: %v", err)
		log.Errorln(errMsg)
		notify.PushError("", errMsg)
		return
	}
	scanner := bufio.NewScanner(content)
	Reg := regexp.MustCompile(`# Yaml : (.*)`)
	Reg2 := regexp.MustCompile(`external-controller: '?(.*:)?(\d+)'?`)
	for scanner.Scan() {
		if Reg.MatchString(scanner.Text()) {
			config = Reg.FindStringSubmatch(scanner.Text())[1]
			break
		} else {
			config = ""
		}
	}
	for scanner.Scan() {
		if Reg2.MatchString(scanner.Text()) {
			controllerPort = Reg2.FindStringSubmatch(scanner.Text())[2]
			break
		} else {
			controllerPort = constant.ControllerPort
		}
	}
	content.Close()

	return
}

func (m *ConfigInfoModel) TaskCron() {
	successNum := 0
	failNum := 0
	for i, v := range m.items {
		if v.Url != "" {
			log.Infoln("[%s] TaskCron Info: %v", profileInfoLogHeader, v)
			successful := p.UpdateConfig(v.Name, v.Url)
			if !successful {
				log.Errorln(fmt.Sprintf("%s: %s", i18n.T(cI18n.MenuConfigCronUpdateFailed), v.Name))
				m.items[i].Url = i18n.T(cI18n.MenuConfigCronUpdateFailed)
				failNum++
			} else {
				log.Errorln(fmt.Sprintf("%s: %s", i18n.T(cI18n.MenuConfigCronUpdateSuccessful), v.Name))
				m.items[i].Url = i18n.T(cI18n.MenuConfigCronUpdateSuccessful)
				successNum++
			}
		}
	}
	var message string
	if failNum > 0 {
		message = fmt.Sprintf("%s[%d] %s\n[%d] %s", message, successNum, i18n.T(cI18n.NotifyMessageCronNumSuccess), failNum, i18n.T(cI18n.NotifyMessageCronNumFail))
	} else {
		message += i18n.T(cI18n.NotifyMessageCronFinishAll)
	}
	walk.MsgBox(nil, i18n.T(cI18n.MsgBoxTitleTips), message, walk.MsgBoxIconInformation)
	m.ResetRows()
}
