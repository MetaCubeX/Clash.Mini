package tray

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Clash-Mini/Clash.Mini/util"
	"github.com/Dreamacro/clash/config"
	"github.com/Dreamacro/clash/tunnel"
	"net/http"
	"os"
	path "path/filepath"
	"time"

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

var (
	_, ControllerPort = controller.CheckConfig()

	NeedLoadSelector = false
)

func LoadSelector(mGroup *stx.MenuItemEx) {
	if NeedLoadSelector && config.GroupsList.Len() > 0 {
		groupNowMap := tunnel.Proxies()
		SelectorMap = make(map[string]SelectorInfo)
		util.JsonUnmarshal(util.IgnoreErrorBytes(json.Marshal(groupNowMap)), &SelectorMap)
		for name, group := range SelectorMap {
			if group.Now != "" {
				proxyNow := SelectorMap[group.Now]
				log.Infoln("load: %s -> %s", name, group.Name)
				SwitchGroupAndProxy(mGroup, group.Name, proxyNow.Name)
				continue
			}
		}
		NeedLoadSelector = false
	}
}

func mConfigProxyFunc(mConfigProxy *stx.MenuItemEx) {
	configGroup := ConfigGroupsMap[mConfigProxy.Parent.GetId()]
	GroupPath := mConfigProxy.Parent.GetTitle()
	ProxyName := configGroup[mConfigProxy.GetId()]

	url := fmt.Sprintf(`http://%s:%s/proxies/%s`, constant.Localhost, ControllerPort, GroupPath)
	body := make(map[string]interface{})
	body["name"] = ProxyName
	bytesData, err := json.Marshal(body)
	if err != nil {
		log.Errorln("putConfig Marshal error: %v", err)
		return
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest(http.MethodPut, url, reader)
	if err != nil {
		log.Errorln("putConfig NewRequest error: %v", err)
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Errorln("putConfig Do error: %v", err)
		return
	}
	if resp != nil && resp.StatusCode != http.StatusNoContent {
		log.Errorln("putConfig Do error[HTTP %d]: %s", resp.StatusCode, url)
		return
	}
	if err := resp.Body.Close(); err != nil {
		return
	}
	log.Infoln("PUT Proxies info:  Group: %s - Proxy: %s", GroupPath, ProxyName)

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
	firstInit = true
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
	firstInit = true
}

func mOtherTaskFunc(mOtherTask *stx.MenuItemEx) {
	if mOtherTask.Checked() {
		controller.TaskCommand(task.OFF)
		if !controller.RegCompare(cmd.Task) {
			notify.DoTrayMenuDelay(startup.OFF, constant.NotifyDelay)
		}
	} else {
		controller.TaskCommand(task.ON)
		defer os.Remove(path.Join(".", "task.xml"))
		if controller.RegCompare(cmd.Task) {
			notify.DoTrayMenuDelay(startup.ON, constant.NotifyDelay)
		}
		time.Sleep(2 * time.Second)
	}
	firstInit = true
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
	firstInit = true
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
	firstInit = true
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
	firstInit = true
}
