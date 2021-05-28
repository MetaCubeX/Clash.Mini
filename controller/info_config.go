package controller

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lxn/walk"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
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

var sizeUnits = []string{"", "K", "M", "G", "T", "P", "E"}

func formatFileSize(fileSize int64) (size string) {
	order := 0
	floatSize := float64(fileSize)
	for {
		if floatSize < 1024 || order >= len(sizeUnits) {
			break
		}
		order++
		floatSize /= 1024
	}
	return fmt.Sprintf("%.02f %s%s", floatSize, sizeUnits[order], "B")
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
				Size: formatFileSize(f.Size()),
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
	_, controller := checkConfig()
	err := copyFileContents("./Profile/"+Name+".yaml", "config.yaml", Name)
	time.Sleep(1 * time.Second)
	if err != nil {
		panic(err)
	}
	str, _ := os.Getwd()
	str = filepath.Join(str, "config.yaml")
	url := `http://127.0.0.1:` + controller + "/configs"
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
	res, _ := http.NewRequest("GET", url, nil)
	res.Header.Add("User-Agent", "clash")
	resp, err := client.Do(res)
	if err != nil {
		return false
	}
	if resp != nil && resp.StatusCode == 200 {
		f, errf := os.OpenFile("./Profile/"+Name+".yaml", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
		if errf != nil {
			panic(err)
			return false
		}
		f.WriteString(`# Clash.Mini : ` + url + "\n")
		io.Copy(f, resp.Body)
		resp.Body.Close()
		f.Close()
		return true
	}
	return false
}

func UserINFO() (UsedINFO, UnUsedINFO, ExpireINFO string) {
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
	defer content.Close()
	if infoURL != "" {
		client := &http.Client{}
		res, _ := http.NewRequest("GET", infoURL, nil)
		res.Header.Add("User-Agent", "clash")
		resp, errdo := client.Do(res)
		if errdo != nil {
			return
		}
		userinfo := resp.Header.Get("Subscription-Userinfo")
		if userinfo != "" {
			reg := regexp.MustCompile(`upload=(\d+);\sdownload=(\d+);\stotal=(\d+)(;\sexpire=(\d+)?)?`)
			info := reg.FindStringSubmatch(userinfo)
			Upload, _ := strconv.ParseInt(info[1], 10, 64)
			Download, _ := strconv.ParseInt(info[2], 10, 64)
			Total, _ := strconv.ParseInt(info[3], 10, 64)
			Unused := Total - Upload - Download
			Used := Upload + Download
			UsedINFO = formatFileSize(Used)
			UnUsedINFO = formatFileSize(Unused)
			if info[5] != "" {
				Expire, _ := strconv.ParseInt(info[5], 10, 64)
				tm := time.Unix(Expire, 0)
				ExpireINFO = tm.Format("2006-01-02")
			} else {
				ExpireINFO = "无法确定"
			}
			return
		}
	} else {
		return
	}
	return
}

func (m *ConfigInfoModel) TaskCorn() {
	//go func() {
	success := 0
	fail := 0
	for i, v := range m.items {
		if v.Url != "" {
			fmt.Println(v)
			err := updateConfig(v.Name, v.Url)
			if err != true {
				fmt.Println(v.Name + "升级失败")
				m.items[i].Url = "更新失败"
				fail = fail + 1
			} else {
				fmt.Println(v.Name + "升级成功")
				m.items[i].Url = "成功更新"
				success = success + 1
			}
		}
	}
	if fail > 0 {
		walk.MsgBox(nil, "提示", "["+strconv.Itoa(success)+"] 个配置升级成功！"+"\n["+strconv.Itoa(fail)+"] 个配置升级失败！", walk.MsgBoxIconInformation)
	} else {
		walk.MsgBox(nil, "提示", "全部配置升级成功！", walk.MsgBoxIconInformation)
	}
	m.ResetRows()
	//}()
}
