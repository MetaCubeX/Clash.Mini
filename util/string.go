package util

import "fmt"

const (
	AppTitle = "Clash.Mini"
)

var ()

func GetSubTitle(subTitle string) string {
	return fmt.Sprintf("%s - %s", subTitle, AppTitle)
}

func SpliceMenuFullTitle(title string, subTitle string) string {
	if len(subTitle) == 0 {
		return title
	} else {
		return fmt.Sprintf("%s\t%s", title, subTitle)
	}
}
