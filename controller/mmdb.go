package controller

import (
	"github.com/MetaCubeX/Clash.Mini/cmd/mmdb"
	"github.com/MetaCubeX/Clash.Mini/config"
	"github.com/MetaCubeX/Clash.Mini/constant"
	commonUtils "github.com/MetaCubeX/Clash.Mini/util/common"
	"github.com/MetaCubeX/Clash.Mini/util/http"
	stringUtils "github.com/MetaCubeX/Clash.Mini/util/string"

	"github.com/lxn/walk"
)

var (
	mmdbMap = map[mmdb.Type]string{
		mmdb.Max:  "Dreamacro/maxmind-geoip",
		mmdb.Lite: "Hackl0us/GeoIP2-CN",
	}
)

func GetMMDB(value mmdb.Type) {
	// 从字典中获取MMDB仓库
	var url string
	if value, ok := mmdbMap[value]; !ok {
		// 不存在返回
		return
	} else {
		// 存在时拼接
		stringUtils.AppendStringTo(&url, constant.GitHubCDN, value, constant.MMDBSuffix)
	}

	err := http.DownloadFile(url, commonUtils.GetExecutablePath(constant.MmdbFile))
	if err != nil {
		walk.MsgBox(nil, constant.UIConfigMsgTitle, err.Error(), walk.MsgBoxIconError)
		return
	}
	config.SetCmd(value)
}
