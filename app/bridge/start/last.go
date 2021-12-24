package start

import (
	_ "github.com/Clash-Mini/Clash.Mini/app/bridge/start/fourth"
	"github.com/Clash-Mini/Clash.Mini/config"
	. "github.com/Clash-Mini/Clash.Mini/constant"
	"github.com/Clash-Mini/Clash.Mini/controller"
	"github.com/Dreamacro/clash/hub/executor"
	path "path/filepath"

	"github.com/Clash-Mini/Clash.Mini/log"
	_ "github.com/Clash-Mini/Clash.Mini/tray"
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
