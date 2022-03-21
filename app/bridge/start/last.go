package start

import (
	_ "github.com/MetaCubeX/Clash.Mini/app/bridge/start/fourth"
	"github.com/MetaCubeX/Clash.Mini/log"
	_ "github.com/MetaCubeX/Clash.Mini/tray"
)

func init() {
	log.Infoln("[bridge] Started")
}
