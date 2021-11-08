package tray

import (
	"github.com/Clash-Mini/Clash.Mini/app"
	"github.com/Clash-Mini/Clash.Mini/constant"
	"github.com/Clash-Mini/Clash.Mini/controller"
	"github.com/Clash-Mini/Clash.Mini/icon"
	"github.com/Clash-Mini/Clash.Mini/log"
	commonUtils "github.com/Clash-Mini/Clash.Mini/util/common"
	stringUtils "github.com/Clash-Mini/Clash.Mini/util/string"

	C "github.com/Dreamacro/clash/constant"
	stx "github.com/getlantern/systray"
)

var (
	mainTitle	string
	mainTooltip string
)

func init() {
	if commonUtils.IsWindows() {
		C.SetHomeDir(constant.Pwd)
	}
	// 装载托盘
	stx.RunEx(onReady, onExit)
}

// onReady 托盘启动时
func onReady() {
	log.Infoln("[tray] Clash.Mini tray menu onReady")
	stx.SetIcon(icon.DateN)
	mainTitle = stringUtils.GetMenuItemFullTitle(app.Name, "v" + app.Version)
	mainTooltip = app.Name + " by Maze"
	stx.SetLeftClickFunc(stx.ClickFunc(controller.Dashboard))
	stx.SetTitle(mainTitle)
	stx.SetTooltip(mainTooltip)

	//updater.CheckUpdate()
	// 初始化托盘菜单
	initTrayMenu()
}
