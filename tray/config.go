package tray

import (
	"github.com/Clash-Mini/Clash.Mini/constant"
	cI18n "github.com/Clash-Mini/Clash.Mini/constant/i18n"
	"github.com/Clash-Mini/Clash.Mini/controller"
	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Clash-Mini/Clash.Mini/notify"
	"github.com/JyCyunMe/go-i18n/i18n"
	stx "github.com/getlantern/systray"
	"github.com/lxn/walk"
	"io/ioutil"
	path "path/filepath"
	"strings"
	"time"
)

func resetProfile(mSwitchProfile *stx.MenuItemEx) {
	fileInfoArr, err := ioutil.ReadDir(constant.ConfigDir)
	if err != nil {
		log.Fatalln("ResetRows ReadDir error: %v", err)
	}
	var profileNames []string
	for _, f := range fileInfoArr {
		if path.Ext(f.Name()) == constant.ConfigSuffix {
			profileName := strings.TrimSuffix(f.Name(), path.Ext(f.Name()))
			profileNames = append(profileNames, profileName)
		}
	}
	for _, profileName := range profileNames {
		mSwitchProfile.AddSubMenuItemEx(profileName, profileName, func(menuItemEx *stx.MenuItemEx) {
			log.Infoln("switch profile to \\%s\\", menuItemEx.GetTitle())
			// TODO: switch
			controller.PutConfig(menuItemEx.GetTitle())
			walk.MsgBox(nil, i18n.T(cI18n.MessageBoxTitleTips),
				i18n.TData(cI18n.MenuConfigMessageEnableConfigSuccess, &i18n.Data{Data: map[string]interface{}{
					"Config": menuItemEx.GetTitle(),
				}}),
				walk.MsgBoxIconInformation)
			menuItemEx.SwitchCheckboxBrother(true)
			go func() {
				time.Sleep(constant.NotifyDelay)
				userInfo := controller.UpdateSubscriptionUserInfo()
				if len(userInfo.UnusedInfo) > 0 {
					notify.PushFlowInfo(userInfo.UsedInfo, userInfo.UnusedInfo, userInfo.ExpireInfo)
				}
			}()
		})
	}
}

func switchProfile(mSwitchProfile *stx.MenuItemEx) {
	configName, _ := controller.CheckConfig()
	for e := mSwitchProfile.Children.Front(); e != nil; e = e.Next() {
		Profile := e.Value.(*stx.MenuItemEx)
		if configName == Profile.GetTitle()+constant.ConfigSuffix {
			Profile.Check()
		}
	}
}
