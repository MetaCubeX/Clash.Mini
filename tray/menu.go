package tray

import (
	"container/list"
	"fmt"
	"time"

	"github.com/Clash-Mini/Clash.Mini/cmd"
	cp "github.com/Clash-Mini/Clash.Mini/cmd/proxy"
	"github.com/Clash-Mini/Clash.Mini/cmd/sys"
	"github.com/Clash-Mini/Clash.Mini/constant"
	cI18n "github.com/Clash-Mini/Clash.Mini/constant/i18n"
	"github.com/Clash-Mini/Clash.Mini/controller"
	"github.com/Clash-Mini/Clash.Mini/icon"
	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Clash-Mini/Clash.Mini/notify"
	"github.com/Clash-Mini/Clash.Mini/sysproxy"
	"github.com/Clash-Mini/Clash.Mini/util"
	"github.com/Dreamacro/clash/config"
	C "github.com/Dreamacro/clash/constant"
	"github.com/Dreamacro/clash/hub/route"
	"github.com/Dreamacro/clash/proxy"
	"github.com/Dreamacro/clash/tunnel"
	. "github.com/JyCyunMe/go-i18n/i18n"
	stx "github.com/getlantern/systray"

	"github.com/MakeNowJust/hotkey"
)

var (
	firstInit = true
)

func init() {
	if constant.IsWindows() {
		C.SetHomeDir(constant.PWD)
	}

	InitI18n(&English, log.Infoln, log.Errorln)
	stx.RunEx(onReady, onExit)
}

func resetI18nMenuItemEx(menuItemEx *stx.MenuItemEx, i18nID string) {
	newValue := T(i18nID)
	menuItemEx.SetTitle(newValue)
	menuItemEx.SetTooltip(newValue)
}

func onReady() {

	log.Infoln("onReady")
	stx.SetIcon(icon.DateN)
	stx.SetTitle(util.AppTitle)
	stx.SetTooltip(util.AppTitle + " by Maze")

	stx.AddMainMenuItemEx(util.AppTitle, "", func(menuItemEx *stx.MenuItemEx) {
		fmt.Println("Hi Clash.Mini")
	})
	stx.AddSeparator()

	// 全局代理
	mGlobal := stx.AddMainMenuItemExI18n(&stx.I18nConfig{
		TitleID: cI18n.TrayMenuGlobalProxy,
		TitleFormat: "\tAlt+G",
		TooltipID: cI18n.TrayMenuGlobalProxy,
	}, func(menuItemEx *stx.MenuItemEx) {
		tunnel.SetMode(tunnel.Global)
		firstInit = true
	})
	// 规则代理
	mRule := stx.AddMainMenuItemExI18n(&stx.I18nConfig{
		TitleID: cI18n.TrayMenuRuleProxy,
		TitleFormat: "\tAlt+R",
		TooltipID: cI18n.TrayMenuRuleProxy,
	}, func(menuItemEx *stx.MenuItemEx) {
		tunnel.SetMode(tunnel.Rule)
		firstInit = true
	})
	// 全局直连
	mDirect := stx.AddMainMenuItemExI18n(&stx.I18nConfig{
		TitleID: cI18n.TrayMenuDirectProxy,
		TitleFormat: "\tAlt+D",
		TooltipID: cI18n.TrayMenuDirectProxy,
	}, func(menuItemEx *stx.MenuItemEx) {
		tunnel.SetMode(tunnel.Direct)
		firstInit = true
	})
	stx.AddSeparator()

	// 切换节点
	mGroup := stx.AddMainMenuItemExI18n(&stx.I18nConfig { TitleID: cI18n.TrayMenuSwitchProxy }, stx.NilCallback)
	if ConfigGroupsMap == nil {
		config.ParsingProxiesCallback = func(groupsList *list.List, proxiesList *list.List) {
			RefreshProxyGroups(mGroup, groupsList, proxiesList)
			NeedLoadSelector = true
		}
		route.SwitchProxiesCallback = func(sGroup string, sProxy string) {
			SwitchGroupAndProxy(mGroup, sGroup, sProxy)
		}
	}
	var mPingTest = &stx.MenuItemEx{}
	// 延迟测速
	// 当前节点延迟
	stx.AddMainMenuItemExBind(TC("延迟测速", "TRAY_MENU.PING_TEST") + "\t10ms",
		TC("延迟测速", "TRAY_MENU.PING_TEST") + "\t1分钟前", stx.NilCallback, mPingTest).
		// 最低延迟:
		AddSubMenuItemEx(TC("最低延迟: ", "TRAY_MENU.PING_TEST.LOWEST_DELAY") + "10ms",
		TC("最低延迟: ", "TRAY_MENU.PING_TEST.LOWEST_DELAY") + "10ms", stx.NilCallback).
		// 最快节点:
		AddMenuItemEx(TC("最快节点: ", "TRAY_MENU.PING_TEST.FAST_PROXY") + "HK-101",
		TC("最快节点: ", "TRAY_MENU.PING_TEST.FAST_PROXY") + "HK-101", stx.NilCallback).
		// 上次更新:
		AddMenuItemEx(TC("上次更新: ", "TRAY_MENU.PING_TEST.LAST_UPDATE") + "1分钟前",
		TC("上次更新: ", "TRAY_MENU.PING_TEST.LAST_UPDATE") + "1分钟前", stx.NilCallback)
	stx.AddSeparator()
	AddSwitchCallback(func() {
		mGlobal.SwitchLanguage()
		mRule.SwitchLanguage()
		mDirect.SwitchLanguage()
		mGroup.SwitchLanguage()
		mPingTest.SwitchLanguageWithChildren()
	})

	// TODO: add config switch
	// 切换订阅
	mSwitchConfig := stx.AddMainMenuItemExI18n(&stx.I18nConfig { TitleID: cI18n.TrayMenuSwitchConfig }, stx.NilCallback)
	stx.AddSeparator()

	// 系统代理
	mEnabled := stx.AddMainMenuItemExI18n(&stx.I18nConfig{
		TitleID: cI18n.TrayMenuSystemProxy,
		TitleFormat: "\tAlt+S",
		TooltipID: cI18n.TrayMenuSystemProxy,
	}, mEnabledFunc)
	// 控制面板
	mDashboard := stx.AddMainMenuItemExI18n(&stx.I18nConfig { TitleID: cI18n.TrayMenuDashboard }, func(menuItemEx *stx.MenuItemEx) {
		go controller.Dashboard()
	})
	// 配置管理
	mConfig := stx.AddMainMenuItemExI18n(&stx.I18nConfig { TitleID: cI18n.TrayMenuConfigManagement }, func(menuItemEx *stx.MenuItemEx) {
		go controller.ShowMenuConfig()
	})
	AddSwitchCallback(func() {
		mSwitchConfig.SwitchLanguage()
		mEnabled.SwitchLanguage()
		mDashboard.SwitchLanguage()
		mConfig.SwitchLanguage()
	})

	var mOthers = &stx.MenuItemEx{}
	var mI18nSwitcher = &stx.MenuItemEx{}
	var mOthersTask = &stx.MenuItemEx{}
	var mOthersAutosys = &stx.MenuItemEx{}
	var mOthersUpdateCron = &stx.MenuItemEx{}
	var maxMindMMDB = &stx.MenuItemEx{}
	var hackl0usMMDB = &stx.MenuItemEx{}
	// 其他设置
	stx.AddMainMenuItemExBindI18n(&stx.I18nConfig { TitleID: cI18n.TrayMenuOtherSettings }, stx.NilCallback, mOthers).
		// 切换语言
		AddSubMenuItemExBindI18n(&stx.I18nConfig { TitleID: cI18n.TrayMenuOtherSettingsSwitchLanguage }, stx.NilCallback, mI18nSwitcher).
		// 设置开机启动
		AddMenuItemExBindI18n(&stx.I18nConfig { TitleID: cI18n.TrayMenuOtherSettingsSystemAutorun }, mOtherTaskFunc, mOthersTask).
		// 设置默认代理
		AddMenuItemExBindI18n(&stx.I18nConfig { TitleID: cI18n.TrayMenuOtherSettingsSystemAutoProxy }, mOtherAutosysFunc, mOthersAutosys).
		// 设置定时更新
		AddMenuItemExBindI18n(&stx.I18nConfig { TitleID: cI18n.TrayMenuOtherSettingsCronUpdateConfigs }, mOtherUpdateCronFunc, mOthersUpdateCron).
		// 设置GeoIP2数据库
		AddMenuItemExI18n(&stx.I18nConfig { TitleID: cI18n.TrayMenuOtherSettingsSetMMDB }, stx.NilCallback).
		// MaxMind数据库
		AddSubMenuItemExBindI18n(&stx.I18nConfig { TitleID: cI18n.TrayMenuOtherSettingsSetMMDBMaxmind }, maxMindMMBDFunc, maxMindMMDB).
		// Hackl0us数据库
		AddMenuItemExBindI18n(&stx.I18nConfig { TitleID: cI18n.TrayMenuOtherSettingsSetMMDBHackl0Us }, hackl0usMMDBFunc, hackl0usMMDB)
	for _, l := range Languages {
		lang := l
		langName := fmt.Sprintf("%s (%s)", lang.Name, lang.Tag.String())
		mLang := mI18nSwitcher.AddSubMenuItemEx(langName, langName, func(menuItemEx *stx.MenuItemEx) {
			log.Infoln("[i18n] switch language to %s", langName)
			SwitchLanguage(lang)
			menuItemEx.SwitchCheckboxBrother(true)
		})
		if Language != nil && Language.Tag == lang.Tag {
			mLang.SwitchCheckboxBrother(true)
		}
	}
	stx.AddSeparator()

	// 退出
	mQuit := stx.AddMainMenuItemExI18n(&stx.I18nConfig{ TitleID: cI18n.TrayMenuQuit }, func(menuItemEx *stx.MenuItemEx) {
		stx.Quit()
		return
	})
	AddSwitchCallback(func() {
		mOthers.SwitchLanguageWithChildren()
		mQuit.SwitchLanguage()
	})

	if !constant.IsWindows() {
		mEnabled.Hide()
		mOthers.Hide()
		mConfig.Hide()
	}

	proxyModeGroup := []*stx.MenuItemEx{mGlobal, mRule, mDirect}
	mmdbGroup := []*stx.MenuItemEx{maxMindMMDB, hackl0usMMDB}
	hotKey(mEnabled)

	go func() {
		t := time.NewTicker(time.Second)
		defer t.Stop()
		SavedPort := proxy.GetPorts().Port
		if controller.RegCompare(cmd.Sys) {
			var Ports int
			if proxy.GetPorts().MixedPort != 0 {
				Ports = proxy.GetPorts().MixedPort
			} else {
				Ports = proxy.GetPorts().Port
			}
			err := sysproxy.SetSystemProxy(
				&sysproxy.ProxyConfig{
					Enable: true,
					Server: fmt.Sprintf("%s:%d", constant.Localhost, Ports),
				})
			if err != nil {
				log.Errorln("SetSystemProxy error: %v", err)
				notify.PushWithLine("❌错误❌", "设置系统代理时出错")
				return
			}
			mEnabled.Check()
			notify.DoTrayMenu(sys.ON)
		}
		if controller.RegCompare(cmd.Cron) {
			mOthersUpdateCron.Check()
			go controller.CronTask()
		}
		//if config.GroupsList.Len() > 0 {
		//	log.Infoln("--")
		//	//log.Infoln(config.GroupsList)
		//	RefreshProxyGroups(mGroup, config.GroupsList, config.ProxiesList)
		//}

		for {
			<-t.C
			switch tunnel.Mode() {
			case tunnel.Global:
				if mGlobal.Checked() {
				} else {
					RefreshProxyGroups(mGroup, nil, config.ProxiesList)
					NeedLoadSelector = true
					stx.SwitchCheckboxGroup(mGlobal, true, proxyModeGroup)
					mGroup.Enable()
					if mEnabled.Checked() {
						stx.SetIcon(icon.DateG)
						notify.DoTrayMenu(cp.Global)
					} else {
						stx.SetIcon(icon.DateN)
					}
				}
			case tunnel.Rule:
				if mRule.Checked() {
				} else {
					RefreshProxyGroups(mGroup, config.GroupsList, config.ProxiesList)
					NeedLoadSelector = true
					stx.SwitchCheckboxGroup(mRule, true, proxyModeGroup)
					mGroup.Enable()
					if mEnabled.Checked() {
						stx.SetIcon(icon.DateS)
						notify.DoTrayMenu(cp.Rule)
					} else {
						stx.SetIcon(icon.DateN)
					}
				}
			case tunnel.Direct:
				if mDirect.Checked() {
				} else {
					RefreshProxyGroups(mGroup, nil, nil)
					mGroup.Disable()
					stx.SwitchCheckboxGroup(mDirect, true, proxyModeGroup)
					if mEnabled.Checked() {
						stx.SetIcon(icon.DateD)
						notify.DoTrayMenu(cp.Direct)
					} else {
						stx.SetIcon(icon.DateN)
					}
				}
			}
			if firstInit {
				if controller.RegCompare(cmd.Task) {
					mOthersTask.Check()
				} else {
					mOthersTask.Uncheck()
				}

				if controller.RegCompare(cmd.MMDB) {
					stx.SwitchCheckboxGroup(hackl0usMMDB, true, mmdbGroup)
				} else {
					stx.SwitchCheckboxGroup(maxMindMMDB, true, mmdbGroup)
				}

				if controller.RegCompare(cmd.Sys) {
					mOthersAutosys.Check()
				} else {
					mOthersAutosys.Uncheck()
				}

				if controller.RegCompare(cmd.Cron) {
					mOthersUpdateCron.Check()
				} else {
					mOthersUpdateCron.Uncheck()
				}

				if mEnabled.Checked() {
					var p int
					if proxy.GetPorts().MixedPort != 0 {
						p = proxy.GetPorts().MixedPort
					} else {
						p = proxy.GetPorts().Port
					}
					if SavedPort != p {
						SavedPort = p
						err := sysproxy.SetSystemProxy(
							&sysproxy.ProxyConfig{
								Enable: true,
								Server: fmt.Sprintf("%s:%d", constant.Localhost, SavedPort),
							})
						if err != nil {
							continue
						}
					}
				}

				p, err := sysproxy.GetCurrentProxy()
				if err != nil {
					continue
				}

				if p.Enable && p.Server == fmt.Sprintf("%s:%d", constant.Localhost, SavedPort) {
					if mEnabled.Checked() {
					} else {
						mEnabled.Check()
					}
				} else {
					if mEnabled.Checked() {
						mEnabled.Uncheck()
					} else {
					}
				}
				firstInit = false
			}
			LoadSelector(mGroup)
		}

	}()

	go func() {
		userInfo := controller.UpdateSubscriptionUserInfo()
		time.Sleep(2 * time.Second)
		if len(userInfo.UnusedInfo) > 0 {
			notify.PushFlowInfo(userInfo.UsedInfo, userInfo.UnusedInfo, userInfo.ExpireInfo)
		}
	}()

}

func onExit() {
	err := sysproxy.SetSystemProxy(sysproxy.GetSavedProxy())
	if err != nil {
		log.Errorln("onExit SetSystemProxy error: %v", err)
	}
}

func hotKey(mEnabled *stx.MenuItemEx) {
	message := ""
	hkey := hotkey.New()
	_, err1 := hkey.Register(hotkey.Alt, 'R', func() {
		tunnel.SetMode(tunnel.Rule)
	})
	if err1 != nil {
		message += "Alt+R热键注册失败\n"
	}
	_, err2 := hkey.Register(hotkey.Alt, 'G', func() {
		tunnel.SetMode(tunnel.Global)
	})
	if err2 != nil {
		message += "Alt+G热键注册失败\n"
	}
	_, err3 := hkey.Register(hotkey.Alt, 'D', func() {
		tunnel.SetMode(tunnel.Direct)
	})
	if err3 != nil {
		message += "Alt+D热键注册失败\n"
	}
	_, err4 := hkey.Register(hotkey.Alt, 'S', func() {
		mEnabledFunc(mEnabled)
	})
	if err4 != nil {
		message += "Alt+S热键注册失败\n"
	}
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		go notify.PushWithLine("📢通知📢", message)
	}
}
