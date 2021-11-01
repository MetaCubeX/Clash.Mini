package tray

import (
	"container/list"
	"fmt"
	"time"

	"github.com/Clash-Mini/Clash.Mini/app"
	"github.com/Clash-Mini/Clash.Mini/cmd"
	cmdP "github.com/Clash-Mini/Clash.Mini/cmd/proxy"
	"github.com/Clash-Mini/Clash.Mini/cmd/sys"
	"github.com/Clash-Mini/Clash.Mini/common"
	"github.com/Clash-Mini/Clash.Mini/config"
	"github.com/Clash-Mini/Clash.Mini/constant"
	cI18n "github.com/Clash-Mini/Clash.Mini/constant/i18n"
	"github.com/Clash-Mini/Clash.Mini/controller"
	"github.com/Clash-Mini/Clash.Mini/icon"
	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Clash-Mini/Clash.Mini/notify"
	"github.com/Clash-Mini/Clash.Mini/proxy"
	"github.com/Clash-Mini/Clash.Mini/sysproxy"
	"github.com/Clash-Mini/Clash.Mini/util"
	. "github.com/Clash-Mini/Clash.Mini/util/maybe"

	clashConfig "github.com/Dreamacro/clash/config"
	C "github.com/Dreamacro/clash/constant"
	"github.com/Dreamacro/clash/hub/route"
	clashP "github.com/Dreamacro/clash/proxy"
	"github.com/Dreamacro/clash/tunnel"
	. "github.com/JyCyunMe/go-i18n/i18n"
	stx "github.com/getlantern/systray"
	"golang.org/x/text/language"
)

var (
	firstInit   = true
	loadProfile = true
)

func init() {
	if constant.IsWindows() {
		C.SetHomeDir(constant.PWD)
	}

	SetDefaultLang(English)
	langName := config.GetOrDefault("lang", DefaultLang.Name).(string)
	var lang *Lang
	tag, err := language.Parse(langName)
	if err != nil {
		log.Errorln("[i18n] language \"%s\" is invalid, will use default: %s (%s)",
			langName, DefaultLang.Name, DefaultLang.Tag.String())
		config.Set("lang", DefaultLang.Tag.String())
	} else {
		lang = &Lang{Tag: tag}
	}

	InitI18n(lang, log.Infoln, log.Errorln)
	stx.RunEx(onReady, onExit)
}

func onReady() {
	log.Infoln("onReady")
	stx.SetIcon(icon.DateN)
	mainTitle := util.GetMenuItemFullTitle(app.Name, "v" + app.Version)
	mainTooltip := app.Name + " by Maze"
	stx.SetTitle(mainTitle)
	stx.SetTooltip(mainTooltip)

	stx.AddMainMenuItemEx(mainTitle, mainTooltip, func(menuItemEx *stx.MenuItemEx) {
		fmt.Println("Hi Clash.Mini")
	})
	stx.AddSeparator()

	// 全局代理
	mGlobal := stx.AddMainMenuItemExI18n(stx.NewI18nConfig(stx.I18nConfig{
		TitleID:     cI18n.TrayMenuGlobalProxy,
		TitleFormat: "\tAlt+G",
		TooltipID:   cI18n.TrayMenuGlobalProxy,
	}), func(menuItemEx *stx.MenuItemEx) {
		tunnel.SetMode(tunnel.Global)
		firstInit = true
	})
	// 规则代理
	mRule := stx.AddMainMenuItemExI18n(stx.NewI18nConfig(stx.I18nConfig{
		TitleID:     cI18n.TrayMenuRuleProxy,
		TitleFormat: "\tAlt+R",
		TooltipID:   cI18n.TrayMenuRuleProxy,
	}), func(menuItemEx *stx.MenuItemEx) {
		tunnel.SetMode(tunnel.Rule)
		firstInit = true
	})
	// 全局直连
	mDirect := stx.AddMainMenuItemExI18n(stx.NewI18nConfig(stx.I18nConfig{
		TitleID:     cI18n.TrayMenuDirectProxy,
		TitleFormat: "\tAlt+D",
		TooltipID:   cI18n.TrayMenuDirectProxy,
	}), func(menuItemEx *stx.MenuItemEx) {
		tunnel.SetMode(tunnel.Direct)
		firstInit = true
	})
	stx.AddSeparator()

	// TEST: showCustomizedTrayMenu
	//stx.AddMainMenuItemExI18n(stx.NewI18nConfig(stx.I18nConfig{ TitleID: cI18n.TrayMenuSwitchProxy }), func(menuItemEx *stx.MenuItemEx) {
	//	controller.TrayMenuInit()
	//})
	// 切换节点
	mGroup := stx.AddMainMenuItemExI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuSwitchProxy}), stx.NilCallback)
	if ConfigGroupsMap == nil {
		clashConfig.ParsingProxiesCallback = func(groupsList *list.List, proxiesList *list.List) {
			RefreshProxyGroups(mGroup, groupsList, proxiesList)
			NeedLoadSelector = true
		}
		route.SwitchProxiesCallback = func(sGroup string, sProxy string) {
			SwitchGroupAndProxy(mGroup, sGroup, sProxy)
		}
	}
	var mPingTest = &stx.MenuItemEx{}
	var mPingTestLowestPing = &stx.MenuItemEx{}
	var mPingTestFastProxy = &stx.MenuItemEx{}
	var mPingTestLastUpdate = &stx.MenuItemEx{}
	// 延迟测速
	// 当前节点延迟
	stx.AddMainMenuItemExBindI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuPingTest}), stx.NilCallback, mPingTest).
		// 最低延迟:
		AddSubMenuItemExBindI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuPingTestLowestDelay}), stx.NilCallback, mPingTestLowestPing).
		// 最快节点:
		AddMenuItemExBindI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuPingTestFastProxy}), stx.NilCallback, mPingTestFastProxy).
		// 上次更新:
		AddMenuItemExBindI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuPingTestLastUpdate}), stx.NilCallback, mPingTestLastUpdate).
		// 立即更新
		AddMenuItemExI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuPingTestDoNow}),
			func(menuItemEx *stx.MenuItemEx) {
				proxy.RefreshAllDelay(func(name string, delay int16) {
					AddSwitchCallbackDo(&CallbackData{Callback: func(params ...interface{}) {
						sList, exist := mProxyMap[name]
						if !exist || len(sList) < 0 {
							return
						}
						for _, pm := range sList {
							if pm.Children.Len() > 0 {
								continue
							}
							var lastDelay string
							if exist && delay > -1 && uint16(delay) < max {
								lastDelay = TData(cI18n.UtilDatetimeShortMilliSeconds,
									&Data{Data: map[string]interface{}{"ms": delay}})
							} else {
								lastDelay = T(cI18n.ProxyTestTimeout)
							}
							pm.SetTitle(util.GetMenuItemFullTitle(pm.GetTooltip(), lastDelay))
							Maybe().OfNullable(pm.ExtraData).IfOk(func(o interface{}) {
								pp := o.(*proxy.Proxy)
								pp.Delay = delay
								go PingTestInfo.SetFastProxy(pp)
							})
						}
					}})
				}, func(delayMap map[string]int16) {
					//RefreshProxyDelay(mGroup, delayMap)
					//RefreshProxyGroups(mGroup, clashConfig.GroupsList, clashConfig.ProxiesList)
				})
			})
	stx.AddSeparator()
	PingTestInfo.Callback = func(pt *PingTest) {
		var lowestPing string
		var fastProxy string
		var lastUpdateDT string
		if pt == nil {
			lowestPing = "-"
			fastProxy = "-"
			lastUpdateDT = "-"
		} else {
			defer func() {
				pt.locker.RUnlock()
			}()
			pt.locker.RLock()
			//lowestPing = fmt.Sprintf("%d", pt.LowestDelay)
			lowestPing = TData(cI18n.UtilDatetimeShortMilliSeconds,
				&Data{Data: map[string]interface{}{"ms": pt.LowestDelay}})
			fastProxy = pt.FastProxy.Name
			lastUpdateDT = util.GetHumanTime(pt.LastUpdateDT)
		}

		mPingTestLowestPing.I18nConfig.TitleConfig.Format = fmt.Sprintf("\t%s", lowestPing)
		mPingTestLowestPing.SwitchLanguage()
		//mPingTest.SwitchLanguage()
		mPingTestFastProxy.I18nConfig.TitleConfig.Format = fmt.Sprintf("\t%s", fastProxy)
		mPingTestFastProxy.SwitchLanguage()
		mPingTestLastUpdate.I18nConfig.TitleConfig.Format = fmt.Sprintf("\t%s", lastUpdateDT)
		mPingTestLastUpdate.SwitchLanguage()
	}
	AddSwitchCallback(&CallbackData{Callback: func(params ...interface{}) {
		mGlobal.SwitchLanguage()
		mRule.SwitchLanguage()
		mDirect.SwitchLanguage()
		mGroup.SwitchLanguage()
		mPingTest.SwitchLanguageWithChildren()
	}})
	PingTestInfo.Callback(nil)

	// 切换订阅
	mSwitchProfile := stx.AddMainMenuItemExI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuSwitchProfile}), stx.NilCallback)
	stx.AddSeparator()
	SetMSwitchProfile(mSwitchProfile)

	// 系统代理
	mEnabled := stx.AddMainMenuItemExI18n(stx.NewI18nConfig(stx.I18nConfig{
		TitleID:     cI18n.TrayMenuSystemProxy,
		TitleFormat: "\tAlt+S",
		TooltipID:   cI18n.TrayMenuSystemProxy,
	}), mEnabledFunc)
	// 控制面板
	mDashboard := stx.AddMainMenuItemExI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuDashboard}), func(menuItemEx *stx.MenuItemEx) {
		go controller.Dashboard()
	})
	// 配置管理
	mConfig := stx.AddMainMenuItemExI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuConfigManagement}), func(menuItemEx *stx.MenuItemEx) {
		go controller.ShowMenuConfig()
	})
	// 查看日志
	mLogger := stx.AddMainMenuItemExI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuShowLog}), func(menuItemEx *stx.MenuItemEx) {
		//go controller.ShowMenuConfig()
	})
	AddSwitchCallback(&CallbackData{Callback: func(params ...interface{}) {
		mSwitchProfile.SwitchLanguage()
		mEnabled.SwitchLanguage()
		mDashboard.SwitchLanguage()
		mConfig.SwitchLanguage()
		mLogger.SwitchLanguage()
	}})

	var mOthers = &stx.MenuItemEx{}
	var mI18nSwitcher = &stx.MenuItemEx{}
	var mOthersTask = &stx.MenuItemEx{}
	var mOthersAutosys = &stx.MenuItemEx{}
	var mOthersUpdateCron = &stx.MenuItemEx{}
	var maxMindMMDB = &stx.MenuItemEx{}
	var hackl0usMMDB = &stx.MenuItemEx{}
	// 其他设置
	stx.AddMainMenuItemExBindI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuOtherSettings}), stx.NilCallback, mOthers).
		// 切换语言
		AddSubMenuItemExBindI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuOtherSettingsSwitchLanguage}), stx.NilCallback, mI18nSwitcher).
		// 设置开机启动
		AddMenuItemExBindI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuOtherSettingsSystemAutorun}), mOtherTaskFunc, mOthersTask).
		// 设置默认代理
		AddMenuItemExBindI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuOtherSettingsSystemAutoProxy}), mOtherAutosysFunc, mOthersAutosys).
		// 设置定时更新
		AddMenuItemExBindI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuOtherSettingsCronUpdateConfigs}), mOtherUpdateCronFunc, mOthersUpdateCron).
		// 设置GeoIP2数据库
		AddMenuItemExI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuOtherSettingsSetMMDB}), stx.NilCallback).
		// MaxMind数据库
		AddSubMenuItemExBindI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuOtherSettingsSetMMDBMaxmind}), maxMindMMBDFunc, maxMindMMDB).
		// Hackl0us数据库
		AddMenuItemExBindI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuOtherSettingsSetMMDBHackl0Us}), hackl0usMMDBFunc, hackl0usMMDB)
	for _, l := range Languages {
		lang := l
		langName := fmt.Sprintf("%s (%s)", lang.Name, lang.Tag.String())
		mLang := mI18nSwitcher.AddSubMenuItemEx(langName, langName, func(menuItemEx *stx.MenuItemEx) {
			log.Infoln("[i18n] switch language to %s", langName)
			SwitchLanguage(lang)
			config.Set("lang", lang.Tag.String())
			menuItemEx.SwitchCheckboxBrother(true)
		})
		if Language != nil && Language.Tag == lang.Tag {
			mLang.SwitchCheckboxBrother(true)
		}
	}
	stx.AddSeparator()

	// 退出
	mQuit := stx.AddMainMenuItemExI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuQuit}), func(menuItemEx *stx.MenuItemEx) {
		stx.Quit()
		return
	})
	AddSwitchCallback(&CallbackData{Callback: func(params ...interface{}) {
		mOthers.SwitchLanguageWithChildren()
		mQuit.SwitchLanguage()
	}})

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
		SavedPort := clashP.GetPorts().Port
		if config.IsCmdPositive(cmd.Sys) {
			var Ports int
			if clashP.GetPorts().MixedPort != 0 {
				Ports = clashP.GetPorts().MixedPort
			} else {
				Ports = clashP.GetPorts().Port
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
		if config.IsCmdPositive(cmd.Cron) {
			mOthersUpdateCron.Check()
			go controller.CronTask()
		}
		//if clashConfig.GroupsList.Len() > 0 {
		//	log.Infoln("--")
		//	//log.Infoln(clashConfig.GroupsList)
		//	RefreshProxyGroups(mGroup, clashConfig.GroupsList, clashConfig.ProxiesList)
		//}

		for {
			<-t.C
			switch tunnel.Mode() {
			case tunnel.Global:
				if mGlobal.Checked() {
				} else {
					RefreshProxyGroups(mGroup, nil, clashConfig.ProxiesList)
					NeedLoadSelector = true
					stx.SwitchCheckboxGroup(mGlobal, true, proxyModeGroup)
					mGroup.Enable()
					if mEnabled.Checked() {
						stx.SetIcon(icon.DateG)
						notify.DoTrayMenu(cmdP.Global)
					} else {
						stx.SetIcon(icon.DateN)
					}
				}
			case tunnel.Rule:
				if mRule.Checked() {
				} else {
					RefreshProxyGroups(mGroup, clashConfig.GroupsList, clashConfig.ProxiesList)
					NeedLoadSelector = true
					stx.SwitchCheckboxGroup(mRule, true, proxyModeGroup)
					mGroup.Enable()
					if mEnabled.Checked() {
						stx.SetIcon(icon.DateS)
						notify.DoTrayMenu(cmdP.Rule)
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
						notify.DoTrayMenu(cmdP.Direct)
					} else {
						stx.SetIcon(icon.DateN)
					}
				}
			}
			if loadProfile {
				common.RefreshProfile()
			}
			loadProfile = false
			if firstInit {
				if config.IsCmdPositive(cmd.Task) {
					mOthersTask.Check()
				} else {
					mOthersTask.Uncheck()
				}
				if config.IsCmdPositive(cmd.MMDB) {
					stx.SwitchCheckboxGroup(hackl0usMMDB, true, mmdbGroup)
				} else {
					stx.SwitchCheckboxGroup(maxMindMMDB, true, mmdbGroup)
				}

				if config.IsCmdPositive(cmd.Sys) {
					mOthersAutosys.Check()
				} else {
					mOthersAutosys.Uncheck()
				}

				if config.IsCmdPositive(cmd.Cron) {
					mOthersUpdateCron.Check()
				} else {
					mOthersUpdateCron.Uncheck()
				}

				if mEnabled.Checked() {
					var p int
					if clashP.GetPorts().MixedPort != 0 {
						p = clashP.GetPorts().MixedPort
					} else {
						p = clashP.GetPorts().Port
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
