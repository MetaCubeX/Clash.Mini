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
	//walk.SorterBase
	//sortColumn int
	//ortOrder  walk.SortOrder
	items []*ConfigInfo
}

func formatFileSize(fileSize int64) (size string) {
	if fileSize < 1024 {
		return fmt.Sprintf("%.2fB", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(fileSize)/float64(1024*1024*1024))
	} else {
		return fmt.Sprintf("%.2fTB", float64(fileSize)/float64(1024*1024*1024*1024))
	}
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

//func (m *ConfigInfoModel) Sort(col int,order walk.SortOrder) error{
//	m.sortColumn,m.sortOrder = col,order
//	sort.Stable(m)
//	return m.SorterBase.Sort(col,order)
//}

//func (m *ConfigInfoModel) Len() int {
//	return len(m.items)
//}

//func (m *ConfigInfoModel) Less(i,j int) bool {
//	a,b := m.items[i],m.items[j]
//	c:=func(ls bool)bool{
//		if m.sortOrder == walk.SortAscending{
//			return ls
//		}
//		return !ls
//	}
//	switch m.sortColumn {
//	case 0:
//		return c(a.Index < b.Index)
//	case 1:
//		return c(a.Name<b.Name)
//	case 2:
//		return c(a.Byte<b.Byte)
//	case 3:
//		return c(a.Time<b.Time)
//	case 4:
//		return c(a.Url<b.Url)
//	}
//	panic("unreachable")
//}

//func (m *ConfigInfoModel) Swap(i,j int) {
//	m.items[i],m.items[j] = m.items[j],m.items[i]
//}

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

func updateConfig(Name, url string) error {
	client := &http.Client{}
	res, _ := http.NewRequest("GET", url, nil)
	res.Header.Add("User-Agent", "clash")
	resp, err := client.Do(res)
	if err != nil {
		return err
	}
	if resp != nil {
		f, errf := os.OpenFile("./Profile/"+Name+".yaml", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
		if errf != nil {
			panic(err)
		}
		f.WriteString(`# Clash.Mini : ` + url + "\n")
		io.Copy(f, resp.Body)
		resp.Body.Close()
		f.Close()
	}
	return err
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
			reg := regexp.MustCompile(`=(\d+);\s.*=(\d+);\s.*=(\d+);\s.*=(\d+)?`)
			info := reg.FindStringSubmatch(userinfo)
			Upload, _ := strconv.ParseInt(info[1], 10, 64)
			Download, _ := strconv.ParseInt(info[2], 10, 64)
			Total, _ := strconv.ParseInt(info[3], 10, 64)
			Unused := Total - Upload - Download
			Used := Upload + Download
			UsedINFO = formatFileSize(Used)
			UnUsedINFO = formatFileSize(Unused)
			if info[4] != "" {
				Expire, _ := strconv.ParseInt(info[4], 10, 64)
				tm := time.Unix(Expire, 0)
				ExpireINFO = tm.Format("2006-01-02")
			} else {
				ExpireINFO = "无限期"
			}
			return
		}
	} else {
		return
	}
	return
}
