package parser

import (
	"github.com/Clash-Mini/Clash.Mini/cmd"
	"github.com/Clash-Mini/Clash.Mini/cmd/auto"
	"github.com/Clash-Mini/Clash.Mini/cmd/cron"
	"github.com/Clash-Mini/Clash.Mini/cmd/mmdb"
	"github.com/Clash-Mini/Clash.Mini/cmd/proxy"
	"github.com/Clash-Mini/Clash.Mini/cmd/startup"
	"github.com/Clash-Mini/Clash.Mini/cmd/sys"
	"github.com/Clash-Mini/Clash.Mini/cmd/task"
	"github.com/Clash-Mini/Clash.Mini/log"
)

func GetCmdDefaultValue(command cmd.CommandType, value string) (defaultValue cmd.GeneralType) {
	defaultValue = GetCmdValue(command, value)
	if defaultValue != cmd.Invalid {
		defaultValue = defaultValue.GetDefault()
	}
	return
}

func GetCmdValue(command cmd.CommandType, value string) cmd.GeneralType {
	switch command {
	case cmd.Task:
		return task.ParseType(value)
	case cmd.Sys:
		return sys.ParseType(value)
	case cmd.MMDB:
		return mmdb.ParseType(value)
	case cmd.Cron:
		return cron.ParseType(value)
	case cmd.Proxy:
		return proxy.ParseType(value)
	case cmd.Startup:
		return startup.ParseType(value)
	case cmd.Auto:
		return auto.ParseType(value)
	default:
		log.Errorln("command \"%s\" is not support\n", command)
		return cmd.Invalid
	}
}
