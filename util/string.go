package util

import (
	"fmt"
	"strings"

	"github.com/Clash-Mini/Clash.Mini/app"
	"github.com/Clash-Mini/Clash.Mini/log"
)

func GetSubTitle(subTitle string) string {
	return fmt.Sprintf("%s - %s", subTitle, app.Name)
}

func GetMenuItemFullTitle(title string, subTitle string) string {
	if len(subTitle) == 0 {
		return title
	} else {
		return fmt.Sprintf("%s\t%s", title, subTitle)
	}
}

func JoinString(sep string, s... string) string {
	return strings.Join(s, sep)
}

func JoinStringTo(ori *string, sep string, s... string) string {
	*ori = JoinString(sep, s...)
	return *ori
}

func AppendString(ori string, s... string) string {
	return JoinString("", append([]string{ori}, s...)...)
}

func AppendStringTo(ori *string, s... string) string {
	JoinStringTo(*&ori, "", append([]string{*ori}, s...)...)
	return *ori
}

func TrinocularString(b bool, trueString, falseString string) string {
	if b {
		return trueString
	}
	return falseString
}

// IgnoreErrorBytes 忽略错误[]byte
func IgnoreErrorBytes(data []byte, err error) []byte {
	if err != nil {
		log.Errorln("IgnoreError: %v", err)
	}
	return data
}

// IgnoreErrorString 忽略错误string
func IgnoreErrorString(data string, err error) string {
	if err != nil {
		log.Errorln("IgnoreError: %v", err)
	}
	return data
}

// ToLowerCamelCase 转小驼峰camelCase
func ToLowerCamelCase(s string) string {
	return toCamelCase(s, false)
}

// ToUpperCamelCase 转大驼峰CamelCase
func ToUpperCamelCase(s string) string {
	return toCamelCase(s, true)
}

func toCamelCase(s string, toUpper bool) string {
	if len(s) == 0 {
		return s
	}
	var camelCaseStr string
	if toUpper {
		camelCaseStr = strings.ToUpper(s[:1])
	} else {
		camelCaseStr = strings.ToLower(s[:1])
	}
	if len(s) == 1 {
		return camelCaseStr
	}
	return camelCaseStr + s[1:]
}
