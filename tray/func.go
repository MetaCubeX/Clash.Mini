package tray

import (
	"fmt"

	"github.com/Clash-Mini/Clash.Mini/cmd"
	"github.com/Clash-Mini/Clash.Mini/cmd/auto"
	"github.com/Clash-Mini/Clash.Mini/cmd/cron"
	"github.com/Clash-Mini/Clash.Mini/cmd/mmdb"
	"github.com/Clash-Mini/Clash.Mini/cmd/startup"
	"github.com/Clash-Mini/Clash.Mini/cmd/sys"
	"github.com/Clash-Mini/Clash.Mini/cmd/task"
	"github.com/Clash-Mini/Clash.Mini/constant"
	"github.com/Clash-Mini/Clash.Mini/controller"
	"github.com/Clash-Mini/Clash.Mini/icon"
	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Clash-Mini/Clash.Mini/notify"
	"github.com/Clash-Mini/Clash.Mini/sysproxy"
	"github.com/Dreamacro/clash/proxy"
	stx "github.com/getlantern/systray"
)

//func mProxyFunc(mEnabled *stx.MenuItemEx, p cp.Type) {
//	if mEnabled.Checked() {
//		err := sysproxy.SetSystemProxy(sysproxy.GetSavedProxy())
//		if err != nil {
//		} else {
//		}
//	}
//}

func mConfigProxyFunc(mConfigProxy *stx.MenuItemEx) {
	log.Infoln(mConfigProxy.GetTitle())
	// get proxy info
	configGroup := ConfigGroupsMap[mConfigProxy.Parent.GetId()]
	configProxy := configGroup[mConfigProxy.GetId()]
	log.Infoln("got proxy info: %s", configProxy)
}

func mEnabledFunc(mEnabled *stx.MenuItemEx) {
	if mEnabled.Checked() {
		err := sysproxy.SetSystemProxy(sysproxy.GetSavedProxy())
		if err != nil {
		} else {
			mEnabled.Uncheck()
			stx.SetIcon(icon.DateN)
			notify.DoTrayMenu(sys.OFF)
		}
	} else {
		var Ports int
		if proxy.GetPorts().MixedPort != 0 {
			Ports = proxy.GetPorts().MixedPort
		} else {
			Ports = proxy.GetPorts().Port
		}
		err := sysproxy.SetSystemProxy(
			&sysproxy.ProxyConfig{
				Enable: true,
				Server: fmt.Sprintf("%s:%d", constant.Localhost, Ports),
			})
		if err != nil {
		} else {
			mEnabled.Check()
			stx.SetIcon(icon.DateS)
			notify.DoTrayMenu(sys.ON)
		}
	}
}

func mOtherAutosysFunc(mOtherAutosys *stx.MenuItemEx) {
	if mOtherAutosys.Checked() {
		controller.RegCmd(sys.OFF)
		if !controller.RegCompare(cmd.Sys) {
			notify.DoTrayMenuDelay(auto.OFF, constant.NotifyDelay)
		}
	} else {
		controller.RegCmd(sys.ON)
		if controller.RegCompare(cmd.Sys) {
			notify.DoTrayMenuDelay(auto.ON, constant.NotifyDelay)
		}
	}
}

func mOtherTaskFunc(mOtherTask *stx.MenuItemEx) {
	if mOtherTask.Checked() {
		controller.TaskCommand(task.OFF)
		if !controller.RegCompare(cmd.Task) {
			notify.DoTrayMenuDelay(startup.OFF, constant.NotifyDelay)
		}
	} else {
		controller.TaskCommand(task.ON)
		if controller.RegCompare(cmd.Task) {
			notify.DoTrayMenuDelay(startup.ON, constant.NotifyDelay)
		}
	}
}

func maxMindMMBDFunc(maxMindMMBD *stx.MenuItemEx) {
	if maxMindMMBD.Checked() {
		return
	} else {
		controller.GetMMDB(mmdb.Max)
		if !controller.RegCompare(cmd.MMDB) {
			notify.DoTrayMenuDelay(mmdb.Max, constant.NotifyDelay)
		}
	}
}

func hackl0usMMDBFunc(hackl0usMMDB *stx.MenuItemEx) {
	if hackl0usMMDB.Checked() {
		return
	} else {
		controller.GetMMDB(mmdb.Lite)
		if controller.RegCompare(cmd.MMDB) {
			notify.DoTrayMenuDelay(mmdb.Lite, constant.NotifyDelay)
		}
	}
}

func mOtherUpdateCronFunc(mOtherUpdateCron *stx.MenuItemEx) {
	if mOtherUpdateCron.Checked() {
		controller.RegCmd(cron.OFF)
		if !controller.RegCompare(cmd.Cron) {
			notify.DoTrayMenuDelay(cron.OFF, constant.NotifyDelay)
		}
	} else {
		controller.RegCmd(cron.ON)
		if !controller.RegCompare(cmd.Cron) {
			notify.DoTrayMenuDelay(cron.ON, constant.NotifyDelay)
		}
	}
}
