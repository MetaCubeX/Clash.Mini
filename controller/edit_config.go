package controller

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
	"io/ioutil"
	"os"
	"path/filepath"
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
									configDir := filepath.Join(".", "Profile", ConfigName+".yaml")
									NewconfigDir := filepath.Join(".", "Profile", oUrlName.Text()+".yaml")
									buf, err := ioutil.ReadFile(configDir)
									if err != nil {
										panic(err)
									}

									content := string(buf)
									subStr := `# Clash.Mini : `

									if strings.Contains(content, subStr) {
										newContent := strings.Replace(content, subStr+ConfigUrl, subStr+oUrl.Text(), 1)
										err = ioutil.WriteFile(configDir, []byte(newContent), 0)
									} else {
										newContent := subStr + oUrl.Text() + "\n" + content
										err = ioutil.WriteFile(configDir, []byte(newContent), 0)
									}
									err = os.Rename(configDir, NewconfigDir)
									if err != nil {
										walk.MsgBox(EditMenuConfig, "配置提示", "配置修改失败！", walk.MsgBoxIconError)
										return
									} else {
										walk.MsgBox(EditMenuConfig, "配置提示", "配置修改成功！", walk.MsgBoxIconInformation)
									}
									err = EditMenuConfig.Close()
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
