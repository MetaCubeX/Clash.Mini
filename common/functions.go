package common

import (
	"net/http"
	_ "net/http/pprof"
)

var (
	RefreshProfile func()
)

func init() {
	// pprof
	go func() {
		http.ListenAndServe("http://127.0.0.1:6060", nil)
	}()
}
