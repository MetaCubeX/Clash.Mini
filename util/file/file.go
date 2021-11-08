package file

import (
	"fmt"
	"math"
	"os"
)

var (
	fileSizeUnits 	= []string{"", "K", "M", "G", "T", "P", "E"}
)

// FormatHumanizedFileSize 格式化为可读文件大小
func FormatHumanizedFileSize(fileSize int64) (size string) {
	i := math.Floor(math.Log(float64(fileSize)) / math.Log(1024))
	return fmt.Sprintf("%.02f %sB", float64(fileSize)/math.Pow(1024, i), fileSizeUnits[int(i)])
}

// IsExists 文件是否存在
func IsExists(fileName string) (exists bool, err error) {
	_, err = os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
