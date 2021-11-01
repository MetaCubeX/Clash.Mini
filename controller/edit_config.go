package controller

import (
	"io/ioutil"
	"os"
	path "path/filepath"
	"strings"

	"github.com/Clash-Mini/Clash.Mini/constant"
	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Clash-Mini/Clash.Mini/util"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
)

func EditConfig(configName, configUrl string) {
	var editMenuConfig *walk.MainWindow
	var oUrl *walk.LineEdit
	var oUrlName *walk.LineEdit
	err := MainWindow{
		Visible:  true,
		AssignTo: &editMenuConfig,
		Name:     "EditConfig",
		Title:    util.GetSubTitle("编辑配置"),
		Icon:     appIcon,
		Font: Font{
			Family:    "Microsoft YaHei",
			PointSize: 9,
		},
		Layout: VBox{Alignment: AlignHCenterVCenter}, //布局
		Children: []Widget{ //不动态添加控件的话，在此布局或者QT设计器设计UI文件，然后加载。
			Composite{
				Layout: VBox{},
				Children: []Widget{
					Label{
						Text: "订阅名称:",
					},
					LineEdit{
						AssignTo: &oUrlName,
						Text:     configName,
					},
					Label{
						Text: "订阅链接:",
					},
					LineEdit{
						AssignTo: &oUrl,
						Text:     configUrl,
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
								if win.IDYES == walk.MsgBox(editMenuConfig, "提示",
									"确认修改该配置？", walk.MsgBoxYesNo) {
									configDir := path.Join(constant.ConfigDir, configName+constant.ConfigSuffix)
									newConfigDir := path.Join(constant.ConfigDir, oUrlName.Text()+constant.ConfigSuffix)
									buf, err := ioutil.ReadFile(configDir)
									if err != nil {
										panic(err)
									}

									content := string(buf)
									subStr := `# Clash.Mini : `

									if strings.Contains(content, subStr) {
										newContent := strings.Replace(content, subStr+configUrl, subStr+oUrl.Text(), 1)
										err = ioutil.WriteFile(configDir, []byte(newContent), 0)
									} else {
										newContent := subStr + oUrl.Text() + "\n" + content
										err = ioutil.WriteFile(configDir, []byte(newContent), 0)
									}
									CacheNameDir := path.Join(constant.CacheDir, configName+constant.ConfigSuffix+constant.CacheFile)
									NewCacheNameDir := path.Join(constant.CacheDir, oUrlName.Text()+constant.ConfigSuffix+constant.CacheFile)
									err = os.Rename(CacheNameDir, NewCacheNameDir)
									if err != nil {
										log.Errorln("无cache配置")
									}
									err = os.Rename(configDir, newConfigDir)
									if err != nil {
										walk.MsgBox(editMenuConfig, constant.UIConfigMsgTitle,
											"配置修改失败！", walk.MsgBoxIconError)
										return
									} else {
										walk.MsgBox(editMenuConfig, constant.UIConfigMsgTitle,
											"配置修改成功！", walk.MsgBoxIconInformation)

									}
									err = editMenuConfig.Close()
								}
							} else {
								walk.MsgBox(editMenuConfig, constant.UIConfigMsgTitle,
									"请输入订阅名称和链接！", walk.MsgBoxIconError)
							}
						},
					},
					PushButton{
						Text: "取消",
						OnClicked: func() {
							err := editMenuConfig.Close()
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
	StyleMenuRun(editMenuConfig, 420, 120)
}
