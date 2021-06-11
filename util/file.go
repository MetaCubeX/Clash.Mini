package util

import (
	"fmt"
	"math"
)

var (
	fileSizeUnits = []string{"", "K", "M", "G", "T", "P", "E"}
)

// FormatHumanizedFileSize 格式化为可读文件大小
func FormatHumanizedFileSize(fileSize int64) (size string) {
	i := math.Floor(math.Log(float64(fileSize)) / math.Log(1024))
	return fmt.Sprintf("%.02f %sB", float64(fileSize)/math.Pow(1024, i), fileSizeUnits[int(i)])
}
