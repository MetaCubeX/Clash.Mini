package start

import (
	"github.com/Dreamacro/clash/hub/executor"
	_ "github.com/MetaCubeX/Clash.Mini/app/bridge/start/fourth"
	"github.com/MetaCubeX/Clash.Mini/config"
	. "github.com/MetaCubeX/Clash.Mini/constant"
	"github.com/MetaCubeX/Clash.Mini/controller"
	path "path/filepath"

	"github.com/MetaCubeX/Clash.Mini/log"
	_ "github.com/MetaCubeX/Clash.Mini/tray"
)

func init() {
	//log.Infoln("[bridge] Step Last: RawConfig Checking...")
	//parseRawConfig()
	log.Infoln("[bridge] Started")
}

func parseRawConfig() {
	Name := config.GetProfile()
	exist, configName := controller.CheckConfig(Name)
	if exist {
		_, err := executor.ParseWithPath(path.Join(ProfileDir, configName))
		if err != nil {
			log.Errorln("[config] Start initial RawConfig failed", err)
		}
	}
}
