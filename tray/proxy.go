package tray

import (
	"container/list"
	"encoding/json"
	"fmt"
	"github.com/Clash-Mini/Clash.Mini/util"
	"github.com/Dreamacro/clash/component/profile/cachefile"
	stx "github.com/getlantern/systray"
)

var (
	ConfigGroupsMap map[uint32]map[uint32]string
)

// TEST
// TODO: not fit standard
type GroupsList struct {
	Name    string   `json:"name"`
	Proxies []string `json:"proxies"`
}

func RefreshProxyGroups(mGroup *stx.MenuItemEx, groupsList *list.List, proxiesList *list.List) {
	mGroup.ClearChildren()
	// TODO: need unmarshal proxy info
	ConfigGroupsMap = make(map[uint32]map[uint32]string)
	lGroup := groupsList
	for e := lGroup.Front(); e != nil; e = e.Next() {
		group := e.Value
		fmt.Println(group)
		jsonString, _ := json.Marshal(group)
		s := GroupsList{}
		if err := json.Unmarshal(jsonString, &s); err != nil {
			return
		}
		//mConfigGroup := mGroup.AddSubMenuItemEx(s.Name, s.Name, mConfigProxyFunc)
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
	println(util.ToJsonString(cachefile.Cache().SelectedMap()))

	isRestoredSelector := false
	for cGroup, cProxy := range cachefile.Cache().SelectedMap() {
		for e := mGroup.Children.Front(); !isRestoredSelector && e != nil; e = e.Next() {
			group := e.Value.(*stx.MenuItemEx)
			if group.GetTitle() == cGroup {
				for e := group.Children.Front(); !isRestoredSelector && e != nil; e = e.Next() {
					proxy := e.Value.(*stx.MenuItemEx)
					if proxy.GetTitle() == cProxy {
						stx.SwitchCheckboxBrother(group, true)
						stx.SwitchCheckboxBrother(proxy, true)
						isRestoredSelector = true
					}
				}

			}
		}
	}
}
