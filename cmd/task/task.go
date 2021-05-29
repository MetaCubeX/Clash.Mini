package task

import "github.com/Clash-Mini/Clash.Mini/cmd"

type Type int8

const (
	ON Type = iota
	OFF

	Invalid = -1
)

var (
	typeMap = map[Type]string{
		ON:  "ON",
		OFF: "OFF",
	}
)

func (t Type) String() string {
	return typeMap[t]
}

func (t Type) GetCommandType() cmd.CommandType {
	return cmd.Task
}

func ParseType(s string) Type {
	for typeEnum, typeName := range typeMap {
		if s == typeName {
			return typeEnum
		}
	}
	return Invalid
}

func (t Type) IsValid() bool {
	return t != Invalid
}

func IsValid(s string) bool {
	return ParseType(s).IsValid()
}

func (t Type) IsON() bool {
	return t == ON
}
