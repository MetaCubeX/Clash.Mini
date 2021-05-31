package controller

import (
	"fmt"
	"os"
	path "path/filepath"
	"time"

	"github.com/Clash-Mini/Clash.Mini/constant"
	"github.com/Clash-Mini/Clash.Mini/notify"
	"github.com/Clash-Mini/Clash.Mini/util"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
	"github.com/skratchdot/open-golang/open"
)

var (
	appIcon, _ = walk.NewIconFromResourceId(2)
	hMenu      win.HMENU
	currStyle  int32
	xScreen    int32
	yScreen    int32
)

func init() {
	xScreen = win.GetSystemMetrics(win.SM_CXSCREEN)
	yScreen = win.GetSystemMetrics(win.SM_CYSCREEN)
}

func StyleMenuRun(w *walk.MainWindow, SizeW int32, SizeH int32) {
	currStyle = win.GetWindowLong(w.Handle(), win.GWL_STYLE)
	//removes default styling
	win.SetWindowLong(w.Handle(), win.GWL_STYLE, currStyle&^win.WS_SIZEBOX&^win.WS_MINIMIZEBOX&^win.WS_MAXIMIZEBOX)
	hMenu = win.GetSystemMenu(w.Handle(), false)
	win.RemoveMenu(hMenu, win.SC_CLOSE, win.MF_BYCOMMAND)
	win.SetWindowPos(w.Handle(), 0, (xScreen-SizeW)/2, (yScreen-SizeH)/2, SizeW, SizeH, win.SWP_FRAMECHANGED)
	win.ShowWindow(w.Handle(), win.SW_SHOW)
	w.Run()
}

func MenuConfig() {
	var (
		model         = NewConfigInfoModel()
		tv            *walk.TableView
		MenuConfig    *walk.MainWindow
		configIni     *walk.Label
		updateConfigs *walk.PushButton
	)
	configName, _ := checkConfig()


	err := MainWindow{
		Visible:  false,
		AssignTo: &MenuConfig,
		Name:     "ok",
		Title:    util.GetSubTitle("配置管理"),
		Icon:     appIcon,
		Layout:   VBox{}, //布局
		Children: []Widget{ //不动态添加控件的话，在此布局或者QT设计器设计UI文件，然后加载。
			Composite{
				Layout: VBox{
					Margins: Margins{Left: 8, Right: 8, Bottom: 0, Top: 0},
				},
				Children: []Widget{
					Label{
						Text:     "当前配置: " + configName,
						AssignTo: &configIni,
					},
					HSpacer{},
				},
			},
			Composite{
				Layout: VBox{
					Margins: Margins{Left: 8, Right: 8, Bottom: 0, Top: 0},
				},
				Children: []Widget{
					TableView{
						AssignTo:         &tv,
						CheckBoxes:       false,
						ColumnsOrderable: true,
						MultiSelection:   true,
						Columns: []TableViewColumn{
							{Title: "配置名称"},
							{Title: "文件大小"},
							{Title: "更新日期", Format: "01-02 15:04:05"},
							{Title: "订阅地址", Width: 295},
						},
						Model: model,
					},
				},
			},
			Composite{
				Layout: HBox{
					Margins: Margins{Left: 8, Right: 8, Bottom: 6, Top: 6},
				},
				Children: []Widget{
					HSpacer{},
					SplitButton{
						Text: "启用配置",
						MenuItems: []MenuItem{
							Action{
								AssignTo: nil,
								Text:     "添加配置",
								OnTriggered: func() {
									MenuConfig.SetVisible(false)
									AddConfig()
									model.ResetRows()
									time.Sleep(100 * time.Millisecond)
									MenuConfig.SetVisible(true)
								},
							},
							Action{
								AssignTo: nil,
								Text:     "升级配置",
								OnTriggered: func() {
									index := tv.CurrentIndex()
									if index != -1 && model.items[index].Url != "" {
										configName := model.items[index].Name
										configUrl := model.items[index].Url
										success := updateConfig(configName, configUrl)
										if !success {
											walk.MsgBox(MenuConfig, "提示",
												"更新配置失败", walk.MsgBoxIconError)
											return
										}
										walk.MsgBox(MenuConfig, "提示",
											fmt.Sprintf("成功更新 %s 配置！", configName), walk.MsgBoxIconInformation)
									} else {
										walk.MsgBox(MenuConfig, "提示",
											"请选择要更新的配置！", walk.MsgBoxIconError)
										return
									}
									model.ResetRows()
								},
							},
							Action{
								AssignTo: nil,
								Text:     "编辑配置",
								OnTriggered: func() {
									index := tv.CurrentIndex()
									if index != -1 {
										EditConfigName := model.items[index].Name
										EditConfigUrl := model.items[index].Url
										MenuConfig.SetVisible(false)
										EditConfig(EditConfigName, EditConfigUrl)
										model.ResetRows()
										time.Sleep(200 * time.Millisecond)
										MenuConfig.SetVisible(true)
									} else {
										walk.MsgBox(MenuConfig, "提示",
											"请选择要编辑的配置！", walk.MsgBoxIconError)
										return
									}
									model.ResetRows()
								},
							},
							Action{
								AssignTo: nil,
								Text:     "删除配置",
								OnTriggered: func() {
									index := tv.CurrentIndex()
									if index != -1 {
										deleteConfigName := model.items[index].Name
										if win.IDYES == walk.MsgBox(MenuConfig, "提示",
											"请确认是否删除该配置？", walk.MsgBoxYesNo) {
											err := os.Remove(path.Join(constant.ConfigDir,
												deleteConfigName+constant.ConfigSuffix))
											if err != nil {
												walk.MsgBox(MenuConfig, "提示",
													"删除配置失败！", walk.MsgBoxIconError)
												return
											} else {
												walk.MsgBox(MenuConfig, "提示",
													fmt.Sprintf("成功删除 %s 配置！", deleteConfigName),
													walk.MsgBoxIconInformation)
											}
										}
									} else {
										walk.MsgBox(MenuConfig, "提示",
											"请选择要删除的配置！", walk.MsgBoxIconError)
										return
									}
									model.ResetRows()
								},
							},
						},
						OnClicked: func() {
							index := tv.CurrentIndex()
							if index != -1 {
								configName := model.items[index].Name
								putConfig(configName)
								walk.MsgBox(MenuConfig, "提示",
									fmt.Sprintf("成功启用 %s 配置！", configName),
									walk.MsgBoxIconInformation)
								configIni.SetText(`当前配置: ` + configName + constant.ConfigSuffix)
								go func() {
									time.Sleep(1 * time.Second)
									userInfo := UpdateSubscriptionUserInfo()
									if len(userInfo.UnusedInfo) > 0 {
										notify.PushFlowInfo(userInfo.UsedInfo, userInfo.UnusedInfo, userInfo.ExpireInfo)
									}
								}()
							} else {
								walk.MsgBox(MenuConfig, "提示", "请选择要启用的配置！", walk.MsgBoxIconError)
								return
							}
							model.ResetRows()
						},
					},
					PushButton{
						Text:     "一键更新",
						AssignTo: &updateConfigs,
						OnClicked: func() {
							updateConfigs.SetText("更新中")
							model.TaskCron()
							updateConfigs.SetText("更新完成")
						},
					},
					PushButton{
						Text: "订阅转换",
						OnClicked: func() {
							err := open.Run(constant.SubConverterUrl)
							if err != nil {
								return
							}
						},
					},
					PushButton{
						Text: "打开目录",
						OnClicked: func() {
							err := open.Run(constant.ConfigDir)
							if err != nil {
								return
							}
						},
					},
					PushButton{
						Text: "关闭窗口",
						OnClicked: func() {
							err := MenuConfig.Close()
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
	StyleMenuRun(MenuConfig, 650, 250)

}
