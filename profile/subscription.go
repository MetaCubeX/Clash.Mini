package profile

import (
	"bufio"
	"fmt"
	"github.com/Clash-Mini/Clash.Mini/config"
	"net/http"
	"os"
	path "path/filepath"
	"strconv"
	"strings"
	"time"

	. "github.com/Clash-Mini/Clash.Mini/constant"
	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Clash-Mini/Clash.Mini/notify"
	fileUtils "github.com/Clash-Mini/Clash.Mini/util/file"
)

const (
	subscriptionLogHeader = logHeader + ".subscription"
)

type SubscriptionUserInfo struct {
	Upload     int64
	Download   int64
	Total      int64
	Unused     int64
	Used       int64
	ExpireUnix int64

	UsedInfo   string
	UnusedInfo string
	ExpireInfo string
}

func UpdateSubscriptionUserInfo() (userInfo SubscriptionUserInfo) {
	configName := config.GetProfile()
	content, err := os.OpenFile(path.Join(ProfileDir, configName+ConfigSuffix), os.O_RDWR, 0666)
	if err != nil {
		errMsg := fmt.Sprintf("updateSubscriptionUserInfo OpenFile error: %v", err)
		log.Errorln(errMsg)
		notify.PushError("", errMsg)
		return
	}

	reader := bufio.NewReader(content)
	lineData, _, err := reader.ReadLine()
	if err != nil {
		log.Errorln("[profile] updateSubscriptionUserInfo ReadLine error: %v", err)
		return
	}
	infoURL := GetTagLineUrl(string(lineData))
	if err = content.Close(); err != nil {
		log.Errorln("[profile] RefreshProfiles CloseFile error: %v", err)
		return
	}

	if infoURL != "" {
		client := &http.Client{Timeout: 5 * time.Second}
		res, _ := http.NewRequest(http.MethodGet, infoURL, nil)
		res.Header.Add("User-Agent", "clash")
		resp, err := client.Do(res)
		if err != nil {
			return
		}
		userInfoStr := resp.Header.Get("Subscription-Userinfo")
		if len(strings.TrimSpace(userInfoStr)) == 0 {
			res2, _ := http.NewRequest(http.MethodGet, infoURL, nil)
			res2.Header.Add("User-Agent", "Quantumultx")
			resp2, err := client.Do(res2)
			if err != nil {
				return
			}
			userInfoStr = resp2.Header.Get("Subscription-Userinfo")

		}
		if len(strings.TrimSpace(userInfoStr)) > 0 {
			flags := strings.Split(userInfoStr, ";")
			for _, value := range flags {
				info := strings.Split(value, "=")
				switch {
				case strings.Contains(value, "upload"):
					userInfo.Upload, _ = strconv.ParseInt(info[1], 10, 64)
				case strings.Contains(value, "download"):
					userInfo.Download, _ = strconv.ParseInt(info[1], 10, 64)
				case strings.Contains(value, "total"):
					userInfo.Total, _ = strconv.ParseInt(info[1], 10, 64)
				case strings.Contains(value, "expire"):
					userInfo.ExpireUnix, _ = strconv.ParseInt(info[1], 10, 64)
				}
			}
			userInfo.Used = userInfo.Upload + userInfo.Download
			userInfo.Unused = userInfo.Total - userInfo.Used
			userInfo.UsedInfo = fileUtils.FormatHumanizedFileSize(userInfo.Used)
			userInfo.UnusedInfo = fileUtils.FormatHumanizedFileSize(userInfo.Unused)
			if userInfo.ExpireUnix > 0 {
				userInfo.ExpireInfo = time.Unix(userInfo.ExpireUnix, 0).Format("2006-01-02")
			} else {
				userInfo.ExpireInfo = "None"
			}
			return
		}
	}
	return
}
