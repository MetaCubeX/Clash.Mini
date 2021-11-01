package controller

import (
	"fmt"
	"os"
	path "path/filepath"
	"time"

	"github.com/Clash-Mini/Clash.Mini/constant"
	cI18n "github.com/Clash-Mini/Clash.Mini/constant/i18n"
	"github.com/Clash-Mini/Clash.Mini/notify"
	"github.com/Clash-Mini/Clash.Mini/util"

	"github.com/JyCyunMe/go-i18n/i18n"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
	"github.com/skratchdot/open-golang/open"
)

var (
	appIcon, _ 	= walk.NewIconFromResourceId(2)
	hMenu      	win.HMENU
	currStyle  	int32
	xScreen    	int32
	yScreen    	int32
	dpiScale   	float64

	WindowMap  	= make(map[string]*walk.MainWindow)
	MenuConfig 	*walk.MainWindow

	titleBar 	*walk.Label
)

func init() {
	xScreen = win.GetSystemMetrics(win.SM_CXSCREEN)
	yScreen = win.GetSystemMetrics(win.SM_CYSCREEN)

}

func StyleMenuRun(w *walk.MainWindow, SizeW int32, SizeH int32) {
	if dpiScale == 0 {
		dpiScale = float64(win.GetDpiForWindow(w.Handle())) / 96.0
	}
	//WindowMap[w.Name()] = w
	currStyle = win.GetWindowLong(w.Handle(), win.GWL_STYLE)
	//removes default styling
	win.SetWindowLong(w.Handle(), win.GWL_STYLE, currStyle&^win.WS_SIZEBOX&^win.WS_MINIMIZEBOX&^win.WS_MAXIMIZEBOX&^win.WS_BORDER)
	hMenu = win.GetSystemMenu(w.Handle(), false)
	win.RemoveMenu(hMenu, win.SC_CLOSE, win.MF_BYCOMMAND)
	SizeW, SizeH = CalcDpiScaledSize(SizeW, SizeH)

	win.SetWindowPos(w.Handle(), 0, (xScreen-SizeW)/2, (yScreen-SizeH)/2, SizeW, SizeH, win.SWP_FRAMECHANGED)
	//win.ShowWindow(w.Handle(), win.SW_SHOW)
	win.ShowWindow(w.Handle(), win.SW_SHOWNORMAL)
	win.SetFocus(w.Handle())
	w.Run()
}

func CalcDpiScaledSize(SizeW int32, SizeH int32) (int32, int32) {
	return int32(float64(SizeW) * dpiScale), int32(float64(SizeH) * dpiScale)
}

func ShowMenuConfig() {
	MenuConfigInit()
}

func MenuConfigInit() {
	var (
		model         = NewConfigInfoModel()
		tv            *walk.TableView
		configIni     *walk.Label
		updateConfigs *walk.PushButton

		actUpdateConfig *walk.Action
		actEditConfig   *walk.Action
		actDeleteConfig *walk.Action
	)
	configName, _ := CheckConfig()

	err := MainWindow{
		Visible:  false,
		AssignTo: &MenuConfig,
		Name:     "MenuSettings",
		Title:    util.GetSubTitle(i18n.TC("配置管理", "MENU_CONFIG.WINDOW.CONFIG_MANAGEMENT")),
		Icon:     appIcon,
		Font: Font{
			Family:    "Microsoft YaHei",
			PointSize: 9,
		},
		//MinSize: Size{Width: 600, Height: 250},
		//MaxSize: Size{Width: 800, Height: 300},
		Layout: VBox{Alignment: AlignHCenterVCenter}, //布局
		Children: []Widget{ //不动态添加控件的话，在此布局或者QT设计器设计UI文件，然后加载。
			Composite{
				Layout: VBox{
					Margins: Margins{Left: 8, Right: 8, Bottom: 0, Top: 0},
				},
				Children: []Widget{
					Label{
						Text:     "😀🐂" + util.GetSubTitle(i18n.TC("配置管理", "MENU_CONFIG.WINDOW.CONFIG_MANAGEMENT")),
						AssignTo: &configIni,
						Font: Font{Family: "Microsoft YaHei"},
					},
					Label{
						Text:     i18n.TC("当前配置: ", "MENU_CONFIG.WINDOW.CURRENT_CONFIG") + configName,
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
						MultiSelection:   false,
						Alignment:        AlignHCenterVCenter,
						Columns: []TableViewColumn{
							{Title: i18n.TC("配置名称", "MENU_CONFIG.WINDOW.CONFIG_NAME")},
							{Title: i18n.TC("文件大小", "MENU_CONFIG.WINDOW.FILE_SIZE")},
							{Title: i18n.TC("更新日期", "MENU_CONFIG.WINDOW.UPDATE_DATETIME"), Format: "01-02 15:04:05"},
							{Title: i18n.TC("订阅地址", "MENU_CONFIG.WINDOW.SUBSCRIPTION_URL"), Width: 295},
						},
						Model: model,
						OnSelectedIndexesChanged: func() {
							hasConfig := len(tv.SelectedIndexes()) == 1
							actUpdateConfig.SetEnabled(hasConfig)
							actEditConfig.SetEnabled(hasConfig)
							actDeleteConfig.SetEnabled(hasConfig)
						},
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
						Text: i18n.T(cI18n.MenuConfigWindowEnableConfig),
						MenuItems: []MenuItem{
							Action{
								AssignTo: nil,
								Text:     i18n.TC("添加配置", "MENU_CONFIG.WINDOW.ADD_CONFIG"),
								OnTriggered: func() {
									MenuConfig.SetVisible(false)
									AddConfig()
									model.ResetRows()
									time.Sleep(100 * time.Millisecond)
									MenuConfig.SetVisible(true)
								},
							},
							Action{
								AssignTo: &actUpdateConfig,
								Text:     i18n.TC("升级配置", "MENU_CONFIG.WINDOW.UPDATE_CONFIG"),
								OnTriggered: func() {
									index := tv.CurrentIndex()
									if index != -1 && model.items[index].Url != "" {
										configName := model.items[index].Name
										configUrl := model.items[index].Url
										success := updateConfig(configName, configUrl)
										if !success {
											walk.MsgBox(MenuConfig, i18n.TC("提示", "MESSAGEBOX.TITLE.TIPS"),
												i18n.TC("更新配置失败", "MENU_CONFIG.WINDOW.UPDATE_ALL"), walk.MsgBoxIconError)
											return
										}
										walk.MsgBox(MenuConfig, i18n.TC("提示", "MESSAGEBOX.TITLE.TIPS"),
											fmt.Sprintf("成功更新 %s 配置！", configName), walk.MsgBoxIconInformation)
									} else {
										walk.MsgBox(MenuConfig, i18n.TC("提示", "MESSAGEBOX.TITLE.TIPS"),
											i18n.TC("请选择要更新的配置！", "MENU_CONFIG.WINDOW.UPDATE_ALL"), walk.MsgBoxIconError)
										return
									}
									model.ResetRows()
								},
							},
							Action{
								AssignTo: &actEditConfig,
								Text:     i18n.T(cI18n.MenuConfigWindowEditConfig),
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
										walk.MsgBox(MenuConfig, i18n.TC("提示", "MESSAGEBOX.TITLE.TIPS"),
											i18n.TC("请选择要编辑的配置！", "MENU_CONFIG.WINDOW.UPDATE_ALL"), walk.MsgBoxIconError)
										return
									}
									model.ResetRows()
								},
							},
							Action{
								AssignTo: &actDeleteConfig,
								Text:     i18n.TC("删除配置", "MENU_CONFIG.WINDOW.DELETE_CONFIG"),
								OnTriggered: func() {
									index := tv.CurrentIndex()
									if index != -1 {
										deleteConfigName := model.items[index].Name
										if win.IDYES == walk.MsgBox(MenuConfig, i18n.TC("提示", "MESSAGEBOX.TITLE.TIPS"),
											i18n.TC("请确认是否删除该配置？", "MENU_CONFIG.WINDOW.UPDATE_ALL"), walk.MsgBoxYesNo) {
											err := os.Remove(path.Join(constant.CacheDir,
												deleteConfigName+constant.ConfigSuffix+constant.CacheFile))
											err = os.Remove(path.Join(constant.ConfigDir,
												deleteConfigName+constant.ConfigSuffix))
											if err != nil {
												walk.MsgBox(MenuConfig, i18n.TC("提示", "MESSAGEBOX.TITLE.TIPS"),
													i18n.TC("删除配置失败！", "MENU_CONFIG.WINDOW.UPDATE_ALL"), walk.MsgBoxIconError)
												return
											} else {
												walk.MsgBox(MenuConfig, i18n.TC("提示", "MESSAGEBOX.TITLE.TIPS"),
													fmt.Sprintf("成功删除 %s 配置！", deleteConfigName),
													walk.MsgBoxIconInformation)
											}
										}
									} else {
										walk.MsgBox(MenuConfig, i18n.TC("提示", "MESSAGEBOX.TITLE.TIPS"),
											i18n.TC("请选择要删除的配置！", "MENU_CONFIG.WINDOW.UPDATE_ALL"), walk.MsgBoxIconError)
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
								PutConfig(configName)
								walk.MsgBox(MenuConfig, i18n.T(cI18n.MessageBoxTitleTips),
									i18n.TData(cI18n.MenuConfigMessageEnableConfigSuccess, &i18n.Data{Data: map[string]interface{}{
										"Config": configName,
									}}),
									walk.MsgBoxIconInformation)
								configIni.SetText(i18n.TC("当前配置: ", "MENU_CONFIG.WINDOW.CURRENT_CONFIG") + configName + constant.ConfigSuffix)
								go func() {
									time.Sleep(1 * time.Second)
									userInfo := UpdateSubscriptionUserInfo()
									if len(userInfo.UnusedInfo) > 0 {
										notify.PushFlowInfo(userInfo.UsedInfo, userInfo.UnusedInfo, userInfo.ExpireInfo)
									}
								}()
							} else {
								walk.MsgBox(MenuConfig, i18n.TC("提示", "MESSAGEBOX.TITLE.TIPS"),
									i18n.TC("请选择要启用的配置！", "MENU_CONFIG.WINDOW.UPDATE_ALL"), walk.MsgBoxIconError)
								return
							}
							model.ResetRows()
						},
					},
					PushButton{
						Text:     i18n.TC("一键更新", "MENU_CONFIG.WINDOW.UPDATE_ALL"),
						AssignTo: &updateConfigs,
						OnClicked: func() {
							updateConfigs.SetEnabled(false)
							updateConfigs.SetText(i18n.TC("更新中", "MENU_CONFIG.WINDOW.UPDATE_ALL"))
							model.TaskCron()
							updateConfigs.SetText(i18n.TC("更新完成", "MENU_CONFIG.WINDOW.UPDATE_ALL"))
							updateConfigs.SetEnabled(true)
						},
					},
					PushButton{
						Text: i18n.TC("订阅转换", "MENU_CONFIG.WINDOW.CONVERT_SUBSCRIPTION"),
						OnClicked: func() {
							err := open.Run(constant.SubConverterUrl)
							if err != nil {
								return
							}
						},
					},
					PushButton{
						Text: i18n.TC("打开目录", "MENU_CONFIG.WINDOW.OPEN_CONFIG_DIR"),
						OnClicked: func() {
							err := open.Run(constant.ConfigDir)
							if err != nil {
								return
							}
						},
					},
					PushButton{
						Text: i18n.TC("关闭窗口", "MENU_CONFIG.WINDOW.CLOSE_WINDOW"),
						OnClicked: func() {
							err := MenuConfig.Close()
							if err != nil {
								return
							}
							MenuConfig.Dispose()
							MenuConfig = nil
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
