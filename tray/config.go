package tray

import (
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
		ResetProfiles(event)
		SwitchProfile()
	}
	mUpdateAll = mSwitchProfile.AddSubMenuItemEx("一键更新", "", func(menuItemEx *stx.MenuItemEx) {
		//for e := p.Profiles.Front(); e != nil; e = e.Next() {
		//	profile := e.Value.(*p.Info)
		//	p.UpdateConfig(profile.Name, profile.Url)
		//}
		go common.RefreshProfile(nil)
	}).AddSeparator()
}

func ResetProfiles(event *fsnotify.Event) {
	if mSwitchProfile == nil {
		return
	}

	p.RefreshProfiles(event)
	if event == nil {
		log.Infoln("[config] loaded %d profile(s)", p.Profiles.Len())
		for e := p.Profiles.Front(); e != nil; e = e.Next() {
			rawData := e.Value.(*p.RawData)
			if rawData.FileInfo != nil {
				addProfileMenuItem(rawData.FileInfo.Name)
			}
		}
	} else if event.Op|fsnotify.Write == fsnotify.Write {
		addProfileMenuItem(event.Name)
	} else if event.Op|fsnotify.Remove == fsnotify.Remove {
		p.RemoveProfile(event.Name)
	}
	////mSwitchProfile.ClearChildren()
	//mSwitchProfile.ForChildrenLoop(true, func(_ int, profile *stx.MenuItemEx) (remove bool) {
	//	if profile.GetId() == mUpdateAll.GetId() {
	//		return false
	//	}
	//	_, exists := p.MenuItemMap.Load(profile.GetTitle())
	//	return !exists
	//})

	if p.Profiles.Len() == 0 {
		mSwitchProfile.Disable()
	} else {
		mSwitchProfile.Enable()
	}
}

func InitProfiles() {
	if mSwitchProfile == nil {
		return
	}
	ResetProfiles(nil)
	if p.Profiles.Len() == 0 {
		mSwitchProfile.Disable()
	} else {
		mSwitchProfile.Enable()
	}
}

func addProfileMenuItem(profileName string) {
	profileName = p.GetConfigName(profileName)
	log.Infoln("[profile] added: %s", profileName)
	var rawData *p.RawData
	v, exists := p.RawDataMap.Load(profileName)
	if exists {
		rawData = v.(*p.RawData)
		if rawData.MenuItemEx != nil {
			//p.RemoveProfile(profileName)
			rawData.MenuItemEx.Delete()
		}
	} else {
		rawData = &p.RawData{}
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

	rawData.MenuItemEx = mP
	p.RawDataMap.Store(profileName, rawData)
	//ResetProfiles()
}

func SwitchProfile() {
	if mSwitchProfile == nil {
		return
	}

	defer p.Locker.Unlock()
	p.Locker.Lock()
	configName, _ := controller.CheckConfig()
	mSwitchProfile.ForChildrenLoop(true, func(_ int, profile *stx.MenuItemEx) (remove bool) {
		if profile.GetId() == mUpdateAll.GetId() {
			return
		}
		//log.Infoln("into:: %s", profile.GetTitle() + constant.ConfigSuffix)
		if configName == profile.GetTitle() + constant.ConfigSuffix {
			profile.Check()
		} else {
			profile.Uncheck()
		}
		return
	})
}
