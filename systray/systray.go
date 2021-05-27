package systray

import (
	"fmt"
	"github.com/Clash-Mini/Clash.Mini/controller"
	"github.com/Clash-Mini/Clash.Mini/icon"
	"github.com/Clash-Mini/Clash.Mini/notify"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"github.com/Clash-Mini/Clash.Mini/sysproxy"
	C "github.com/Dreamacro/clash/constant"
	"github.com/Dreamacro/clash/proxy"
	"github.com/Dreamacro/clash/tunnel"
	"github.com/getlantern/systray"
)

func init() {
	if runtime.GOOS == "windows" {
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
	systray.SetTitle("Clash.Mini")
	systray.SetTooltip("Clash.Mini by Maze")

	mTitle := systray.AddMenuItem("Clash.Mini", "")
	systray.AddSeparator()

	mGlobal := systray.AddMenuItem("全局代理", "Set as Global")
	mRule := systray.AddMenuItem("规则代理", "Set as Rule")
	mDirect := systray.AddMenuItem("全局直连", "Set as Direct")
	systray.AddSeparator()

	mEnabled := systray.AddMenuItem("系统代理", "")
	mURL := systray.AddMenuItem("控制面板", "")
	mConfig := systray.AddMenuItem("配置管理", "")
	mOther := systray.AddMenuItem("其他设置", "")
	mOtherTask := mOther.AddSubMenuItem("设置开机启动", "")
	mOtherAutosys := mOther.AddSubMenuItem("设置系统代理", "")
	mOtherMMBD := mOther.AddSubMenuItem("设置GeoIP2数据库", "")
	MaxMindMMBD := mOtherMMBD.AddSubMenuItem("MaxMind数据库", "")
	Hackl0usMMBD := mOtherMMBD.AddSubMenuItem("Hackl0us数据库", "")

	systray.AddSeparator()

	mQuit := systray.AddMenuItem("退出", "Quit Clash.Mini")

	if runtime.GOOS != "windows" {
		mEnabled.Hide()
		mOther.Hide()
		mConfig.Hide()
	}

	go func() {
		t := time.NewTicker(time.Second)
		defer t.Stop()
		SavedPort := proxy.GetPorts().Port
		if controller.RegCompare("Sys") == true {
			var Ports int
			if proxy.GetPorts().MixedPort != 0 {
				Ports = proxy.GetPorts().MixedPort
			} else {
				Ports = proxy.GetPorts().Port
			}
			sysproxy.SetSystemProxy(
				&sysproxy.ProxyConfig{
					Enable: true,
					Server: "127.0.0.1:" + strconv.Itoa(Ports),
				})
			mEnabled.Check()
			notify.Notify("SysON")
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
						notify.Notify("Global")
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
						notify.Notify("Rule")
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
						notify.Notify("Direct")
					} else {
						systray.SetIcon(icon.DateN)
					}
				}
			}
			if controller.RegCompare("Task") == true {
				mOtherTask.Check()
			} else {
				mOtherTask.Uncheck()
			}

			if controller.RegCompare("MMBD") == true {
				MaxMindMMBD.Uncheck()
				Hackl0usMMBD.Check()
			} else {
				MaxMindMMBD.Check()
				Hackl0usMMBD.Uncheck()
			}

			if controller.RegCompare("Sys") == true {
				mOtherAutosys.Check()
			} else {
				mOtherAutosys.Uncheck()
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
							Server: "127.0.0.1:" + strconv.Itoa(SavedPort),
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

			if p.Enable && p.Server == "127.0.0.1:"+strconv.Itoa(SavedPort) {
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
			UnUsedINFO, TotalINFO, ExpireINFO := controller.UserINFO()
			time.Sleep(2 * time.Second)
			if UnUsedINFO != "" {
				notify.NotifyINFO(UnUsedINFO, TotalINFO, ExpireINFO)
			}
		}()
		for {
			select {
			case <-mTitle.ClickedCh:
				fmt.Println("Title Clicked")
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
						notify.Notify("SysOFF")
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
							Server: "127.0.0.1:" + strconv.Itoa(Ports),
						})
					if err != nil {
					} else {
						mEnabled.Check()
						systray.SetIcon(icon.DateS)
						notify.Notify("SysON")
					}
				}
			case <-mURL.ClickedCh:
				go controller.Dashboard()
			case <-mConfig.ClickedCh:
				go controller.MenuConfig()
			case <-mOtherAutosys.ClickedCh:
				if mOtherAutosys.Checked() {
					controller.Regcmd("Sys", "OFF")
					time.Sleep(2 * time.Second)
					if controller.RegCompare("Sys") == false {
						notify.Notify("SysAutoOFF")
					}
				} else {
					controller.Regcmd("Sys", "ON")
					time.Sleep(2 * time.Second)
					if controller.RegCompare("Sys") == true {
						notify.Notify("SysAutoON")
					}
				}
			case <-mOtherTask.ClickedCh:
				if mOtherTask.Checked() {
					controller.TaskCommand("delete")
					time.Sleep(2 * time.Second)
					if controller.RegCompare("Task") == false {
						notify.Notify("StartupOff")
					}
				} else {
					controller.TaskCommand("create")
					time.Sleep(2 * time.Second)
					taskFile := filepath.Join(".", "task.xml")
					taskPath, _ := os.Getwd()
					Filepath := filepath.Join(taskPath, taskFile)
					os.Remove(Filepath)
					if controller.RegCompare("Task") == true {
						notify.Notify("Startup")
					}
				}
			case <-MaxMindMMBD.ClickedCh:
				if MaxMindMMBD.Checked() {
					return
				} else {
					controller.GetMMDB("Max")
					if controller.RegCompare("MMBD") == false {
						time.Sleep(2 * time.Second)
						notify.Notify("Max")
					}
				}
			case <-Hackl0usMMBD.ClickedCh:
				if Hackl0usMMBD.Checked() {
					return
				} else {
					controller.GetMMDB("Lite")
					if controller.RegCompare("MMBD") == true {
						time.Sleep(2 * time.Second)
						notify.Notify("Lite")
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
