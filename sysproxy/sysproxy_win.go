//go:build windows
// +build windows

package sysproxy

import (
	"errors"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
)

var (
	wininet            = windows.MustLoadDLL("Wininet.dll")
	internetSetOptionW = wininet.MustFindProc("InternetSetOptionW")
)

func GetCurrentProxy() (*ProxyConfig, error) {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.QUERY_VALUE)
	if err != nil {
		return nil, err
	}

	Enable, _, err := k.GetIntegerValue("ProxyEnable")
	if err != nil {
	}

	Server, _, err := k.GetStringValue("ProxyServer")
	if err != nil {
	}

	err = k.Close()
	if err != nil {
		return nil, err
	}

	return &ProxyConfig{
		Enable: Enable > 0,
		Server: Server,
	}, nil
}

func SetSystemProxy(p *ProxyConfig) error {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer k.Close()

	if p.Enable {
		err = k.SetDWordValue("ProxyEnable", 0x00000001)
	} else {
		err = k.SetDWordValue("ProxyEnable", 0x00000000)
	}
	if err != nil {
		return err
	}

	err = k.SetStringValue("ProxyServer", p.Server)
	if err != nil {
		return err
	}

	err = func() error {
		ret, _, errno := internetSetOptionW.Call(0, 39, 0, 0)
		if ret != 1 {
			return errors.New(errno.Error())
		}
		ret, _, errno = internetSetOptionW.Call(0, 37, 0, 0)
		if ret != 1 {
			return errors.New(errno.Error())
		}

		return nil
	}()
	if err != nil {
		return err
	}

	return nil
}
