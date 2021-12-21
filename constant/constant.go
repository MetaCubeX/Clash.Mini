package constant

import (
	path "path/filepath"
	"time"

	cConfig "github.com/Clash-Mini/Clash.Mini/constant/config"
	commonUtils "github.com/Clash-Mini/Clash.Mini/util/common"
)

const (
	ConfigSuffix   = ".yaml"
	ConfigFile     = "config.yaml"
	MixGeneralFile = "general.yaml"
	MixTunFile     = "tun.yaml"
	MixDnsFile     = "dns.yaml"

	CacheFile = "cache.db"
	MmdbFile  = "Country.mmdb"

	Localhost      = "127.0.0.1"
	ControllerPort = "9090"
	DashboardPort  = "8070"

	NotifyDelay = 2 * time.Second

	GitHubCDN  = "https://cdn.jsdelivr.net/gh/"
	MMDBSuffix = "@release/" + MmdbFile

	UIConfigMsgTitle = "配置提示"

	SubConverterUrl = "https://id9.cc"
)

var (
	Pwd           = commonUtils.GetPwdPath()
	Executable    = commonUtils.GetExecutable()
	ExecutableDir = commonUtils.GetExecutablePath()
	ProfileDir    = ".core/.profile"
	CacheDir      = ".core/.cache"
	MixinDir      = ".core/.mixin"
	TaskFile      = "task.xml"
)

func init() {
	cConfig.DirPath = commonUtils.GetExecutablePath(cConfig.DirPath)
	cConfig.DashboardDir = commonUtils.GetExecutablePath(cConfig.DashboardDir)
	ProfileDir = commonUtils.GetExecutablePath(ProfileDir)
	CacheDir = commonUtils.GetExecutablePath(CacheDir)
	MixinDir = commonUtils.GetExecutablePath(MixinDir)
	TaskFile = path.Join(cConfig.DirPath, TaskFile)
}
