package tray

import (
	"container/list"
	"fmt"
	cI18n "github.com/Clash-Mini/Clash.Mini/constant/i18n"
	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Clash-Mini/Clash.Mini/util"
	"github.com/Dreamacro/clash/config"
	"github.com/Dreamacro/clash/tunnel"
	"github.com/JyCyunMe/go-i18n/i18n"
	stx "github.com/getlantern/systray"
)

const (
	max uint16 = 0xffff
)

var (
	ConfigGroupsMap map[uint32]map[uint32]string
	SelectorMap     map[string]SelectorInfo

	mProxyMap		map[string][]*stx.MenuItemEx
)

func init() {
	mProxyMap = make(map[string][]*stx.MenuItemEx)
}

// TEST
// TODO: not fit standard
type GroupsList struct {
	Name    string   `json:"name"`
	Proxies []string `json:"proxies"`
	Type    string   `json:"type"`
}

type SelectorInfo struct {
	All     []string      `json:"all,omitempty"`
	History []interface{} `json:"history,omitempty"`
	Name    string        `json:"name"`
	Now     string        `json:"now"`
	Type    string        `json:"type"`
}

func SwitchGroupAndProxy(mGroup *stx.MenuItemEx, sGroup string, sProxy string) {
	log.Infoln("switch: %s :: %s", sGroup, sProxy)
	for e := mGroup.Children.Front(); e != nil; e = e.Next() {
		group := e.Value.(*stx.MenuItemEx)
		if group.GetTitle() == sGroup {
			for e := group.Children.Front(); e != nil; e = e.Next() {
				proxy := e.Value.(*stx.MenuItemEx)
				if proxy.GetTitle() == sProxy {
					proxy.SwitchCheckboxBrother(true)
				}
			}
		}
	}
}

func RefreshProxyGroups(mGroup *stx.MenuItemEx, groupsList *list.List, proxiesList *list.List) {
	mGroup.ClearChildren()
	mProxyMap = make(map[string][]*stx.MenuItemEx)
	// TODO: need unmarshal proxy info
	ConfigGroupsMap = make(map[uint32]map[uint32]string)
	if groupsList == nil {
		if proxiesList != nil {
			groupsList = list.New()
			groupsList.PushFront(GroupsList{
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
		s := GroupsList{}
		if err := util.ConvertForceByJson(&s, e.Value); err != nil {
			return
		}
		mConfigGroup := mGroup.AddSubMenuItemCheckboxEx(s.Name, s.Name, false, mConfigProxyFunc)
		configProxiesMap := make(map[uint32]string)
		for _, configProxy := range s.Proxies {
			proxy, exist := tunnel.Proxies()[configProxy]
			var lastDelay string
			if exist {
				if proxy.LastDelay() != max {
					lastDelay = i18n.TData("", cI18n.UtilDatetimeShortMilliSeconds,
						&i18n.Data{Data: map[string]interface{}{ "ms": proxy.LastDelay() }})
				}
			} else {
				lastDelay = i18n.T(cI18n.ProxyTestTimeout)
			}
			mConfigProxy := mConfigGroup.AddSubMenuItemCheckboxEx(fmt.Sprintf("%s\t%s", configProxy, lastDelay),
				configProxy, false, mConfigProxyFunc)
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

func RefreshProxyDelay(mGroup *stx.MenuItemEx, delayMap map[string]int16) {
	for e := mGroup.Children.Front(); e != nil; e = e.Next() {
		//println(util.ToJsonString(e.Value))
		s := e.Value.(*stx.MenuItemEx)
		if s.Children.Len() > 0 {
			RefreshProxyDelay(s, delayMap)
		} else {
			delay, exist := delayMap[s.GetTooltip()]
			var lastDelay string
			if exist {
				if delay == -1 || uint16(delay) == max {
					lastDelay = "Timeout"
				} else {
					lastDelay = fmt.Sprintf("\t%d ms", delay)
				}
			} else {
				lastDelay = "Timeout"
			}
	//		proxy, exist := tunnel.Proxies()[s.GetTooltip()]
	//		var lastDelay string
	//		if exist {
	//			if proxy.LastDelay() != max {
	//				lastDelay = fmt.Sprintf("\t%d ms", proxy.LastDelay())
	//			} else {
	//				lastDelay = "\tTimeout"
	//			}
	//		}
			s.SetTitle(fmt.Sprintf("%s\t%s", s.GetTooltip(), lastDelay))
		}
	}
}
