package start

import (
	"fmt"

	_ "github.com/Clash-Mini/Clash.Mini/app/bridge/start/fourth"

	_ "github.com/Clash-Mini/Clash.Mini/tray"
)

func init() {
	fmt.Println("[bridge] started")
}