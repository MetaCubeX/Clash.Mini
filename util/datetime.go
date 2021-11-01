package util

import (
	"fmt"
	"time"
)

// TODO:
func GetHumanTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	diff := time.Now().Unix() - t.Unix()

	if diff < 60 {
		return "刚刚"
	}
	min := diff / 60
	if min < 30 {
		return fmt.Sprintf("%d 分钟前", min)
	} else if min < 60 {
		return "半小时前"
	}
	hour := min / 60
	if hour < 12 {
		return fmt.Sprintf("%d 小时前", hour)
	} else if hour < 24 {
		return "半天前"
	}
	return ""
}

