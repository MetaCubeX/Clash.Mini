////go:generate go-bindata -pkg static -ignore .../.DS_Store -o gh-pages.go gh-pages/...

package static

import (
	"embed"
)

var (
	//go:embed gh-pages
	ghPages embed.FS
)

func init() {
	go func() {
		//subFs, err := fs.Sub(ghPages, "gh-pages")
		//if err != nil {
		//	log.Fatalln("open sub directory in embed.FS error: %v", err)
		//}
		//dashboardBindUrl := fmt.Sprintf("%s:%s", constant.Localhost, constant.DashboardPort)
		//if err := http.ListenAndServe(dashboardBindUrl, http.FileServer(http.FS(subFs))); err != nil {
		//	log.Fatalln("ListenAndServe error: %v", err)
		//}
	}()
}
