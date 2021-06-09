//go:generate go-bindata -pkg static -ignore .../.DS_Store -o gh-pages.go gh-pages/...

package static

import (
	"fmt"
	"net/http"

	"github.com/Clash-Mini/Clash.Mini/constant"
	"github.com/Clash-Mini/Clash.Mini/log"

	"github.com/elazarl/go-bindata-assetfs"
)

func init() {
	go func() {
		handler := http.FileServer(&assetfs.AssetFS{
			Asset:     Asset,
			AssetDir:  AssetDir,
			AssetInfo: AssetInfo,
			Prefix:    "gh-pages",
		})
		if err := http.ListenAndServe(fmt.Sprintf("%s:%s", constant.Localhost, constant.DashboardPort), handler); err != nil {
			log.Fatalln("ListenAndServe error: %v", err)
		}

	}()
}
