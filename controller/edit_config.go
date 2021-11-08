package controller

import (
	"io/ioutil"
	"os"
	path "path/filepath"
	"strings"

	"github.com/Clash-Mini/Clash.Mini/constant"
	"github.com/Clash-Mini/Clash.Mini/log"
	cI18n "github.com/Clash-Mini/Clash.Mini/constant/i18n"
	stringUtils "github.com/Clash-Mini/Clash.Mini/util/string"

	"github.com/JyCyunMe/go-i18n/i18n"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/walk"
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
		Title:    stringUtils.GetSubTitle(i18n.T(cI18n.EditConfigWindowTitle)),
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
						Text: i18n.T(cI18n.EditConfigWindowSubscriptionName),
					},
					LineEdit{
						AssignTo: &oUrlName,
						Text:     configName,
					},
					Label{
						Text: i18n.T(cI18n.EditConfigWindowSubscriptionUrl),
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
						Text: i18n.T(cI18n.ButtonSubmit),
						OnClicked: func() {
							if oUrlName != nil {
								if win.IDYES == walk.MsgBox(editMenuConfig, i18n.T(cI18n.MsgBoxTitleTips),
									i18n.TData(cI18n.EditConfigMessageEditConfigConfirmMsg, &i18n.Data{Data: map[string]interface{}{
										"Config": configName,
									}}), walk.MsgBoxYesNo) {
									configDir := path.Join(constant.ProfileDir, configName + constant.ConfigSuffix)
									newConfigDir := path.Join(constant.ProfileDir, oUrlName.Text() + constant.ConfigSuffix)
									buf, err := ioutil.ReadFile(configDir)
									if err != nil {
										panic(err)
									}

									content := string(buf)
									subStr := `# Clash.Mini : `

									if strings.Contains(content, subStr) {
										newContent := strings.Replace(content, subStr+configUrl, subStr + oUrl.Text(), 1)
										err = ioutil.WriteFile(configDir, []byte(newContent), 0)
									} else {
										newContent := subStr + oUrl.Text() + "\n" + content
										err = ioutil.WriteFile(configDir, []byte(newContent), 0)
									}
									CacheNameDir := path.Join(constant.CacheDir, configName+constant.ConfigSuffix+constant.CacheFile)
									NewCacheNameDir := path.Join(constant.CacheDir, oUrlName.Text()+constant.ConfigSuffix+constant.CacheFile)
									err = os.Rename(CacheNameDir, NewCacheNameDir)
									if err != nil {
										log.Errorln("not found cache file")
									}
									err = os.Rename(configDir, newConfigDir)
									if err != nil {
										walk.MsgBox(editMenuConfig, constant.UIConfigMsgTitle,
											i18n.TData(cI18n.EditConfigMessageEditConfigFailure, &i18n.Data{Data: map[string]interface{}{
												"Config": configName,
											}}), walk.MsgBoxIconError)
										return
									} else {
										walk.MsgBox(editMenuConfig, constant.UIConfigMsgTitle,
											i18n.TData(cI18n.EditConfigMessageEditConfigSuccess, &i18n.Data{Data: map[string]interface{}{
												"Config": configName,
											}}), walk.MsgBoxIconInformation)

									}
									err = editMenuConfig.Close()
								}
							} else {
								walk.MsgBox(editMenuConfig, constant.UIConfigMsgTitle,
									i18n.T(cI18n.EditConfigMessageEditConfigNothing), walk.MsgBoxIconError)
							}
						},
					},
					PushButton{
						Text: i18n.T(cI18n.ButtonCancel),
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
