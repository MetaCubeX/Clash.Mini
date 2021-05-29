package cmd

type CommandType string

const (
	Task CommandType = "Task"
	Sys              = "Sys"
	MMDB             = "MMDB"
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
