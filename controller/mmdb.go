package controller

import (
	"github.com/lxn/walk"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func GetMMDB(arg string) {
	var url string
	var value string
	client := &http.Client{}
	switch arg {
	case "Max":
		url = "https://cdn.jsdelivr.net/gh/Dreamacro/maxmind-geoip@release/Country.mmdb"
		value = "Max"
	case "Lite":
		url = "https://cdn.jsdelivr.net/gh/Hackl0us/GeoIP2-CN@release/Country.mmdb"
		value = "Lite"
	}
	res, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(res)
	defer resp.Body.Close()
	if err != nil {
		walk.MsgBox(nil, "配置提示", "请检查订阅链接是否正确！", walk.MsgBoxIconError)
	}
	ConfigDir := filepath.Join(".", "Country.mmdb")
	f, err := os.OpenFile(ConfigDir, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	io.Copy(f, resp.Body)
	Regcmd("MMBD", value)
}
