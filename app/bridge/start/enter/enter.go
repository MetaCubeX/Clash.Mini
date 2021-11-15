package enter

import (
	"github.com/Clash-Mini/Clash.Mini/app/bridge/mq"
)

func init() {
	mq.WriteMsg("bridge", "Start...")
	mq.WriteMsg("bridge", "Step First: Checking...")
}
