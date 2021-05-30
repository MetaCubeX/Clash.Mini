package controller

import (
	"io"
	"net/http"
	"os"
	path "path/filepath"

	"github.com/Clash-Mini/Clash.Mini/cmd/mmdb"

	"github.com/lxn/walk"
)

func GetMMDB(value mmdb.Type) {
	var url string
	client := &http.Client{}
	switch value {
	case mmdb.Max:
		url = "https://cdn.jsdelivr.net/gh/Dreamacro/maxmind-geoip@release/Country.mmdb"
		break
	case mmdb.Lite:
		url = "https://cdn.jsdelivr.net/gh/Hackl0us/GeoIP2-CN@release/Country.mmdb"
		break
	}
	res, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, err := client.Do(res)
	defer resp.Body.Close()
	if err != nil {
		walk.MsgBox(nil, "配置提示", "请检查订阅链接是否正确！", walk.MsgBoxIconError)
	}
	ConfigDir := path.Join(".", "Country.mmdb")
	f, err := os.OpenFile(ConfigDir, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	io.Copy(f, resp.Body)
	RegCmd(value)
}
