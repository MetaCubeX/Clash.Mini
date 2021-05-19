package controller

import (
	"github.com/lxn/win"
	"github.com/skratchdot/open-golang/open"
	"github.com/zserge/lorca"
)

func Dashboard() {
	_, controller := checkConfig()
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
	ui, err := lorca.New("http://127.0.0.1:8070/?hostname=127.0.0.1&port="+controller+"&secret=", "", 0, 0)
	if err != nil {
		open.Run("http://127.0.0.1:8070/?hostname=127.0.0.1&port=" + controller + "&secret=")
	} else {
		ui.SetBounds(PageInit)
		defer ui.Close()
		// Wait until UI window is closed
		<-ui.Done()
	}
}
