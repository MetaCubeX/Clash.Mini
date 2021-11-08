package profile

import (
	"bufio"
	"io/ioutil"
	"os"
	path "path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/Clash-Mini/Clash.Mini/common"
	"github.com/Clash-Mini/Clash.Mini/constant"
	"github.com/Clash-Mini/Clash.Mini/log"
	fileUtils "github.com/Clash-Mini/Clash.Mini/util/file"

	"github.com/fsnotify/fsnotify"
)

type Info struct {
	//Index   		int
	Name    		string
	FileSize    	string
	UpdateTime    	time.Time
	Url     		string

	Enabled 		bool
}

var (
	Profiles	[]*Info
	// TODO: others depend on it
	ProfileMap	= make(map[string]*Info)
	Locker		= new(sync.RWMutex)

	ProfileTagRegexp = regexp.MustCompile(`# Clash.Mini : (http.*)`)
)

func init() {
	go func() {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Errorln("[profile] profiles watcher create error: %v", err)
		}
		defer func(watcher *fsnotify.Watcher) {
			err := watcher.Close()
			if err != nil {
				log.Errorln("[profile] profiles watcher close error: %v", err)
			}
		}(watcher)

		done := make(chan bool)
		go func() {
			for {
				select {
				case event, ok := <-watcher.Events:
					if !ok {
						return
					}
					log.Infoln("[profile] watcher event: %v", event)
					if event.Op|fsnotify.Write|fsnotify.Remove == fsnotify.Write|fsnotify.Remove {
						log.Infoln("[profile] modified file: %s", event.Name)
						go func() {
							RefreshProfiles()
						}()
					}
				case err, ok := <-watcher.Errors:
					if !ok {
						return
					}
					log.Errorln("[profile] watcher error:", err)
				}
			}
		}()

		err = watcher.Add(constant.ProfileDir)
		if err != nil {
			log.Errorln("[profile] watch profile dir error:", err)
			return
		}
		<-done
	}()
}

func RefreshProfiles() {
	defer func() {
		Locker.Unlock()
	}()
	Locker.Lock()

	fileInfos, err := ioutil.ReadDir(constant.ProfileDir)
	if err != nil {
		log.Errorln("[profile] RefreshProfiles ReadDir error: %v", err)
		return
	}
	var profiles []*Info
	profileMap := make(map[string]*Info)
	for _, f := range fileInfos {
		extName := path.Ext(f.Name())
		if extName != constant.ConfigSuffix {
			continue
		}
		profileName := strings.TrimSuffix(f.Name(), extName)
		content, err := os.OpenFile(path.Join(constant.ProfileDir, f.Name()), os.O_RDWR, 0666)
		if err != nil {
			log.Errorln("[profile] RefreshProfiles OpenFile error: %v", err)
			continue
		}
		reader := bufio.NewReader(content)
		lineData, _, err := reader.ReadLine()
		if err != nil {
			log.Errorln("[profile] RefreshProfiles ReadLine error: %v", err)
			continue
		}
		line := GetTagLineUrl(string(lineData))
		if err = content.Close(); err != nil {
			continue
		}
		profile := &Info{
			Name: profileName,
			FileSize: fileUtils.FormatHumanizedFileSize(f.Size()),
			UpdateTime: f.ModTime(),
			Url: line,
		}
		profiles = append(profiles, profile)
		profileMap[profileName] = profile
	}
	Profiles = profiles
	ProfileMap = profileMap
	if common.RefreshProfile != nil {
		common.RefreshProfile()
	}
}

func GetTagLineUrl(line string) string {
	if ProfileTagRegexp.MatchString(line) {
		return ProfileTagRegexp.FindStringSubmatch(line)[1]
	}
	return ""
}
