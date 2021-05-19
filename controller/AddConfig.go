package controller

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"io"
	"net/http"
	"os"
	"strings"
)

func AddConfig() {
	var AddMenuConfig *walk.MainWindow
	var oUrl *walk.LineEdit
	var oUrlName *walk.LineEdit
	err := MainWindow{
		Visible:  true,
		AssignTo: &AddMenuConfig,
		Title:    "添加配置 - Clash.Mini",
		Icon:     "./icon/Clash.Mini.ico",
		Layout:   VBox{}, //布局
		Children: []Widget{ //不动态添加控件的话，在此布局或者QT设计器设计UI文件，然后加载。
			Composite{
				Layout: VBox{},
				Children: []Widget{
					Label{
						Text: "请输入订阅链接:",
					},
					LineEdit{
						AssignTo: &oUrl,
					},
					Label{
						Text: "请输入订阅名称:",
					},
					LineEdit{
						AssignTo: &oUrlName,
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
								res, _ := http.NewRequest("GET", oUrl.Text(), nil)
								res.Header.Add("User-Agent", "clash")
								resp, err := client.Do(res)
								//Reg := regexp.MustCompile(`proxies`)
								if err != nil {
									walk.MsgBox(AddMenuConfig, "配置提示", "请检查订阅链接是否正确！", walk.MsgBoxIconError)
								}
								if resp != nil {
									//body, _ := ioutil.ReadAll(resp.Body)
									//if Reg.MatchString(string(body)) {
									f, err := os.Create("./Profile/" + oUrlName.Text() + ".yaml")
									if err != nil {
										panic(err)
									}
									f.WriteString(`# Clash.Mini : ` + oUrl.Text() + "\n")
									io.Copy(f, resp.Body)
									resp.Body.Close()
									f.Close()
									walk.MsgBox(AddMenuConfig, "配置提示", "添加配置成功！", walk.MsgBoxIconInformation)
									AddMenuConfig.Close()
									MenuConfig()
									//} else {
									//	walk.MsgBox(AddMenuConfig, "配置提示", "检测为非Clash配置，添加配置失败！", walk.MsgBoxIconError)
									//}
								}
							} else {
								walk.MsgBox(AddMenuConfig, "配置提示", "请输入订阅链接和名称！", walk.MsgBoxIconError)
							}

						},
					},
					PushButton{
						Text: "取消",
						OnClicked: func() {
							defer MenuConfig()
							err := AddMenuConfig.Close()
							if err != nil {
								return
							}
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
