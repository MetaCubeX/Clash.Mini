package notify

import (
	"fmt"
	"time"

	"github.com/Clash-Mini/Clash.Mini/app"
	"github.com/Clash-Mini/Clash.Mini/cmd"
	"github.com/Clash-Mini/Clash.Mini/cmd/auto"
	"github.com/Clash-Mini/Clash.Mini/cmd/cron"
	"github.com/Clash-Mini/Clash.Mini/cmd/mmdb"
	"github.com/Clash-Mini/Clash.Mini/cmd/proxy"
	"github.com/Clash-Mini/Clash.Mini/cmd/startup"
	"github.com/Clash-Mini/Clash.Mini/cmd/sys"
	cI18n "github.com/Clash-Mini/Clash.Mini/constant/i18n"
	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Clash-Mini/Clash.Mini/static"

	"github.com/JyCyunMe/go-i18n/i18n"
	"github.com/go-toast/toast"
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

func DoTrayMenu(value cmd.GeneralType) {
	var message string
	switch value.GetCommandType() {
	case cmd.Sys: {
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
	case cmd.Proxy: {
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
	case cmd.Startup: {
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
	case cmd.Auto: {
		switch value {
		case auto.ON:
			message = i18n.T(cI18n.NotifyMessageAutoOn)
			break
		case auto.OFF:
			message = i18n.T(cI18n.NotifyMessageAutoOff)
			break
		}
		break
	}
	case cmd.MMDB: {
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
	case cmd.Cron: {
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
		AppID:   app.Name,
		Title:   title,
		Icon:    static.NotifyIconPath,
		Message: message,
	}
	err := notification.Push()
	if err != nil {
		log.Errorln("Notify Push error: %v", err)
	}
}
