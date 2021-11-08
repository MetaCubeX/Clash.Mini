package controller

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	path "path/filepath"
	"strings"

	"github.com/Clash-Mini/Clash.Mini/constant"
	cI18n "github.com/Clash-Mini/Clash.Mini/constant/i18n"
	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Clash-Mini/Clash.Mini/notify"
	"github.com/Clash-Mini/Clash.Mini/profile"

	"github.com/JyCyunMe/go-i18n/i18n"
	"github.com/robfig/cron/v3"
)

func CronTask() {
	c := cron.New()
	// TODO: custom
	c.AddFunc("@every 3h", func() {
		type cronInfo struct {
			Name string
			Url  string
		}
		currentName, _ := CheckConfig()
		InfoArr, err := ioutil.ReadDir(constant.ProfileDir)
		if err != nil {
			errMsg := fmt.Sprintf("CronTask ReadDir error: %v", err)
			log.Errorln(errMsg)
			notify.PushError("", errMsg)
			return
		}
		var match string
		items := make([]*cronInfo, 0)
		for _, cf := range InfoArr {
			if path.Ext(cf.Name()) == constant.ConfigSuffix {
				content, err := os.OpenFile(path.Join(constant.ProfileDir, cf.Name()), os.O_RDWR, 0666)
				if err != nil {
					errMsg := fmt.Sprintf("CronTask OpenFile error: %v", err)
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
				match = profile.GetTagLineUrl(string(lineData))
				if err = content.Close(); err != nil {
					log.Errorln("[profile] RefreshProfiles CloseFile error: %v", err)
					return
				}

				items = append(items, &cronInfo{
					Name: strings.TrimSuffix(cf.Name(), path.Ext(cf.Name())),
					Url:  match,
				})
			}
		}
		success := 0
		fail := 0
		for i, v := range items {
			if v.Url != "" {
				log.Infoln("CronTask Info: %v", v)
				successful := updateConfig(v.Name, v.Url)
				if !successful {
					log.Errorln(fmt.Sprintf("%s: %s", i18n.T(cI18n.MenuConfigCronUpdateFailed), v.Name))
					items[i].Url = i18n.T(cI18n.MenuConfigCronUpdateFailed)
					fail++
				} else {
					log.Infoln(fmt.Sprintf("%s: %s", i18n.T(cI18n.MenuConfigCronUpdateSuccessful), v.Name))
					items[i].Url = i18n.T(cI18n.MenuConfigCronUpdateSuccessful)
					success++
					if v.Name == currentName {
						PutConfig(v.Name)
					}
				}

			}
		}
		notify.PushProfileCronFinished(success, fail)
	})
	c.Start()
	select {}
}
