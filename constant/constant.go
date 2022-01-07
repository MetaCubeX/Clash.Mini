package constant

import (
	"net"
	path "path/filepath"
	"strings"
	"time"

	cConfig "github.com/MetaCubeX/Clash.Mini/constant/config"
	commonUtils "github.com/MetaCubeX/Clash.Mini/util/common"
)

const (
	ConfigSuffix   = ".yaml"
	ConfigFile     = "config.yaml"
	MixGeneralFile = "general.yaml"
	MixTunFile     = "tun.yaml"
	MixDnsFile     = "dns.yaml"

	CacheFile = "cache.db"
	MmdbFile  = "Country.mmdb"

	DashboardPort = "8070"

	NotifyDelay = 2 * time.Second
	LocalHost   = "127.0.0.1"
	GitHubCDN   = "https://cdn.jsdelivr.net/gh/"
	MMDBSuffix  = "@release/" + MmdbFile

	UIConfigMsgTitle = "配置提示"

	SubConverterUrl = "https://id9.cc"
)

var (
	ControllerPort   = "9090"
	ControllerHost   = "127.0.0.1"
	ControllerSecret = ""
	Pwd              = commonUtils.GetPwdPath()
	Executable       = commonUtils.GetExecutable()
	ExecutableDir    = commonUtils.GetExecutablePath()
	ProfileDir       = ".core/.profile"
	CacheDir         = ".core/.cache"
	MixinDir         = ".core/.mixin"
	TaskFile         = "task.xml"
)

func init() {
	cConfig.DirPath = commonUtils.GetExecutablePath(cConfig.DirPath)
	cConfig.DashboardDir = commonUtils.GetExecutablePath(cConfig.DashboardDir)
	ProfileDir = commonUtils.GetExecutablePath(ProfileDir)
	CacheDir = commonUtils.GetExecutablePath(CacheDir)
	MixinDir = commonUtils.GetExecutablePath(MixinDir)
	TaskFile = path.Join(cConfig.DirPath, TaskFile)
}

func SetController(address, secret string) {
	result := strings.Split(address, ":")
	if len(result) == 0 {
		ControllerHost = "127.0.0.1"
		ControllerPort = "9090"
	} else if len(result) == 1 {
		ControllerHost = strings.TrimSpace(result[0])
		ControllerPort = "9090"
	} else if len(strings.TrimSpace(result[0])) == 0 {
		ControllerHost = "127.0.0.1"
		ControllerPort = result[1]
	} else {
		if !net.ParseIP(result[0]).Equal(net.IPv4zero) {
			ControllerHost = result[0]
		}

		ControllerPort = result[1]
	}
	ControllerSecret = secret
}
