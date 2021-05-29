package controller

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/lxn/walk"

	"github.com/Clash-Mini/Clash.Mini/util"
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

const (
	localHost = `http://127.0.0.1`
)

var (
	fileSizeUnits = []string{"", "K", "M", "G", "T", "P", "E"}
)

// 格式化为可读文件大小
func formatHumanizationFileSize(fileSize int64) (size string) {
	order := 0
	floatSize := float64(fileSize)
	for {
		if floatSize < 1024 || order >= len(fileSizeUnits) {
			break
		}
		order++
		floatSize /= 1024
	}
	return fmt.Sprintf("%.02f %sB", floatSize, fileSizeUnits[order])
}

func (m *ConfigInfoModel) ResetRows() {
	fileInfoArr, err := ioutil.ReadDir("./Profile")
	if err != nil {
		log.Fatal(err)
	}
	var match string
	m.items = make([]*ConfigInfo, 0)
	for _, f := range fileInfoArr {
		if path.Ext(f.Name()) == ".yaml" {
			content, err := os.OpenFile("./Profile/"+f.Name(), os.O_RDWR, 0666)
			if err != nil {
				log.Fatal(err)
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
			content.Close()
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
	out.WriteString("# Yaml : " + name + ".yaml\n")
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

func putConfig(Name string) {
	_, controllerPort := checkConfig()
	err := copyFileContents("./Profile/"+Name+".yaml", "config.yaml", Name)
	time.Sleep(1 * time.Second)
	if err != nil {
		panic(err)
	}
	str, _ := os.Getwd()
	str = filepath.Join(str, "config.yaml")
	url := fmt.Sprintf("%s:%s/configs", localHost, controllerPort)
	body := make(map[string]interface{})
	body["path"] = str
	bytesData, err := json.Marshal(body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("PUT", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	resp.Body.Close()
}

func checkConfig() (config, controller string) {
	controller = "9090"
	config = "config.yaml"
	content, err := os.OpenFile("./config.yaml", os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
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
			controller = Reg2.FindStringSubmatch(scanner.Text())[2]
			break
		} else {
			controller = "9090"
		}
	}
	content.Close()
	return config, controller
}

func updateConfig(Name, url string) bool {
	client := &http.Client{}
	res, _ := http.NewRequest(http.MethodGet, url, nil)
	res.Header.Add("User-Agent", "clash")
	resp, err := client.Do(res)
	if err != nil {
		return false
	}
	if resp != nil && resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		Reg, _ := regexp.MatchString(`port`, string(body))
		if Reg != true {
			fmt.Println("错误的内容")
			return false
		}
		rebody := ioutil.NopCloser(bytes.NewReader(body))
		f, errf := os.OpenFile(fmt.Sprintf("./Profile/%s.yaml", Name), os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
		if errf != nil {
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
	content, err := os.OpenFile("./config.yaml", os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
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
			fmt.Println(err)
		}
	}(content)
	if infoURL != "" {
		client := &http.Client{}
		res, _ := http.NewRequest(http.MethodGet, infoURL, nil)
		res.Header.Add("User-Agent", "clash")
		resp, err := client.Do(res)
		if err != nil {
			return
		}
		userInfoStr := resp.Header.Get("Subscription-Userinfo")
		if len(strings.TrimSpace(userInfoStr)) > 0 {
			err = util.UnmarshalByValues(userInfoStr, &userInfo)
			if err != nil {
				fmt.Println(err)
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

func (m *ConfigInfoModel) TaskCorn() {
	successNum := 0
	failNum := 0
	for i, v := range m.items {
		if v.Url != "" {
			fmt.Println(v)
			err := updateConfig(v.Name, v.Url)
			if err != true {
				fmt.Println(v.Name + "更新失败")
				m.items[i].Url = "更新失败"
				failNum++
			} else {
				fmt.Println(v.Name + "更新成功")
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
