package util

import "fmt"

const (
	AppTitle = "Clash.Mini"
)

var ()

func GetSubTitle(subTitle string) string {
	return fmt.Sprintf("%s - %s", subTitle, AppTitle)
}
