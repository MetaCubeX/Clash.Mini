package tray

import (
	"container/list"
	"fmt"
	"time"

	"github.com/Clash-Mini/Clash.Mini/cmd"
	cp "github.com/Clash-Mini/Clash.Mini/cmd/proxy"
	"github.com/Clash-Mini/Clash.Mini/cmd/sys"
	"github.com/Clash-Mini/Clash.Mini/constant"
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
	stx "github.com/getlantern/systray"
)

func init() {
	if constant.IsWindows() {
		C.SetHomeDir(constant.PWD)
	}

	stx.RunEx(onReady, onExit)
}

func NilCallback(menuItem *stx.MenuItemEx) {
	log.Infoln("clicked %s, id: %d", menuItem.GetTitle(), menuItem.GetId())
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

	mGlobal := stx.AddMainMenuItemEx("全局代理", "全局代理", func(menuItemEx *stx.MenuItemEx) {
		tunnel.SetMode(tunnel.Global)
	})
	mRule := stx.AddMainMenuItemEx("规则代理", "规则代理", func(menuItemEx *stx.MenuItemEx) {
		tunnel.SetMode(tunnel.Rule)
	})
	mDirect := stx.AddMainMenuItemEx("全局直连", "全局直连", func(menuItemEx *stx.MenuItemEx) {
		tunnel.SetMode(tunnel.Direct)
	})
	stx.AddSeparator()

	mGroup := stx.AddMainMenuItemEx("切换节点", "切换节点", stx.NilCallback)
	if ConfigGroupsMap == nil {
		config.ParsingProxiesCallback = func(groupsList *list.List, proxiesList *list.List) {
			RefreshProxyGroups(mGroup, groupsList, proxiesList)
			NeedLoadSelector = true
		}
		route.SwitchProxiesCallback = func(sGroup string, sProxy string) {
			SwitchGroupAndProxy(mGroup, sGroup, sProxy)
		}
	}
	stx.AddSeparator()

	mEnabled := stx.AddMainMenuItemEx("系统代理", "系统代理", mEnabledFunc)
	stx.AddMainMenuItemEx("控制面板", "控制面板", func(menuItemEx *stx.MenuItemEx) {
		go controller.Dashboard()
	})
	mConfig := stx.AddMainMenuItemEx("配置管理", "配置管理", func(menuItemEx *stx.MenuItemEx) {
		go controller.ShowMenuConfig()
	})

	var mOtherTask = &stx.MenuItemEx{}
	var mOtherAutosys = &stx.MenuItemEx{}
	var mOtherUpdateCron = &stx.MenuItemEx{}
	var maxMindMMDB = &stx.MenuItemEx{}
	var hackl0usMMDB = &stx.MenuItemEx{}
	mOther := stx.AddMainMenuItemEx("其他设置", "其他设置", stx.NilCallback).
		AddSubMenuItemExBind("设置开机启动", "设置开机启动", mOtherTaskFunc, mOtherTask).
		AddMenuItemExBind("设置默认代理", "设置默认代理", mOtherAutosysFunc, mOtherAutosys).
		AddMenuItemExBind("设置定时更新", "设置定时更新", mOtherUpdateCronFunc, mOtherUpdateCron).
		AddMenuItemEx("设置GeoIP2数据库", "设置GeoIP2数据库", stx.NilCallback).
		AddSubMenuItemExBind("MaxMind数据库", "MaxMind数据库", maxMindMMBDFunc, maxMindMMDB).
		AddMenuItemExBind("Hackl0us数据库", "Hackl0us数据库", hackl0usMMDBFunc, hackl0usMMDB)
	stx.AddSeparator()

	stx.AddMainMenuItemEx("退出", "Quit Clash.Mini", func(menuItemEx *stx.MenuItemEx) {
		stx.Quit()
		return
	})

	if !constant.IsWindows() {
		mEnabled.Hide()
		mOther.Hide()
		mConfig.Hide()
	}

	proxyModeGroup := []*stx.MenuItemEx{mGlobal, mRule, mDirect}
	mmdbGroup := []*stx.MenuItemEx{maxMindMMDB, hackl0usMMDB}

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
			mOtherUpdateCron.Check()
			go controller.CronTask()
		}
		//if config.GroupsList.Len() > 0 {
		//	log.Infoln("--")
		//	//log.Infoln(config.GroupsList)
		//	RefreshProxyGroups(mGroup, config.GroupsList, config.ProxiesList)
		//}

		firstInit := true
		NeedLoadSelector = true
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
					mOtherTask.Check()
				} else {
					mOtherTask.Uncheck()
				}

				if controller.RegCompare(cmd.MMDB) {
					stx.SwitchCheckboxGroup(hackl0usMMDB, true, mmdbGroup)
				} else {
					stx.SwitchCheckboxGroup(maxMindMMDB, true, mmdbGroup)
				}

				if controller.RegCompare(cmd.Sys) {
					mOtherAutosys.Check()
				} else {
					mOtherAutosys.Uncheck()
				}

				if controller.RegCompare(cmd.Cron) {
					mOtherUpdateCron.Check()
				} else {
					mOtherUpdateCron.Uncheck()
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
