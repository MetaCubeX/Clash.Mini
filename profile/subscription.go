package profile

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Clash-Mini/Clash.Mini/constant"
	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Clash-Mini/Clash.Mini/notify"
	"github.com/Clash-Mini/Clash.Mini/util"
	commonUtils "github.com/Clash-Mini/Clash.Mini/util/common"
	fileUtils "github.com/Clash-Mini/Clash.Mini/util/file"
)

const (
	subscriptionLogHeader = logHeader + ".subscription"
)

type SubscriptionUserInfo struct {
	Upload     int64 `query:"upload"`
	Download   int64 `query:"download"`
	Total      int64 `query:"total"`
	Unused     int64
	Used       int64
	ExpireUnix int64 `query:"expire"`

	UsedInfo   string
	UnusedInfo string
	ExpireInfo string
}

func UpdateSubscriptionUserInfo() (userInfo SubscriptionUserInfo) {
	content, err := os.OpenFile(commonUtils.GetExecutablePath(constant.ConfigFile), os.O_RDWR, 0666)
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
			userInfo = SubscriptionUserInfo{}
			err = util.UnmarshalByValues(userInfoStr, &userInfo)
			if err != nil {
				log.Errorln("[%s] UpdateSubscriptionUserInfo UnmarshalByValues error: %v", subscriptionLogHeader, err)
				return
			}
			userInfo.Used = userInfo.Upload + userInfo.Download
			userInfo.Unused = userInfo.Total - userInfo.Used
			userInfo.UsedInfo = fileUtils.FormatHumanizedFileSize(userInfo.Used)
			userInfo.UnusedInfo = fileUtils.FormatHumanizedFileSize(userInfo.Unused)
			if userInfo.ExpireUnix > 0 {
				userInfo.ExpireInfo = time.Unix(userInfo.ExpireUnix, 0).Format("2006-01-02")
			} else {
				userInfo.ExpireInfo = "暂无"
			}
			return
		}
	}
	return
}