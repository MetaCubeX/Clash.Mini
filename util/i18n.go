package util

import (
	"time"

	c18n "github.com/Clash-Mini/Clash.Mini/constant/i18n"
	"github.com/JyCyunMe/go-i18n/i18n"
)

// TODO:
// Deprecated:
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
// Deprecated:
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
// Deprecated:
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
// Deprecated:
func getHumanizationTime(duration time.Duration) (s string) {
	s = getHumanizationHours(duration)
	s += getHumanizationMinutes(duration)
	s += getHumanizationSeconds(duration)
	return
}

func GetHumanTimeI18n(t time.Time) string {
	return getHumanTime(t) + i18n.T(c18n.UtilDatetimeAgo)
}

func getHumanTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	diff := time.Now().Unix() - t.Unix()
	if diff < 0 {
		return ""
	}

	month := diff / Month
	if month >= 1 {
		return i18n.TData(c18n.UtilDatetimeMonths, &i18n.Data{
			Data: map[string]interface{}{
				"Months": month,
			},
			PluralCount: int(month),
		})
	}
	if diff >= Month {
		return i18n.TData(c18n.UtilDatetimeMonths, &i18n.Data{
			Data: map[string]interface{}{
				"Months": i18n.T(c18n.UtilDatetimeHalf),
			},
			PluralCount: 1,
		})
	}
	week := diff / (7 * Day)
	if week >= 1 {
		return i18n.TData(c18n.UtilDatetimeWeeks, &i18n.Data{
			Data: map[string]interface{}{
				"Weeks": week,
			},
			PluralCount: int(week),
		})
	}
	day := diff / Day
	if day >= 1 {
		return i18n.TData(c18n.UtilDatetimeDays, &i18n.Data{
			Data: map[string]interface{}{
				"Days": day,
			},
			PluralCount: int(day),
		})
	}
	if diff >= HalfADay {
		return i18n.TData(c18n.UtilDatetimeDays, &i18n.Data{
			Data: map[string]interface{}{
				"Days": i18n.T(c18n.UtilDatetimeHalf),
			},
			PluralCount: 1,
		})
	}
	hour := diff / Hour
	if hour >= 1 {
		return i18n.TData(c18n.UtilDatetimeHours, &i18n.Data{
			Data: map[string]interface{}{
				"Hours": hour,
			},
			PluralCount: int(hour),
		})
	}
	if diff >= HalfAHour {
		return i18n.TData(c18n.UtilDatetimeHours, &i18n.Data{
			Data: map[string]interface{}{
				"Hours": i18n.T(c18n.UtilDatetimeHalf),
			},
			PluralCount: 1,
		})
	}
	min := diff / Minute
	if diff < HalfAHour && diff > Minute {
		return i18n.TData(c18n.UtilDatetimeMinutes, &i18n.Data{
			Data: map[string]interface{}{
				"Minutes": min,
			},
			PluralCount: int(min),
		})
	}
	if diff < Minute {
		return i18n.T(c18n.UtilDatetimeMoments)
	}
	return ""
}
