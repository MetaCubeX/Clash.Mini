package start

import (
	_ "github.com/Clash-Mini/Clash.Mini/app/bridge/start/fourth"
	"github.com/Dreamacro/clash/hub/executor"

	"github.com/Clash-Mini/Clash.Mini/log"
	_ "github.com/Clash-Mini/Clash.Mini/tray"
)

func init() {
	log.Infoln("[bridge] Step Last: RawConfig Checking...")
	err := parseRawConfig()
	if err != nil {
		log.Errorln("[config] Start initial RawConfig failed", err)
	}
	log.Infoln("[bridge] Started")
}

func parseRawConfig() error {
	_, err := executor.Parse()
	return err
}
