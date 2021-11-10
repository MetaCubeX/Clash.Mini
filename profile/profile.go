package profile

import (
	"bufio"
	"bytes"
	"container/list"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	path "path/filepath"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Clash-Mini/Clash.Mini/common"
	"github.com/Clash-Mini/Clash.Mini/constant"
	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Clash-Mini/Clash.Mini/static"
	_ "github.com/Clash-Mini/Clash.Mini/static"
	fileUtils "github.com/Clash-Mini/Clash.Mini/util/file"
	httpUtils "github.com/Clash-Mini/Clash.Mini/util/http"

	"github.com/fsnotify/fsnotify"
)

const (
	logHeader = "profile"
)

type Info struct {
	//Index   		int
	Name    		string
	FileSize    	string
	UpdateTime    	time.Time
	Url     		string

	FileHash    	string
	Enabled 		bool
}

var (
	Profiles	= new(list.List)
	//Profiles	[]*Info
	// TODO: others depend on it
	ProfileMap	sync.Map
	//ProfileMap	= make(map[string]*Info)
	counter		int32
	//wg			sync.WaitGroup
	Locker		= new(sync.RWMutex)

	// TODO: others depend on it
	MenuItemMap		sync.Map

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

		//wg = sync.Pool{}
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
						//wg.Add()
						go RefreshProfiles(&event)
					}
				case err, ok := <-watcher.Errors:
					if !ok {
						return
					}
					log.Errorln("[profile] watcher error: %v", err)
				}
			}
		}()

		err = watcher.Add(constant.ProfileDir)
		if err != nil {
			log.Errorln("[profile] watch profile dir error: %v", err)
			return
		}
		<-done
	}()
}

func RefreshProfiles(event *fsnotify.Event) {
	defer func() {
		Locker.Unlock()
		atomic.AddInt32(&counter, -1)
	}()
	Locker.Lock()
	atomic.AddInt32(&counter, 1)

	var err	error
	var fileInfos []fs.FileInfo
	var isRemove bool
	if event == nil {
		fileInfos, err = ioutil.ReadDir(constant.ProfileDir)
		if err != nil {
			log.Errorln("[profile] RefreshProfiles ReadDir error: %v", err)
			return
		}
	} else {
		fileInfos = []fs.FileInfo{static.NewFakeFile(event.Name)}
		isRemove = event.Op|fsnotify.Remove == fsnotify.Remove
	}
	//var profiles []*Info
	for _, f := range fileInfos {
		extName := path.Ext(f.Name())
		if extName != constant.ConfigSuffix {
			continue
		}
		fileName := f.Name()
		profileName := f.Name()
		if path.IsAbs(f.Name()) {
			profileName = path.Base(f.Name())
		} else {
			fileName = path.Join(constant.ProfileDir, fileName)
		}
		profileName = strings.TrimSuffix(profileName, extName)
		if isRemove {
			ProfileMap.Delete(profileName)
			for e := Profiles.Front(); e != nil; e = e.Next() {
				if e.Value.(*Info).Name == profileName {
					Profiles.Remove(e)
					break
				}
			}
			continue
		}
		profile := &Info{
			Name: profileName,
			FileSize: fileUtils.FormatHumanizedFileSize(f.Size()),
			UpdateTime: f.ModTime(),
		}
		content, err := os.OpenFile(fileName, os.O_RDWR, 0666)
		if err != nil {
			log.Errorln("[profile] RefreshProfiles OpenFile error: %v", err)
			continue
		}
		reader := bufio.NewReader(content)
		v, exists := ProfileMap.Load(profileName)
		hash := fileUtils.GetHash(reader, 32)
		if exists && v.(*Info).FileHash == hash {
			continue
		} else {
			profile.FileHash = hash
		}
		lineData, _, err := reader.ReadLine()
		if err != nil {
			log.Errorln("[profile] RefreshProfiles ReadLine error: %v", err)
			continue
		}
		profile.Url = GetTagLineUrl(string(lineData))
		if err = content.Close(); err != nil {
			continue
		}
		//profiles = append(profiles, profile)
		v, loaded := ProfileMap.LoadOrStore(profileName, profile)
		log.Infoln("[profile] loaded: %t", loaded)
		if loaded {
			original := v.(*Info)
			original.Url = profile.Url
			original.FileSize = profile.FileSize
			original.FileHash = profile.FileHash
			original.UpdateTime = profile.UpdateTime
		} else {
			Profiles.PushBack(profile)
		}
		//if loaded {
		//
		//}
	}
	// TODO:
	//Profiles = profiles
	//if atomic.LoadInt32(&counter) == 1 {
	go common.RefreshProfile(event)
	//}
}

// UpdateConfig 更新订阅配置
func UpdateConfig(name, url string) (successful bool) {
	client := &http.Client{Timeout: 5 * time.Second}
	res, _ := http.NewRequest(http.MethodGet, url, nil)
	res.Header.Add("User-Agent", "clash")
	rsp, err := client.Do(res)
	if err != nil {
		return false
	}
	if rsp != nil && rsp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(rsp.Body)
		matched, _ := regexp.MatchString(`proxy-groups`, string(body))
		if !matched {
			log.Errorln("[profile] format is not supported")
			return false
		}
		rebody := ioutil.NopCloser(bytes.NewReader(body))

		f, err := os.OpenFile(path.Join(constant.ProfileDir, name+constant.ConfigSuffix),
			os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
		if err != nil {
			panic(err)
			return false
		}
		defer func() {
			httpUtils.DeferSafeCloseResponseBody(rsp)
			if f != nil {
				f.Close()
			}
		}()
		_, err = f.WriteString(fmt.Sprintf("# Clash.Mini : %s\n", url))
		if err != nil {
			log.Errorln("[profile] writeString error: %v", err)
			return false
		}

		// TODO:
		//parser.DoParse(url)
		//config.Config

		_, err = io.Copy(f, rebody)
		if err != nil {
			log.Errorln("[profile] copy error: %v", err)
			return false
		}
		return true
	}
	return false
}

func GetTagLineUrl(line string) string {
	if ProfileTagRegexp.MatchString(line) {
		return ProfileTagRegexp.FindStringSubmatch(line)[1]
	}
	return ""
}
