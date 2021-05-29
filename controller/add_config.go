package controller

import (
	"bytes"
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func AddConfig() {
	var AddMenuConfig *walk.MainWindow
	var oUrl *walk.LineEdit
	var oUrlName *walk.LineEdit
	err := MainWindow{
		Visible:  false,
		AssignTo: &AddMenuConfig,
		Title:    "添加配置 - Clash.Mini",
		Icon:     appIcon,
		Layout:   VBox{}, //布局
		Children: []Widget{ //不动态添加控件的话，在此布局或者QT设计器设计UI文件，然后加载。
			Composite{
				Layout: VBox{},
				Children: []Widget{
					Label{
						Text: "订阅名称:",
					},
					LineEdit{
						AssignTo: &oUrlName,
					},
					Label{
						Text: "订阅链接:",
					},
					LineEdit{
						AssignTo: &oUrl,
					},
				},
			},
			Composite{
				Layout: HBox{MarginsZero: true},
				Children: []Widget{
					HSpacer{},
					PushButton{
						Text: "添加",
						OnClicked: func() {
							if oUrlName != nil && oUrl != nil && strings.HasPrefix(oUrl.Text(), "http") {
								client := &http.Client{}
								res, _ := http.NewRequest(http.MethodGet, oUrl.Text(), nil)
								res.Header.Add("User-Agent", "clash")
								resp, err := client.Do(res)
								defer resp.Body.Close()
								if err != nil {
									walk.MsgBox(AddMenuConfig, "配置提示", "请检查订阅链接是否正确！", walk.MsgBoxIconError)
									return
								}
								if resp != nil && resp.StatusCode == 200 {
									body, _ := ioutil.ReadAll(resp.Body)
									Reg, _ := regexp.MatchString(`port`, string(body))
									if Reg != true {
										fmt.Println("错误的内容")
										walk.MsgBox(AddMenuConfig, "配置提示", "检测为非Clash配置，添加配置失败！", walk.MsgBoxIconError)
										return
									}
									rebody := ioutil.NopCloser(bytes.NewReader(body))
									ConfigDir := filepath.Join(".", "Profile", oUrlName.Text()+".yaml")
									f, err := os.Create(ConfigDir)
									if err != nil {
										panic(err)
									}
									_, err = f.WriteString(`# Clash.Mini : ` + oUrl.Text() + "\n")
									_, err = io.Copy(f, rebody)
									err = f.Close()
									walk.MsgBox(AddMenuConfig, "配置提示", "添加配置成功！", walk.MsgBoxIconInformation)
									AddMenuConfig.Close()
								} else {
									walk.MsgBox(AddMenuConfig, "配置提示", "请检查订阅链接是否正确！", walk.MsgBoxIconError)
								}
							} else {
								walk.MsgBox(AddMenuConfig, "配置提示", "请输入订阅名称和链接！", walk.MsgBoxIconError)
							}
						},
					},
					PushButton{
						Text: "取消",
						OnClicked: func() {
							AddMenuConfig.Close()
						},
					},
				},
			},
		},
	}.Create()
	if err != nil {
		return
	}
	StyleMenuRun(AddMenuConfig, 420, 120)
}
