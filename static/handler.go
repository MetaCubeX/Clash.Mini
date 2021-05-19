//go:generate go-bindata -pkg static -ignore .../.DS_Store -o gh-pages.go gh-pages/...

package static

import (
	"log"
	"net/http"

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
		if err := http.ListenAndServe("127.0.0.1:8070", handler); err != nil {
			log.Panicln(err)
		}

	}()
}
