package systray

import (
	"fmt"
	"github.com/Clash-Mini/Clash.Mini/controller"
	"github.com/Clash-Mini/Clash.Mini/icon"
	"github.com/Clash-Mini/Clash.Mini/notify"
	"os"
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
	systray.SetIcon(icon.Date)
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
	mOtherStartup := mOther.AddSubMenuItem("设置开机启动(UAC)", "")

	systray.AddSeparator()

	mQuit := systray.AddMenuItem("退出", "Quit Clash.Mini")

	go func() {
		t := time.NewTicker(time.Second)
		defer t.Stop()
		SavedPort := proxy.GetPorts().Port

		for {
			<-t.C
			switch tunnel.Mode() {
			case tunnel.Global:
				if mGlobal.Checked() {
				} else {
					mGlobal.Check()
					mRule.Uncheck()
					mDirect.Uncheck()
					systray.SetIcon(icon.Date2)
					notify.GlobalNotify()
				}
			case tunnel.Rule:
				if mRule.Checked() {
				} else {
					mGlobal.Uncheck()
					mRule.Check()
					mDirect.Uncheck()
					systray.SetIcon(icon.Date)
					notify.RuleNotify()
				}
			case tunnel.Direct:
				if mDirect.Checked() {
				} else {
					mGlobal.Uncheck()
					mRule.Uncheck()
					mDirect.Check()
					systray.SetIcon(icon.Date)
					notify.DirectNotify()
				}
			}

			if controller.RegCompare() == true {
				mOtherStartup.Check()
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
						notify.SysNotify()
					}
				}
			case <-mURL.ClickedCh:
				go controller.Dashboard()
			case <-mConfig.ClickedCh:
				go controller.MenuConfig()
			case <-mOtherStartup.ClickedCh:
				if mOtherStartup.Checked() {
					go controller.Command("delete")
					mOtherStartup.Uncheck()
				} else {
					go controller.Command("add")
					mOtherStartup.Check()
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
