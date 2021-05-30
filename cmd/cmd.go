package cmd

type CommandType string

const (
	Task CommandType = "Task"
	Sys              = "Sys"
	MMDB             = "MMDB"
	Cron             = "Cron"

	// TODO: extract general value

	ON      = 0
	OFF     = 1
	Invalid = -1

	OnName  = "ON"
	OffName = "OFF"
)

type GeneralType interface {
	String() string
	GetCommandType() CommandType
	IsON() bool
}

func (t CommandType) GetName() string {
	return string(t)
}

func (t CommandType) IsValid(value GeneralType) bool {
	return t == value.GetCommandType()
}
