package controller

import (
	"errors"
	C "github.com/Clash-Mini/Clash.Mini/config"
	"github.com/Clash-Mini/Clash.Mini/constant"
	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Clash-Mini/Clash.Mini/mixin"
	"github.com/Clash-Mini/Clash.Mini/util/file"
	"github.com/Dreamacro/clash/config"
	"github.com/Dreamacro/clash/hub/executor"
	"github.com/Dreamacro/clash/hub/route"
	"gopkg.in/yaml.v3"
	"os"
	path "path/filepath"
)

// Option 配置文件前置处理
// 当返回nil则停止调用，并取消应用
///**

var MixinGeneralPath = path.Join(constant.MixinDir, constant.MixGeneralFile)
var MixinTunPath = path.Join(constant.MixinDir, constant.MixTunFile)
var MixinDnsPath = path.Join(constant.MixinDir, constant.MixDnsFile)

type PreHandle func(map[string]interface{}) (map[string]interface{}, error)
type PostHandle func(*config.Config) (*config.Config, error)

var ErrorNoSuchConfigFile = errors.New("No such config file on the path")

type Core struct {
	mixinConfigPath  string
	configPath       string
	preHandleChains  []PreHandle
	postHandleChains []PostHandle
}

func NewCoreWithoutHandle(configPath string) *Core {
	return NewCore(configPath, make([]PreHandle, 0, 2), make([]PostHandle, 0, 2))
}

func NewCore(configPath string, preHandles []PreHandle, postHandles []PostHandle) *Core {
	var core = &Core{
		configPath:       configPath,
		preHandleChains:  preHandles,
		postHandleChains: postHandles,
	}
	if C.IsMixinPositive(mixin.General) {
		core.AddPreHandle(true, mixinConfigToOrigin(MixinGeneralPath))
	}
	if C.IsMixinPositive(mixin.Tun) {
		core.AddPreHandle(true, mixinConfigToOrigin(MixinTunPath))
		core.AddPreHandle(true, mixinConfigToOrigin(MixinDnsPath))
	}
	if !C.IsMixinPositive(mixin.Tun) && C.IsMixinPositive(mixin.Dns) {
		core.AddPreHandle(true, mixinConfigToOrigin(MixinDnsPath))
	}
	core.AddPreHandle(true, loadConfig(configPath))
	return core
}

func (core *Core) AddPreHandle(addFirst bool, handle ...PreHandle) {
	if addFirst {
		core.preHandleChains = append(handle, core.preHandleChains...)
	} else {
		core.preHandleChains = append(core.preHandleChains, handle...)
	}
}

func (core *Core) AddPostHandle(addFirst bool, handle ...PostHandle) {
	if addFirst {
		core.postHandleChains = append(handle, core.postHandleChains...)
	} else {
		core.postHandleChains = append(core.postHandleChains, handle...)
	}
}

func GetConfig(path string) (*config.Config, error) {
	if bytes, err := readFile(path); err == nil {
		return config.Parse(bytes)
	} else {
		return nil, err
	}
}

func readFile(path string) ([]byte, error) {
	if ok, err := file.IsExists(path); err == nil && ok {
		if fileBytes, err := os.ReadFile(path); err == nil {
			return fileBytes, nil
		} else {
			return nil, err
		}
	} else {
		if err != nil {
			return nil, err
		}

		return nil, ErrorNoSuchConfigFile
	}
}

func readYamlWithMap(path string) (map[string]interface{}, error) {
	var (
		err       error
		fileBytes []byte
	)

	if fileBytes, err = readFile(path); err == nil {
		var cfgMap map[string]interface{}
		if err = yaml.Unmarshal(fileBytes, &cfgMap); err == nil {
			return cfgMap, nil
		}
	}

	return nil, err
}

func loadConfig(configPath string) PreHandle {
	return func(cfg map[string]interface{}) (map[string]interface{}, error) {
		if cfgMap, err := readYamlWithMap(configPath); err == nil {
			return cfgMap, nil
		} else {
			return nil, err
		}
	}
}

func mixinConfigToOrigin(mixinConfigPath string) PreHandle {
	return func(cfg map[string]interface{}) (map[string]interface{}, error) {
		if cfgMap, err := readYamlWithMap(mixinConfigPath); err == nil {
			for name, content := range cfgMap {
				cfg[name] = content
			}

			return cfg, nil
		} else {
			return cfg, err
		}
	}
}

func (core *Core) ApplyConfig(isUpdate bool) error {
	var configMap map[string]interface{}
	var err error
	var bytes []byte
	var cfg *config.Config
	for _, handle := range core.preHandleChains {
		configMap, err = handle(configMap)
		if err != nil {
			log.Errorln("load config failed, error:%v", err)
			return err
		}
	}

	if bytes, err = yaml.Marshal(configMap); err != nil {
		log.Errorln("convert config failed,error:%v", err)
		return err
	}

	if cfg, err = config.Parse(bytes); err != nil {
		log.Errorln("config file error after mixing,error:%v", err)
		return err
	}

	for _, handle := range core.postHandleChains {
		cfg, err = handle(cfg)
		if err != nil {
			log.Errorln("handle config an error occurred, error:%v", err)
			return err
		}
	}

	if !isUpdate {
		if cfg.General.ExternalUI != "" {
			route.SetUIPath(cfg.General.ExternalUI)
		}

		if cfg.General.ExternalController != "" {
			constant.SetController(cfg.General.ExternalController, cfg.General.Secret)
			go route.Start(cfg.General.ExternalController, cfg.General.Secret)
		}
	}

	executor.ApplyConfig(cfg, !isUpdate)
	return nil
}

func CoreStart(configFile string) error {
	core := NewCoreWithoutHandle(configFile)
	err := core.ApplyConfig(false)
	if err != nil {
		return err
	}
	return nil
}

func CoreUpdate(configFile string) error {
	core := NewCoreWithoutHandle(configFile)
	err := core.ApplyConfig(true)
	if err != nil {
		return err
	}
	return nil
}
