package mq

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"sync"
)

var (
	msgQueueBuffer = bytes.NewBuffer([]byte{})
	msgQueue       = bufio.NewReadWriter(bufio.NewReader(msgQueueBuffer), bufio.NewWriter(msgQueueBuffer))
	msgQueueLocker = new(sync.Mutex)
)

func WriteMsg(logHeader, format string, v ...interface{}) {
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
