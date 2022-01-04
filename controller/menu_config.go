package controller

import (
	"fmt"
	"github.com/Clash-Mini/Clash.Mini/config"
	"os"
	path "path/filepath"
	"time"

	"github.com/Clash-Mini/Clash.Mini/constant"
	cI18n "github.com/Clash-Mini/Clash.Mini/constant/i18n"
	"github.com/Clash-Mini/Clash.Mini/notify"
	p "github.com/Clash-Mini/Clash.Mini/profile"
	stringUtils "github.com/Clash-Mini/Clash.Mini/util/string"

	"github.com/JyCyunMe/go-i18n/i18n"
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
	dpiScale   float64

	WindowMap  = make(map[string]*walk.MainWindow)
	MenuConfig *walk.MainWindow

	titleBar *walk.Label

	firstInit bool
	t         = time.NewTicker(time.Second)
)

func StyleMenuRun(w *walk.MainWindow, SizeW int32, SizeH int32) {
	//WindowMap[w.Name()] = w
	currStyle = win.GetWindowLong(w.Handle(), win.GWL_STYLE)
	//removes default styling
	win.SetWindowLong(w.Handle(), win.GWL_STYLE, currStyle&^win.WS_SIZEBOX&^win.WS_MINIMIZEBOX&^win.WS_MAXIMIZEBOX)
	hMenu = win.GetSystemMenu(w.Handle(), false)
	//win.RemoveMenu(hMenu, win.SC_CLOSE, win.MF_BYCOMMAND)
	SizeW, SizeH = CalcDpiScaledSize(SizeW, SizeH)

	win.SetWindowPos(w.Handle(), 0, (xScreen-SizeW)/2, (yScreen-SizeH)/2, SizeW, SizeH, win.SWP_FRAMECHANGED)
	//win.ShowWindow(w.Handle(), win.SW_SHOW)
	win.ShowWindow(w.Handle(), win.SW_SHOWNORMAL)
	win.SetFocus(w.Handle())
	w.Run()
}

func ShowMenuConfig() {
	MenuConfigInit()
}

func MenuConfigInit() {
	var (
		model            = NewConfigInfoModel()
		tv               *walk.TableView
		configIni        *walk.Label
		configInfo       *walk.Label
		enableConfig     *walk.Action
		updateAllConfigs *walk.PushButton

		actUpdateConfig *walk.Action
		actEditConfig   *walk.Action
		actDeleteConfig *walk.Action
	)
	configName := config.GetProfile()
	currentName := configName
	CurrentConfigInfo := i18n.T(cI18n.MenuConfigWindowCurrentConfig) + ` : ` + configName + constant.ConfigSuffix
	FlowExpiration := i18n.T(cI18n.NotifyMessageFlowExpiration) + ` : ` + p.UserInfo.ExpireInfo
	FlowUsed := i18n.T(cI18n.NotifyMessageFlowUsed) + ` : ` + p.UserInfo.UsedInfo
	FlowUnused := i18n.T(cI18n.NotifyMessageFlowUnused) + ` : ` + p.UserInfo.UnusedInfo

	err := MainWindow{
		Visible:  false,
		AssignTo: &MenuConfig,
		Name:     "MenuSettings",
		Title:    stringUtils.GetSubTitle(i18n.T(cI18n.MenuConfigWindowConfigManagement)),
		Icon:     appIcon,
		Font: Font{
			Family:    "Microsoft YaHei",
			PointSize: 9,
		},

		Layout: VBox{Alignment: AlignHCenterVCenter}, //布局
		Children: []Widget{ //不动态添加控件的话，在此布局或者QT设计器设计UI文件，然后加载。
			Composite{
				Layout: HBox{
					Margins: Margins{Left: 8, Right: 8, Bottom: 0, Top: 0},
				},
				Children: []Widget{
					Label{
						Text:     CurrentConfigInfo,
						AssignTo: &configIni,
						MinSize:  Size{Width: 200},
					},
					Label{
						Text:     FlowUsed,
						AssignTo: &configInfo,
					},
					HSpacer{},
				},
			},
			Composite{
				Layout: HBox{
					Margins: Margins{Left: 8, Right: 8, Bottom: 0, Top: 0},
				},
				Children: []Widget{
					Label{
						Text:     FlowExpiration,
						AssignTo: &configInfo,
						MinSize:  Size{Width: 200},
					},
					Label{
						Text:     FlowUnused,
						AssignTo: &configInfo,
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
						AssignTo:                    &tv,
						CheckBoxes:                  false,
						ColumnsOrderable:            false,
						MultiSelection:              false,
						SelectionHiddenWithoutFocus: true,
						//AlternatingRowBG: true,
						// TODO: 高亮启用的配置
						Alignment: AlignHCenterVCenter,
						Columns: []TableViewColumn{
							{Title: "-", Width: 20},
							{Title: i18n.T(cI18n.MenuConfigWindowConfigName)},
							{Title: i18n.T(cI18n.MenuConfigWindowFileSize)},
							{Title: i18n.T(cI18n.MenuConfigWindowUpdateDatetime), Format: "01-02 15:04:05"},
							{Title: i18n.T(cI18n.MenuConfigWindowSubscriptionUrl), Width: 275},
						},
						Model: model,
						OnCurrentIndexChanged: func() {
							hasConfig := tv.CurrentIndex() > -1
							//walk.MsgBox(MenuConfig, i18n.T(cI18n.MsgBoxTitleTips),
							//	fmt.Sprintf("IndexChanged: %d, %t", tv.CurrentIndex(), hasConfig), walk.MsgBoxIconInformation)
							enableConfig.SetVisible(hasConfig && !model.items[tv.CurrentIndex()].checked)
							actUpdateConfig.SetVisible(hasConfig)
							actEditConfig.SetVisible(hasConfig)
							actDeleteConfig.SetVisible(hasConfig)
						},
						ContextMenuItems: []MenuItem{
							Action{
								AssignTo: &enableConfig,
								Text:     i18n.T(cI18n.MenuConfigWindowEnableConfig),
								Shortcut: Shortcut{Modifiers: walk.ModAlt, Key: walk.KeyN},
								OnTriggered: func() {
									index := tv.CurrentIndex()
									if index != -1 {
										configName := model.items[index].Name
										if currentName == configName {
											ApplyConfig(configName, true)
										} else {
											ApplyConfig(configName, false)
										}
										walk.MsgBox(MenuConfig, i18n.T(cI18n.MsgBoxTitleTips),
											i18n.TData(cI18n.MenuConfigMessageEnableConfigSuccess, &i18n.Data{Data: map[string]interface{}{
												"Config": configName,
											}}),
											walk.MsgBoxIconInformation)
										configIni.SetText(fmt.Sprintf("%s: %s%s", i18n.T(cI18n.MenuConfigWindowCurrentConfig), configName, constant.ConfigSuffix))
										go func() {
											time.Sleep(1 * time.Second)
											userInfo := p.UpdateSubscriptionUserInfo()
											if len(userInfo.UnusedInfo) > 0 {
												notify.PushFlowInfo(userInfo.UsedInfo, userInfo.UnusedInfo, userInfo.ExpireInfo)
											}
										}()
										CurrentProfile = config.GetProfile()
									}
									model.ResetRows()
									firstInit = true
								},
							},
							Action{
								AssignTo: &actUpdateConfig,
								Text:     i18n.T(cI18n.MenuConfigWindowUpdateConfig),
								Shortcut: Shortcut{Modifiers: walk.ModAlt, Key: walk.KeyU},
								Visible:  true,
								OnTriggered: func() {
									index := tv.CurrentIndex()
									if index != -1 && model.items[index].Url != "" {
										configName := model.items[index].Name
										configUrl := model.items[index].Url
										successful := p.UpdateConfig(configName, configUrl)
										if !successful {
											walk.MsgBox(MenuConfig, i18n.T(cI18n.MsgBoxTitleTips),
												i18n.TData(cI18n.MenuConfigMessageUpdateConfigFailure, &i18n.Data{Data: map[string]interface{}{
													"Config": configName,
												}}), walk.MsgBoxIconError)
											return
										}
										if currentName == configName {
											ApplyConfig(configName, true)
										}
										walk.MsgBox(MenuConfig, i18n.T(cI18n.MsgBoxTitleTips),
											i18n.TData(cI18n.MenuConfigMessageUpdateConfigSuccess, &i18n.Data{Data: map[string]interface{}{
												"Config": configName,
											}}), walk.MsgBoxIconInformation)
									}
									model.ResetRows()
									firstInit = true
								},
							},
							Action{
								AssignTo: &actEditConfig,
								Text:     i18n.T(cI18n.MenuConfigWindowEditConfig),
								Shortcut: Shortcut{Modifiers: walk.ModAlt, Key: walk.KeyE},
								Visible:  true,
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
									}
									model.ResetRows()
									firstInit = true
								},
							},
							Action{
								AssignTo: &actDeleteConfig,
								Text:     i18n.T(cI18n.MenuConfigWindowDeleteConfig),
								Shortcut: Shortcut{Modifiers: walk.ModAlt, Key: walk.KeyD},
								Visible:  true,
								OnTriggered: func() {
									index := tv.CurrentIndex()
									if index != -1 {
										deleteConfigName := model.items[index].Name
										if win.IDYES == walk.MsgBox(MenuConfig, i18n.T(cI18n.MsgBoxTitleTips),
											i18n.TData(cI18n.MenuConfigMessageDeleteConfigConfirmMsg, &i18n.Data{Data: map[string]interface{}{
												"Config": deleteConfigName,
											}}), walk.MsgBoxYesNo) {
											err := os.Remove(path.Join(constant.CacheDir, deleteConfigName+constant.ConfigSuffix+constant.CacheFile))
											err = os.Remove(path.Join(constant.ProfileDir, deleteConfigName+constant.ConfigSuffix))
											if err != nil {
												walk.MsgBox(MenuConfig, i18n.T(cI18n.MsgBoxTitleTips),
													i18n.TData(cI18n.MenuConfigMessageDeleteConfigFailure, &i18n.Data{Data: map[string]interface{}{
														"Config": deleteConfigName,
													}}), walk.MsgBoxIconError)
												return
											} else {
												walk.MsgBox(MenuConfig, i18n.T(cI18n.MsgBoxTitleTips),
													i18n.TData(cI18n.MenuConfigMessageDeleteConfigSuccess, &i18n.Data{Data: map[string]interface{}{
														"Config": deleteConfigName,
													}}), walk.MsgBoxIconInformation)
											}
										}
									}
									model.ResetRows()
									firstInit = true
								},
							},
							Separator{},
							Action{
								AssignTo: nil,
								Text:     i18n.T(cI18n.MenuConfigWindowAddConfig),
								Shortcut: Shortcut{Modifiers: walk.ModAlt, Key: walk.KeyA},
								OnTriggered: func() {
									MenuConfig.SetVisible(false)
									AddConfig()
									model.ResetRows()
									time.Sleep(100 * time.Millisecond)
									firstInit = true
									MenuConfig.SetVisible(true)
								},
							},
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
					PushButton{
						Text:     i18n.T(cI18n.MenuConfigWindowUpdateAll),
						AssignTo: &updateAllConfigs,
						OnClicked: func() {
							updateAllConfigs.SetEnabled(false)
							updateAllConfigs.SetText(i18n.T(cI18n.MenuConfigWindowUpdating))
							model.TaskCron()
							updateAllConfigs.SetText(i18n.T(cI18n.MenuConfigWindowUpdateFinished))
							updateAllConfigs.SetEnabled(true)
						},
					},
					PushButton{
						Text: i18n.T(cI18n.MenuConfigWindowConvertSubscription),
						OnClicked: func() {
							err := open.Run(constant.SubConverterUrl)
							if err != nil {
								return
							}
						},
					},
					PushButton{
						Text: i18n.T(cI18n.MenuConfigWindowOpenConfigDir),
						OnClicked: func() {
							err := open.Run(constant.ProfileDir)
							if err != nil {
								return
							}
						},
					},
					//PushButton{
					//	Text: i18n.T(cI18n.MenuConfigWindowCloseWindow),
					//	OnClicked: func() {
					//		err := MenuConfig.Close()
					//		if err != nil {
					//			return
					//		}
					//		MenuConfig.Dispose()
					//		MenuConfig = nil
					//	},
					//},
				},
			},
		},
	}.Create()
	if err != nil {
		return
	}
	firstInit = true
	go func() {
		for {
			<-t.C
			if firstInit {
				configName = config.GetProfile()
				CurrentProfile = configName
				for _, item := range model.items {
					if item.Name == CurrentProfile {
						item.checked = true
						break
					}
				}
				firstInit = false
			}
		}
	}()
	StyleMenuRun(MenuConfig, 650, 250)
}
