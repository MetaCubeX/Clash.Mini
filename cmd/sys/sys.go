package sys

import (
	"github.com/Clash-Mini/Clash.Mini/cmd"
)

type Type int8

const (
	ON Type = iota + 5
	OFF

	Invalid = -1
)

var (
	typeMap = map[Type]string{
		ON:  cmd.OnName,
		OFF: cmd.OffName,
	}
)

// String implements cmd.GeneralType
func (t Type) String() string {
	return typeMap[t]
}

// GetCommandType implements cmd.GeneralType
func (t Type) GetCommandType() cmd.CommandType {
	return cmd.Sys
}

// GetDefault implements cmd.GeneralType
func (t Type) GetDefault() cmd.GeneralType {
	return OFF
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

// IsON implements cmd.GeneralType
func (t Type) IsON() bool {
	return t == ON
}
