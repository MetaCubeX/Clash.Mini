package tray

import (
	"sync"
	"time"

	"github.com/Clash-Mini/Clash.Mini/proxy"
)

type PingTest struct {
	LowestDelay		int16
	FastProxy		*proxy.Proxy
	LastUpdateDT	time.Time

	Callback		func(pt *PingTest)
	locker			*sync.RWMutex
}

func (pt *PingTest) SetFastProxy(p *proxy.Proxy)  {
	pt.locker.Lock()
	defer pt.locker.Unlock()
	if pt.LowestDelay == -1 || (p.Delay != -1 && p.Delay <= pt.LowestDelay) {
		pt.FastProxy = p
		pt.LowestDelay = p.Delay
		pt.LastUpdateDT = time.Now()
		if pt.Callback != nil {
			go pt.Callback(pt)
		}
	}
}

