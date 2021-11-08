package loopback

import (
	"fmt"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
	"syscall"
	"time"

	"github.com/Clash-Mini/Clash.Mini/log"
	stringUtils "github.com/Clash-Mini/Clash.Mini/util/string"
)

const (
	rate 					= 2 * time.Second
	appContainerMappingKey	= registry.CURRENT_USER
	appContainerMappingPath	= `Software\Classes\Local Settings\Software\Microsoft\Windows\CurrentVersion\AppContainer\Mappings`

	appDisplayNamPath		= "DisplayName"
	appMonikerPath			= "Moniker"
)

var (
	watcherTicker			*time.Ticker

	verbPtr					*uint16
	exePtr					*uint16
)

func init() {
	verbPtr, _ 	= syscall.UTF16PtrFromString("")
	// TODO: lookup exe
	exePtr, _ 	= syscall.UTF16PtrFromString("CheckNetIsolation")
}

// TODO: load check

func enableLoopback(appIDs []string, enable bool) {
	for _, id := range appIDs {
		argPtr, _ := syscall.UTF16PtrFromString(fmt.Sprintf("LoopbackExempt -%s -p=%s",
			stringUtils.TrinocularString(enable, "a", "d"), id))

		err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, nil, 0)
		log.Infoln("[loopback] enableLoopback: %s", id)
		if err != nil {
			log.Errorln("Cmd exec failed: %s", err)
		}
	}
}

func StartBreaker() {
	if watcherTicker != nil {
		return
	}
	log.Infoln("[loopback] Loopback Breaker is starting...")
	watcherTicker = time.NewTicker(rate)
	go func() {
		k, err := registry.OpenKey(appContainerMappingKey, appContainerMappingPath, registry.READ)
		if err != nil {
			log.Errorln("[loopback] openKey failed: %s", err.Error())
			deleteTicker()
			return
		}
		defer func(k registry.Key) {
			err := k.Close()
			if err != nil {
				log.Errorln("[loopback] closeKey failed: %s", err.Error())
				deleteTicker()
			}
		}(k)

		for i:=0; true; i++ {
			select {
			case <-watcherTicker.C:
				//log.Infoln("Checking...")
				stat, err := k.Stat()
				if i > 0 && (err != nil || time.Since(stat.ModTime()) > rate) {
					continue
				}
				appIDs, err := k.ReadSubKeyNames(0)
				log.Infoln("[loopback] detected UWP %d app(s)", len(appIDs))
				//apps := getUwpApps(appIDs)
				//log.Infoln("[loopback] detected UWP apps: %s", stringUtils.JoinString(", ", apps...))
				if err != nil {
					log.Errorln("[loopback] readSubKey failed: %s", err.Error())
				}
				go enableLoopback(appIDs, true)
			}
		}
	}()
}

func StopBreaker() {
	if watcherTicker == nil {
		return
	}
	log.Infoln("[loopback] Loopback Breaker is stopping...")
	deleteTicker()
	go func() {
		k, err := registry.OpenKey(appContainerMappingKey, appContainerMappingPath, registry.READ)
		if err != nil {
			log.Errorln("[loopback] openKey failed: %s", err.Error())
			return
		}
		defer func(k registry.Key) {
			err := k.Close()
			if err != nil {
				log.Errorln("[loopback] closeKey failed: %s", err.Error())
			}
		}(k)

		_, err = k.Stat()
		if err != nil {
			log.Errorln("[loopback] statKey failed: %s", err.Error())
			return
		}
		appIDs, err := k.ReadSubKeyNames(0)
		log.Infoln("[loopback] delete UWP %d app(s)", len(appIDs))
		//apps := getUwpApps(appIDs)
		//log.Infoln("[loopback] delete UWPs: %s [%s]", stringUtils.JoinString(", ", apps...))
		if err != nil {
			log.Errorln("[loopback] readSubKey failed: %s", err.Error())
		}
		go enableLoopback(appIDs, false)
	}()

}

func getUwpApps(appIDs []string) (apps []string) {
	for _, sid := range appIDs {
		key, err := registry.OpenKey(appContainerMappingKey, fmt.Sprintf(`%s\%s`, appContainerMappingPath, sid), registry.READ)
		if err != nil {
			continue
		}
		appDisplayName, _, err := key.GetStringValue(appDisplayNamPath)
		if err != nil {
			continue
		}
		appMoniker, _, err := key.GetStringValue(appMonikerPath)
		if err != nil {
			continue
		}
		apps = append(apps, fmt.Sprintf("(%s <%s> [%s])", appDisplayName, appMoniker, sid))
	}
	return
}

func deleteTicker() {
	watcherTicker.Stop()
	watcherTicker = nil
}
