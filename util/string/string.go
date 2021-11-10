package string

import (
	"fmt"
	"strings"

	"github.com/Clash-Mini/Clash.Mini/app"
	"github.com/Clash-Mini/Clash.Mini/log"
)

const (
	logHeader = "util.string"
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
		log.Errorln("[%s] IgnoreError: %v", logHeader, err)
	}
	return data
}

// IgnoreErrorString 忽略错误string
func IgnoreErrorString(data string, err error) string {
	if err != nil {
		log.Errorln("[%s] IgnoreError: %v", logHeader, err)
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

func StartsWith(s, starts string) bool {
	sLen := len(s)
	startsLen := len(starts)
	if sLen < startsLen {
		return false
	} else if sLen == startsLen {
		return s == starts
	} else if sLen > startsLen {
		return s[:startsLen] == starts
	}
	return false
}

func EndsWith(s, ends string) bool {
	sLen := len(s)
	endsLen := len(ends)
	if sLen < endsLen {
		return false
	} else if sLen == endsLen {
		return s == ends
	} else if sLen > endsLen {
		return s[sLen - endsLen:] == ends
	}
	return false
}

func UnescapeArgQuote(s string) string {
	sLen := len(s)
	if sLen < 1 {
		return s
	} else if sLen < 2 && s == `"` {
		return ""
	}
	s = strings.ReplaceAll(s, `\"`, `"`)
	if s[:1] == `"` && s[sLen - 1:] == `"` {
		return s[1:sLen - 1]
	}
	return s
}
