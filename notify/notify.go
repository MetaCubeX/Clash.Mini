package notify

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	cI18n "github.com/Clash-Mini/Clash.Mini/constant/i18n"
	"github.com/JyCyunMe/go-i18n/i18n"
	"io/ioutil"
	"os"
	path "path/filepath"
	"time"

	"github.com/Clash-Mini/Clash.Mini/app"
	"github.com/Clash-Mini/Clash.Mini/cmd"
	"github.com/Clash-Mini/Clash.Mini/cmd/auto"
	"github.com/Clash-Mini/Clash.Mini/cmd/cron"
	"github.com/Clash-Mini/Clash.Mini/cmd/mmdb"
	"github.com/Clash-Mini/Clash.Mini/cmd/proxy"
	"github.com/Clash-Mini/Clash.Mini/cmd/startup"
	"github.com/Clash-Mini/Clash.Mini/cmd/sys"
	"github.com/Clash-Mini/Clash.Mini/icon"
	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/go-toast/toast"
)

const (
	notifyLine = "--------------------\n"
)

var (
	iconPath, _ = iconBytesToFilePath(icon.DateS)
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
	switch value {
	case sys.ON:
		message = i18n.T(cI18n.NotifyMessageSysOn)
		break
	case sys.OFF:
		message = i18n.T(cI18n.NotifyMessageSysOff)
		break
	case proxy.Direct:
		message = i18n.T(cI18n.NotifyMessageModeDirect)
		break
	case proxy.Rule:
		message = i18n.T(cI18n.NotifyMessageModeRULE)
		break
	case proxy.Global:
		message = i18n.T(cI18n.NotifyMessageModeGLOBAL)
		break
	case startup.ON:
		message = i18n.T(cI18n.NotifyMessageStartupOn)
		break
	case startup.OFF:
		message = i18n.T(cI18n.NotifyMessageStartupOff)
		break
	case auto.ON:
		message = i18n.T(cI18n.NotifyMessageAutoOn)
		break
	case auto.OFF:
		message = i18n.T(cI18n.NotifyMessageAutoOff)
		break
	case mmdb.Max:
		message = i18n.T(cI18n.NotifyMessageMmdbMax)
		break
	case mmdb.Lite:
		message = i18n.T(cI18n.NotifyMessageMmdbLite)
		break
	case cron.ON:
		message = i18n.T(cI18n.NotifyMessageCronOn)
		break
	case cron.OFF:
		message = i18n.T(cI18n.NotifyMessageCronOff)
		break
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

func PushWithLine(title string, message string) {
	PushMessage(title, getNotifyContent(message))
}

func PushMessage(title string, message string) {
	notification := toast.Notification{
		AppID:   app.Name,
		Title:   title,
		Icon:    iconPath,
		Message: message,
	}
	err := notification.Push()
	if err != nil {
		log.Errorln("Notify Push error: %v", err)
	}
}

func iconBytesToFilePath(iconBytes []byte) (string, error) {
	bh := md5.Sum(iconBytes)
	dataHash := hex.EncodeToString(bh[:])
	iconFilePath := path.Join(os.TempDir(), "systray_temp_icon_"+dataHash)

	if _, err := os.Stat(iconFilePath); os.IsNotExist(err) {
		if err := ioutil.WriteFile(iconFilePath, iconBytes, 0644); err != nil {
			return "", err
		}
	}
	return iconFilePath, nil
}
