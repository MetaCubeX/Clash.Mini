package controller

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
	"io/ioutil"
	"os"
	"strings"
)

func EditConfig(ConfigName, ConfigUrl string) {
	var EditMenuConfig *walk.MainWindow
	var oUrl *walk.LineEdit
	var oUrlName *walk.LineEdit
	err := MainWindow{
		Visible:  true,
		AssignTo: &EditMenuConfig,
		Title:    "编辑配置 - Clash.Mini",
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
						Text:     ConfigName,
					},
					Label{
						Text: "订阅链接:",
					},
					LineEdit{
						AssignTo: &oUrl,
						Text:     ConfigUrl,
					},
				},
			},
			Composite{
				Layout: HBox{MarginsZero: true},
				Children: []Widget{
					HSpacer{},
					PushButton{
						Text: "确认修改",
						OnClicked: func() {
							if oUrlName != nil {
								if win.IDYES == walk.MsgBox(EditMenuConfig, "提示", "确认修改该配置？", walk.MsgBoxYesNo) {
									if oUrl != nil && strings.HasPrefix(oUrl.Text(), "http") {
										buf, err := ioutil.ReadFile("./Profile/" + ConfigName + ".yaml")
										if err != nil {
											panic(err)
										}
										content := string(buf)
										if strings.Contains(content, `# Clash.Mini : `) {
											newContent := strings.Replace(content, `# Clash.Mini : `+ConfigUrl, `# Clash.Mini : `+oUrl.Text(), 1)
											ioutil.WriteFile("./Profile/"+ConfigName+".yaml", []byte(newContent), 0)
										} else {
											newContent := `# Clash.Mini : ` + oUrl.Text() + "\n" + content
											ioutil.WriteFile("./Profile/"+ConfigName+".yaml", []byte(newContent), 0)
										}
									}
									err1 := os.Rename("./Profile/"+ConfigName+".yaml", "./Profile/"+oUrlName.Text()+".yaml")
									if err1 != nil {
										walk.MsgBox(EditMenuConfig, "配置提示", "配置修改失败！", walk.MsgBoxIconError)
										return
									} else {
										walk.MsgBox(EditMenuConfig, "配置提示", "配置修改成功！", walk.MsgBoxIconInformation)
									}
									EditMenuConfig.Close()
								}
							} else {
								walk.MsgBox(EditMenuConfig, "配置提示", "请输入订阅名称和链接！", walk.MsgBoxIconError)
							}
						},
					},
					PushButton{
						Text: "取消",
						OnClicked: func() {
							err := EditMenuConfig.Close()
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
	StyleMenuRun(EditMenuConfig, 420, 120)
}
