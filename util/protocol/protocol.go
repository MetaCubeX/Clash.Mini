package protocol

import (
	"errors"
	"fmt"
	"golang.org/x/sys/windows/registry"
	"io/fs"

	"github.com/Clash-Mini/Clash.Mini/constant"
)

func RegistryOpenOrCreateKey(k registry.Key, path string, access uint32) (registry.Key, error) {
	k, exists, err := registry.CreateKey(k, path, access)
	if exists {
		k, err := registry.OpenKey(k, path, access)
		return k, err
	}
	if err != nil {
		return k, err
	}
	return k, err
}

func RegisterCommandProtocol(enable bool) error {
	if enable {
		// 菜单配置关联
		exe := constant.Executable
		k, err := RegistryOpenOrCreateKey(registry.CLASSES_ROOT, `clash`, registry.QUERY_VALUE|registry.SET_VALUE|registry.CREATE_SUB_KEY) //registry.ALL_ACCESS) //
		err = k.SetStringValue("", "Clash Protocol")
		err = k.SetStringValue("URL Protocol", exe)
		err = k.Close()
		if err != nil {
			return err
		}
		k, err = RegistryOpenOrCreateKey(registry.CLASSES_ROOT, `clash\DefaultIcon`, registry.QUERY_VALUE|registry.SET_VALUE|registry.CREATE_SUB_KEY) //registry.ALL_ACCESS) //
		err = k.SetStringValue("", fmt.Sprintf("%s,1", exe))
		err = k.Close()
		if err != nil {
			return err
		}
		k, err = RegistryOpenOrCreateKey(registry.CLASSES_ROOT, `clash\shell\open\command`, registry.QUERY_VALUE|registry.SET_VALUE|registry.CREATE_SUB_KEY) //registry.ALL_ACCESS) //
		err = k.SetStringValue("", fmt.Sprintf(`"%s" --uac-protocol="%s"`, exe, "%1"))
		err = k.Close()
		if err != nil {
			return err
		}
		return nil
	} else {
		// 菜单配置取消关联
		_, err := registry.OpenKey(registry.CLASSES_ROOT, `clash`, registry.QUERY_VALUE)
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				return nil
			}
			return err
		}
		err = registry.DeleteKey(registry.CLASSES_ROOT, `clash`)
		if err != nil {
			return err
		}
	}
	return nil
}
