package parser

import (
	"github.com/MetaCubeX/Clash.Mini/cmd"
	"github.com/MetaCubeX/Clash.Mini/cmd/autosys"
	"github.com/MetaCubeX/Clash.Mini/cmd/breaker"
	"github.com/MetaCubeX/Clash.Mini/cmd/cron"
	"github.com/MetaCubeX/Clash.Mini/cmd/hotkey"
	"github.com/MetaCubeX/Clash.Mini/cmd/mmdb"
	"github.com/MetaCubeX/Clash.Mini/cmd/protocol"
	"github.com/MetaCubeX/Clash.Mini/cmd/proxy"
	"github.com/MetaCubeX/Clash.Mini/cmd/startup"
	"github.com/MetaCubeX/Clash.Mini/cmd/task"
	"github.com/MetaCubeX/Clash.Mini/log"
	"github.com/MetaCubeX/Clash.Mini/mixin"
	"github.com/MetaCubeX/Clash.Mini/mixin/dns"
	"github.com/MetaCubeX/Clash.Mini/mixin/general"
	"github.com/MetaCubeX/Clash.Mini/mixin/tun"
)

const (
	logHeader = "cmd.parser"
)

// GetCmdOrDefaultValue 获取命令值或默认值
func GetCmdOrDefaultValue(command cmd.CommandType, defaultValue string) (value cmd.GeneralType) {
	value = GetCmdValue(command, defaultValue)
	if !value.IsValid() {
		value = value.GetDefault()
	}
	return
}

func GetMixinOrDefaultValue(command mixin.CommandType, defaultValue string) (value mixin.GeneralType) {
	value = GetMixinValue(command, defaultValue)
	if !value.IsValid() {
		value = value.GetDefault()
	}
	return
}

// GetCmdValue 获取命令值
func GetCmdValue(command cmd.CommandType, value string) cmd.GeneralType {
	switch command {
	case cmd.Task:
		return task.ParseType(value)
	case cmd.Protocol:
		return protocol.ParseType(value)
	case cmd.Autosys:
		return autosys.ParseType(value)
	case cmd.MMDB:
		return mmdb.ParseType(value)
	case cmd.Cron:
		return cron.ParseType(value)
	case cmd.Proxy:
		return proxy.ParseType(value)
	case cmd.Startup:
		return startup.ParseType(value)
	case cmd.Breaker:
		return breaker.ParseType(value)
	case cmd.Hotkey:
		return hotkey.ParseType(value)
	default:
		log.Errorln("[%s] command \"%s\" is not support\n", logHeader, command)
		return cmd.Invalid
	}
}

func GetMixinValue(command mixin.CommandType, value string) mixin.GeneralType {
	switch command {
	case mixin.General:
		return general.ParseType(value)
	case mixin.Tun:
		return tun.ParseType(value)
	case mixin.Dns:
		return dns.ParseType(value)

	default:
		log.Errorln("[%s] command \"%s\" is not support\n", logHeader, command)
		return mixin.Invalid
	}
}
