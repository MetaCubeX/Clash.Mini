package constant

import (
	"os"
	path "path/filepath"
	"runtime"
	"time"
)

const (
	ConfigFile   = "config.yaml"
	ConfigSuffix = ".yaml"

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

	osWindows bool
)

func init() {
	var err error
	PWD, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	ConfigDir = path.Join(".", ConfigDir)
	osWindows = runtime.GOOS == "windows"
}

func IsWindows() bool {
	return osWindows
}
