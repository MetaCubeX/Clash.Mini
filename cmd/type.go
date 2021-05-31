package cmd

type Type int8

var (
	typeMap = map[Type]string{
		ON:  OnName,
		OFF: OffName,
	}
)

// String implements cmd.GeneralType
func (t Type) String() string {
	return typeMap[t]
}

// GetCommandType implements cmd.GeneralType
func (t Type) GetCommandType() CommandType {
	return Sys
}

// GetDefault implements cmd.GeneralType
func (t Type) GetDefault() GeneralType {
	return OFF
}

func (t Type) IsValid() bool {
	return t != Invalid
}

// IsON implements cmd.GeneralType
func (t Type) IsON() bool {
	return t == ON
}
