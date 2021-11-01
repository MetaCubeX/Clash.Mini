package util

import (
	"container/list"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/Clash-Mini/Clash.Mini/log"
)

var (
	boolValuesMap = map[interface{}]interface{}{
		"true":  true,
		"false": false,
		"1":     true,
		"0":     false,
	}
)

// TODO: support custom UnmarshalOption

// ConvertForceByJson 通过JSON强制转换
func ConvertForceByJson(dv interface{}, ov interface{}) (err error) {
	jsonString, _ := json.Marshal(ov)
	if err = json.Unmarshal(jsonString, dv); err != nil {
		return
	}
	return
}

// unmarshalValues 解码为UrlValues
func unmarshalValues(str string) (subInfoMap url.Values, err error) {
	subInfoMap, err = url.ParseQuery(str)
	var trimField string
	for field, _ := range subInfoMap {
		trimField = strings.TrimSpace(field)
		if trimField != field {
			subInfoMap.Add(trimField, subInfoMap.Get(field))
			subInfoMap.Del(field)
		}
	}
	return subInfoMap, err
}

// UnmarshalByValues 解码为struct
func UnmarshalByValues(str string, v interface{}) error {
	return UnmarshalByValuesWithTag(str, "query", v)
}

// UnmarshalByValuesWithTag 解码为struct
func UnmarshalByValuesWithTag(str string, fieldTag string, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		return fmt.Errorf("unmarshal non-pointer \"%s\"", rv.Type().String())
	}
	if rv.IsNil() {
		return fmt.Errorf("unmarshall by reflect failed, because the interface ptr is nil")
	}

	subInfoMap, err := unmarshalValues(str)
	if err != nil {
		return err
	}
	rv = rv.Elem()
	isInterface := rv.Kind() == reflect.Interface
	var interfaceElem reflect.Value
	if isInterface {
		interfaceElem = rv
		rv = reflect.New(rv.Elem().Type()).Elem()
		rv.Set(interfaceElem.Elem())
	}
	rvt := rv.Type()
	fieldNum := rv.NumField()

	for i := 0; i < fieldNum; i++ {
		rvf := rvt.Field(i)

		var isOmitempty bool
		var tagEx string
		var tags []string
		if len(fieldTag) > 0 {
			tagEx = rvf.Tag.Get(fieldTag)
			tags = strings.Split(tagEx, ",")
			if len(tags) > 1 {
				isOmitempty = tags[len(tags)-1] == "omitempty"
				if isOmitempty {
					tags = tags[:len(tags)-1]
				}
			}
		}
		fieldName := rvf.Name
		rfv := rv.Field(i)
		if tags == nil || len(tags) == 0 {
			tags = []string{rvf.Name}
		}
		var fieldVal []string
		for _, tag := range tags {
			fieldVal = subInfoMap[tag]
			if fieldVal == nil {
				continue
			}
			log.Debugln("%s %s(%s)=%s [%s]", rfv.Kind(), fieldName, tag, fieldVal, isOmitempty)
		}
		if fieldVal == nil || len(fieldVal) < 1 {
			continue
		}
		switch rfv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			intVal, err := strconv.ParseInt(fieldVal[0], 10, 64)
			if err != nil {
				return err
			}
			rfv.SetInt(intVal)
			break
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			intVal, err := strconv.ParseUint(fieldVal[0], 10, 64)
			if err != nil {
				return err
			}
			rfv.SetUint(intVal)
			break
		case reflect.Bool:
			boolValue := getReflectValue(fieldVal[0], boolValuesMap)
			if boolValue == nil {
				return fmt.Errorf("unmarshall by reflect failed, field \"%s\" kind \"bool\" must be [true, false], but it's \"%s\"", fieldName, fieldVal)
			}
			rfv.SetBool(boolValue.(bool))
			break
		case reflect.String:
			rfv.SetString(fieldVal[0])
			break
		case reflect.Struct, reflect.Interface:
			// TODO: use recursion
			return fmt.Errorf("unmarshall by reflect failed, field \"%s\" kind \"%s\" is not support", fieldName, rfv.Kind())
		case reflect.Slice, reflect.Array, reflect.TypeOf(list.List{}).Kind():
			// TODO: use recursion inside loop
			return fmt.Errorf("unmarshall by reflect failed, field \"%s\" kind \"%s\" is not support", fieldName, rfv.Kind())
		default:
			rfv.Set(reflect.ValueOf(fieldVal))
		}
	}
	if isInterface {
		interfaceElem.Set(rv)
	}
	return nil
}

// getReflectValue 获取指定反射类型的映射值
func getReflectValue(v interface{}, valuesMap interface{}) interface{} {
	for key, value := range valuesMap.(map[interface{}]interface{}) {
		if key == v {
			return value
		}
	}
	return nil
}

// ToJsonString struct转为JSON字符串
func ToJsonString(v interface{}) string {
	jsonBytes, _ := json.MarshalIndent(v, "", "\t")
	return string(jsonBytes)
}

// JsonUnmarshal JSON字节数组转struct
func JsonUnmarshal(data []byte, v interface{}) {
	if err := json.Unmarshal(data, v); err != nil {
		log.Errorln("JsonUnmarshal error: %v", err)
	}
}
