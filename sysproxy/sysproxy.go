package sysproxy

// ProxyConfig 系统代理信息
type ProxyConfig struct {

	// 启用
	Enable bool

	// 代理地址
	Server string

}

func (c *ProxyConfig) String() string {
	if c == nil {
		return "nil"
	}

	if c.Enable {
		return "Enabled: True" + "; Server: " + c.Server
	}

	return "Enabled: False" + "; Server: " + c.Server
}

// GetNilProxy 获取空系统代理信息
func GetNilProxy() *ProxyConfig {
	return &ProxyConfig{
		Enable: false,
		Server: "",
	}
}

// ClearSystemProxy 清除系统代理
func ClearSystemProxy() error {
	return SetSystemProxy(GetNilProxy())
}
