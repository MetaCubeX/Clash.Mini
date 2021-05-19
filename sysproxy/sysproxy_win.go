//+build windows

package sysproxy

import (
	"errors"
	"syscall"

	"golang.org/x/sys/windows/registry"
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
	if SavedProxy == nil {
		GetSavedProxy()
	}

	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.SET_VALUE)
	if err != nil {
		return err
	}
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

	err = k.Close()
	if err != nil {
		return err
	}

	err = func() error {
		wininet, err := syscall.LoadLibrary("Wininet.dll")
		if err != nil {
			return err
		}

		internetSetOptionW, err := syscall.GetProcAddress(wininet, "InternetSetOptionW")
		if err != nil {
			return err
		}

		ret, _, errno := syscall.Syscall6(uintptr(internetSetOptionW), 4, 0, 39, 0, 0, 0, 0)
		if ret != 1 {
			return errors.New(errno.Error())
		}
		ret, _, errno = syscall.Syscall6(uintptr(internetSetOptionW), 4, 0, 37, 0, 0, 0, 0)
		if ret != 1 {
			return errors.New(errno.Error())
		}

		err = syscall.FreeLibrary(wininet)
		if err != nil {
			return err
		}

		return nil
	}()
	if err != nil {
		return err
	}

	return nil
}
