package app

import (
	"github.com/MetaCubeX/Clash.Mini/log"
	commonUtils "github.com/MetaCubeX/Clash.Mini/util/common"
	"io/ioutil"
	"os"
)

const (
	Name     = "Clash.Mini"
	Version  = "0.1.5"
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
