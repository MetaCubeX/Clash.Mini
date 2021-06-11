package proxy

import (
	"context"
	"sync"
	"time"

	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Dreamacro/clash/tunnel"
)

const (
	TestUrl = "http://www.gstatic.com/generate_204"
)

var (
	maxGoroutinesNum 	 = 20
	locker				 *sync.RWMutex
)

func init() {
	locker = new(sync.RWMutex)
}

func RefreshAllDelay(singleCallback func(name string, delay int16), doneCallback func(delayMap map[string]int16)) {
	locker.Lock()
	log.Infoln("lock")
	ch := make(chan bool, maxGoroutinesNum)
	defer func() {
		log.Infoln("unlock")
		locker.Unlock()
		delayMap := make(map[string]int16)
		if doneCallback != nil {
			log.Infoln("done")
			doneCallback(delayMap)
		}
	}()
	for _ , p := range tunnel.Proxies() {
		ch <- true
		proxy := p
		go func() {
			if proxy.Name() == "DIRECT" || proxy.Name() == "GLOBAL" || proxy.Name() == "REJECT"{
				return
			}
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(10))
			var delay uint16
			var err error
			var timeout bool
			defer func() {
				cancel()
				<-ch
				if singleCallback != nil {
					if timeout {
						singleCallback(proxy.Name(), int16(-1))
					} else {
						singleCallback(proxy.Name(), int16(delay))
					}
				}
			}()
			delay, err = proxy.URLTest(ctx, TestUrl)
			if ctx.Err() != nil {
				timeout = true
				log.Warnln("[Delay] Timeout: %s", proxy.Name())
				//log.Errorln("[Delay] Timeout: %s", proxy.Name())
				return
			}
			if err != nil || delay == 0 {
				timeout = true
				log.Warnln("An error occurred in the delay test: %s", proxy.Name())
				//log.Errorln("An error occurred in the delay test: %s", proxy.Name())
				return
			}
			log.Infoln("test finished: %s :: %d ms", proxy.Name(), delay)
		}()
	}
}
