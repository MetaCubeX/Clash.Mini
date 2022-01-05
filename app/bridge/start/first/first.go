package first

import (
	"github.com/MetaCubeX/Clash.Mini/app/bridge/mq"
	_ "github.com/MetaCubeX/Clash.Mini/app/bridge/start/enter"
	_ "github.com/MetaCubeX/Clash.Mini/config"
	"github.com/MetaCubeX/Clash.Mini/log"
)

func init() {
	mq.PrintMsg(log.Infoln)
}
