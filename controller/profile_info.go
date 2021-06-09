package controller

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	path "path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/Clash-Mini/Clash.Mini/constant"
	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Clash-Mini/Clash.Mini/util"

	"github.com/lxn/walk"
)

type ConfigInfo struct {
	Index   int
	Name    string
	Size    string
	Time    time.Time
	Url     string
	checked bool
}

type ConfigInfoModel struct {
	walk.TableModelBase
	items []*ConfigInfo
}

var (
	fileSizeUnits = []string{"", "K", "M", "G", "T", "P", "E"}
)

// 格式化为可读文件大小
func formatHumanizationFileSize(fileSize int64) (size string) {
	i := math.Floor(math.Log(float64(fileSize)) / math.Log(1024))
	return fmt.Sprintf("%.02f %sB", float64(fileSize)/math.Pow(1024, i), fileSizeUnits[int(i)])
}

func (m *ConfigInfoModel) ResetRows() {
	fileInfoArr, err := ioutil.ReadDir(constant.ConfigDir)
	if err != nil {
		log.Fatalln("ResetRows ReadDir error: %v", err)
	}
	var match string
	m.items = make([]*ConfigInfo, 0)
	for _, f := range fileInfoArr {
		if path.Ext(f.Name()) == constant.ConfigSuffix {
			content, err := os.OpenFile(path.Join(constant.ConfigDir, f.Name()), os.O_RDWR, 0666)
			if err != nil {
				log.Fatalln("ResetRows OpenFile error: %v", err)
			}
			scanner := bufio.NewScanner(content)
			Reg := regexp.MustCompile(`# Clash.Mini : (http.*)`)
			for scanner.Scan() {
				if Reg.MatchString(scanner.Text()) {
					match = Reg.FindStringSubmatch(scanner.Text())[1]
					break
				} else {
					match = ""
				}
			}
			if err = content.Close(); err != nil {
				return
			}
			m.items = append(m.items, &ConfigInfo{
				Name: strings.TrimSuffix(f.Name(), path.Ext(f.Name())),
				Size: formatHumanizationFileSize(f.Size()),
				Time: f.ModTime(),
				Url:  match,
			})
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
	case 0:
		return item.Name
	case 1:
		return item.Size
	case 2:
		return item.Time
	case 3:
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

func putConfig(name string) {
	cacheName, controllerPort := CheckConfig()
	err := copyCacheFile(constant.CacheFile, path.Join(constant.CacheDir, cacheName+constant.CacheFile))
	if err != nil {
		log.Errorln("putConfig copyCacheFile1 error: %v", err)
	}
	err = copyFileContents(path.Join(constant.ConfigDir, name+constant.ConfigSuffix), constant.ConfigFile, name)
	if err != nil {
		panic(err)
	}
	err = copyCacheFile(path.Join(constant.CacheDir, name+constant.ConfigSuffix+constant.CacheFile), constant.CacheFile)
	if err != nil {
		log.Errorln("putConfig copyCacheFile2 error: %v", err)
	}
	time.Sleep(1 * time.Second)
	str := path.Join(constant.PWD, constant.ConfigFile)
	url := fmt.Sprintf("http://%s:%s/configs", constant.Localhost, controllerPort)
	body := make(map[string]interface{})
	body["path"] = str
	bytesData, err := json.Marshal(body)
	if err != nil {
		log.Errorln("putConfig Marshal error: %v", err)
		return
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest(http.MethodPut, url, reader)
	if err != nil {
		log.Errorln("putConfig NewRequest error: %v", err)
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Errorln("putConfig Do error: %v", err)
		return
	}

	if err := resp.Body.Close(); err != nil {
		return
	}
}

func CheckConfig() (config, controllerPort string) {
	controllerPort = constant.ControllerPort
	config = constant.ConfigFile
	content, err := os.OpenFile(path.Join(".", constant.ConfigFile), os.O_RDWR, 0666)
	if err != nil {
		log.Fatalln("CheckConfig error: %v", err)
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

func updateConfig(name, url string) bool {
	client := &http.Client{Timeout: 5 * time.Second}
	res, _ := http.NewRequest(http.MethodGet, url, nil)
	res.Header.Add("User-Agent", "clash")
	resp, err := client.Do(res)
	if err != nil {
		return false
	}
	if resp != nil && resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		Reg, _ := regexp.MatchString(`proxy-groups`, string(body))
		if Reg != true {
			log.Errorln("错误的内容")
			return false
		}
		rebody := ioutil.NopCloser(bytes.NewReader(body))

		f, err := os.OpenFile(path.Join(constant.ConfigDir, name+constant.ConfigSuffix),
			os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
		if err != nil {
			panic(err)
			return false
		}
		f.WriteString(fmt.Sprintf("# Clash.Mini : %s\n", url))
		io.Copy(f, rebody)
		resp.Body.Close()
		f.Close()
		return true
	}
	return false
}

type SubscriptionUserInfo struct {
	Upload     int64 `query:"upload"`
	Download   int64 `query:"download"`
	Total      int64 `query:"total"`
	Unused     int64
	Used       int64
	ExpireUnix int64 `query:"expire"`

	UsedInfo   string
	UnusedInfo string
	ExpireInfo string
}

func UpdateSubscriptionUserInfo() (userInfo SubscriptionUserInfo) {
	var (
		infoURL = ""
	)
	content, err := os.OpenFile(path.Join(".", constant.ConfigFile), os.O_RDWR, 0666)
	if err != nil {
		log.Fatalln("updateSubscriptionUserInfo error", err)
	}
	scanner := bufio.NewScanner(content)
	Reg := regexp.MustCompile(`# Clash.Mini : (http.*)`)
	for scanner.Scan() {
		if Reg.MatchString(scanner.Text()) {
			infoURL = Reg.FindStringSubmatch(scanner.Text())[1]
			break
		} else {
			infoURL = ""
		}
	}
	defer func(content *os.File) {
		err := content.Close()
		if err != nil {
			log.Errorln("UpdateSubscriptionUserInfo close error", err)
		}
	}(content)
	if infoURL != "" {
		client := &http.Client{Timeout: 5 * time.Second}
		res, _ := http.NewRequest(http.MethodGet, infoURL, nil)
		res.Header.Add("User-Agent", "clash")
		resp, err := client.Do(res)
		if err != nil {
			return
		}
		userInfoStr := resp.Header.Get("Subscription-Userinfo")
		if len(strings.TrimSpace(userInfoStr)) == 0 {
			res2, _ := http.NewRequest(http.MethodGet, infoURL, nil)
			res2.Header.Add("User-Agent", "Quantumultx")
			resp2, err := client.Do(res2)
			if err != nil {
				return
			}
			userInfoStr = resp2.Header.Get("Subscription-Userinfo")
		}
		if len(strings.TrimSpace(userInfoStr)) > 0 {
			userInfo = SubscriptionUserInfo{}
			err = util.UnmarshalByValues(userInfoStr, &userInfo)
			if err != nil {
				log.Errorln("UpdateSubscriptionUserInfo UnmarshalByValues error: %v", err)
				return
			}
			userInfo.Used = userInfo.Upload + userInfo.Download
			userInfo.Unused = userInfo.Total - userInfo.Used
			userInfo.UsedInfo = formatHumanizationFileSize(userInfo.Used)
			userInfo.UnusedInfo = formatHumanizationFileSize(userInfo.Unused)
			if userInfo.ExpireUnix > 0 {
				userInfo.ExpireInfo = time.Unix(userInfo.ExpireUnix, 0).Format("2006-01-02")
			} else {
				userInfo.ExpireInfo = "暂无"
			}
			return
		}
	} else {
		return
	}
	return
}

func (m *ConfigInfoModel) TaskCron() {
	successNum := 0
	failNum := 0
	for i, v := range m.items {
		if v.Url != "" {
			log.Infoln("TaskCron Info: %v", v)
			err := updateConfig(v.Name, v.Url)
			if err != true {
				log.Errorln(v.Name + "更新失败")
				m.items[i].Url = "更新失败"
				failNum++
			} else {
				log.Infoln(v.Name + "更新成功")
				m.items[i].Url = "成功更新"
				successNum++
			}
		}
	}
	if failNum > 0 {
		walk.MsgBox(nil, "提示", fmt.Sprintf("[%d] 个配置更新成功！\n[%d] 个配置更新失败！", successNum, failNum),
			walk.MsgBoxIconInformation)
	} else {
		walk.MsgBox(nil, "提示", "全部配置更新成功！", walk.MsgBoxIconInformation)
	}
	m.ResetRows()
}
