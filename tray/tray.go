package tray

import (
	C "github.com/Dreamacro/clash/constant"
	"github.com/MetaCubeX/Clash.Mini/app"
	"github.com/MetaCubeX/Clash.Mini/constant"
	"github.com/MetaCubeX/Clash.Mini/controller"
	"github.com/MetaCubeX/Clash.Mini/icon"
	"github.com/MetaCubeX/Clash.Mini/log"
	commonUtils "github.com/MetaCubeX/Clash.Mini/util/common"
	stringUtils "github.com/MetaCubeX/Clash.Mini/util/string"
	stx "github.com/getlantern/systray"
)

const (
	logHeader = "tray"
)

var (
	mainTitle   string
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
	mainTitle = stringUtils.GetMenuItemFullTitle(app.Name, app.Version)
	mainTooltip = app.Name + " by Maze"
	stx.SetLeftDoubleClickFunc(stx.ClickFunc(controller.Dashboard))
	stx.SetTitle(mainTitle)
	stx.SetTooltip(mainTooltip)

	//updater.CheckUpdate()
	// 初始化托盘菜单
	initTrayMenu()
}
