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
	"time"

	"github.com/MetaCubeX/Clash.Mini/common"
	"github.com/MetaCubeX/Clash.Mini/constant"
	"github.com/MetaCubeX/Clash.Mini/log"
	"github.com/MetaCubeX/Clash.Mini/static"
	_ "github.com/MetaCubeX/Clash.Mini/static"
	fileUtils "github.com/MetaCubeX/Clash.Mini/util/file"
	httpUtils "github.com/MetaCubeX/Clash.Mini/util/http"
	stringUtils "github.com/MetaCubeX/Clash.Mini/util/string"

	"github.com/fsnotify/fsnotify"
	stx "github.com/getlantern/systray"
)

const (
	logHeader = "profile"
)

type Info struct {
	Name       string
	FileSize   string
	UpdateTime time.Time
	Url        string

	FileHash string
	Enabled  bool
}

var (
	Profiles   = new(list.List)
	RawDataMap sync.Map
	Locker     = new(sync.Mutex)

	// for fs watcher
	watcherDataMap = make(map[string]*WatcherData)

	TagRegexp = regexp.MustCompile(`# Clash.Mini : (http.*)`)
)

type WatcherData struct {
	once     *sync.Once
	onceData fsnotify.Event
	timer    *time.Timer
}

type RawData struct {
	FileInfo         *Info
	FileInfoListElem *list.Element
	MenuItemEx       *stx.MenuItemEx
}

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

		eventsWG := sync.WaitGroup{}
		eventsWG.Add(1)
		go func() {
			for {
				select {
				case event, ok := <-watcher.Events:
					if !ok {
						return
					}
					log.Infoln("[profile] watcher event: %v", event)
					if event.Op&(fsnotify.Write|fsnotify.Remove|fsnotify.Rename) != 0 {
						//event := event
						log.Infoln("[profile] modified file: %s", event.Name)
						log.Warnln("[profile] waiting 100ms and even multi changes will be once in 50ms: %s", event.Name)

						var watcherData *WatcherData
						var exists bool
						if watcherData, exists = watcherDataMap[event.Name]; !exists {
							watcherData = &WatcherData{
								once:     new(sync.Once),
								onceData: event,
								timer: time.AfterFunc(50*time.Millisecond,
									func() {
										if v, exists := watcherDataMap[event.Name]; exists {
											v.timer.Reset(50 * time.Millisecond)
										} else {
											delete(watcherDataMap, event.Name)
										}
									}),
							}
							watcherDataMap[event.Name] = watcherData
						} else {
							watcherData.timer.Reset(50 * time.Millisecond)
						}
						watcherData.once.Do(func() {
							time.AfterFunc(100*time.Millisecond, func() {
								common.RefreshProfile(&(watcherDataMap[event.Name].onceData))
								delete(watcherDataMap, event.Name)
							})
						})
					}
				case err, ok := <-watcher.Errors:
					if !ok {
						return
					}
					log.Errorln("[profile] watcher error: %v", err)
					return
				}
			}
		}()
		err = watcher.Add(constant.ProfileDir)
		if err != nil {
			log.Errorln("[profile] watch profile dir error: %v", err)
			return
		}
		eventsWG.Wait()
	}()

}

func RemoveProfile(name string) (exists bool) {
	original, exists := RawDataMap.LoadAndDelete(name)
	if exists {
		rawData := original.(*RawData)
		if rawData.MenuItemEx != nil {
			log.Infoln("[profile] removed: %s", name)
			original.(*RawData).MenuItemEx.Delete()
		}
		if rawData.FileInfoListElem != nil {
			Profiles.Remove(original.(*RawData).FileInfoListElem)
		}
		if rawData.FileInfo != nil {
			Profiles.Remove(original.(*RawData).FileInfoListElem)
		}
	}
	return exists
}

func RefreshProfiles(event *fsnotify.Event) {
	defer func() {
		Locker.Unlock()
	}()
	Locker.Lock()

	var err error
	var fileInfos []fs.FileInfo
	var isRemove, isRename bool
	if event == nil {
		fileInfos, err = ioutil.ReadDir(constant.ProfileDir)
		if err != nil {
			log.Errorln("[profile] RefreshProfiles ReadDir error: %v", err)
			return
		}
	} else {
		fileInfos = []fs.FileInfo{static.NewFakeFile(event.Name)}
		isRemove = event.Op|fsnotify.Remove == fsnotify.Remove
		isRename = event.Op|fsnotify.Rename == fsnotify.Rename
	}
	//var profiles []*Info
	for _, f := range fileInfos {
		extName := path.Ext(f.Name())
		if extName != constant.ConfigSuffix {
			continue
		}
		fileName := f.Name()
		profileName := f.Name()
		if path.IsAbs(profileName) {
			profileName = path.Base(profileName)
		} else {
			fileName = path.Join(constant.ProfileDir, fileName)
		}
		profileName = GetConfigName(profileName)
		if isRemove || isRename {
			RemoveProfile(profileName)
		} else {
			profile := &Info{
				Name:       profileName,
				FileSize:   fileUtils.FormatHumanizedFileSize(f.Size()),
				UpdateTime: f.ModTime(),
			}
			content, err := os.OpenFile(fileName, os.O_RDWR, 0666)
			if err != nil {
				log.Errorln("[profile] RefreshProfiles OpenFile error: %v", err)
				continue
			}
			reader := bufio.NewReader(content)
			v, exists := RawDataMap.Load(profileName)
			log.Infoln("[profile] profile <%s> is %s", profileName,
				stringUtils.TrinocularString(exists, "exists", "not exists"))
			hash := fileUtils.GetHash(reader, 32)

			var original *RawData
			if exists {
				original = v.(*RawData)
				if original.FileInfo != nil && original.FileInfo.FileHash == hash {
					continue
				}
			} else {
				original = &RawData{
					FileInfo: profile,
				}
			}
			profile.FileHash = hash
			lineData, _, err := reader.ReadLine()
			if err != nil {
				log.Errorln("[profile] RefreshProfiles ReadLine error: %v", err)
				continue
			}
			profile.Url = GetTagLineUrl(string(lineData))
			if err = content.Close(); err != nil {
				continue
			}
			if !exists {
				original.FileInfoListElem = Profiles.PushBack(original)
				RawDataMap.Store(profileName, original)
			}
		}
	}
}

// UpdateConfig 更新订阅配置
func UpdateConfig(name, url string) (successful bool) {
	client := &http.Client{Timeout: 20 * time.Second}
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
		_, err = f.WriteString(fmt.Sprintf("# Clash.Mini : %s\n\n", url))
		if err != nil {
			log.Errorln("[profile] writeString error: %v", err)
			return false
		}

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
	if TagRegexp.MatchString(line) {
		return TagRegexp.FindStringSubmatch(line)[1]
	}
	return ""
}

func GetConfigName(fileName string) string {
	startIdx := strings.LastIndexByte(fileName, os.PathSeparator)
	endIdx := strings.LastIndex(fileName, constant.ConfigSuffix)
	if startIdx > -1 && endIdx >= startIdx {
		return fileName[startIdx+1 : endIdx]
	}
	return strings.TrimSuffix(fileName, constant.ConfigSuffix)
}
