package controller

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	path "path/filepath"
	"strings"
	"time"

	"github.com/Clash-Mini/Clash.Mini/cmd/mmdb"
	"github.com/Clash-Mini/Clash.Mini/constant"
	"github.com/Clash-Mini/Clash.Mini/log"

	"github.com/lxn/walk"
)

func GetMMDB(value mmdb.Type) {
	// 检查是否存在
	var url string
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	switch value {
	case mmdb.Max:
		url = strings.Join([]string{constant.GitHubCDN, "Dreamacro/maxmind-geoip", constant.MMDBSuffix}, "")
		break
	case mmdb.Lite:
		url = strings.Join([]string{constant.GitHubCDN, "Hackl0us/GeoIP2-CN", constant.MMDBSuffix}, "")
		break
	}
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	rsp, err := client.Do(req)
	defer rsp.Body.Close()
	if err != nil || (rsp != nil && rsp.StatusCode != http.StatusOK) {
		rspBody, _ := ioutil.ReadAll(rsp.Body)
		log.Warnln("GetMMDB Do error: %v, request url: %s, response: [%s] %s",
			err, req.URL.String(), rsp.StatusCode, string(rspBody))
		var errMsg string
		if err == http.ErrHandlerTimeout || (rsp != nil && rsp.StatusCode == http.StatusInternalServerError ||
			rsp.StatusCode == http.StatusServiceUnavailable) {
			errMsg = "无法访问到链接！"
		} else if err == http.ErrNoLocation || err == http.ErrMissingFile ||
			(rsp != nil && rsp.StatusCode == http.StatusNotFound) {
			errMsg = "链接已失效！"
		} else {
			errMsg = "下载失败！"
		}
		walk.MsgBox(nil, constant.UIConfigMsgTitle, errMsg, walk.MsgBoxIconError)
	}
	mmdbFile := path.Join(".", "Country.mmdb")
	f, err := os.OpenFile(mmdbFile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatalln("GetMMDB OpenFile error: %v", err)
		}
	}
	io.Copy(f, rsp.Body)
	RegCmd(value)
}
