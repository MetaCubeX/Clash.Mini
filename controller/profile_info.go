package controller

import (
	"bufio"
	"fmt"
	"github.com/Clash-Mini/Clash.Mini/config"
	"io"
	"io/ioutil"
	"os"
	path "path/filepath"
	"strings"
	"time"

	"github.com/Clash-Mini/Clash.Mini/common"
	"github.com/Clash-Mini/Clash.Mini/constant"
	cI18n "github.com/Clash-Mini/Clash.Mini/constant/i18n"
	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Clash-Mini/Clash.Mini/notify"
	p "github.com/Clash-Mini/Clash.Mini/profile"
	fileUtils "github.com/Clash-Mini/Clash.Mini/util/file"
	stringUtils "github.com/Clash-Mini/Clash.Mini/util/string"

	"github.com/JyCyunMe/go-i18n/i18n"
	"github.com/lxn/walk"
)

const (
	profileInfoLogHeader = logHeader + ".ProfileInfo"
)

type ConfigInfo struct {
	Index int
	Name  string
	Size  string
	Time  time.Time
	Url   string

	checked bool
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
				Name:    profileName,
				Size:    fileUtils.FormatHumanizedFileSize(f.Size()),
				Time:    f.ModTime(),
				Url:     match,
				checked: profileName == CurrentProfile,
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

func CopyFileContents(src, dst, name string) (err error) {
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

func ApplyConfig(name string, isUpdate bool) bool {
	//oldProfile := config.GetProfile()
	exist, configName := CheckConfig(name)

	if exist {
		//if oldProfile != name {
		// Archive cache
		//err := copyCacheFile(constant.CacheFile, path.Join(constant.CacheDir, oldProfile+"-"+constant.CacheFile))
		//if err != nil {
		//	log.Errorln("[%s] ApplyConfig archive cache error: %v", profileInfoLogHeader, err)
		//}

		// Replace cache
		//err = copyCacheFile(path.Join(constant.CacheDir, name+"-"+constant.CacheFile), constant.CacheFile)
		//if err != nil {
		//	log.Errorln("[%s] ApplyConfig replace cache error: %v", profileInfoLogHeader, err)
		//}
		//}
		//time.Sleep(1 * time.Second)

		str := path.Join(constant.ProfileDir, configName)

		// Load configuration file, mix configuration in memory, and start
		if err := CoreStart(str, isUpdate); err != nil {
			errString := fmt.Sprintf("Parse config error: %s", err.Error())
			log.Errorln(errString)
			common.SetStatus(false)
			return false
		}

		common.SetStatus(true)
		config.SetProfile(name)
		return true
	} else {
		log.Errorln("cannot found %s", name)
		return false
	}

}

func CheckConfig(name string) (exits bool, configName string) {
	if files, err := ioutil.ReadDir(constant.ProfileDir); err == nil {
		for _, file := range files {
			//log.Infoln(file.Name())
			if !file.IsDir() {
				if file.Name() == name+constant.ConfigSuffix {
					return true, file.Name()
				}
			}
		}
		log.Warnln("not found config by name[%s]", name)
		return false, ""
	} else {
		log.Errorln("read config list error:%v", err)
		return false, ""
	}
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
				log.Infoln(fmt.Sprintf("%s: %s", i18n.T(cI18n.MenuConfigCronUpdateSuccessful), v.Name))
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
