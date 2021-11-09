package proxy

import (
	"strings"

	"github.com/Clash-Mini/Clash.Mini/cmd"
)

// 代理模式

type Type string

const (
	Direct 	Type = "direct"
	Rule	Type = "rule"
	Global 	Type = "global"

	Invalid Type = ""
)

var (
	typeMap = map[string]Type{
		Direct.String(): Direct,
		Rule.String():   Rule,
		Global.String(): Global,
	}
)

// String implements cmd.GeneralType
func (t Type) String() string {
	return string(t)
}

// GetCommandType implements cmd.GeneralType
func (t Type) GetCommandType() cmd.CommandType {
	return cmd.Sys
}

// GetDefault implements cmd.GeneralType
func (t Type) GetDefault() cmd.GeneralType {
	return Rule
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
	return t == Rule || t == Global
}
