package notify

import (
	"github.com/Clash-Mini/Clash.Mini/icon"
	"github.com/go-toast/toast"
	"github.com/lxn/walk"
)

var (
	content    string
	appPath, _ = walk.IconBytesToFilePath(icon.Date)
)

func Notify(info string) {

	switch info {
	case "Sys":
		content = "开-✅ 成功设置系统代理"
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
