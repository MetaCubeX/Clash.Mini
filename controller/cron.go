package controller

import (
	"bufio"
	"fmt"
	"github.com/Clash-Mini/Clash.Mini/notify"
	"github.com/robfig/cron/v3"
	"io/ioutil"
	"log"
	"os"
	path "path/filepath"
	"regexp"
	"strings"
)

func Corntask() {
	c := cron.New()
	c.AddFunc("@every 6h", func() {
		type cronInfo struct {
			Name string
			Url  string
		}

		InfoArr, err := ioutil.ReadDir("./Profile")
		if err != nil {
			log.Fatal(err)
		}
		var match string
		items := make([]*cronInfo, 0)
		for _, cf := range InfoArr {
			if path.Ext(cf.Name()) == ".yaml" {
				content, err := os.OpenFile("./Profile/"+cf.Name(), os.O_RDWR, 0666)
				if err != nil {
					log.Fatal(err)
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
				fmt.Println(v)
				err := updateConfig(v.Name, v.Url)
				if err != true {
					fmt.Println(v.Name + "更新失败")
					items[i].Url = "更新失败"
					fail++
				} else {
					fmt.Println(v.Name + "更新成功")
					items[i].Url = "成功更新"
					success++
				}
			}
		}
		notify.NotifyCorn(success, fail)
	})
	c.Start()
	select {}
}
