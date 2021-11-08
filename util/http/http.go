package http

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Clash-Mini/Clash.Mini/util"
)

var (
	client = &http.Client{
		Timeout: 10 * time.Second,
	}
)

func NewRequest(method string, url string, body io.Reader) (req *http.Request, err error) {
	req, err = http.NewRequest(method, url, body)
	return
}

func AddClashHeader(req *http.Request) *http.Request {
	req.Header.Add("User-Agent", "clash")
	return req
}

func DoRequest(req *http.Request) (rsp *http.Response, err error) {
	rsp, err = client.Do(req)
	return
}

// DownloadFile 下载文件
func DownloadFile(url string, destPath string) error {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	rsp, err := client.Do(req)
	defer DeferSafeCloseResponseBody(rsp)
	var errMsg string
	statusCode := -1
	if rsp != nil {
		statusCode = rsp.StatusCode
	}
	if err != nil || (statusCode != http.StatusOK) {
		rspBody, _ := ioutil.ReadAll(rsp.Body)
		log.Warnln("[http] download error: %v, request url: %s, response: [%d] %s",
			err, req.URL.String(), statusCode, string(rspBody))
		if err == http.ErrHandlerTimeout ||
			(util.EqualsAny(statusCode, http.StatusInternalServerError, http.StatusServiceUnavailable)) {
			errMsg = fmt.Sprintf("the service is unavailable [%d]", statusCode)
		} else if err == http.ErrNoLocation || err == http.ErrMissingFile ||
			(statusCode == http.StatusNotFound) {
			errMsg = "cannot access the url [404]"
		} else {
			errMsg = "download failed！"
		}
	}
	if len(errMsg) > 0 {
		return fmt.Errorf("下载出错: %s", errMsg)
	}
	f, err := os.OpenFile(destPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Warnln("download closing file io error: %v, request url: %s, file path: %s", err, url, f.Name())
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

// RequestGet 发起Get请求
func RequestGet(url string) (*[]byte, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	rsp, err := client.Do(req)
	defer DeferSafeCloseResponseBody(rsp)
	var errMsg string
	statusCode := -1
	if rsp != nil {
		statusCode = rsp.StatusCode
	}
	if err != nil || (statusCode != http.StatusOK) {
		if rsp != nil {
			rspBody, err := ioutil.ReadAll(rsp.Body)
			var rspBodyInfo string
			if err != nil {
				rspBodyInfo = fmt.Sprintf("[%d b] %s", len(rspBody), string(rspBody))
			}
			errMsg = fmt.Sprintf("[http] request error: %v, url: %s, response: [%d]%s",
				err, req.URL.String(), statusCode, rspBodyInfo)
			log.Warnln(errMsg)
		} else {
			log.Warnln("[http] request error: %v, url: %s", err, url)
		}
		if err == http.ErrHandlerTimeout ||
			(util.EqualsAny(statusCode, http.StatusInternalServerError, http.StatusServiceUnavailable)) {
			errMsg = fmt.Sprintf("the service is unavailable [%d]", statusCode)
		} else if err == http.ErrNoLocation || err == http.ErrMissingFile || statusCode == http.StatusNotFound {
			errMsg = "cannot access the url [404]"
		} else {
			errMsg = "request failed！" + errMsg
		}
	}
	if len(errMsg) > 0 {
		return nil, fmt.Errorf("[http] download error[%d]: %s", statusCode, errMsg)
	}
	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func SafeCloseResponseBody(rsp *http.Response) error {
	if rsp == nil || rsp.Close {
		return nil
	}
	err := rsp.Body.Close()
	if err != nil {
		return fmt.Errorf("[http] close response error: %v, url: %s", err, rsp.Request.RequestURI)
	}
	return nil
}

func DeferSafeCloseResponseBody(rsp *http.Response) {
	err := SafeCloseResponseBody(rsp)
	if err != nil {
		log.Warnln(err.Error())
	}
}
