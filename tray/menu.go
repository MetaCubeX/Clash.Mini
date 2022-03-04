package tray

import (
	"container/list"
	"fmt"
	"github.com/MetaCubeX/Clash.Mini/cmd/autosys"
	"github.com/MetaCubeX/Clash.Mini/mixin"
	"github.com/MetaCubeX/Clash.Mini/util/powershell"
	"time"

	"github.com/MetaCubeX/Clash.Mini/app"
	"github.com/MetaCubeX/Clash.Mini/cmd"
	cmdP "github.com/MetaCubeX/Clash.Mini/cmd/proxy"
	"github.com/MetaCubeX/Clash.Mini/common"
	"github.com/MetaCubeX/Clash.Mini/config"
	"github.com/MetaCubeX/Clash.Mini/constant"
	cI18n "github.com/MetaCubeX/Clash.Mini/constant/i18n"
	"github.com/MetaCubeX/Clash.Mini/controller"
	"github.com/MetaCubeX/Clash.Mini/icon"
	"github.com/MetaCubeX/Clash.Mini/log"
	"github.com/MetaCubeX/Clash.Mini/notify"
	p "github.com/MetaCubeX/Clash.Mini/profile"
	"github.com/MetaCubeX/Clash.Mini/proxy"
	"github.com/MetaCubeX/Clash.Mini/sysproxy"
	"github.com/MetaCubeX/Clash.Mini/util"
	commonUtils "github.com/MetaCubeX/Clash.Mini/util/common"
	. "github.com/MetaCubeX/Clash.Mini/util/maybe"
	stringUtils "github.com/MetaCubeX/Clash.Mini/util/string"

	clashConfig "github.com/Dreamacro/clash/config"
	"github.com/Dreamacro/clash/hub/route"
	clashP "github.com/Dreamacro/clash/listener"
	"github.com/Dreamacro/clash/tunnel"
	. "github.com/JyCyunMe/go-i18n/i18n"
	stx "github.com/getlantern/systray"
	"github.com/skratchdot/open-golang/open"
)

const (
	menuLogHeader = logHeader + ".menu"
)

var (
	firstInit   = true
	loadProfile = true

	coreTrayMenuEnabled      = true
	dashboardTrayMenuEnabled = true

	mCoreProxyMode = &stx.MenuItemEx{}
	mGlobal        = &stx.MenuItemEx{}
	mRule          = &stx.MenuItemEx{}
	mDirect        = &stx.MenuItemEx{}

	mGroup     = &stx.MenuItemEx{}
	mPingTest  = &stx.MenuItemEx{}
	mConfig    = &stx.MenuItemEx{}
	mEnabled   = &stx.MenuItemEx{}
	mDashboard = &stx.MenuItemEx{}

	mOthers       = &stx.MenuItemEx{}
	mI18nSwitcher = &stx.MenuItemEx{}

	mOthersMixinDir     = &stx.MenuItemEx{}
	mOthersMixinGeneral = &stx.MenuItemEx{}
	mOthersMixinTun     = &stx.MenuItemEx{}
	mOthersMixinDns     = &stx.MenuItemEx{}

	mOthersProtocol    = &stx.MenuItemEx{}
	mOthersUwpLoopback = &stx.MenuItemEx{}
	mOthersTask        = &stx.MenuItemEx{}
	mOthersAutosys     = &stx.MenuItemEx{}
	mOthersHotkey      = &stx.MenuItemEx{}
	mOthersUpdateCron  = &stx.MenuItemEx{}
	maxMindMMDB        = &stx.MenuItemEx{}
	hackl0usMMDB       = &stx.MenuItemEx{}
)

// addMenuProxyModes
func addMenuProxyModes() {
	// 核心开关
	//mCoreSwitcher := stx.AddMainMenuItemExI18n(stx.NewI18nConfig(stx.I18nConfig{
	//	TitleID:     cI18n.TrayMenuGlobalProxy,
	//	TitleFormat: "\tAlt+G",
	//	TooltipID:   cI18n.TrayMenuGlobalProxy,
	//}), func(menuItemEx *stx.MenuItemEx) {
	//	//tunnel.SetMode(tunnel.Global)
	//	//firstInit = true
	//})

	// 代理模式
	stx.AddMainMenuItemExBindI18n(stx.NewI18nConfig(stx.I18nConfig{
		TitleID:   cI18n.TrayMenuCoreStopped,
		TooltipID: cI18n.TrayMenuCoreStopped,
	}), stx.NilCallback, mCoreProxyMode).
		// 全局代理
		AddSubMenuItemExBindI18n(stx.NewI18nConfig(stx.I18nConfig{
			TitleID:   cI18n.TrayMenuGlobalProxy,
			TooltipID: cI18n.TrayMenuGlobalProxy,
		}), func(menuItemEx *stx.MenuItemEx) {
			tunnel.SetMode(tunnel.Global)
			firstInit = true
		}, mGlobal).
		// 规则代理
		AddMenuItemExBindI18n(stx.NewI18nConfig(stx.I18nConfig{
			TitleID:   cI18n.TrayMenuRuleProxy,
			TooltipID: cI18n.TrayMenuRuleProxy,
		}), func(menuItemEx *stx.MenuItemEx) {
			tunnel.SetMode(tunnel.Rule)
			firstInit = true
		}, mRule).
		// 全局直连
		AddMenuItemExBindI18n(stx.NewI18nConfig(stx.I18nConfig{
			TitleID:   cI18n.TrayMenuDirectProxy,
			TooltipID: cI18n.TrayMenuDirectProxy,
		}), func(menuItemEx *stx.MenuItemEx) {
			tunnel.SetMode(tunnel.Direct)
			firstInit = true
		}, mDirect)
	stx.AddSeparator()
	mCoreProxyMode.Disabled()
	if common.DisabledCore {
		mCoreProxyMode.I18nConfig = stx.NewI18nConfig(stx.I18nConfig{
			TitleID:   cI18n.TrayMenuCoreDisabled,
			TooltipID: cI18n.TrayMenuCoreDisabled,
		})
		mCoreProxyMode.SwitchLanguage()
	}
	AddSwitchCallback(&CallbackData{Callback: func(params ...interface{}) {
		mCoreProxyMode.SwitchLanguageWithChildren()
	}})
}

func addMenuEndpoints() {
	// 切换节点
	mGroup = stx.AddMainMenuItemExI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuSwitchProxy}), stx.NilCallback)
	if ConfigGroupsMap == nil {
		clashConfig.ParsingProxiesCallback = func(groupsList *list.List, proxiesList *list.List) {
			RefreshProxyGroups(mGroup, groupsList, proxiesList)
			NeedLoadSelector = true
		}
		route.SwitchProxiesCallback = func(sGroup string, sProxy string) {
			SwitchGroupAndProxy(mGroup, sGroup, sProxy)
		}
	}
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
							pm.SetTitle(stringUtils.GetMenuItemFullTitle(pm.GetTooltip(), lastDelay))
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
			lastUpdateDT = util.GetHumanTimeI18n(pt.LastUpdateDT)
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
		mGroup.SwitchLanguage()
		mPingTest.SwitchLanguageWithChildren()
	}})
	PingTestInfo.Callback(nil)
}

func initTrayMenu() {
	stx.AddMainMenuItemEx(mainTitle, mainTooltip, func(menuItemEx *stx.MenuItemEx) {
		log.Infoln("[%s] Hi Clash.Mini, %s", menuLogHeader, app.Version)
		_ = open.Run("https://github.com/Clash-Mini/Clash.Mini")
	})
	stx.AddSeparator()

	// TEST: 配置关联订阅
	addMenuProxyModes()
	// TEST: showCustomizedTrayMenu
	//stx.AddMainMenuItemExI18n(stx.NewI18nConfig(stx.I18nConfig{ TitleID: cI18n.TrayMenuSwitchProxy }), func(menuItemEx *stx.MenuItemEx) {
	//	controller.TrayMenuInit()
	//})
	addMenuEndpoints()

	// 切换订阅
	mSwitchProfile := stx.AddMainMenuItemExI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuSwitchProfile}), stx.NilCallback)
	stx.AddSeparator()
	SetMSwitchProfile(mSwitchProfile)

	// 系统代理
	mEnabled = stx.AddMainMenuItemExI18n(stx.NewI18nConfig(stx.I18nConfig{
		TitleID:   cI18n.TrayMenuSystemProxy,
		TooltipID: cI18n.TrayMenuSystemProxy,
	}), mEnabledFunc)
	// 控制面板
	mDashboard = stx.AddMainMenuItemExI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuDashboard}), func(menuItemEx *stx.MenuItemEx) {
		go controller.Dashboard()
	})
	if common.DisabledDashboard {
		mDashboard.Disable()
	}
	// 配置管理
	mConfig = stx.AddMainMenuItemExI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuConfigManagement}), func(menuItemEx *stx.MenuItemEx) {
		go controller.ShowMenuConfig()
	})
	// 查看日志
	mLogger := stx.AddMainMenuItemExI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuShowLog}), func(menuItemEx *stx.MenuItemEx) {

		err := powershell.ShowCmd()
		if err != nil {
			log.Errorln("[%s] ShowLog exec failed: %s", logHeader, err)
		}
		//powershell.ShowCmd()
		// TODO: new ui
		//go controller.ShowMenuConfig()
	})
	//mLogger.Disable()
	AddSwitchCallback(&CallbackData{Callback: func(params ...interface{}) {
		mSwitchProfile.SwitchLanguage()
		mSwitchProfile.SwitchLanguageWithChildren()
		mEnabled.SwitchLanguage()
		mDashboard.SwitchLanguage()
		mConfig.SwitchLanguage()
		mLogger.SwitchLanguage()
	}})

	// 其他设置
	stx.AddMainMenuItemExBindI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuOtherSettings}), stx.NilCallback, mOthers).
		// 切换语言
		AddSubMenuItemExBindI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuOtherSettingsSwitchLanguage}), stx.NilCallback, mI18nSwitcher).
		// Mixin
		AddMenuItemExI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuOtherSettingsMixin}), stx.NilCallback).
		AddSubMenuItemExBindI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuOtherSettingsMixinDir}), mOthersMixinDirFunc, mOthersMixinDir).
		AddSeparator().
		AddMenuItemExBindI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuOtherSettingsMixinGeneral}), mOthersMixinGeneralFunc, mOthersMixinGeneral).
		AddMenuItemExBindI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuOtherSettingsMixinTun}), mOthersMixinTunFunc, mOthersMixinTun).
		AddMenuItemExBindI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuOtherSettingsMixinDns}), mOthersMixinDnsFunc, mOthersMixinDns).Parent.
		// 设置开机启动
		AddMenuItemExBindI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuOtherSettingsSystemAutorun}), mOtherTaskFunc, mOthersTask).
		// 默认系统代理
		AddMenuItemExBindI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuOtherSettingsSystemAutoProxy}), mOtherAutosysFunc, mOthersAutosys).
		// 设置定时更新
		AddMenuItemExBindI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuOtherSettingsCronUpdateConfigs}), mOtherUpdateCronFunc, mOthersUpdateCron).
		// 设置GeoIP2数据库
		AddMenuItemExI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuOtherSettingsSetMMDB}), stx.NilCallback).
		// MaxMind数据库
		AddSubMenuItemExBindI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuOtherSettingsSetMMDBMaxmind}), maxMindMMBDFunc, maxMindMMDB).
		// Hackl0us数据库
		AddMenuItemExBindI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuOtherSettingsSetMMDBHackl0Us}), hackl0usMMDBFunc, hackl0usMMDB).Parent.
		AddSeparator().
		// 注册快捷键
		AddMenuItemExBindI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuOtherSettingsHotkey}), mOtherHotkeyFunc, mOthersHotkey).
		// 关联URL协议
		AddMenuItemExBindI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuOtherSettingsRegisterProtocol}), mOtherProtocolFunc, mOthersProtocol).
		// 全局UWP回环
		AddMenuItemExBindI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuOtherSettingsUwpLoopback}), mOtherUwpLoopbackFunc, mOthersUwpLoopback)

	for _, l := range Languages {
		lang := l
		langName := fmt.Sprintf("%s (%s)", lang.Name, lang.Tag.String())
		mLang := mI18nSwitcher.AddSubMenuItemEx(langName, langName, func(menuItemEx *stx.MenuItemEx) {
			log.Infoln("[i18n] switch language to %s", langName)
			err := SwitchLanguage(lang)
			if err != nil {
				log.Errorln("[i18n] switch language to %s failed: %v", langName, err)
				return
			}
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
		_ = controller.CloseDashboard()
		// 等待清理托盘图标
		//time.AfterFunc(500 * time.Millisecond, func() {
		//	os.Exit(0)
		//})
	})
	AddSwitchCallback(&CallbackData{Callback: func(params ...interface{}) {
		mOthers.SwitchLanguageWithChildren()
		mQuit.SwitchLanguage()
	}})

	if !commonUtils.IsWindows() {
		mEnabled.Hide()
		mOthers.Hide()
		mConfig.Hide()
	}

	proxyModeGroup := []*stx.MenuItemEx{mGlobal, mRule, mDirect}
	mmdbGroup := []*stx.MenuItemEx{maxMindMMDB, hackl0usMMDB}

	go func() {
		t := time.NewTicker(time.Second)
		defer t.Stop()
		var SavedPort int
		if config.IsCmdPositive(cmd.Autosys) || config.IsCmdPositive(cmd.Breaker) {
			var Ports int
			if clashP.GetPorts().MixedPort != 0 {
				Ports = clashP.GetPorts().MixedPort
			} else {
				Ports = clashP.GetPorts().Port
			}
			SavedPort = Ports
			err := sysproxy.SetSystemProxy(
				&sysproxy.ProxyConfig{
					Enable: true,
					Server: fmt.Sprintf("%s:%d", constant.LocalHost, Ports),
				})
			if err != nil {
				log.Errorln("[%s] SetSystemProxy error: %v", menuLogHeader, err)
				notify.PushError("", "设置系统代理时出错")
				return
			}
			mEnabled.Check()
			notify.DoTrayMenu(autosys.ON)
		}
		if config.IsCmdPositive(cmd.Cron) {
			mOthersUpdateCron.Check()
			go controller.CronTask()
		}

		if config.IsCmdPositive(cmd.Hotkey) {
			hotKey(true)
		}

		//if clashConfig.GroupsList.Len() > 0 {
		//	log.Infoln("--")
		//	//log.Infoln(clashConfig.GroupsList)
		//	RefreshProxyGroups(mGroup, clashConfig.GroupsList, clashConfig.ProxiesList)
		//}
		//p.RefreshProfiles(nil)

		for {
			<-t.C
			SwitchDashboardTrayMenu(!common.DisabledDashboard)
			if !common.CoreRunningStatus {
				SwitchCoreTrayMenu(false)
			} else {
				SwitchCoreTrayMenu(true)
				switch tunnel.Mode() {
				case tunnel.Global:
					if mGlobal.Checked() {
					} else {
						RefreshProxyGroups(mGroup, nil, clashConfig.ProxiesList)
						NeedLoadSelector = true
						ChangeCoreProxyMode(mGlobal)
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
						ChangeCoreProxyMode(mRule)
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
						ChangeCoreProxyMode(mDirect)
						stx.SwitchCheckboxGroup(mDirect, true, proxyModeGroup)
						if mEnabled.Checked() {
							stx.SetIcon(icon.DateD)
							notify.DoTrayMenu(cmdP.Direct)
						} else {
							stx.SetIcon(icon.DateN)
						}
					}
				}
			}
			if loadProfile {
				//InitProfiles()
				common.RefreshProfile(nil)
			}
			loadProfile = false

			if firstInit {
				if config.IsCmdPositive(cmd.Hotkey) {
					mOthersHotkey.Check()
				} else {
					mOthersHotkey.Uncheck()
				}
				if config.IsCmdPositive(cmd.Startup) {
					mOthersTask.Check()
				} else {
					mOthersTask.Uncheck()
				}
				if config.IsCmdPositive(cmd.MMDB) {
					stx.SwitchCheckboxGroup(hackl0usMMDB, true, mmdbGroup)
				} else {
					stx.SwitchCheckboxGroup(maxMindMMDB, true, mmdbGroup)
				}
				if config.IsCmdPositive(cmd.Protocol) {
					mOthersProtocol.Check()
				} else {
					mOthersProtocol.Uncheck()
				}
				if config.IsCmdPositive(cmd.Autosys) {
					mOthersAutosys.Check()
				} else {
					mOthersAutosys.Uncheck()
				}
				if config.IsCmdPositive(cmd.Breaker) {
					mOthersUwpLoopback.Check()
					mOthersAutosys.Disable()
				} else {
					mOthersUwpLoopback.Uncheck()
					mOthersAutosys.Enable()
				}
				if config.IsCmdPositive(cmd.Cron) {
					mOthersUpdateCron.Check()
				} else {
					mOthersUpdateCron.Uncheck()
				}
				if config.IsMixinPositive(mixin.General) {
					mOthersMixinGeneral.Check()
				} else {
					mOthersMixinGeneral.Uncheck()
				}
				if config.IsMixinPositive(mixin.Tun) {
					mOthersMixinTun.Check()
				} else {
					mOthersMixinTun.Uncheck()
				}
				if config.IsMixinPositive(mixin.Dns) {
					mOthersMixinDns.Check()
				} else {
					mOthersMixinDns.Uncheck()
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
								Server: fmt.Sprintf("%s:%d", constant.LocalHost, SavedPort),
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

				if p.Enable && p.Server == fmt.Sprintf("%s:%d", constant.LocalHost, SavedPort) {
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
		time.Sleep(4 * time.Second)
		userInfo := p.UpdateSubscriptionUserInfo()
		time.Sleep(2 * time.Second)
		if len(userInfo.UnusedInfo) > 0 {
			notify.PushFlowInfo(userInfo.UsedInfo, userInfo.UnusedInfo, userInfo.ExpireInfo)
		}
	}()

}

// onReady 托盘退出时
func onExit() {
	log.Infoln("[tray] exiting")
	//loopback.StopBreaker()
	if mEnabled.Checked() {
		err := sysproxy.ClearSystemProxy()
		if err != nil {
			log.Errorln("[tray] onExit cancel sysproxy error: %v", err)
			return
		}
	}
	log.Infoln("[tray] exited")
	// logger bug?
	log.Debugln("")
}

func ChangeCoreProxyMode(mie *stx.MenuItemEx) {
	mCoreProxyMode.I18nConfig = mie.I18nConfig
	//mCoreProxyMode.I18nConfig.TitleConfig.Format = "\tAlt+P"
	mCoreProxyMode.SwitchLanguage()
}

func SwitchCoreTrayMenu(enabled bool) {
	if enabled == coreTrayMenuEnabled {
		return
	}
	coreTrayMenuEnabled = enabled
	if enabled {
		mCoreProxyMode.Enable()
		mGroup.Enable()
		//mSwitchProfile.Enable()
		mPingTest.Enable()
		mConfig.Enable()
		mEnabled.Enable()
	} else {
		mCoreProxyMode.Disable()
		mGroup.Disable()
		//mSwitchProfile.Disable()
		mPingTest.Disable()
		//mConfig.Disable()
		mEnabled.Disable()
	}
}

func SwitchDashboardTrayMenu(enabled bool) {
	if enabled == dashboardTrayMenuEnabled {
		return
	}
	dashboardTrayMenuEnabled = enabled
	if enabled {
		mDashboard.Enable()
	} else {
		mDashboard.Disable()
	}
}
