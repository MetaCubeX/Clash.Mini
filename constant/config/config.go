package config

import (
	"github.com/Dreamacro/clash/hub/executor"
)

const (
	FileName   = "config"
	FileFormat = "yaml"
)

var (
	DirPath      = ".cm"
	DashboardDir = ".cm/dashboard"
	RawConfig, _ = executor.Parse()
)
