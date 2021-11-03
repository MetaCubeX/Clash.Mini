package tray

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	"github.com/Clash-Mini/Clash.Mini/config"
	"github.com/Clash-Mini/Clash.Mini/constant"
	"github.com/Clash-Mini/Clash.Mini/controller"
	"github.com/Clash-Mini/Clash.Mini/icon"
	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Clash-Mini/Clash.Mini/notify"
	"github.com/Clash-Mini/Clash.Mini/proxy"
	"github.com/Clash-Mini/Clash.Mini/sysproxy"
	"github.com/Clash-Mini/Clash.Mini/util"

	clashConfig "github.com/Dreamacro/clash/config"
	cP "github.com/Dreamacro/clash/proxy"
	"github.com/Dreamacro/clash/tunnel"
	"github.com/MakeNowJust/hotkey"
	stx "github.com/getlantern/systray"
)

var (
	_, ControllerPort = controller.CheckConfig()
	NeedLoadSelector  = false
)

func LoadSelector(mGroup *stx.MenuItemEx) {
	if NeedLoadSelector && clashConfig.GroupsList.Len() > 0 {
		groupNowMap := tunnel.Proxies()
		SelectorMap = make(map[string]proxy.SelectorInfo)
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
		// 取消系统代理
		err := sysproxy.ClearSystemProxy()
		if err != nil {
			log.Errorln("cancel sysproxy failed: %v", err)
		} else {
			mEnabled.Uncheck()
			stx.SetIcon(icon.DateN)
			notify.DoTrayMenu(sys.OFF)
		}
	} else {
		// 设置系统代理
		var port int
		if cP.GetPorts().MixedPort != 0 {
			port = cP.GetPorts().MixedPort
		} else {
			port = cP.GetPorts().Port
		}
		err := sysproxy.SetSystemProxy(
			&sysproxy.ProxyConfig{
				Enable: true,
				Server: fmt.Sprintf("%s:%d", constant.Localhost, port),
			})
		if err != nil {
			log.Errorln("setting sysproxy failed: %v", err)
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
		config.SetCmd(sys.OFF)
		if !config.IsCmdPositive(cmd.Sys) {
			notify.DoTrayMenuDelay(auto.OFF, constant.NotifyDelay)
		}
	} else {
		config.SetCmd(sys.ON)
		if config.IsCmdPositive(cmd.Sys) {
			notify.DoTrayMenuDelay(auto.ON, constant.NotifyDelay)
		}
	}
	firstInit = true
}

func mOtherTaskFunc(mOtherTask *stx.MenuItemEx) {
	if mOtherTask.Checked() {
		controller.TaskCommand(task.OFF)
		if !config.IsCmdPositive(cmd.Task) {
			notify.DoTrayMenuDelay(startup.OFF, constant.NotifyDelay)
		}
	} else {
		controller.TaskCommand(task.ON)
		defer os.Remove(path.Join(".", "task.xml"))
		if config.IsCmdPositive(cmd.Task) {
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
		if !config.IsCmdPositive(cmd.MMDB) {
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
		if config.IsCmdPositive(cmd.MMDB) {
			notify.DoTrayMenuDelay(mmdb.Lite, constant.NotifyDelay)
		}
	}
	firstInit = true
}

func mOtherUpdateCronFunc(mOtherUpdateCron *stx.MenuItemEx) {
	if mOtherUpdateCron.Checked() {
		config.SetCmd(cron.OFF)
		if !config.IsCmdPositive(cmd.Cron) {
			notify.DoTrayMenuDelay(cron.OFF, constant.NotifyDelay)
		}
	} else {
		config.SetCmd(cron.ON)
		if !config.IsCmdPositive(cmd.Cron) {
			notify.DoTrayMenuDelay(cron.ON, constant.NotifyDelay)
		}
	}
	firstInit = true
}

func hotKey(mEnabled *stx.MenuItemEx) {
	message := ""
	hkey := hotkey.New()
	_, err1 := hkey.Register(hotkey.Alt, 'R', func() {
		tunnel.SetMode(tunnel.Rule)
	})
	if err1 != nil {
		message += "Alt+R热键注册失败\n"
	}
	_, err2 := hkey.Register(hotkey.Alt, 'G', func() {
		tunnel.SetMode(tunnel.Global)
	})
	if err2 != nil {
		message += "Alt+G热键注册失败\n"
	}
	_, err3 := hkey.Register(hotkey.Alt, 'D', func() {
		tunnel.SetMode(tunnel.Direct)
	})
	if err3 != nil {
		message += "Alt+D热键注册失败\n"
	}
	_, err4 := hkey.Register(hotkey.Alt, 'S', func() {
		mEnabledFunc(mEnabled)
	})
	if err4 != nil {
		message += "Alt+S热键注册失败\n"
	}
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		go notify.PushWithLine("📢通知📢", message)
	}
}
