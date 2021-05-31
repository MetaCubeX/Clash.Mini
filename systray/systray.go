package systray

import (
	"encoding/json"
	"fmt"
	"github.com/Dreamacro/clash/config"
	"os"
	path "path/filepath"
	"runtime"
	"time"

	"github.com/Clash-Mini/Clash.Mini/cmd"
	"github.com/Clash-Mini/Clash.Mini/cmd/auto"
	"github.com/Clash-Mini/Clash.Mini/cmd/cron"
	"github.com/Clash-Mini/Clash.Mini/cmd/mmdb"
	cp "github.com/Clash-Mini/Clash.Mini/cmd/proxy"
	"github.com/Clash-Mini/Clash.Mini/cmd/startup"
	"github.com/Clash-Mini/Clash.Mini/cmd/sys"
	"github.com/Clash-Mini/Clash.Mini/cmd/task"
	"github.com/Clash-Mini/Clash.Mini/constant"
	"github.com/Clash-Mini/Clash.Mini/controller"
	"github.com/Clash-Mini/Clash.Mini/icon"
	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Clash-Mini/Clash.Mini/notify"
	"github.com/Clash-Mini/Clash.Mini/sysproxy"
	"github.com/Clash-Mini/Clash.Mini/util"

	C "github.com/Dreamacro/clash/constant"
	"github.com/Dreamacro/clash/proxy"
	"github.com/Dreamacro/clash/tunnel"
	"github.com/getlantern/systray"
)

func init() {
	if constant.IsWindows() {
		currentDir, _ := os.Getwd()
		C.SetHomeDir(currentDir)
	}
	go func() {
		runtime.LockOSThread()
		systray.Run(onReady, onExit)
		runtime.UnlockOSThread()
	}()
}

func onReady() {
	systray.SetIcon(icon.DateN)
	systray.SetTitle(util.AppTitle)
	systray.SetTooltip(util.AppTitle + " by Maze")

	mTitle := systray.AddMenuItem(util.AppTitle, "")
	systray.AddSeparator()

	mGlobal := systray.AddMenuItem("全局代理", "Set as Global")
	mRule := systray.AddMenuItem("规则代理", "Set as Rule")
	mDirect := systray.AddMenuItem("全局直连", "Set as Direct")
	systray.AddSeparator()

	type GroupsList struct {
		Name    string   `json:"name"`
		Proxies []string `json:"proxies"`
	}

	mGroup := systray.AddMenuItem("切换节点", "Proxies Control")
	data := config.GroupsList
	for _, group := range data {
		jsonString, _ := json.Marshal(group)
		s := GroupsList{}
		json.Unmarshal(jsonString, &s)
		groups := mGroup.AddSubMenuItem(s.Name, s.Name)
		for _, Proxy := range s.Proxies {
			_ = groups.AddSubMenuItem(Proxy, Proxy)
		}
	}

	systray.AddSeparator()
	mEnabled := systray.AddMenuItem("系统代理", "")
	mURL := systray.AddMenuItem("控制面板", "")
	mConfig := systray.AddMenuItem("配置管理", "")
	mOther := systray.AddMenuItem("其他设置", "")
	mOtherTask := mOther.AddSubMenuItem("设置开机启动", "")
	mOtherAutosys := mOther.AddSubMenuItem("设置默认代理", "")
	mOtherUpdateCron := mOther.AddSubMenuItem("设置定时更新", "")
	mOtherMMBD := mOther.AddSubMenuItem("设置GeoIP2数据库", "")
	MaxMindMMBD := mOtherMMBD.AddSubMenuItem("MaxMind数据库", "")
	Hackl0usMMBD := mOtherMMBD.AddSubMenuItem("Hackl0us数据库", "")

	systray.AddSeparator()

	mQuit := systray.AddMenuItem("退出", "Quit Clash.Mini")

	if !constant.IsWindows() {
		mEnabled.Hide()
		mOther.Hide()
		mConfig.Hide()
	}

	go func() {
		t := time.NewTicker(time.Second)
		defer t.Stop()
		SavedPort := proxy.GetPorts().Port
		if controller.RegCompare(cmd.Sys) {
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
				log.Errorln("SetSystemProxy error: %v", err)
				notify.PushWithLine("❌错误❌", "设置系统代理时出错")
				return
			}
			mEnabled.Check()
			notify.DoTrayMenu(sys.ON)
		}
		if controller.RegCompare(cmd.Cron) {
			mOtherUpdateCron.Check()
			go controller.CronTask()
		}
		for {
			<-t.C
			switch tunnel.Mode() {
			case tunnel.Global:
				if mGlobal.Checked() {
				} else {
					mGlobal.Check()
					mRule.Uncheck()
					mDirect.Uncheck()
					if mEnabled.Checked() {
						systray.SetIcon(icon.DateG)
						notify.DoTrayMenu(cp.Global)
					} else {
						systray.SetIcon(icon.DateN)
					}
				}
			case tunnel.Rule:
				if mRule.Checked() {
				} else {
					mGlobal.Uncheck()
					mRule.Check()
					mDirect.Uncheck()
					if mEnabled.Checked() {
						systray.SetIcon(icon.DateS)
						notify.DoTrayMenu(cp.Rule)
					} else {
						systray.SetIcon(icon.DateN)
					}
				}
			case tunnel.Direct:
				if mDirect.Checked() {
				} else {
					mGlobal.Uncheck()
					mRule.Uncheck()
					mDirect.Check()

					if mEnabled.Checked() {
						systray.SetIcon(icon.DateD)
						notify.DoTrayMenu(cp.Direct)
					} else {
						systray.SetIcon(icon.DateN)
					}
				}
			}
			if controller.RegCompare(cmd.Task) {
				mOtherTask.Check()
			} else {
				mOtherTask.Uncheck()
			}

			if controller.RegCompare(cmd.MMDB) {
				MaxMindMMBD.Uncheck()
				Hackl0usMMBD.Check()
			} else {
				MaxMindMMBD.Check()
				Hackl0usMMBD.Uncheck()
			}

			if controller.RegCompare(cmd.Sys) {
				mOtherAutosys.Check()
			} else {
				mOtherAutosys.Uncheck()
			}

			if controller.RegCompare(cmd.Cron) {
				mOtherUpdateCron.Check()
			} else {
				mOtherUpdateCron.Uncheck()
			}

			if mEnabled.Checked() {
				var p int
				if proxy.GetPorts().MixedPort != 0 {
					p = proxy.GetPorts().MixedPort
				} else {
					p = proxy.GetPorts().Port
				}
				if SavedPort != p {
					SavedPort = p
					err := sysproxy.SetSystemProxy(
						&sysproxy.ProxyConfig{
							Enable: true,
							Server: fmt.Sprintf("%s:%d", constant.Localhost, SavedPort),
						})
					if err != nil {
						continue
					}
				}
			}

			p, err := sysproxy.GetCurrentProxy()
			if err != nil {
				continue
			}

			if p.Enable && p.Server == fmt.Sprintf("%s:%d", constant.Localhost, SavedPort) {
				if mEnabled.Checked() {
				} else {
					mEnabled.Check()
				}
			} else {
				if mEnabled.Checked() {
					mEnabled.Uncheck()
				} else {
				}
			}

		}
	}()

	go func() {
		go func() {
			userInfo := controller.UpdateSubscriptionUserInfo()
			time.Sleep(2 * time.Second)
			if len(userInfo.UnusedInfo) > 0 {
				notify.PushFlowInfo(userInfo.UsedInfo, userInfo.UnusedInfo, userInfo.ExpireInfo)
			}
		}()
		for {
			select {
			case <-mTitle.ClickedCh:
				proxies := tunnel.Proxies()
				log.Debugln(util.ToJsonString(proxies))
				log.Debugln("Title Clicked")
			case <-mGlobal.ClickedCh:
				tunnel.SetMode(tunnel.Global)
			case <-mRule.ClickedCh:
				tunnel.SetMode(tunnel.Rule)
			case <-mDirect.ClickedCh:
				tunnel.SetMode(tunnel.Direct)
			case <-mEnabled.ClickedCh:
				if mEnabled.Checked() {
					err := sysproxy.SetSystemProxy(sysproxy.GetSavedProxy())
					if err != nil {
					} else {
						mEnabled.Uncheck()
						systray.SetIcon(icon.DateN)
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
						systray.SetIcon(icon.DateS)
						notify.DoTrayMenu(sys.ON)
					}
				}
			case <-mURL.ClickedCh:
				go controller.Dashboard()
			case <-mConfig.ClickedCh:
				go controller.MenuConfig()
			case <-mOtherAutosys.ClickedCh:
				if mOtherAutosys.Checked() {
					controller.RegCmd(sys.OFF)
					time.Sleep(2 * time.Second)
					if !controller.RegCompare(cmd.Sys) {
						notify.DoTrayMenu(auto.OFF)
					}
				} else {
					controller.RegCmd(sys.ON)
					time.Sleep(2 * time.Second)
					if controller.RegCompare(cmd.Sys) {
						notify.DoTrayMenu(auto.ON)
					}
				}
			case <-mOtherTask.ClickedCh:
				if mOtherTask.Checked() {
					controller.TaskCommand(task.OFF)
					time.Sleep(2 * time.Second)
					if !controller.RegCompare(cmd.Task) {
						notify.DoTrayMenu(startup.OFF)
					}
				} else {
					controller.TaskCommand(task.ON)
					time.Sleep(2 * time.Second)
					os.Remove(path.Join(".", "task.xml"))
					if controller.RegCompare(cmd.Task) {
						notify.DoTrayMenu(startup.ON)
					}
				}
			case <-MaxMindMMBD.ClickedCh:
				if MaxMindMMBD.Checked() {
					return
				} else {
					controller.GetMMDB(mmdb.Max)
					if !controller.RegCompare(cmd.MMDB) {
						time.Sleep(2 * time.Second)
						notify.DoTrayMenu(mmdb.Max)
					}
				}
			case <-Hackl0usMMBD.ClickedCh:
				if Hackl0usMMBD.Checked() {
					return
				} else {
					controller.GetMMDB(mmdb.Lite)
					if controller.RegCompare(cmd.MMDB) {
						time.Sleep(2 * time.Second)
						notify.DoTrayMenu(mmdb.Lite)
					}
				}
			case <-mOtherUpdateCron.ClickedCh:
				if mOtherUpdateCron.Checked() {
					controller.RegCmd(cron.OFF)
					time.Sleep(2 * time.Second)
					if !controller.RegCompare(cmd.Cron) {
						notify.DoTrayMenu(cron.OFF)
					}
				} else {
					controller.RegCmd(cron.ON)
					time.Sleep(2 * time.Second)
					if !controller.RegCompare(cmd.Cron) {
						notify.DoTrayMenu(cron.ON)
					}
				}

			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	for {
		err := sysproxy.SetSystemProxy(sysproxy.GetSavedProxy())
		if err != nil {
			continue
		} else {
			break
		}
	}

	os.Exit(1)
}
