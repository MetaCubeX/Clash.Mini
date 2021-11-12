package controller

import (
	"fmt"
	"sync"

	"github.com/Clash-Mini/Clash.Mini/common"
	"github.com/Clash-Mini/Clash.Mini/constant"
	"github.com/Clash-Mini/Clash.Mini/log"

	"github.com/skratchdot/open-golang/open"
	"github.com/zserge/lorca"
)

const (
	dashboardLogHeader = logHeader + ".addConfig"

	localUIPattern = `http://%s:%s/?hostname=%s&port=%s&secret=`
)

var (
	dashboardLocker = new(sync.Mutex)
	dashboardUI     lorca.UI
)

func Dashboard() {
	if common.DisabledDashboard {
		return
	}
	defer func() {
		dashboardLocker.Unlock()
	}()
	dashboardLocker.Lock()
	_, controllerPort := CheckConfig()

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
	localUIUrl := fmt.Sprintf(localUIPattern, constant.Localhost, constant.DashboardPort,
		constant.Localhost, controllerPort)
	var err error
	dashboardUI, err = lorca.New("", "", 0, 0,
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
