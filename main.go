//go:generate goversioninfo -manifest=./resource/Clash.Mini_x64.exe.manifest -64 -o ./resource_amd64.syso
//go:generate goversioninfo -manifest=./resource/Clash.Mini_x86.exe.manifest -o ./resource_386.syso

//GOOS=windows GOARCH=amd64 go build -ldflags "-H=windowsgui -s -w" -o ./Clash.Mini_x64.exe
//GOOS=windows GOARCH=386 go build -ldflags "-H=windowsgui -s -w" -o ./Clash.Mini_x86.exe
package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"

	_ "github.com/Clash-Mini/Clash.Mini/common"
	_ "github.com/Clash-Mini/Clash.Mini/config"
	"github.com/Clash-Mini/Clash.Mini/log"
	_ "github.com/Clash-Mini/Clash.Mini/static"
	_ "github.com/Clash-Mini/Clash.Mini/tray"

	"github.com/Dreamacro/clash/config"
	C "github.com/Dreamacro/clash/constant"
	"github.com/Dreamacro/clash/hub"
	"github.com/Dreamacro/clash/hub/executor"
)

var (
	flagSet            map[string]bool
	version            bool
	testConfig         bool
	homeDir            string
	configFile         string
	externalUI         string
	externalController string
	secret             string

	logLevel           string
	guiOnly            bool
)

func init() {
	flag.StringVar(&logLevel, "log-level", "info", "set log level")
	flag.BoolVar(&guiOnly, "gui-only", false, "run gui only (except clash core)")
	flag.StringVar(&homeDir, "d", "", "set configuration directory")
	flag.StringVar(&configFile, "f", "", "specify configuration file")
	flag.StringVar(&externalUI, "ext-ui", "", "override external ui directory")
	flag.StringVar(&externalController, "ext-ctl", "", "override external controller address")
	flag.StringVar(&secret, "secret", "", "override secret for RESTful API")
	flag.BoolVar(&version, "v", false, "show current version of clash")
	flag.BoolVar(&testConfig, "t", false, "test configuration and exit")
	flag.Parse()

	flagSet = map[string]bool{}
	flag.Visit(func(f *flag.Flag) {
		flagSet[f.Name] = true
	})
}

func main() {

	if version {
		fmt.Printf("Clash %s %s %s %s\n", C.Version, runtime.GOOS, runtime.GOARCH, C.BuildTime)
		return
	}

	if homeDir != "" {
		if !filepath.IsAbs(homeDir) {
			currentDir, _ := os.Getwd()
			homeDir = filepath.Join(currentDir, homeDir)
		}
		C.SetHomeDir(homeDir)
	}

	if configFile != "" {
		if !filepath.IsAbs(configFile) {
			currentDir, _ := os.Getwd()
			configFile = filepath.Join(currentDir, configFile)
		}
		C.SetConfig(configFile)
	} else {
		configFile := filepath.Join(C.Path.HomeDir(), C.Path.Config())
		C.SetConfig(configFile)
	}

	if err := config.Init(C.Path.HomeDir()); err != nil {
		log.Fatalln("Initial configuration directory error: %s", err.Error())
	}

	if testConfig {
		if _, err := executor.Parse(); err != nil {
			log.Errorln(err.Error())
			fmt.Printf("configuration file %s test failed\n", C.Path.Config())
			os.Exit(1)
		}
		fmt.Printf("configuration file %s test is successful\n", C.Path.Config())
		return
	}

	var options []hub.Option
	if flagSet["ext-ui"] {
		options = append(options, hub.WithExternalUI(externalUI))
	}
	if flagSet["ext-ctl"] {
		options = append(options, hub.WithExternalController(externalController))
	}
	if flagSet["secret"] {
		options = append(options, hub.WithSecret(secret))
	}

	if !guiOnly {
		go func() {
			defer func() {
				if recover() != nil {
					log.Warnln("[recovery] Clash core is down")
				}
			}()
			if err := hub.Parse(options...); err != nil {
				errString := fmt.Sprintf("Parse config error: %s", err.Error())
				log.Errorln(errString)
				panic(errString)
			}
		}()
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}
