//+build linux darwin

package sysproxy

// GetCurrentProxy is ...
func GetCurrentProxy() (*ProxyConfig, error) {
	return &ProxyConfig{
		Enable: false,
		Server: ":80",
	}, nil
}

// SetSystemProxy is ...
func SetSystemProxy(p *ProxyConfig) error {
	return nil
}
