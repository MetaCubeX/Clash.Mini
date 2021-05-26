package notify

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/Clash-Mini/Clash.Mini/icon"
	"github.com/go-toast/toast"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	content    string
	appPath, _ = iconBytesToFilePath(icon.DateS)
)

func Notify(info string) {

	switch info {
	case "SysON":
		content = "--------------------\n系统代理：✅"
	case "SysOFF":
		content = "--------------------\n系统代理：❎"
	case "Direct":
		content = "--------------------\n已切换为：直连模式-✅"
	case "Rule":
		content = "--------------------\n已切换为：规则模式-✅"
	case "Global":
		content = "--------------------\n已切换为：全局模式-✅"
	case "Startup":
		content = "--------------------\n开机启动：✅"
	case "StartupOff":
		content = "--------------------\n开机启动：❎"
	case "SysAutoON":
		content = "--------------------\n默认代理：✅"
	case "SysAutoOFF":
		content = "--------------------\n默认代理：❎"
	case "Max":
		content = "--------------------\n成功切换：Maxmind数据库"
	case "Lite":
		content = "--------------------\n成功切换：Hackl0us数据库"
	}
	notification := toast.Notification{
		AppID:   "Clash.Mini",
		Title:   "📢通知📢",
		Icon:    appPath,
		Message: content,
	}
	err := notification.Push()
	if err != nil {
	}
}

func NotifyINFO(UsedINFO, UnUsedINFO, ExpireINFO string) {
	content = "--------------------\n已用流量：" + UsedINFO + "\n剩余流量：" + UnUsedINFO + "\n到期时间：" + ExpireINFO
	notification := toast.Notification{
		AppID:   "Clash.Mini",
		Title:   "📢流量信息📢",
		Icon:    appPath,
		Message: content,
	}
	err := notification.Push()
	if err != nil {
	}
}

func iconBytesToFilePath(iconBytes []byte) (string, error) {
	bh := md5.Sum(iconBytes)
	dataHash := hex.EncodeToString(bh[:])
	iconFilePath := filepath.Join(os.TempDir(), "systray_temp_icon_"+dataHash)

	if _, err := os.Stat(iconFilePath); os.IsNotExist(err) {
		if err := ioutil.WriteFile(iconFilePath, iconBytes, 0644); err != nil {
			return "", err
		}
	}
	return iconFilePath, nil
}
