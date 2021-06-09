package tray

import (
	"container/list"
	"github.com/Clash-Mini/Clash.Mini/util"

	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Dreamacro/clash/config"
	stx "github.com/getlantern/systray"
)

var (
	ConfigGroupsMap map[uint32]map[uint32]string
	SelectorMap     map[string]SelectorInfo
)

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
					stx.SwitchCheckboxBrother(proxy, true)
				}
			}
		}
	}
}

func RefreshProxyGroups(mGroup *stx.MenuItemEx, groupsList *list.List, proxiesList *list.List) {
	mGroup.ClearChildren()
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
			mConfigProxy := mConfigGroup.AddSubMenuItemCheckboxEx(configProxy, configProxy, false, mConfigProxyFunc)
			configProxiesMap[mConfigProxy.GetId()] = configProxy
		}
		ConfigGroupsMap[mConfigGroup.GetId()] = configProxiesMap
	}
	if mGroup.Children.Len() == 0 {
		mGroup.Disable()
	} else {
		mGroup.Enable()
	}
}
