package common

import (
	"fmt"
	"github.com/Clash-Mini/Clash.Mini/app"
	"net/http"
	_ "net/http/pprof"
)

var (
	RefreshProfile func()
)

func InitFunctionsAfterGetVarFlags()  {
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