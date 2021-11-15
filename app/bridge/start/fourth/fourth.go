package fourth

import (
	"fmt"
	"os"

	_ "github.com/Clash-Mini/Clash.Mini/app/bridge/start/third"

	"github.com/Clash-Mini/Clash.Mini/common"
	. "github.com/Clash-Mini/Clash.Mini/constant"
	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Clash-Mini/Clash.Mini/notify"
)

func init() {
	log.Infoln("[bridge] Step Fourth: Checking...")

	//common.GetVarFlags()
	//common.InitVariablesAfterGetVarFlags()
	//common.InitFunctionsAfterGetVarFlags()

	if _, err := os.Stat(ProfileDir); err != nil {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(ProfileDir, 0666); err != nil {
				errMsg := fmt.Sprintf("cannot create config dir: %v", err)
				log.Errorln(errMsg)
				notify.PushError("", errMsg)
				common.DisabledCore = true
				return
			}
		}
	}
	if _, err := os.Stat(CacheDir); err != nil {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(CacheDir, 0666); err != nil {
				errMsg := fmt.Sprintf("cannot create cache dir: %v", err)
				log.Errorln(errMsg)
				notify.PushError("", errMsg)
				common.DisabledCore = true
				return
			}
		}
	}
}
