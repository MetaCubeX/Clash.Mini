package controller

import (
	"fmt"
	"github.com/Clash-Mini/Clash.Mini/common"
	"github.com/Clash-Mini/Clash.Mini/constant"
	cConfig "github.com/Clash-Mini/Clash.Mini/constant/config"
	"github.com/Clash-Mini/Clash.Mini/log"
	C "github.com/Dreamacro/clash/constant"
	"github.com/skratchdot/open-golang/open"
	"mime"
	"strings"
	"sync"

	"github.com/zserge/lorca"
)

const (
	dashboardLogHeader = logHeader + ".addConfig"

	localUIPattern = `http://%s:%s/?hostname=%s&port=%s&secret=%s`
)

var (
	dashboardLocker = new(sync.Mutex)
	dashboardUI     lorca.UI
	localUIUrl      string
)

func Dashboard() {

	_ = mime.AddExtensionType(".js", "application/javascript")
	if common.DisabledDashboard {
		return
	}
	defer func() {
		dashboardLocker.Unlock()
	}()
	dashboardLocker.Lock()
	RawConfig, _ := GetConfig(C.Path.Config())
	externalController := RawConfig.General.ExternalController
	secret := RawConfig.General.Secret

	host := strings.Split(externalController, ":")
	if len(host) == 1 {
		localUIUrl = fmt.Sprintf(localUIPattern, constant.Localhost, constant.DashboardPort,
			constant.Localhost, host[0], secret)
	} else {
		if host[0] == "" {
			host[0] = constant.Localhost
		}
		localUIUrl = fmt.Sprintf(localUIPattern, constant.Localhost, constant.DashboardPort,
			host[0], host[1], secret)
	}

	pageWidth := 800
	pageHeight := 580
	RefreshWindowResolution()
	pageInit := lorca.Bounds{
		Left:        int(CalcDpiCenterScaledSize(xScreen, int32(pageWidth))),
		Top:         int(CalcDpiCenterScaledSize(yScreen, int32(pageHeight)) + GetTaskbarHeight()),
		Width:       pageWidth,
		Height:      pageHeight,
		WindowState: "normal",
	}

	var err error
	dashboardUI, err = lorca.New("", cConfig.DashboardDir, 0, 0,
		fmt.Sprintf("--window-position=-%d,-%d", xScreen, yScreen))
	if err != nil {
		log.Errorln("[%s] create dashboard failed, it will call system browser: %v", dashboardLogHeader, err)
		err := open.Run(localUIUrl)
		if err != nil {
			log.Errorln("[%s] call dashboard on system browser failed %v", dashboardLogHeader, err)
		}
		return
	}
	err = dashboardUI.Load(localUIUrl)
	if err != nil {
		return
	}
	err = dashboardUI.SetBounds(pageInit)
	if err != nil {
		log.Errorln("[%s] SetBounds dashboard failed %v", dashboardLogHeader, err)
		return
	}
	defer func(ui lorca.UI) {
		err := ui.Close()
		if err != nil {
			log.Errorln("[%s] close dashboard failed %v", dashboardLogHeader, err)
		}
	}(dashboardUI)
	// Wait until UI window is closed
	select {
	case <-dashboardUI.Done():
	}
}

func CloseDashboard() error {
	if dashboardUI != nil {
		return dashboardUI.Close()
	}
	return nil
}
