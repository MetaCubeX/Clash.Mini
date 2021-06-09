package constant

import (
	"os"
	path "path/filepath"
	"runtime"
	"time"

	"github.com/Clash-Mini/Clash.Mini/log"
)

const (
	ConfigFile   = "config.yaml"
	ConfigSuffix = ".yaml"
	CacheFile    = ".cache"

	Localhost      = "127.0.0.1"
	ControllerPort = "9090"
	DashboardPort  = "8070"

	NotifyDelay = 2 * time.Second

	GitHubCDN  = "https://cdn.jsdelivr.net/gh/"
	MMDBSuffix = "@release/Country.mmdb"

	UIConfigMsgTitle = "配置提示"

	SubConverterUrl = "https://id9.cc"
)

var (
	PWD       string
	ConfigDir = "profile"
	CacheDir  = "cache"

	osWindows bool
)

func init() {
	var err error
	PWD, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	ConfigDir = path.Join(PWD, ConfigDir)
	CacheDir = path.Join(PWD, CacheDir)
	osWindows = runtime.GOOS == "windows"
	if _, err := os.Stat(ConfigDir); err != nil {
		if os.IsNotExist(err) {
			if err = os.Mkdir(ConfigDir, 0666); err != nil {
				log.Fatalln("cannot create config dir: %v", err)
			}
		}
	}
	if _, err := os.Stat(CacheDir); err != nil {
		if os.IsNotExist(err) {
			if err = os.Mkdir(CacheDir, 0666); err != nil {
				log.Fatalln("cannot create cache dir: %v", err)
			}
		}
	}
}

func IsWindows() bool {
	return osWindows
}
