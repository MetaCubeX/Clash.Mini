package tray

import (
	"os"
	"strings"
	"time"

	"github.com/Clash-Mini/Clash.Mini/common"
	"github.com/Clash-Mini/Clash.Mini/constant"
	cI18n "github.com/Clash-Mini/Clash.Mini/constant/i18n"
	"github.com/Clash-Mini/Clash.Mini/controller"
	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Clash-Mini/Clash.Mini/notify"
	p "github.com/Clash-Mini/Clash.Mini/profile"

	"github.com/JyCyunMe/go-i18n/i18n"
	"github.com/fsnotify/fsnotify"
	stx "github.com/getlantern/systray"
)

var (
	mSwitchProfile 	*stx.MenuItemEx
	mUpdateAll		*stx.MenuItemEx
)

func SetMSwitchProfile(mie *stx.MenuItemEx) {
	mSwitchProfile = mie
	common.RefreshProfile = func(event *fsnotify.Event) {
		if event != nil {
			startIdx := strings.LastIndexByte(event.Name, os.PathSeparator)
			endIdx := strings.LastIndex(event.Name, constant.ConfigSuffix)
			if startIdx > -1 && endIdx >= startIdx {
				event.Name = event.Name[startIdx + 1:endIdx]
			}
		}
		ResetProfiles(event)
		SwitchProfile()
	}
	mUpdateAll = mSwitchProfile.AddSubMenuItemEx("一键更新", "", func(menuItemEx *stx.MenuItemEx) {
		for e := p.Profiles.Front(); e != nil; e = e.Next() {
			profile := e.Value.(*p.Info)
			p.UpdateConfig(profile.Name, profile.Url)
		}
		go common.RefreshProfile(nil)
	}).AddSeparator()
}

func ResetProfiles(event *fsnotify.Event) {
	if mSwitchProfile == nil {
		return
	}

	if event == nil {
		InitProfiles()
		return
	}

	if event.Op|fsnotify.Write == fsnotify.Write {
		addProfileMenuItem(event.Name)
	}
	////mSwitchProfile.ClearChildren()
	//mSwitchProfile.ForChildrenLoop(true, func(_ int, profile *stx.MenuItemEx) (remove bool) {
	//	if profile.GetId() == mUpdateAll.GetId() {
	//		return false
	//	}
	//	_, exists := p.MenuItemMap.Load(profile.GetTitle())
	//	return !exists
	//})
}

func InitProfiles() {
	if mSwitchProfile == nil {
		return
	}

	if p.Profiles.Len() == 0 {
		mSwitchProfile.Disable()
		return
	}
	mSwitchProfile.Enable()
	for e := p.Profiles.Front(); e != nil; e = e.Next() {
		profile := e.Value.(*p.Info)
		addProfileMenuItem(profile.Name)
	}
}

func addProfileMenuItem(profileName string) {
	_, exists := p.MenuItemMap.LoadAndDelete(profileName)
	if exists {
		return
		//v.(*stx.MenuItemEx).Delete()
	}
	mP := mSwitchProfile.AddSubMenuItemEx(profileName, profileName, func(menuItemEx *stx.MenuItemEx) {
		log.Infoln("switch profile to \\%s\\", menuItemEx.GetTitle())
		// TODO: switch
		controller.PutConfig(menuItemEx.GetTitle())
		//walk.MsgBox(nil, i18n.T(cI18n.MsgBoxTitleTips),
		//	i18n.TData(cI18n.MenuConfigMessageEnableConfigSuccess, &i18n.Data{Data: map[string]interface{}{
		//		"Config": menuItemEx.GetTitle(),
		//	}}),
		//	walk.MsgBoxIconInformation)
		message := i18n.TData(cI18n.MenuConfigMessageEnableConfigSuccess, &i18n.Data{Data: map[string]interface{}{
			"Config": menuItemEx.GetTitle(),
		}})
		notify.PushWithLine(cI18n.NotifyMessageTitle, message)
		menuItemEx.SwitchCheckboxBrother(true)
		go func() {
			time.Sleep(constant.NotifyDelay)
			userInfo := p.UpdateSubscriptionUserInfo()
			if len(userInfo.UnusedInfo) > 0 {
				notify.PushFlowInfo(userInfo.UsedInfo, userInfo.UnusedInfo, userInfo.ExpireInfo)
			}
		}()
	})
	p.MenuItemMap.Store(profileName, mP)
	//ResetProfiles()
}

func SwitchProfile() {
	if mSwitchProfile == nil {
		return
	}

	configName, _ := controller.CheckConfig()
	mSwitchProfile.ForChildrenLoop(true, func(_ int, profile *stx.MenuItemEx) (remove bool) {
		if profile.GetId() == mUpdateAll.GetId() {
			return
		}
		if configName == profile.GetTitle() + constant.ConfigSuffix {
			profile.Check()
		} else {
			profile.Uncheck()
		}
		return
	})
}
