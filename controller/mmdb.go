package controller

import (
	"github.com/Clash-Mini/Clash.Mini/cmd/mmdb"
	"github.com/Clash-Mini/Clash.Mini/config"
	"github.com/Clash-Mini/Clash.Mini/constant"
	"github.com/Clash-Mini/Clash.Mini/util"

	"github.com/lxn/walk"
)

var (
	mmdbMap = map[mmdb.Type]string {
		mmdb.Max: "Dreamacro/maxmind-geoip",
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
		util.AppendStringTo(&url, constant.GitHubCDN, value, constant.MMDBSuffix)
	}

	err := util.DownloadFile(url, util.GetPwdPath(constant.MmdbFile))
	if err != nil {
		walk.MsgBox(nil, constant.UIConfigMsgTitle, err.Error(), walk.MsgBoxIconError)
		return
	}
	config.SetCmd(value)
}

