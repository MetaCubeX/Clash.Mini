package controller

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
	"github.com/skratchdot/open-golang/open"
	"os"
)

var hMenu win.HMENU
var currStyle int32
var xScreen int32
var yScreen int32

func init() {
	xScreen = win.GetSystemMetrics(win.SM_CXSCREEN)
	yScreen = win.GetSystemMetrics(win.SM_CYSCREEN)
}

func StyleMenuRun(w *walk.MainWindow, SizeW int32, SizeH int32) {
	currStyle = win.GetWindowLong(w.Handle(), win.GWL_STYLE)
	win.SetWindowLong(w.Handle(), win.GWL_STYLE, currStyle&^win.WS_SIZEBOX&^win.WS_MINIMIZEBOX&^win.WS_MAXIMIZEBOX) //removes default styling
	hMenu = win.GetSystemMenu(w.Handle(), false)
	win.RemoveMenu(hMenu, win.SC_CLOSE, win.MF_BYCOMMAND)
	win.SetWindowPos(w.Handle(), 0, (xScreen-SizeW)/2, (yScreen-SizeH)/2, SizeW, SizeH, win.SWP_FRAMECHANGED)
	win.ShowWindow(w.Handle(), win.SW_SHOW)
	w.Run()
}
func StyleMenu2Run(w *walk.MainWindow, SizeW int32, SizeH int32) {
	win.SetWindowLong(w.Handle(), win.GWL_STYLE, currStyle&^win.WS_SIZEBOX&^win.WS_MINIMIZEBOX&^win.WS_MAXIMIZEBOX) //removes default styling
	win.RemoveMenu(hMenu, win.SC_CLOSE, win.MF_BYCOMMAND)
	win.SetWindowPos(w.Handle(), 0, (xScreen-SizeW)/2, (yScreen-SizeH)/2, SizeW, SizeH, win.SWP_FRAMECHANGED)
	win.ShowWindow(w.Handle(), win.SW_SHOW)
	w.Run()
}

func MenuConfig() {
	var model = NewConfigInfoModel()
	var tv *walk.TableView
	var MenuConfig *walk.MainWindow
	var configIni *walk.Label
	configName, _ := checkConfig()
	err := MainWindow{
		Visible:  false,
		AssignTo: &MenuConfig,
		Name:     "ok",
		Title:    "配置管理 - Clash.Mini",
		Icon:     "icon/Clash.Mini.ico",
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
									go AddConfig()
									//model.ResetRows()
									err := MenuConfig.Close()
									if err != nil {
										return
									}
								},
							},
							Action{
								AssignTo: nil,
								Text:     "升级配置",
								OnTriggered: func() {
									index := tv.CurrentIndex()
									if index != -1 && model.items[index].Url != "" {
										ConfigName := model.items[index].Name
										ConfigUrl := model.items[index].Url
										err := updateConfig(ConfigName, ConfigUrl)
										if err != nil {
											return
											walk.MsgBox(MenuConfig, "提示", "升级配置失败", walk.MsgBoxIconInformation)
										}
										walk.MsgBox(MenuConfig, "提示", "成功升级"+ConfigName+"配置！", walk.MsgBoxIconInformation)
									} else {
										walk.MsgBox(MenuConfig, "提示", "请选择要升级的配置！", walk.MsgBoxIconError)
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
										ConfigName := model.items[index].Name
										ConfigUrl := model.items[index].Url
										go EditConfig(ConfigName, ConfigUrl)
										MenuConfig.Close()
									} else {
										walk.MsgBox(MenuConfig, "提示", "请选择要编辑的配置！", walk.MsgBoxIconError)
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
										ConfigName := model.items[index].Name
										if win.IDYES == walk.MsgBox(MenuConfig, "提示", "请确认是否删除该配置？", walk.MsgBoxYesNo) {
											err := os.Remove("./Profile/" + ConfigName + ".yaml")
											if err != nil {
												walk.MsgBox(MenuConfig, "提示", "删除配置失败！", walk.MsgBoxIconError)
												return
											} else {
												walk.MsgBox(MenuConfig, "提示", "成功删除"+ConfigName+"配置！", walk.MsgBoxIconInformation)
											}
										}
									} else {
										walk.MsgBox(MenuConfig, "提示", "请选择要删除的配置！", walk.MsgBoxIconError)
										return
									}
									model.ResetRows()
								},
							},
						},
						OnClicked: func() {
							index := tv.CurrentIndex()
							if index != -1 {
								ConfigName := model.items[index].Name
								putConfig(ConfigName)
								walk.MsgBox(MenuConfig, "提示", "成功启用"+model.items[index].Name+"配置！", walk.MsgBoxIconInformation)
								configIni.SetText(`当前配置: ` + model.items[index].Name + `.yaml`)
							} else {
								walk.MsgBox(MenuConfig, "提示", "请选择要启用的配置！", walk.MsgBoxIconError)
								return
							}
							model.ResetRows()
						},
					},
					PushButton{
						Text: "订阅转换",
						OnClicked: func() {
							err := open.Run("https://id9.cc")
							if err != nil {
								return
							}
						},
					},
					PushButton{
						Text: "打开目录",
						OnClicked: func() {
							exPath, _ := os.Getwd()
							err := open.Run(exPath + `/Profile`)
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
