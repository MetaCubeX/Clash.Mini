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
	stx "github.com/getlantern/systray"
)

var (
	mSwitchProfile 	*stx.MenuItemEx
	mUpdateAll		*stx.MenuItemEx
)

func SetMSwitchProfile(mie *stx.MenuItemEx) {
	mSwitchProfile = mie
	common.RefreshProfile = func() {
		ResetProfiles()
		SwitchProfile()
	}
	mUpdateAll = mSwitchProfile.AddSubMenuItemEx("一键更新", "", func(menuItemEx *stx.MenuItemEx) {
		for _, profile := range p.Profiles {
			p.UpdateConfig(profile.Name, profile.Url)
		}
		go common.RefreshProfile()
	}).AddSeparator()
}

func ResetProfiles() {
	if mSwitchProfile == nil {
		return
	}

	//mSwitchProfile.ClearChildren()
	mSwitchProfile.ForChildrenLoop(true, func(_ int, profile *stx.MenuItemEx) (remove bool) {
		return profile.GetId() != mUpdateAll.GetId()
	})

	if len(p.Profiles) == 0 {
		mSwitchProfile.Disable()
		return
	}
	mSwitchProfile.Enable()
	for _, profile := range p.Profiles {
		mSwitchProfile.AddSubMenuItemEx(profile.Name, profile.Name, func(menuItemEx *stx.MenuItemEx) {
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
	}
}

func SwitchProfile() {
	if mSwitchProfile == nil {
		return
	}

	configName, _ := controller.CheckConfig()
	mSwitchProfile.ForChildrenLoop(true, func(_ int, profile *stx.MenuItemEx) (remove bool) {
		if configName == profile.GetTitle() + constant.ConfigSuffix {
			profile.Check()
		} else {
			profile.Uncheck()
		}
		return
	})
}
