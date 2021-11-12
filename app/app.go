package app

import (
	"io/ioutil"
	"os"

	"github.com/Clash-Mini/Clash.Mini/log"
	commonUtils "github.com/Clash-Mini/Clash.Mini/util/common"
)

const (
	Name     = "Clash.Mini"
	Version  = "0.1.4-dev"
	CommitId = "{{COMMIT_ID}}"
)

var (
	Debug bool
)

func InitBugger() {
	if !Debug {
		buggerLock := commonUtils.GetExecutablePath("bugger.lock")
		os.Remove(buggerLock)
		InitBugsnag()
		ioutil.WriteFile(buggerLock, []byte{}, 0644)
	} else {
		log.Infoln("[app] skipped init bug reporter")
	}
}
