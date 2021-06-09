package controller

import (
	"bufio"
	"io/ioutil"
	"os"
	path "path/filepath"
	"regexp"
	"strings"

	"github.com/Clash-Mini/Clash.Mini/constant"
	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Clash-Mini/Clash.Mini/notify"

	"github.com/robfig/cron/v3"
)

func CronTask() {
	c := cron.New()
	c.AddFunc("@every 3h", func() {
		type cronInfo struct {
			Name string
			Url  string
		}
		currentName, _ := CheckConfig()
		InfoArr, err := ioutil.ReadDir(constant.ConfigDir)
		if err != nil {
			log.Fatalln("CronTask ReadDir error: %v", err)
		}
		var match string
		items := make([]*cronInfo, 0)
		for _, cf := range InfoArr {
			if path.Ext(cf.Name()) == constant.ConfigSuffix {
				content, err := os.OpenFile(path.Join(constant.ConfigDir, cf.Name()), os.O_RDWR, 0666)
				if err != nil {
					log.Fatalln("CronTask OpenFile error: %v", err)
				}
				scanner := bufio.NewScanner(content)
				Reg := regexp.MustCompile(`# Clash.Mini : (http.*)`)
				for scanner.Scan() {
					if Reg.MatchString(scanner.Text()) {
						match = Reg.FindStringSubmatch(scanner.Text())[1]
						break
					} else {
						match = ""
					}
				}
				content.Close()
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
				err := updateConfig(v.Name, v.Url)
				if err != true {
					log.Errorln(v.Name + "更新失败")
					items[i].Url = "更新失败"
					fail++
				} else {
					log.Infoln(v.Name + "更新成功")
					items[i].Url = "成功更新"
					success++
					if v.Name == currentName {
						putConfig(v.Name)
					}
				}

			}
		}
		notify.PushProfileCronFinished(success, fail)
	})
	c.Start()
	select {}
}
