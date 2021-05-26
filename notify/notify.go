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
		content = "系统代理：✅"
	case "SysOFF":
		content = "系统代理：❎"
	case "Direct":
		content = "已切换为：直连模式-✅"
	case "Rule":
		content = "已切换为：规则模式-✅"
	case "Global":
		content = "已切换为：全局模式-✅"
	case "Startup":
		content = "开机启动：✅"
	case "StartupOff":
		content = "开机启动：❎"
	case "SysAutoON":
		content = "默认代理：✅"
	case "SysAutoOFF":
		content = "默认代理：❎"
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
