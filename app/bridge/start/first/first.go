package first

import (
	"fmt"

	_ "github.com/Clash-Mini/Clash.Mini/app/bridge/start/enter"

	"github.com/Clash-Mini/Clash.Mini/app"
	_ "github.com/Clash-Mini/Clash.Mini/config"
	"github.com/Clash-Mini/Clash.Mini/log"
)

func init() {
	fmt.Println("[bridge] first")

	app.PrintMsg(log.Infoln)
}
