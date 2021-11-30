package start

import (
	_ "github.com/Clash-Mini/Clash.Mini/app/bridge/start/fourth"

	"github.com/Clash-Mini/Clash.Mini/log"
	_ "github.com/Clash-Mini/Clash.Mini/tray"
)

func init() {
	log.Infoln("[bridge] Started")
}
