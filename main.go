//go:generate goversioninfo -manifest=./resource/Clash.Mini_x64.exe.manifest -64 -o ./resource_amd64.syso
//go:generate goversioninfo -manifest=./resource/Clash.Mini_x86.exe.manifest -o ./resource_386.syso

//GOOS=windows GOARCH=amd64 go build -ldflags "-H=windowsgui -s -w" -o ./Clash.Mini_x64.exe
//GOOS=windows GOARCH=386 go build -ldflags "-H=windowsgui -s -w" -o ./Clash.Mini_x86.exe
package main

import (
	cConfig "github.com/Clash-Mini/Clash.Mini/config"
	"github.com/Clash-Mini/Clash.Mini/constant"
	cI18n "github.com/Clash-Mini/Clash.Mini/constant/i18n"
	"github.com/Clash-Mini/Clash.Mini/controller"
	"github.com/Clash-Mini/Clash.Mini/mixin"
	"github.com/Clash-Mini/Clash.Mini/util/common"
	"github.com/Clash-Mini/Clash.Mini/util/uac"
	C "github.com/Dreamacro/clash/constant"
	"github.com/JyCyunMe/go-i18n/i18n"
	"github.com/lxn/walk"
	"github.com/lxn/win"
	"os"
	"os/signal"
	"path"
	"syscall"

	_ "github.com/Clash-Mini/Clash.Mini/app/bridge/start"
	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Dreamacro/clash/config"
)

func main() {

	Restart()

	Name := cConfig.GetProfile()
	exist, configName := controller.CheckConfig(Name)
	if exist {
		configFile := path.Join(constant.ProfileDir, configName)
		C.SetConfig(configFile)
	}
	// init mmdb and geosite
	if err := config.Init(common.GetExecutablePath()); err != nil {
		log.Errorln("Initial configuration directory error: %s", err.Error())
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}

func Restart() (done bool) {
	done = true
	if cConfig.IsMixinPositive(mixin.Tun) {
		if !uac.AmAdmin {
			rlt := walk.MsgBox(nil, i18n.T(cI18n.UacMsgBoxTitle),
				i18n.T(cI18n.UacMsgBoxTunFailedMsg), walk.MsgBoxIconQuestion|walk.MsgBoxOKCancel)
			if rlt != win.IDOK {
				log.Infoln("[winTun] user skipped restart")
				return false
			}
			uac.RunAsElevate(constant.Executable, "")
			os.Exit(0)
		}
	}
	return done
}
