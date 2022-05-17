package tray

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	cP "github.com/Dreamacro/clash/listener"
	"github.com/Dreamacro/clash/tunnel"
	"github.com/JyCyunMe/go-i18n/i18n"
	"github.com/MakeNowJust/hotkey"
	"github.com/MetaCubeX/Clash.Mini/cmd"
	"github.com/MetaCubeX/Clash.Mini/cmd/autosys"
	"github.com/MetaCubeX/Clash.Mini/cmd/breaker"
	"github.com/MetaCubeX/Clash.Mini/cmd/cron"
	Hkey "github.com/MetaCubeX/Clash.Mini/cmd/hotkey"
	"github.com/MetaCubeX/Clash.Mini/cmd/protocol"
	"github.com/MetaCubeX/Clash.Mini/cmd/startup"
	"github.com/MetaCubeX/Clash.Mini/cmd/sys"
	"github.com/MetaCubeX/Clash.Mini/cmd/task"
	"github.com/MetaCubeX/Clash.Mini/common"
	"github.com/MetaCubeX/Clash.Mini/config"
	"github.com/MetaCubeX/Clash.Mini/constant"
	cI18n "github.com/MetaCubeX/Clash.Mini/constant/i18n"
	"github.com/MetaCubeX/Clash.Mini/controller"
	"github.com/MetaCubeX/Clash.Mini/icon"
	"github.com/MetaCubeX/Clash.Mini/log"
	"github.com/MetaCubeX/Clash.Mini/mixin"
	"github.com/MetaCubeX/Clash.Mini/mixin/dns"
	"github.com/MetaCubeX/Clash.Mini/mixin/general"
	"github.com/MetaCubeX/Clash.Mini/mixin/tun"
	"github.com/MetaCubeX/Clash.Mini/notify"
	"github.com/MetaCubeX/Clash.Mini/sysproxy"
	httpUtils "github.com/MetaCubeX/Clash.Mini/util/http"
	"github.com/MetaCubeX/Clash.Mini/util/loopback"
	protocolUtils "github.com/MetaCubeX/Clash.Mini/util/protocol"
	stringUtils "github.com/MetaCubeX/Clash.Mini/util/string"
	"github.com/MetaCubeX/Clash.Mini/util/uac"
	stx "github.com/getlantern/systray"
	"github.com/lxn/walk"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/sys/windows"
	"net/http"
	"strings"
)

const (
	funcLogHeader = logHeader + ".func"
)

var (
	configName       = config.GetProfile()
	NeedLoadSelector = false
)

//func LoadSelector(mGroup *stx.MenuItemEx) {
//	if NeedLoadSelector && clashConfig.GroupsList.Len() > 0 {
//		groupNowMap := tunnel.Proxies()
//		SelectorMap = make(map[string]proxy.SelectorInfo)
//		util.JsonUnmarshal(stringUtils.IgnoreErrorBytes(json.Marshal(groupNowMap)), &SelectorMap)
//		for name, group := range SelectorMap {
//			if group.Now != "" {
//				proxyNow := SelectorMap[group.Now]
//				log.Infoln("[tray] load: %s -> %s", name, group.Name)
//				SwitchGroupAndProxy(mGroup, group.Name, proxyNow.Name)
//				continue
//			}
//		}
//		NeedLoadSelector = false
//	}
//}

func mConfigProxyFunc(mConfigProxy *stx.MenuItemEx) {
	configGroup := ConfigGroupsMap[mConfigProxy.Parent.GetId()]
	GroupPath := mConfigProxy.Parent.GetTitle()
	ProxyName := configGroup[mConfigProxy.GetId()]
	host := constant.ControllerHost
	port := constant.ControllerPort
	secret := constant.ControllerSecret
	url := fmt.Sprintf(`http://%s:%s/proxies/%s`, host, port, GroupPath)
	body := make(map[string]interface{})
	body["name"] = ProxyName
	bytesData, err := json.Marshal(body)
	if err != nil {
		log.Errorln("[tray]putConfig Marshal error: %v", err)
		return
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest(http.MethodPut, url, reader)
	if err != nil {
		log.Errorln("[%s] putConfig NewRequest error: %v", funcLogHeader, err)
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", secret))
	client := http.Client{}
	rsp, err := client.Do(request)
	defer httpUtils.DeferSafeCloseResponseBody(rsp)
	if err != nil {
		log.Errorln("[%s] putConfig Do error: %v", funcLogHeader, err)
		return
	}
	if rsp != nil && rsp.StatusCode != http.StatusNoContent {
		log.Errorln("[%s] putConfig Do error[HTTP %d]: %s \n %s", funcLogHeader, rsp.StatusCode)
		return
	}
	log.Infoln("[%s] PUT Proxies info:  Group: %s - Proxy: %s", funcLogHeader, GroupPath, ProxyName)
}

func mEnabledFunc(mEnabled *stx.MenuItemEx) {
	if mEnabled.Checked() {
		// 取消系统代理
		err := sysproxy.ClearSystemProxy()
		if err != nil {
			log.Errorln("[%s] cancel sysproxy failed: %v", funcLogHeader, err)
		} else {
			mEnabled.Uncheck()
			stx.SetIcon(icon.DateN)
			notify.DoTrayMenu(sys.OFF)
		}
	} else {
		// 设置系统代理
		var port int
		if cP.GetPorts().MixedPort != 0 {
			port = cP.GetPorts().MixedPort
		} else {
			port = cP.GetPorts().Port
		}
		err := sysproxy.SetSystemProxy(
			&sysproxy.ProxyConfig{
				Enable: true,
				Server: fmt.Sprintf("%s:%d", constant.LocalHost, port),
			})
		if err != nil {
			log.Errorln("[%s] setting sysproxy failed: %v", funcLogHeader, err)
		} else {
			mEnabled.Check()
			stx.SetIcon(icon.DateS)
			notify.DoTrayMenu(sys.ON)
		}
	}
	firstInit = true
}

func mOtherAutosysFunc(mOtherAutosys *stx.MenuItemEx) {
	var autosysType autosys.Type
	if mOtherAutosys.Checked() {
		autosysType = autosys.OFF
		if config.IsCmdPositive(cmd.Autosys) {
			notify.DoTrayMenuDelay(autosys.OFF, constant.NotifyDelay)
		}
	} else {
		autosysType = autosys.ON
		if !config.IsCmdPositive(cmd.Autosys) {
			notify.DoTrayMenuDelay(autosys.ON, constant.NotifyDelay)
		}
	}
	config.SetCmd(autosysType)
	firstInit = true

}

func mOtherProtocolFunc(mOthersProtocol *stx.MenuItemEx) {
	var protocolValue protocol.Type
	if mOthersProtocol.Checked() {
		mOthersProtocol.Uncheck()
		protocolValue = protocol.OFF
	} else {
		protocolValue = protocol.ON
	}

	registerStr := stringUtils.TrinocularString(protocolValue.IsPositive(), "register", "unregister")
	// TODO: agent mode
	if uac.AmAdmin {
		err := protocolUtils.RegisterCommandProtocol(mOthersProtocol.Checked())
		if err != nil {
			return
		}
	} else {
		err := uac.RunMeWithArg(uac.GetCallArg(stringUtils.TrinocularString(mOthersProtocol.Checked(),
			"--uac-protocol-disable", "--uac-protocol-enable")), "")
		if err != nil {
			if errors.Is(err, windows.ERROR_CANCELLED) {
				log.Warnln("[%s] %s protocol cancelled: %v", funcLogHeader, registerStr, err)
				return
			}
			//if errors.Is(err, &exec.ExitError{}) {
			//	return
			//}
			log.Errorln("[%s] %s protocol failed: %v", funcLogHeader, registerStr, err)
			return
		}
	}
	log.Infoln("[%s] %s protocol success", funcLogHeader, registerStr)

	config.SetCmd(protocolValue)
	firstInit = true
}

func mOtherUwpLoopbackFunc(mOthersUwpLoopback *stx.MenuItemEx) {
	var loopbackValue breaker.Type
	if mOthersUwpLoopback.Checked() {
		loopbackValue = breaker.OFF
	} else {
		loopbackValue = breaker.ON
	}

	if uac.AmAdmin {
		go loopback.Breaker(loopbackValue)
	} else {
		err := uac.RunMeWithArg(uac.GetCallArg(stringUtils.TrinocularString(mOthersUwpLoopback.Checked(),
			"--uac-loopback-disable", "--uac-loopback-enable")), "")
		if err != nil {
			if errors.Is(err, windows.ERROR_CANCELLED) {
				log.Warnln("[%s] operate loopback breaker cancelled: %v", funcLogHeader, err)
				return
			}
			//if errors.Is(err, &exec.ExitError{}) {
			//	return
			//}
			log.Errorln("[%s] operate loopback breaker failed: %v", funcLogHeader, err)
			return
		}
	}
	log.Infoln("[%s] operate loopback breaker success: %s", funcLogHeader, loopbackValue.String())

	if mOthersUwpLoopback.Checked() {
		mOthersUwpLoopback.Uncheck()
		mOthersAutosys.Enable()
	} else {
		mOthersUwpLoopback.Check()
		mEnabledFunc(mEnabled)
	}
	config.SetCmd(loopbackValue)
	firstInit = true
}

func mOtherTaskFunc(mOtherTask *stx.MenuItemEx) {
	var err error
	var startupType startup.Type
	if mOtherTask.Checked() {
		startupType = startup.OFF
		err = task.DoCommand(task.OFF)
		if config.IsCmdPositive(cmd.Startup) {
			notify.DoTrayMenuDelay(startupType, constant.NotifyDelay)
		}
	} else {
		startupType = startup.ON
		err = task.DoCommand(task.ON)
		if !config.IsCmdPositive(cmd.Startup) {
			notify.DoTrayMenuDelay(startupType, constant.NotifyDelay)
		}
	}
	err = config.SetCmd(startupType)
	if err != nil {
		return
	}
	firstInit = true
}

func mOthersMixinDirFunc(mOthersMixinDir *stx.MenuItemEx) {
	if mOthersMixinDir.Checked() {
		return
	} else {
		err := open.Run(constant.MixinDir)
		if err != nil {
			return
		}
	}
}

func mOthersMixinGeneralFunc(mOthersMixinGeneral *stx.MenuItemEx) {
	var generalType general.Type
	configName := config.GetProfile()
	if mOthersMixinGeneral.Checked() {
		generalType = general.OFF
		if config.IsMixinPositive(mixin.General) {
			notify.DoTrayMenuMixinDelay(general.OFF, constant.NotifyDelay)
		}
	} else {
		generalType = general.ON
		if !config.IsMixinPositive(mixin.General) {
			notify.DoTrayMenuMixinDelay(general.ON, constant.NotifyDelay)
		}
	}
	config.SetMixin(generalType)
	controller.ApplyConfig(strings.TrimSuffix(configName, constant.ConfigSuffix), false)
	firstInit = true
}

func mOthersMixinTunFunc(mOthersMixinTun *stx.MenuItemEx) {
	var tunType tun.Type

	if mOthersMixinTun.Checked() {
		tunType = tun.OFF
		if config.IsMixinPositive(mixin.Tun) {
			notify.DoTrayMenuMixinDelay(tun.OFF, constant.NotifyDelay)
		}
	} else {
		tunType = tun.ON
		if !config.IsMixinPositive(mixin.Tun) {
			notify.DoTrayMenuMixinDelay(tun.ON, constant.NotifyDelay)
		}
	}
	config.SetMixin(tunType)

	if !uac.AmAdmin {
		msg := "Please quit & restart the software in administrator mode!"
		walk.MsgBox(nil, i18n.T(cI18n.MsgBoxTitleTips), msg, walk.MsgBoxIconInformation)
	} else {
		controller.ApplyConfig(strings.TrimSuffix(configName, constant.ConfigSuffix), false)
	}
	firstInit = true
}

func mOthersMixinDnsFunc(mOthersMixinDns *stx.MenuItemEx) {
	var dnsType dns.Type
	configName := config.GetProfile()
	if mOthersMixinDns.Checked() {
		dnsType = dns.OFF
		if config.IsMixinPositive(mixin.Dns) {
			notify.DoTrayMenuMixinDelay(dns.OFF, constant.NotifyDelay)
		}
	} else {
		dnsType = dns.ON
		if !config.IsMixinPositive(mixin.Dns) {
			notify.DoTrayMenuMixinDelay(dns.ON, constant.NotifyDelay)
		}
	}
	config.SetMixin(dnsType)
	controller.ApplyConfig(strings.TrimSuffix(configName, constant.ConfigSuffix), false)
	firstInit = true
}

func mOtherUpdateCronFunc(mOtherUpdateCron *stx.MenuItemEx) {
	var cronType cron.Type
	if mOtherUpdateCron.Checked() {
		cronType = cron.OFF
		if config.IsCmdPositive(cmd.Cron) {
			notify.DoTrayMenuDelay(cron.OFF, constant.NotifyDelay)
		}
	} else {
		cronType = cron.ON
		if !config.IsCmdPositive(cmd.Cron) {
			notify.DoTrayMenuDelay(cron.ON, constant.NotifyDelay)
		}
	}
	config.SetCmd(cronType)
	firstInit = true
}

func mOtherHotkeyFunc(mOthersHotkey *stx.MenuItemEx) {
	var HotkeyType Hkey.Type
	if mOthersHotkey.Checked() {
		HotkeyType = Hkey.OFF
		if config.IsCmdPositive(cmd.Hotkey) {
			hotKey(false)
		}
	} else {
		HotkeyType = Hkey.ON
		if !config.IsCmdPositive(cmd.Hotkey) {
			hotKey(true)
		}
	}
	config.SetCmd(HotkeyType)
	firstInit = true
}

var (
	id1, id2, id3, id4     hotkey.Id
	err1, err2, err3, err4 error
	HotKeys                = hotkey.New()
)

func hotKey(b bool) {
	message := ""
	if b {
		if common.DisabledCore {
			mCoreProxyMode.I18nConfig.TitleConfig.Format = "\tAlt+P"
		}
		id1, err1 = HotKeys.Register(hotkey.Alt, 'R', func() {
			tunnel.SetMode(tunnel.Rule)
		})
		if err1 != nil {
			message += "Alt+R热键注册失败\n"
		} else {
			mRule.I18nConfig.TitleConfig.Format = "\tAlt+R"
		}
		id2, err2 = HotKeys.Register(hotkey.Alt, 'G', func() {
			tunnel.SetMode(tunnel.Global)
		})
		if err2 != nil {
			message += "Alt+G热键注册失败\n"
		} else {
			mGlobal.I18nConfig.TitleConfig.Format = "\tAlt+G"
		}
		id3, err3 = HotKeys.Register(hotkey.Alt, 'D', func() {
			tunnel.SetMode(tunnel.Direct)
		})
		if err3 != nil {
			message += "Alt+D热键注册失败\n"
		} else {
			mDirect.I18nConfig.TitleConfig.Format = "\tAlt+D"
		}
		id4, err4 = HotKeys.Register(hotkey.Alt, 'S', func() {
			mEnabledFunc(mEnabled)
		})
		if err4 != nil {
			message += "Alt+S热键注册失败\n"
		} else {
			mEnabled.I18nConfig.TitleConfig.Format = "\tAlt+S"
		}
		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			notify.PushWithLine(cI18n.NotifyMessageTitle, message)
		}
	} else {
		if common.DisabledCore {
			mCoreProxyMode.I18nConfig.TitleConfig.Format = ""
		}
		mRule.I18nConfig.TitleConfig.Format = ""
		mGlobal.I18nConfig.TitleConfig.Format = ""
		mDirect.I18nConfig.TitleConfig.Format = ""
		mEnabled.I18nConfig.TitleConfig.Format = ""
		HotKeys.Unregister(id1)
		HotKeys.Unregister(id2)
		HotKeys.Unregister(id3)
		HotKeys.Unregister(id4)

	}
	mEnabled.SwitchLanguage()
	mDirect.SwitchLanguage()
	mGlobal.SwitchLanguage()
	mRule.SwitchLanguage()
	mCoreProxyMode.SwitchLanguageWithChildren()
}
