package controller

import (
	"fmt"
	"time"

	"github.com/lxn/win"
	"github.com/skratchdot/open-golang/open"
	"github.com/zserge/lorca"
)

const (
	localUIPattern = `%s:8070/?hostname=127.0.0.1&port=%s&secret=`
)

func Dashboard() {
	_, controllerPort := checkConfig()
	xScreen := int(win.GetSystemMetrics(win.SM_CXSCREEN))
	yScreen := int(win.GetSystemMetrics(win.SM_CYSCREEN))
	pageWidth := 800
	pageHeight := 580
	PageInit := lorca.Bounds{
		Left:        (xScreen - pageWidth) / 2,
		Top:         (yScreen - pageHeight) / 2,
		Width:       pageWidth,
		Height:      pageHeight,
		WindowState: "normal",
	}
	localUIUrl := fmt.Sprintf(localUIPattern, localHost, controllerPort)
	ui, err := lorca.New(localUIUrl, "", 0, 0)
	if err != nil {
		open.Run(localUIUrl)
	} else {
		defer ui.Close()
		ui.SetBounds(PageInit)
		// Wait until UI window is closed
		<-ui.Done()
		time.Sleep(time.Duration(2) * time.Second)

	}
}
