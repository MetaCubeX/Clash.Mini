package http

import (
	"regexp"
)

var (
	UrlRegexp = regexp.MustCompile(`^https?://(\w+(?:\.\w+)*)(/(?:[^\s]+?)?)?$`)
)
