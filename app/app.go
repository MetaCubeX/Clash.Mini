package app

import (
	"fmt"
	"os"
	"strings"
)

const (
	appVersion 		= "v0.1.3-pre"

	logLevelFlag 	= "-log-level="
)

var (
	Debug 		bool
)

func init() {
	debugMap := map[string]bool{
		"debug": true,
		"info":  false,
		"warn":  false,
		"error": false,
		"fatal": false,
	}
	for _, arg := range os.Args {
		if strings.Index(arg, logLevelFlag) == 0 {
			logLevel := arg[len(logLevelFlag):]
			var exist bool
			if Debug, exist = debugMap[logLevel]; !exist {
				fmt.Errorf("invalid value for -log-level is \"%s\"", logLevel)
			}
			break
		}
	}
	if !Debug {
		BuggerInit()
	}
}
