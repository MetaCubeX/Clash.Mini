////go:generate go-bindata -pkg static -ignore .../.DS_Store -o gh-pages.go gh-pages/...

package static

import (
	"embed"
	"fmt"
	"github.com/Clash-Mini/Clash.Mini/common"
	"github.com/Clash-Mini/Clash.Mini/constant"
	"io/fs"
	"net/http"
	path "path/filepath"

	"github.com/Clash-Mini/Clash.Mini/log"
)

var (
	//go:embed gh-pages
	ghPages embed.FS

	//go:embed lang/*.lang
	langPackages embed.FS
)

func init() {
	if !common.DisabledDashboard {
		go func() {
			subFs, err := fs.Sub(ghPages, "gh-pages")
			if err != nil {
				log.Fatalln("open sub directory in embed.FS error: %v", err)
			}
			dashboardBindUrl := fmt.Sprintf("%s:%s", constant.Localhost, constant.DashboardPort)
			if err := http.ListenAndServe(dashboardBindUrl, http.FileServer(http.FS(subFs))); err != nil {
				log.Fatalln("ListenAndServe error: %v", err)
			}
		}()
	}
}

func LoadEmbedLanguages(ignoreError bool) ([]*fs.FileInfo, error) {
	var packageBytes []*fs.FileInfo
	fileNames, err := GetAllFileNames(&langPackages, "lang", false, true)
	if err != nil {
		return nil, err
	}
	fmt.Println(fileNames)
	for _, fileName := range fileNames {
		data, err := langPackages.ReadFile(fileName)
		if err != nil {
			if !ignoreError {
				return nil, err
			} else {
				log.Warnln("[embed] read language file \"%s\" error: %v", fileName, err)
				continue
			}
		}
		fakeFile := fs.FileInfo(FakeFile{name: fileName, data: &data})
		packageBytes = append(packageBytes, &(fakeFile))
	}
	return packageBytes, nil
}

// GetAllFileNames 获取全部文件名 (fs, 目录, 是否递归, 是否忽略错误)
func GetAllFileNames(fs *embed.FS, dir string, deep bool, ignoreError bool) (fileNames []string, err error) {
	if len(dir) == 0 {
		dir = "."
	}
	entries, err := fs.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		fp := path.ToSlash(path.Join(dir, entry.Name()))
		if entry.IsDir() {
			if deep {
				res, err := GetAllFileNames(fs, fp, deep, ignoreError)
				if err != nil {
					if !ignoreError {
						return nil, err
					} else {
						continue
					}
				}
				fileNames = append(fileNames, res...)
				continue
			}
		} else {
			fileNames = append(fileNames, fp)
		}
	}
	return
}
