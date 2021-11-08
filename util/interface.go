package util

import (
	"reflect"
)

func EqualsAny(compareV interface{}, vs... interface{}) bool {
	t := reflect.TypeOf(compareV)
	for _, v := range vs {
		if t != reflect.TypeOf(v) {
			continue
		}
		if compareV == v {
			return true
		}
	}
	return false
}
