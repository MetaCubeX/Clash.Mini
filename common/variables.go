package common

import (
	"flag"
	"fmt"
	"github.com/Clash-Mini/Clash.Mini/app"
)

type coreVarFlags struct {
	Version            bool
	TestConfig         bool
	HomeDir            string
	ConfigFile         string
	ExternalUI         string
	ExternalController string
	Secret             string
}

var (
	LogLevel           string
	DisabledCore       bool
	DisabledDashboard  bool

	FlagSet            map[string]bool
	CoreFlags          coreVarFlags
)

func init() {
	getVarFlags()
	InitVariablesAfterGetVarFlags()
	InitFunctionsAfterGetVarFlags()
}

func InitVariablesAfterGetVarFlags()  {
	debugMap := map[string]bool{
		"debug": true,
		"info":  false,
		"warn":  false,
		"error": false,
		"fatal": false,
	}
	var exist bool
	if app.Debug, exist = debugMap[LogLevel]; !exist {
		panic(fmt.Errorf("invalid value for log-level is \"%s\"", LogLevel))
	}
	if !app.Debug {
		app.BuggerInit()
	}
}

func getVarFlags() {
	flag.StringVar(&LogLevel, "log-level", "info", "set log level")
	flag.BoolVar(&DisabledCore, "disabled-core", false, "running without clash core")
	flag.BoolVar(&DisabledDashboard, "disabled-dashboard", false, "running without dashboard")
	flag.StringVar(&CoreFlags.HomeDir, "d", "", "set configuration directory")
	flag.StringVar(&CoreFlags.ConfigFile, "f", "", "specify configuration file")
	flag.StringVar(&CoreFlags.ExternalUI, "ext-ui", "", "override external ui directory")
	flag.StringVar(&CoreFlags.ExternalController, "ext-ctl", "", "override external controller address")
	flag.StringVar(&CoreFlags.Secret, "secret", "", "override secret for RESTful API")
	flag.BoolVar(&CoreFlags.Version, "v", false, "show current version of clash")
	flag.BoolVar(&CoreFlags.TestConfig, "t", false, "test configuration and exit")
	flag.Parse()

	FlagSet = map[string]bool{}
	flag.Visit(func(f *flag.Flag) {
		FlagSet[f.Name] = true
	})
}