package script

import (
	"github.com/Clash-Mini/Clash.Mini/mixin"
	"strings"
)

type Type string

const (
	ON  Type = "on"
	OFF Type = "off"

	Invalid Type = ""
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
func (t Type) GetCommandType() mixin.CommandType {
	return mixin.Script
}

// GetDefault implements cmd.GeneralType
func (t Type) GetDefault() mixin.GeneralType {
	return OFF
}

func ParseType(s string) Type {
	typeEnum, ok := typeMap[s]
	if !ok {
		return Invalid
	}
	return typeEnum
}

func ParseTypeWeak(s string) Type {
	s = strings.ToLower(s)
	return ParseType(s)
}

func (t Type) IsValid() bool {
	return t != Invalid && string(t) != ""
}

func IsValid(s string) bool {
	return ParseType(s).IsValid()
}

// IsPositive implements cmd.GeneralType
func (t Type) IsPositive() bool {
	return t == ON
}
