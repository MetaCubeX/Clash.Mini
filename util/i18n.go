package util

import (
	"time"

	c18n "github.com/Clash-Mini/Clash.Mini/constant/i18n"
	"github.com/JyCyunMe/go-i18n/i18n"
)

// TODO:
func getHumanizationSeconds(duration time.Duration) string {
	seconds := duration.Seconds()
	return i18n.TData(c18n.UtilDatetimeSeconds, &i18n.Data{
		Data: map[string]interface{}{
			"Seconds": seconds,
		},
		PluralCount: int(seconds),
	})
}

// TODO:
func getHumanizationMinutes(duration time.Duration) string {
	minutes := duration.Minutes()
	return i18n.TData(c18n.UtilDatetimeMinutes, &i18n.Data{
		Data: map[string]interface{}{
			"Minutes": minutes,
		},
		PluralCount: int(minutes),
	})
}

// TODO:
func getHumanizationHours(duration time.Duration) string {
	hours := duration.Hours()
	return i18n.TData(c18n.UtilDatetimeHours, &i18n.Data{
		Data: map[string]interface{}{
			"Hours": hours,
		},
		PluralCount: int(hours),
	})
}

// TODO:
func getHumanizationTime(duration time.Duration) (s string) {
	s = getHumanizationHours(duration)
	s += getHumanizationMinutes(duration)
	s += getHumanizationSeconds(duration)
	return
}
