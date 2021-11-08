package app

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"

	commonUtils "github.com/Clash-Mini/Clash.Mini/util/common"
)

const (
	Name 			= "Clash.Mini"
	Version 		= "0.1.4-dev"
	CommitId		= "{{COMMIT_ID}}"
)

const (
	appLogHeader 	= "app"
)

var (
	Debug 			bool

	msgQueueBuffer  = bytes.NewBuffer([]byte{})
	msgQueue		= bufio.NewReadWriter(bufio.NewReader(msgQueueBuffer), bufio.NewWriter(msgQueueBuffer))
	msgQueueLocker	= new(sync.Mutex)
)

func InitBugger() {
	buggerLock := commonUtils.GetExecutablePath("bugger.lock")
	os.Remove(buggerLock)
	if !Debug {
		InitBugsnag()
		ioutil.WriteFile(buggerLock, []byte{}, 0644)
	} else {
		writeMsg(appLogHeader, "skipped init bug reporter")
	}
}

func writeMsg(logHeader, format string, v... interface{}) {
	msgQueue.WriteString(fmt.Sprintf(fmt.Sprintf("[%s] %s\r\n", logHeader, format), v...))
	msgQueue.Flush()
}

func PrintMsg(logFunc func(string, ...interface{})) {
	defer func() {
		msgQueueLocker.Unlock()
	}()
	msgQueueLocker.Lock()
	for msgQueue.Available() > 0 {
		line, _, err := msgQueue.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			logFunc("[msgQueue] error: %v", err)
			continue
		}
		logFunc(string(line))
	}
	msgQueueBuffer.Reset()
	msgQueue.Flush()
}