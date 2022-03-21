package notify

import (
	"fmt"
	"github.com/MetaCubeX/Clash.Mini/cmd/autosys"
	"github.com/MetaCubeX/Clash.Mini/cmd/sys"
	"github.com/MetaCubeX/Clash.Mini/mixin"
	"github.com/MetaCubeX/Clash.Mini/mixin/dns"
	"github.com/MetaCubeX/Clash.Mini/mixin/general"
	"github.com/MetaCubeX/Clash.Mini/mixin/tun"
	"os"
	"time"

	"github.com/MetaCubeX/Clash.Mini/cmd"
	"github.com/MetaCubeX/Clash.Mini/cmd/cron"
	"github.com/MetaCubeX/Clash.Mini/cmd/mmdb"
	"github.com/MetaCubeX/Clash.Mini/cmd/proxy"
	"github.com/MetaCubeX/Clash.Mini/cmd/startup"
	cI18n "github.com/MetaCubeX/Clash.Mini/constant/i18n"
	"github.com/MetaCubeX/Clash.Mini/log"
	"github.com/MetaCubeX/Clash.Mini/static"

	"github.com/JyCyunMe/go-i18n/i18n"
	"github.com/go-toast/toast"
)

const (
	logHeader = "notify"
)

const (
	notifyLine = "--------------------\n"
)

func getNotifyContent(s string) string {
	return notifyLine + s
}

func DoTrayMenuDelay(value cmd.GeneralType, delay time.Duration) {
	time.AfterFunc(delay, func() {
		DoTrayMenu(value)
	})
}

func DoTrayMenuMixinDelay(value mixin.GeneralType, delay time.Duration) {
	time.AfterFunc(delay, func() {
		DoTrayMenuMixin(value)
	})
}

func DoTrayMenuMixin(value mixin.GeneralType) {
	var message string
	switch value.GetCommandType() {
	case mixin.General:
		switch value {
		case general.ON:
			message = i18n.T(cI18n.NotifyMessageMixinGeneralOn)
			break
		case general.OFF:
			message = i18n.T(cI18n.NotifyMessageMixinGeneralOff)
			break
		}
	case mixin.Tun:
		switch value {
		case tun.ON:
			message = i18n.T(cI18n.NotifyMessageMixinTunOn)
			break
		case tun.OFF:
			message = i18n.T(cI18n.NotifyMessageMixinTunOff)
			break
		}
	case mixin.Dns:
		switch value {
		case dns.ON:
			message = i18n.T(cI18n.NotifyMessageMixinDnsOn)
			break
		case dns.OFF:
			message = i18n.T(cI18n.NotifyMessageMixinDnsOff)
			break
		}
	}
	PushWithLine(i18n.T(cI18n.NotifyMessageTitle), message)
}

func DoTrayMenu(value cmd.GeneralType) {
	var message string
	switch value.GetCommandType() {

	case cmd.Sys:
		{
			switch value {
			case sys.ON:
				message = i18n.T(cI18n.NotifyMessageSysOn)
				break
			case sys.OFF:
				message = i18n.T(cI18n.NotifyMessageSysOff)
				break
			}
			break
		}
	case cmd.Proxy:
		{
			switch value {
			case proxy.Direct:
				message = i18n.T(cI18n.NotifyMessageModeDirect)
				break
			case proxy.Rule:
				message = i18n.T(cI18n.NotifyMessageModeRULE)
				break
			case proxy.Global:
				message = i18n.T(cI18n.NotifyMessageModeGLOBAL)
				break
			}
			break
		}
	case cmd.Startup:
		{
			switch value {
			case startup.ON:
				message = i18n.T(cI18n.NotifyMessageStartupOn)
				break
			case startup.OFF:
				message = i18n.T(cI18n.NotifyMessageStartupOff)
				break
			}
			break
		}
	case cmd.Autosys:
		{
			switch value {
			case autosys.ON:
				message = i18n.T(cI18n.NotifyMessageAutosysOn)
				break
			case autosys.OFF:
				message = i18n.T(cI18n.NotifyMessageAutosysOff)
				break
			}
			break
		}
	case cmd.MMDB:
		{
			switch value {
			case mmdb.Max:
				message = i18n.T(cI18n.NotifyMessageMmdbMax)
				break
			case mmdb.Lite:
				message = i18n.T(cI18n.NotifyMessageMmdbLite)
				break
			}
			break
		}
	case cmd.Cron:
		{
			switch value {
			case cron.ON:
				message = i18n.T(cI18n.NotifyMessageCronOn)
				break
			case cron.OFF:
				message = i18n.T(cI18n.NotifyMessageCronOff)
				break
			}
			break
		}
	}
	PushWithLine(i18n.T(cI18n.NotifyMessageTitle), message)
}

func PushFlowInfo(usedInfo, unUsedInfo, expireInfo string) {
	PushWithLine(i18n.T(cI18n.NotifyMessageFlowTitle),
		fmt.Sprintf("%s：%s\n%s：%s\n%s：%s", i18n.T(cI18n.NotifyMessageFlowUsed), usedInfo, i18n.T(cI18n.NotifyMessageFlowUnused), unUsedInfo, i18n.T(cI18n.NotifyMessageFlowExpiration), expireInfo))
}

func PushProfileUpdateFinished(successNum, failNum int) {
	message := i18n.T(cI18n.NotifyMessageUpdateFinish) + "\n"
	if failNum > 0 {
		message = fmt.Sprintf("%s[%d] %s\n[%d] %s", message, successNum, i18n.T(cI18n.NotifyMessageCronNumSuccess), failNum, i18n.T(cI18n.NotifyMessageCronNumFail))
	} else {
		message += i18n.T(cI18n.NotifyMessageCronFinishAll)
	}
	PushWithLine(i18n.T(cI18n.NotifyMessageCronTitle), message)
}

func PushProfileCronFinished(successNum, failNum int) {
	message := i18n.T(cI18n.NotifyMessageCronFinish) + "\n"
	if failNum > 0 {
		message = fmt.Sprintf("%s[%d] %s\n[%d] %s", message, successNum, i18n.T(cI18n.NotifyMessageCronNumSuccess), failNum, i18n.T(cI18n.NotifyMessageCronNumFail))
	} else {
		message += i18n.T(cI18n.NotifyMessageCronFinishAll)
	}
	PushWithLine(i18n.T(cI18n.NotifyMessageCronTitle), message)
}

func PushError(title string, message string) {
	if len(title) == 0 {
		title = i18n.T(cI18n.NotifyMessageErrorTitle)
	}
	go PushMessage(title, getNotifyContent(message))
}

func PushWithLine(title string, message string) {
	go PushMessage(title, getNotifyContent(message))
}

func PushMessage(title string, message string) {
	notification := toast.Notification{
		AppID:   os.Args[0],
		Title:   title,
		Icon:    static.NotifyIconPath,
		Message: message,
	}
	err := notification.Push()
	if err != nil {
		log.Errorln("[%s] Notify Push error: %v", logHeader, err)
	}
}
