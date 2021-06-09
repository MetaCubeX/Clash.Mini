package cmd

type CommandType string

const (
	Task    CommandType = "Task"
	Sys                 = "Sys"
	MMDB                = "MMDB"
	Cron                = "Cron"
	Proxy               = "Proxy"
	Startup             = "Startup"
	Auto                = "Auto"

	OnName  = "ON"
	OffName = "OFF"

	// TODO: extract general value

	ON      Type = 0
	OFF     Type = 1
	Invalid Type = -1
)

type GeneralType interface {
	String() string
	GetCommandType() CommandType
	GetDefault() GeneralType
	IsON() bool
}

func (t CommandType) GetName() string {
	return string(t)
}

func (t CommandType) IsValid(value GeneralType) bool {
	return t == value.GetCommandType()
}
