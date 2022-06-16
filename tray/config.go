package tray

import (
	"fmt"
	"github.com/MetaCubeX/Clash.Mini/config"
	"time"

	"github.com/MetaCubeX/Clash.Mini/common"
	"github.com/MetaCubeX/Clash.Mini/constant"
	cI18n "github.com/MetaCubeX/Clash.Mini/constant/i18n"
	"github.com/MetaCubeX/Clash.Mini/controller"
	"github.com/MetaCubeX/Clash.Mini/log"
	"github.com/MetaCubeX/Clash.Mini/notify"
	p "github.com/MetaCubeX/Clash.Mini/profile"

	"github.com/JyCyunMe/go-i18n/i18n"
	"github.com/fsnotify/fsnotify"
	stx "github.com/getlantern/systray"
)

var (
	mSwitchProfile *stx.MenuItemEx
	mUpdateAll     *stx.MenuItemEx
)

func SetMSwitchProfile(mie *stx.MenuItemEx) {
	mSwitchProfile = mie
	common.RefreshProfile = func(event *fsnotify.Event) {
		ResetProfiles(event)
		SwitchProfile()
	}
	mUpdateAll = mSwitchProfile.AddSubMenuItemExI18n(stx.NewI18nConfig(stx.I18nConfig{TitleID: cI18n.TrayMenuSwitchProfileUpdate}), func(menuItemEx *stx.MenuItemEx) {
		go func() {
			successNum := 0
			failNum := 0
			for e := p.Profiles.Front(); e != nil; e = e.Next() {
				profile := e.Value.(*p.RawData).FileInfo
				successful := p.UpdateConfig(profile.Name, profile.Url)
				if !successful {
					log.Errorln(fmt.Sprintf("%s: %s", i18n.T(cI18n.MenuConfigCronUpdateFailed), profile.Name))
					failNum++
				} else {
					log.Infoln(fmt.Sprintf("%s: %s", i18n.T(cI18n.MenuConfigCronUpdateSuccessful), profile.Name))
					successNum++
				}
			}
			notify.PushProfileUpdateFinished(successNum, failNum)
		}()
		//go common.RefreshProfile(nil)
	}).AddSeparator()
}

func ResetProfiles(event *fsnotify.Event) {
	if mSwitchProfile == nil {
		return
	}

	p.RefreshProfiles(event)
	if event == nil {
		log.Infoln("[profile] loaded %d profile(s)", p.Profiles.Len())
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
	//mSwitchProfile.ClearChildren()
	mSwitchProfile.ForChildrenLoop(true, func(_ int, profile *stx.MenuItemEx) (remove bool) {
		_, exists := p.RawDataMap.Load(profile.GetTitle())
		return !exists
	})

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
			return
			//p.RemoveProfile(profileName)
			//rawData.MenuItemEx.Delete()
		}
	} else {
		rawData = &p.RawData{}
	}

	mP := mSwitchProfile.AddSubMenuItemEx(profileName, profileName, func(menuItemEx *stx.MenuItemEx) {
		log.Infoln("switch profile to \\%s\\", menuItemEx.GetTitle())
		// TODO: switch
		message := i18n.TData(cI18n.MenuConfigMessageEnableConfigSuccess, &i18n.Data{Data: map[string]interface{}{
			"Config": menuItemEx.GetTitle(),
		}})

		if !controller.ApplyConfig(menuItemEx.GetTitle(), false) {
			message = i18n.TData(cI18n.MenuConfigMessageEnableConfigFailure, &i18n.Data{Data: map[string]interface{}{
				"Config": menuItemEx.GetTitle(),
			}})
		} else {
			menuItemEx.SwitchCheckboxBrother(true)
		}

		notify.PushWithLine(i18n.T(cI18n.NotifyMessageTitle), message)
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
	//ResetProfiles(nil)
}

func SwitchProfile() {
	if mSwitchProfile == nil {
		return
	}

	defer p.Locker.Unlock()
	p.Locker.Lock()
	configName := config.GetProfile()
	mSwitchProfile.ForChildrenLoop(true, func(_ int, profile *stx.MenuItemEx) (remove bool) {
		if profile.GetId() == mUpdateAll.GetId() {
			return
		}
		//log.Infoln("into:: %s", profile.GetTitle() + constant.ConfigSuffix)
		if configName == profile.GetTitle() {
			// Initial start
			if controller.ApplyConfig(configName, false) {
				profile.Check()
			} else {
				profile.Uncheck()

				message := i18n.TData(cI18n.MenuConfigMessageEnableConfigFailure, &i18n.Data{Data: map[string]interface{}{
					"Config": configName,
				}})

				notify.PushWithLine(i18n.T(cI18n.NotifyMessageTitle), message)
				go func() {
					time.Sleep(constant.NotifyDelay)
					userInfo := p.UpdateSubscriptionUserInfo()
					if len(userInfo.UnusedInfo) > 0 {
						notify.PushFlowInfo(userInfo.UsedInfo, userInfo.UnusedInfo, userInfo.ExpireInfo)
					}
				}()
			}
		} else {
			profile.Uncheck()
		}
		return
	})
}
