package notify

import "github.com/go-toast/toast"

func SysNotify() {
	notification := toast.Notification{
		AppID:   "Clash.Mini",
		Title:   "📢通知📢",
		Message: "开-✅ 成功设置系统代理",
	}
	err := notification.Push()
	if err != nil {
	}
}

func RuleNotify() {
	notification := toast.Notification{
		AppID:   "Clash.Mini",
		Title:   "📢通知📢",
		Message: "已切换为：规则模式-✅",
	}
	err := notification.Push()
	if err != nil {
	}
}

func DirectNotify() {
	notification := toast.Notification{
		AppID:   "Clash.Mini",
		Title:   "📢通知📢",
		Message: "已切换为：直连模式-✅",
	}
	err := notification.Push()
	if err != nil {
	}
}

func GlobalNotify() {
	notification := toast.Notification{
		AppID:   "Clash.Mini",
		Title:   "📢通知📢",
		Message: "已切换为：全局模式-✅",
	}
	err := notification.Push()
	if err != nil {
	}
}
