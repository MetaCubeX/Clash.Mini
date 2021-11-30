package util

const (
	Second     = 1
	Minute     = 60 * Second
	HalfAHour  = 30 * Minute
	Hour       = 2 * HalfAHour
	HalfADay   = 12 * Hour
	Day        = 2 * HalfADay
	HalfAMonth = 15 * Day
	Month      = 2 * HalfAMonth
)

//func GetHumanTime(t time.Time) string {
//	if t.IsZero() {
//		return ""
//	}
//
//	diff := time.Now().Unix() - t.Unix()
//	if diff < 0 {
//		return ""
//	}
//
//	month := diff / Month
//	if month >= 1 {
//		return fmt.Sprintf(`%d 个月前`, month)
//	}
//	if diff >= Month {
//		return "半个月前"
//	}
//	week := diff / (7 * Day)
//	if week >= 1 {
//		return fmt.Sprintf(`%d 星期前`, week)
//	}
//	day := diff / Day
//	if day >= 1 {
//		return fmt.Sprintf(`%d 天前`, day)
//	}
//	if diff >= HalfADay {
//		return "半天前"
//	}
//	hour := diff / Hour
//	if hour >= 1 {
//		return fmt.Sprintf(`%d 小时前`, hour)
//	}
//	if diff >= HalfAHour {
//		return "半小时前"
//	}
//	min := diff / Minute
//	if diff < HalfAHour && diff > Minute {
//		return fmt.Sprintf(`%d 分钟前`, min)
//	}
//	if diff < Minute {
//		return "刚刚"
//	}
//	return ""
//}
