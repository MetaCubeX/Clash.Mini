package sysproxy

import (
	"fmt"
	"strconv"

	"github.com/Clash-Mini/Clash.Mini/constant"

	"github.com/Dreamacro/clash/proxy"
)

// ProxyConfig is ...
type ProxyConfig struct {
	Enable bool
	Server string
}

// SavedProxy is ...
var SavedProxy *ProxyConfig

func (c *ProxyConfig) String() string {
	if c == nil {
		return "nil"
	}

	if c.Enable {
		return "Enabled: True" + "; Server: " + c.Server
	}

	return "Enabled: False" + "; Server: " + c.Server
}

// GetSavedProxy is ...
func GetSavedProxy() *ProxyConfig {
	if SavedProxy == nil {
		err0 := func() error {
			p, err := GetCurrentProxy()
			if err != nil {
				return err
			}
			var Ports int
			if proxy.GetPorts().MixedPort != 0 {
				Ports = proxy.GetPorts().MixedPort
			} else {
				Ports = proxy.GetPorts().Port
			}
			if p.Enable && p.Server == fmt.Sprintf("%s:%d", constant.Localhost, Ports) {
				SavedProxy = &ProxyConfig{
					Enable: false,
					Server: ":" + strconv.Itoa(proxy.GetPorts().MixedPort),
				}
			} else {
				SavedProxy = p
			}
			return nil
		}()

		if err0 != nil {
			SavedProxy = &ProxyConfig{
				Enable: false,
				Server: ":80",
			}
			return SavedProxy
		}

		return SavedProxy
	}

	return SavedProxy
}
