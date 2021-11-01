package mmdb

import (
	"strings"

	"github.com/Clash-Mini/Clash.Mini/cmd"
)

// MMDB

type Type string

const (
	Lite 	Type = "lite"
	Max		Type = "max"

	Invalid Type = ""
)

var (
	typeMap = map[string]Type{
		Lite.String():	Lite,
		Max.String():	Max,
	}
)

// String implements cmd.GeneralType
func (t Type) String() string {
	return string(t)
}

// GetCommandType implements cmd.GeneralType
func (t Type) GetCommandType() cmd.CommandType {
	return cmd.MMDB
}

// GetDefault implements cmd.GeneralType
func (t Type) GetDefault() cmd.GeneralType {
	return Max
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
	return t == Lite
}
