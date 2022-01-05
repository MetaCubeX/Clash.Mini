package common

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/MetaCubeX/Clash.Mini/app"

	"github.com/fsnotify/fsnotify"
)

var (
	RefreshProfile = func(event *fsnotify.Event) {}
)

func InitFunctionsAfterGetVarFlags() {
	runPprof(app.Debug)
}

func runPprof(run bool) {
	if run {
		// pprof
		go func() {
			fmt.Println(http.ListenAndServe("127.0.0.1:6060", nil))
		}()
	}
}
