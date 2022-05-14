package controller

import (
	"fmt"
	"github.com/MetaCubeX/Clash.Mini/common"
	"github.com/MetaCubeX/Clash.Mini/constant"
	"github.com/MetaCubeX/Clash.Mini/log"
	"github.com/jchv/go-webview2"
	"github.com/skratchdot/open-golang/open"
	"mime"
	"os"
	"sync"
)

const (
	dashboardLogHeader = logHeader + ".addConfig"

	localUIPattern = `http://%s:%s/?hostname=%s&port=%s&secret=%s`
)

var (
	dashboardLocker = new(sync.Mutex)
	dashboardUI     webview2.WebView
	localUIUrl      string
	IsOpen          bool
)

func Dashboard() {

	_ = mime.AddExtensionType(".js", "application/javascript")

	if common.DisabledDashboard || IsOpen {
		return
	}

	IsOpen = true

	defer func() {
		dashboardLocker.Unlock()
		IsOpen = false
	}()

	dashboardLocker.Lock()

	secret := constant.ControllerSecret
	localUIUrl = fmt.Sprintf(localUIPattern, constant.LocalHost, constant.DashboardPort,
		constant.ControllerHost, constant.ControllerPort, secret)
	RefreshWindowResolution()
	pageWidth, pageHeight := CalcDpiScaledSize(800, 580)

	dashboardUI = webview2.NewWithOptions(webview2.WebViewOptions{
		Debug:     true,
		AutoFocus: true,
		WindowOptions: webview2.WindowOptions{
			Title:  "Dashboard",
			Center: true,
			IconId: 2,
			Height: uint(pageHeight),
			Width:  uint(pageWidth),
		},
	})

	defer func(ui webview2.WebView) {
		ui.Destroy()
	}(dashboardUI)

	if dashboardUI == nil {
		log.Warnln("[%s] create dashboard failed, it will call system browser", dashboardLogHeader)
		err := open.Run(localUIUrl)
		if err != nil {
			log.Warnln("[%s] call dashboard on system browser failed %v", dashboardLogHeader, err)
		}
		return
	}

	SendMessage(dashboardUI.Window(), 0x0080, 1, ExtractIcon(os.Args[0], 0))
	dashboardUI.Navigate(localUIUrl)
	dashboardUI.Run()
}

func CloseDashboard() error {
	dashboardLocker.Unlock()
	dashboardUI.Destroy()
	return nil
}
