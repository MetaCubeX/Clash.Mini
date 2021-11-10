package protocol

import (
	"errors"
	"fmt"
	"golang.org/x/sys/windows/registry"
	"io/fs"

	"github.com/Clash-Mini/Clash.Mini/constant"
	stringUtils "github.com/Clash-Mini/Clash.Mini/util/string"
)

const (
	protocolName 		= "clash"
	registrySeparator 	= `\`
)

func GetRegistryPath(path... string) string {
	return stringUtils.JoinString(registrySeparator, path...)
}

func RegistryOpenOrCreateKey(k registry.Key, path string, access uint32) (registry.Key, error) {
	k, _, err := registry.CreateKey(k, path, access)
	//if exists && k == 0 {
	//	k, err := registry.OpenKey(k, path, access)
	//	return k, err
	//}
	if err != nil {
		return k, err
	}
	return k, err
}

func DeleteKeyWithSub(key registry.Key, path string) error {
	k, err := registry.OpenKey(key, path, registry.READ | registry.SET_VALUE)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil
		}
		return err
	}
	defer func() {
		err = k.Close()
	}()

	var info *registry.KeyInfo
	if info, err = k.Stat(); info.SubKeyCount > 0 {
		// needs recursive delete from subkey to key
		keys, err := k.ReadSubKeyNames(0)
		if err != nil {
			return err
		}
		for _, subKey := range keys {
			if err = DeleteKeyWithSub(key, GetRegistryPath(path, subKey)); err != nil {
				return err
			}
		}
	}
	if err = registry.DeleteKey(key, path); err != nil {
		return err
	}
	return nil
}

func RegisterCommandProtocol(enable bool) (err error) {
	if enable {
		// 菜单配置关联
		exe := constant.Executable
		k, err := RegistryOpenOrCreateKey(registry.CLASSES_ROOT, protocolName, registry.QUERY_VALUE|registry.SET_VALUE|registry.CREATE_SUB_KEY) //registry.ALL_ACCESS) //
		err = k.SetStringValue("", "Clash Protocol")
		err = k.SetStringValue("URL Protocol", exe)
		if err = k.Close(); err != nil {
			return err
		}
		k, err = RegistryOpenOrCreateKey(registry.CLASSES_ROOT, GetRegistryPath(protocolName, `DefaultIcon`), registry.QUERY_VALUE|registry.SET_VALUE|registry.CREATE_SUB_KEY) //registry.ALL_ACCESS) //
		err = k.SetStringValue("", fmt.Sprintf("%s,1", exe))
		if err = k.Close(); err != nil {
			return err
		}
		k, err = RegistryOpenOrCreateKey(registry.CLASSES_ROOT, GetRegistryPath(protocolName, `shell\open\command`), registry.QUERY_VALUE|registry.SET_VALUE|registry.CREATE_SUB_KEY) //registry.ALL_ACCESS) //
		err = k.SetStringValue("", fmt.Sprintf(`"%s" call --protocol="%s"`, exe, "%1"))
		if err = k.Close(); err != nil {
			return err
		}
		return nil
	} else {
		// 菜单配置取消关联
		return DeleteKeyWithSub(registry.CLASSES_ROOT, protocolName)
	}
}
