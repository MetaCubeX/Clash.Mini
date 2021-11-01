package util

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/Clash-Mini/Clash.Mini/log"
)

var (
	client = &http.Client{
		Timeout: 10 * time.Second,
	}
)

// DownloadFile 下载文件
func DownloadFile(url string, destPath string) error {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	rsp, err := client.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Warnln("download ending error: %v, request url: %s, error: %s", err.Error())
		}
	}(rsp.Body)
	var errMsg string
	if err != nil || (rsp != nil && rsp.StatusCode != http.StatusOK) {
		rspBody, _ := ioutil.ReadAll(rsp.Body)
		log.Warnln("download error: %v, request url: %s, response: [%s] %s",
			err, req.URL.String(), rsp.StatusCode, string(rspBody))
		if err == http.ErrHandlerTimeout || (rsp != nil && rsp.StatusCode == http.StatusInternalServerError ||
			rsp.StatusCode == http.StatusServiceUnavailable) {
			errMsg = "无法访问链接！"
		} else if err == http.ErrNoLocation || err == http.ErrMissingFile ||
			(rsp != nil && rsp.StatusCode == http.StatusNotFound) {
			errMsg = "链接已失效！"
		} else {
			errMsg = "下载失败！"
		}
	}
	if len(errMsg) > 0 {
		return fmt.Errorf("下载出错: %s", errMsg)
	}
	f, err := os.OpenFile(destPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Warnln("download closing file io error: %v, request url: %s, file path: %s, error: %s",
				f.Name(), err.Error())
		}
	}(f)
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("下载出错: 打开文件失败: %s", err.Error())
		}
	}
	_, err = io.Copy(f, rsp.Body)
	if err != nil {
		return fmt.Errorf("下载出错: 写入文件失败: %s", err.Error())
	}
	return nil
}
