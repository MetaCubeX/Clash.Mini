package cmd

import (
	"strings"
)

type Type string

const (
	ON      Type = "on"
	OFF     Type = "off"
)

var (
	typeMap = map[string]Type{
		ON.String():  ON,
		OFF.String(): OFF,
	}
)

// String implements cmd.GeneralType
func (t Type) String() string {
	return string(t)
}

// GetCommandType implements cmd.GeneralType
func (t Type) GetCommandType() CommandType {
	return ""
}

// GetDefault implements cmd.GeneralType
func (t Type) GetDefault() GeneralType {
	return OFF
}

func (t Type) IsValid() bool {
	return t.String() != ""
}

// IsPositive implements cmd.GeneralType
func (t Type) IsPositive() bool {
	return t == ON
}

//func FindInMap(t Type) string {
//	return typeMap[t.String()].String()
//}

func ParseType(s string) Type {
	typeEnum, ok := typeMap[s]
	if !ok {
		return ""
	}
	return typeEnum
}

func ParseTypeWeak(s string) Type {
	s = strings.ToLower(s)
	return ParseType(s)
}
