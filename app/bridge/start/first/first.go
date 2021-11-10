package first

import (
	"github.com/Clash-Mini/Clash.Mini/app/bridge/mq"
	_ "github.com/Clash-Mini/Clash.Mini/app/bridge/start/enter"

	_ "github.com/Clash-Mini/Clash.Mini/config"
	"github.com/Clash-Mini/Clash.Mini/log"
)

func init() {
	log.Infoln("[bridge] first")

	mq.PrintMsg(log.Infoln)
}
