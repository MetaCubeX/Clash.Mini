package controller

import (
	"github.com/Clash-Mini/Clash.Mini/config"
	"github.com/Clash-Mini/Clash.Mini/constant"
	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Clash-Mini/Clash.Mini/mixin"
	. "github.com/Dreamacro/clash/config"
	"github.com/Dreamacro/clash/hub/executor"
	path "path/filepath"
)

func ParseMixin(RawPath *Config) (MixinRaw *Config, err error) {
	MixinRaw = RawPath
	if config.IsMixinPositive(mixin.Dns) && !config.IsMixinPositive(mixin.Tun) {
		DnsPath, err := executor.ParseWithPath(path.Join(constant.MixinDir, constant.MixDnsFile))
		if err != nil {
			log.Errorln("[%s] PutConfig ParseMixinDnsConfig error: %v", profileInfoLogHeader, err)
			return RawPath, err
		}
		MixinRaw.DNS = DnsPath.DNS
	}

	if config.IsMixinPositive(mixin.Tun) {
		TunPath, err := executor.ParseWithPath(path.Join(constant.MixinDir, constant.MixTunFile))
		if err != nil {
			log.Errorln("[%s] PutConfig ParseMixinTunConfig error: %v", profileInfoLogHeader, err)
			return RawPath, err
		}
		DnsPath, err := executor.ParseWithPath(path.Join(constant.MixinDir, constant.MixDnsFile))
		if err != nil {
			log.Errorln("[%s] PutConfig ParseMixinDnsConfig error: %v", profileInfoLogHeader, err)
			return RawPath, err
		}
		MixinRaw.Tun = TunPath.Tun
		MixinRaw.DNS = DnsPath.DNS
	}
	return MixinRaw, nil
}
