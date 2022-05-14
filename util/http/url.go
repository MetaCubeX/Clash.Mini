package http

import (
	"regexp"
)

var (
	UrlRegexp = regexp.MustCompile(`^(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|.]`)
)
