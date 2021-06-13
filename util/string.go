package util

import (
	"fmt"

	"github.com/Clash-Mini/Clash.Mini/app"
)

var ()

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
