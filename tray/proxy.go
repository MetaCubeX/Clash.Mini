package tray

import (
	"container/list"
	"fmt"
	"github.com/Dreamacro/clash/config"
	"github.com/Dreamacro/clash/tunnel"
	"github.com/JyCyunMe/go-i18n/i18n"
	cI18n "github.com/MetaCubeX/Clash.Mini/constant/i18n"
	"github.com/MetaCubeX/Clash.Mini/log"
	"github.com/MetaCubeX/Clash.Mini/proxy"
	"github.com/MetaCubeX/Clash.Mini/util"
	. "github.com/MetaCubeX/Clash.Mini/util/maybe"
	stringUtils "github.com/MetaCubeX/Clash.Mini/util/string"
	stx "github.com/getlantern/systray"
	"sync"
	"time"
)

const (
	proxyLogHeader = logHeader + ".proxy"

	max uint16 = 0xffff
)

var (
	ConfigGroupsMap map[uint32]map[uint32]string
	SelectorMap     map[string]proxy.SelectorInfo

	mProxyMap    map[string][]*stx.MenuItemEx
	PingTestInfo *PingTest
)

func init() {
	mProxyMap = make(map[string][]*stx.MenuItemEx)
	PingTestInfo = &PingTest{locker: new(sync.RWMutex), LowestDelay: -1, LastUpdateDT: time.Unix(0, 0)}
}

func SwitchGroupAndProxy(mGroup *stx.MenuItemEx, sGroup string, sProxy string) {
	log.Infoln("[%s] switch: %s :: %s", proxyLogHeader, sGroup, sProxy)
	mGroup.ForChildrenLoop(true, func(_ int, group *stx.MenuItemEx) (remove bool) {
		if Maybe().OfNullable(group.ExtraData).IfOkString(func(o interface{}) string {
			return o.(*proxy.Proxy).Name
		}) == sGroup {
			group.ForChildrenLoop(true, func(_ int, p *stx.MenuItemEx) (remove bool) {
				if Maybe().OfNullable(p.ExtraData).IfOkString(func(o interface{}) string {
					return o.(*proxy.Proxy).Name
				}) == sProxy {
					p.SwitchCheckboxBrother(true)
				}
				return
			})
		}
		return
	})
}

func RefreshProxyGroups(mGroup *stx.MenuItemEx, groupsList *list.List, proxiesList *list.List) {
	mGroup.ClearChildren()
	mProxyMap = make(map[string][]*stx.MenuItemEx)
	//// TODO: need unmarshal proxy info
	ConfigGroupsMap = make(map[uint32]map[uint32]string)
	if groupsList == nil {
		if proxiesList != nil {
			groupsList = list.New()
			groupsList.PushFront(proxy.GroupsList{
				Name:    "GLOBAL",
				Proxies: SelectorMap["GLOBAL"].All,
			})
			groupsList.PushBackList(config.GroupsList)
		} else {
			groupsList = list.New()
		}
	}
	for e := groupsList.Front(); e != nil; e = e.Next() {
		//println(util.ToJsonString(e.Value))
		s := proxy.GroupsList{}
		if err := util.ConvertForceByJson(&s, e.Value); err != nil {
			return
		}
		mConfigGroup := mGroup.AddSubMenuItemCheckboxEx(s.Name, s.Name, false, mConfigProxyFunc)
		configProxiesMap := make(map[uint32]string)
		proxyGroup := proxy.Proxy{
			Name: s.Name,
		}
		mConfigGroup.ExtraData = &proxy.Proxy{
			Name: s.Name,
		}
		for _, configProxy := range s.Proxies {
			p, exist := tunnel.Proxies()[configProxy]
			var lastDelay string
			var delay int16
			if exist {
				if p.LastDelay() != max {
					delay = int16(p.LastDelay())
					lastDelay = i18n.TData(cI18n.UtilDatetimeShortMilliSeconds,
						&i18n.Data{Data: map[string]interface{}{"ms": p.LastDelay()}})
				}
			} else {
				lastDelay = i18n.T(cI18n.ProxyTestTimeout)
				delay = -1
			}
			mConfigProxy := mConfigGroup.AddSubMenuItemCheckboxEx(stringUtils.GetMenuItemFullTitle(configProxy, lastDelay),
				configProxy, false, mConfigProxyFunc)
			mConfigProxy.ExtraData = &proxy.Proxy{
				Name:   configProxy,
				Parent: &proxyGroup,
				Delay:  delay,
			}
			configProxiesMap[mConfigProxy.GetId()] = configProxy
			mProxyMap[configProxy] = append(mProxyMap[configProxy], mConfigProxy)
		}
		ConfigGroupsMap[mConfigGroup.GetId()] = configProxiesMap
		mProxyMap[s.Name] = append(mProxyMap[s.Name], mConfigGroup)
	}
	if mGroup.Children.Len() == 0 {
		mGroup.Disable()
	} else {
		mGroup.Enable()
	}
}

// TODO:
func RefreshProxyDelay(mGroup *stx.MenuItemEx, delayMap map[string]int16) {
	mGroup.ForChildrenLoop(true, func(_ int, s *stx.MenuItemEx) (remove bool) {
		//println(util.ToJsonString(e.Value))
		if s.Children.Len() > 0 {
			RefreshProxyDelay(s, delayMap)
		} else {
			delay, exist := delayMap[s.GetTooltip()]
			var lastDelay string
			if exist {
				if delay == -1 || uint16(delay) == max {
					lastDelay = "Timeout"
				} else {
					lastDelay = fmt.Sprintf("%d ms", delay)
				}
			} else {
				lastDelay = "Timeout"
			}
			//		proxy, exist := tunnel.Proxies()[s.GetTooltip()]
			//		var lastDelay string
			//		if exist {
			//			if proxy.LastDelay() != max {
			//				lastDelay = fmt.Sprintf("%d ms", proxy.LastDelay())
			//			} else {
			//				lastDelay = "Timeout"
			//			}
			//		}
			s.SetTitle(stringUtils.GetMenuItemFullTitle(s.GetTooltip(), lastDelay))
		}
		return
	})
	//	//println(util.ToJsonString(e.Value))
	//	s := e.Value.(*stx.MenuItemEx)
	//	if s.Children.Len() > 0 {
	//		RefreshProxyDelay(s, delayMap)
	//	} else {
	//		delay, exist := delayMap[s.GetTooltip()]
	//		var lastDelay string
	//		if exist {
	//			if delay == -1 || uint16(delay) == max {
	//				lastDelay = "Timeout"
	//			} else {
	//				lastDelay = fmt.Sprintf("%d ms", delay)
	//			}
	//		} else {
	//			lastDelay = "Timeout"
	//		}
	////		proxy, exist := tunnel.Proxies()[s.GetTooltip()]
	////		var lastDelay string
	////		if exist {
	////			if proxy.LastDelay() != max {
	////				lastDelay = fmt.Sprintf("%d ms", proxy.LastDelay())
	////			} else {
	////				lastDelay = "Timeout"
	////			}
	////		}
	//		s.SetTitle(util.GetMenuItemFullTitle(s.GetTooltip(), lastDelay))
	//	}
	//}
}
