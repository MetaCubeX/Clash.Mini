//go:generate goversioninfo -manifest=./resource/Clash.Mini_x64.exe.manifest -64 -o ./resource_amd64.syso
//go:generate goversioninfo -manifest=./resource/Clash.Mini_x86.exe.manifest -o ./resource_386.syso

//GOOS=windows GOARCH=amd64 go build -ldflags "-H=windowsgui -s -w" -o ./Clash.Mini_x64.exe
//GOOS=windows GOARCH=386 go build -ldflags "-H=windowsgui -s -w" -o ./Clash.Mini_x86.exe
package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"

	. "github.com/Clash-Mini/Clash.Mini/common"
	_ "github.com/Clash-Mini/Clash.Mini/config"
	"github.com/Clash-Mini/Clash.Mini/log"
	_ "github.com/Clash-Mini/Clash.Mini/static"
	_ "github.com/Clash-Mini/Clash.Mini/tray"

	"github.com/Dreamacro/clash/config"
	C "github.com/Dreamacro/clash/constant"
	"github.com/Dreamacro/clash/hub"
	"github.com/Dreamacro/clash/hub/executor"
)

func main() {

	if CoreFlags.Version {
		fmt.Printf("Clash %s %s %s %s\n", C.Version, runtime.GOOS, runtime.GOARCH, C.BuildTime)
		return
	}

	if CoreFlags.HomeDir != "" {
		if !filepath.IsAbs(CoreFlags.HomeDir) {
			currentDir, _ := os.Getwd()
			CoreFlags.HomeDir = filepath.Join(currentDir, CoreFlags.HomeDir)
		}
		C.SetHomeDir(CoreFlags.HomeDir)
	}

	if CoreFlags.ConfigFile != "" {
		if !filepath.IsAbs(CoreFlags.ConfigFile) {
			currentDir, _ := os.Getwd()
			CoreFlags.ConfigFile = filepath.Join(currentDir, CoreFlags.ConfigFile)
		}
		C.SetConfig(CoreFlags.ConfigFile)
	} else {
		configFile := filepath.Join(C.Path.HomeDir(), C.Path.Config())
		C.SetConfig(configFile)
	}

	if err := config.Init(C.Path.HomeDir()); err != nil {
		log.Fatalln("Initial configuration directory error: %s", err.Error())
	}

	if CoreFlags.TestConfig {
		if _, err := executor.Parse(); err != nil {
			log.Errorln(err.Error())
			fmt.Printf("configuration file %s test failed\n", C.Path.Config())
			os.Exit(1)
		}
		fmt.Printf("configuration file %s test is successful\n", C.Path.Config())
		return
	}

	var options []hub.Option
	if FlagSet["ext-ui"] {
		options = append(options, hub.WithExternalUI(CoreFlags.ExternalUI))
	}
	if FlagSet["ext-ctl"] {
		options = append(options, hub.WithExternalController(CoreFlags.ExternalController))
	}
	if FlagSet["secret"] {
		options = append(options, hub.WithSecret(CoreFlags.Secret))
	}

	if !DisabledCore {
		go func() {
			defer func() {
				if recover() != nil {
					log.Warnln("[recovery] Clash core is down")
					CoreRunningStatus = false
				}
			}()
			if err := hub.Parse(options...); err != nil {
				errString := fmt.Sprintf("Parse config error: %s", err.Error())
				log.Errorln(errString)
				panic(errString)
			}
			CoreRunningStatus = true
		}()
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}
