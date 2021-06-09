package notify

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	path "path/filepath"
	"time"

	"github.com/Clash-Mini/Clash.Mini/cmd"
	"github.com/Clash-Mini/Clash.Mini/cmd/auto"
	"github.com/Clash-Mini/Clash.Mini/cmd/cron"
	"github.com/Clash-Mini/Clash.Mini/cmd/mmdb"
	"github.com/Clash-Mini/Clash.Mini/cmd/proxy"
	"github.com/Clash-Mini/Clash.Mini/cmd/startup"
	"github.com/Clash-Mini/Clash.Mini/cmd/sys"
	"github.com/Clash-Mini/Clash.Mini/icon"
	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Clash-Mini/Clash.Mini/util"

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
		message = "系统代理：✅"
		break
	case sys.OFF:
		message = "系统代理：❎"
		break
	case proxy.Direct:
		message = "已切换为：直连模式-✅"
		break
	case proxy.Rule:
		message = "已切换为：规则模式-✅"
		break
	case proxy.Global:
		message = "已切换为：全局模式-✅"
		break
	case startup.ON:
		message = "开机启动：✅"
		break
	case startup.OFF:
		message = "开机启动：❎"
		break
	case auto.ON:
		message = "默认代理：✅"
		break
	case auto.OFF:
		message = "默认代理：❎"
		break
	case mmdb.Max:
		message = "成功切换：Maxmind数据库"
		break
	case mmdb.Lite:
		message = "成功切换：Hackl0us数据库"
		break
	case cron.ON:
		message = "定时更新：✅"
		break
	case cron.OFF:
		message = "定时更新：❎"
		break
	}
	PushWithLine("📢通知📢", message)
}

func PushFlowInfo(usedInfo, unUsedInfo, expireInfo string) {
	PushWithLine("📢流量信息📢",
		fmt.Sprintf("已用流量：%s\n剩余流量：%s\n到期时间：%s", usedInfo, unUsedInfo, expireInfo))
}

func PushProfileCronFinished(successNum, failNum int) {
	message := "定时更新完成：✅\n"
	if failNum > 0 {
		message = fmt.Sprintf("%s定时更新完成：✅\n[%d] 个配置更新成功！\n[%d] 个配置更新失败！", message, successNum, failNum)
	} else {
		message += "全部配置更新成功！"
	}
	PushWithLine("📢更新通知📢", message)
}

func PushWithLine(title string, message string) {
	PushMessage(title, getNotifyContent(message))
}

func PushMessage(title string, message string) {
	notification := toast.Notification{
		AppID:   util.AppTitle,
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
