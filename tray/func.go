package tray

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Clash-Mini/Clash.Mini/cmd"
	"github.com/Clash-Mini/Clash.Mini/cmd/autosys"
	"github.com/Clash-Mini/Clash.Mini/cmd/breaker"
	"github.com/Clash-Mini/Clash.Mini/cmd/cron"
	"github.com/Clash-Mini/Clash.Mini/cmd/mmdb"
	"github.com/Clash-Mini/Clash.Mini/cmd/protocol"
	"github.com/Clash-Mini/Clash.Mini/cmd/startup"
	"github.com/Clash-Mini/Clash.Mini/cmd/sys"
	"github.com/Clash-Mini/Clash.Mini/cmd/task"
	"github.com/Clash-Mini/Clash.Mini/config"
	"github.com/Clash-Mini/Clash.Mini/constant"
	cI18n "github.com/Clash-Mini/Clash.Mini/constant/i18n"
	"github.com/Clash-Mini/Clash.Mini/controller"
	"github.com/Clash-Mini/Clash.Mini/icon"
	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Clash-Mini/Clash.Mini/mixin"
	"github.com/Clash-Mini/Clash.Mini/mixin/dns"
	"github.com/Clash-Mini/Clash.Mini/mixin/script"
	"github.com/Clash-Mini/Clash.Mini/mixin/tun"
	"github.com/Clash-Mini/Clash.Mini/notify"
	"github.com/Clash-Mini/Clash.Mini/proxy"
	"github.com/Clash-Mini/Clash.Mini/sysproxy"
	"github.com/Clash-Mini/Clash.Mini/util"
	httpUtils "github.com/Clash-Mini/Clash.Mini/util/http"
	"github.com/Clash-Mini/Clash.Mini/util/loopback"
	protocolUtils "github.com/Clash-Mini/Clash.Mini/util/protocol"
	stringUtils "github.com/Clash-Mini/Clash.Mini/util/string"
	"github.com/Clash-Mini/Clash.Mini/util/uac"
	clashConfig "github.com/Dreamacro/clash/config"
	cP "github.com/Dreamacro/clash/listener"
	"github.com/Dreamacro/clash/tunnel"
	"github.com/JyCyunMe/go-i18n/i18n"
	"github.com/MakeNowJust/hotkey"
	stx "github.com/getlantern/systray"
	"github.com/lxn/walk"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/sys/windows"
	"net/http"
)

const (
	funcLogHeader = logHeader + ".func"
)

var (
	ControllerPort   = constant.ControllerPort
	NeedLoadSelector = false
	configName, _    = controller.CheckConfig()
)

func LoadSelector(mGroup *stx.MenuItemEx) {
	if NeedLoadSelector && clashConfig.GroupsList.Len() > 0 {
		groupNowMap := tunnel.Proxies()
		SelectorMap = make(map[string]proxy.SelectorInfo)
		util.JsonUnmarshal(stringUtils.IgnoreErrorBytes(json.Marshal(groupNowMap)), &SelectorMap)
		for name, group := range SelectorMap {
			if group.Now != "" {
				proxyNow := SelectorMap[group.Now]
				log.Infoln("[tray] load: %s -> %s", name, group.Name)
				SwitchGroupAndProxy(mGroup, group.Name, proxyNow.Name)
				continue
			}
		}
		NeedLoadSelector = false
	}
}

func mConfigProxyFunc(mConfigProxy *stx.MenuItemEx) {
	configGroup := ConfigGroupsMap[mConfigProxy.Parent.GetId()]
	GroupPath := mConfigProxy.Parent.GetTitle()
	ProxyName := configGroup[mConfigProxy.GetId()]

	url := fmt.Sprintf(`http://%s:%s/proxies/%s`, constant.Localhost, ControllerPort, GroupPath)
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
	client := http.Client{}
	rsp, err := client.Do(request)
	defer httpUtils.DeferSafeCloseResponseBody(rsp)
	if err != nil {
		log.Errorln("[%s] putConfig Do error: %v", funcLogHeader, err)
		return
	}
	if rsp != nil && rsp.StatusCode != http.StatusNoContent {
		log.Errorln("[%s] putConfig Do error[HTTP %d]: %s", funcLogHeader, rsp.StatusCode, url)
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
				Server: fmt.Sprintf("%s:%d", constant.Localhost, port),
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
	controller.PutConfig(configName)
	if !uac.AmAdmin {
		msg := "Please quit & restart the software in administrator mode!"
		walk.MsgBox(nil, i18n.T(cI18n.MsgBoxTitleTips), msg, walk.MsgBoxIconInformation)
	}
	firstInit = true
}

func mOthersMixinDnsFunc(mOthersMixinDns *stx.MenuItemEx) {
	var dnsType dns.Type
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
	controller.PutConfig(configName)
	firstInit = true
}

func mOthersMixinScriptFunc(mOthersMixinScript *stx.MenuItemEx) {
	var scriptType script.Type
	if mOthersMixinScript.Checked() {
		scriptType = script.OFF
		if config.IsMixinPositive(mixin.Script) {
			notify.DoTrayMenuMixinDelay(script.OFF, constant.NotifyDelay)
		}
	} else {
		scriptType = script.ON
		if !config.IsMixinPositive(mixin.Script) {
			notify.DoTrayMenuMixinDelay(script.ON, constant.NotifyDelay)
		}
	}
	config.SetMixin(scriptType)
	controller.PutConfig(configName)
	firstInit = true
}

func maxMindMMBDFunc(maxMindMMBD *stx.MenuItemEx) {
	if maxMindMMBD.Checked() {
		return
	} else {
		controller.GetMMDB(mmdb.Max)
		if !config.IsCmdPositive(cmd.MMDB) {
			notify.DoTrayMenuDelay(mmdb.Max, constant.NotifyDelay)
		}
	}
	firstInit = true
}

func hackl0usMMDBFunc(hackl0usMMDB *stx.MenuItemEx) {
	if hackl0usMMDB.Checked() {
		return
	} else {
		controller.GetMMDB(mmdb.Lite)
		if config.IsCmdPositive(cmd.MMDB) {
			notify.DoTrayMenuDelay(mmdb.Lite, constant.NotifyDelay)
		}
	}
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

func hotKey(mEnabled *stx.MenuItemEx) {
	message := ""
	hkey := hotkey.New()
	_, err1 := hkey.Register(hotkey.Alt, 'R', func() {
		tunnel.SetMode(tunnel.Rule)
	})
	if err1 != nil {
		message += "Alt+R热键注册失败\n"
	}
	_, err2 := hkey.Register(hotkey.Alt, 'G', func() {
		tunnel.SetMode(tunnel.Global)
	})
	if err2 != nil {
		message += "Alt+G热键注册失败\n"
	}
	_, err3 := hkey.Register(hotkey.Alt, 'D', func() {
		tunnel.SetMode(tunnel.Direct)
	})
	if err3 != nil {
		message += "Alt+D热键注册失败\n"
	}
	_, err4 := hkey.Register(hotkey.Alt, 'S', func() {
		mEnabledFunc(mEnabled)
	})
	if err4 != nil {
		message += "Alt+S热键注册失败\n"
	}
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		notify.PushWithLine(cI18n.NotifyMessageTitle, message)
	}
}
