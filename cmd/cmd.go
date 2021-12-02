package cmd

// 命令类型

type CommandType string

const (
	Task     CommandType = "task"
	Sys      CommandType = "sys"
	Autosys  CommandType = "autosys"
	MMDB     CommandType = "mmdb"
	Cron     CommandType = "cron"
	Proxy    CommandType = "proxy"
	Startup  CommandType = "startup"
	Breaker  CommandType = "breaker"
	Protocol CommandType = "protocol"
	Hotkey   CommandType = "hotkey"
	Invalid  Type        = ""
)

// GeneralType 通用类型
type GeneralType interface {

	// String 字符串
	String() string

	// GetCommandType 获取命令类型
	GetCommandType() CommandType

	// GetDefault 获取该类型默认值
	GetDefault() GeneralType

	// IsPositive 是否为活动值
	IsPositive() bool

	// IsValid 是否为有效值
	IsValid() bool
}

// GetName 获取命令名
func (t CommandType) GetName() string {
	return string(t)
}

// IsValid 是否为有效命令
func (t CommandType) IsValid(value GeneralType) bool {
	return t == value.GetCommandType()
	//return t != Invalid && t == value.GetCommandType()
}
