package config

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/Clash-Mini/Clash.Mini/cmd/breaker"
	"github.com/Clash-Mini/Clash.Mini/cmd/protocol"
	"io/ioutil"
	"os"
	path "path/filepath"
	"reflect"
	"time"

	"github.com/Clash-Mini/Clash.Mini/app"
	"github.com/Clash-Mini/Clash.Mini/cmd"
	"github.com/Clash-Mini/Clash.Mini/cmd/auto"
	"github.com/Clash-Mini/Clash.Mini/cmd/cron"
	"github.com/Clash-Mini/Clash.Mini/cmd/mmdb"
	"github.com/Clash-Mini/Clash.Mini/cmd/parser"
	"github.com/Clash-Mini/Clash.Mini/cmd/proxy"
	"github.com/Clash-Mini/Clash.Mini/cmd/startup"
	"github.com/Clash-Mini/Clash.Mini/cmd/sys"
	"github.com/Clash-Mini/Clash.Mini/cmd/task"
	cConfig "github.com/Clash-Mini/Clash.Mini/constant/config"
	"github.com/Clash-Mini/Clash.Mini/log"
	"github.com/Clash-Mini/Clash.Mini/notify"
	fileUtils "github.com/Clash-Mini/Clash.Mini/util/file"
	stringUtils "github.com/Clash-Mini/Clash.Mini/util/string"
	"github.com/JyCyunMe/go-i18n/i18n"

	"github.com/imdario/mergo"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Lang    string    `mapstructure:"lang"`
	Cmd     CmdConfig `mapstructure:"cmd"`
	Profile string    `mapstructure:"profile"`
}

type CmdConfig struct {
	Auto     auto.Type     `mapstructure:"auto"`
	Cron     cron.Type     `mapstructure:"cron"`
	MMDB     mmdb.Type     `mapstructure:"mmdb"`
	Proxy    proxy.Type    `mapstructure:"proxy"`
	Startup  startup.Type  `mapstructure:"startup"`
	Sys      sys.Type      `mapstructure:"sys"`
	Task     task.Type     `mapstructure:"task"`
	Breaker  breaker.Type  `mapstructure:"breaker"`
	Protocol protocol.Type `mapstructure:"protocol"`
}

var (
	config     *viper.Viper
	configPath = path.Join(cConfig.DirPath, cConfig.FileName+"."+cConfig.FileFormat)
)

func init() {
	// 加载配置
	LoadConfig()
}

func getDefaultConfig() *Config {
	return &Config{
		Lang: i18n.English.Tag.String(),
		Cmd: CmdConfig{
			MMDB: mmdb.Max,
			Cron: cron.ON,
			Auto: auto.OFF,
			//cmd.Task.GetName(): 	cmd.OffName,	//开机启动
			//cmd.Sys.GetName(): 		cmd.OffName,	//默认代理
			//cmd.Proxy.GetName(): 	cmd.OffName,
			Breaker:  breaker.OFF,
			Protocol: protocol.OFF,
			Startup:  startup.OFF, //开机启动
			Proxy:    proxy.Rule,  //代理模式
		},
		Profile: "config",
	}
}

// InitConfig 初始化配置
func InitConfig() {
	appConfig := getDefaultConfig()
	var m *map[string]interface{}
	var originalConfig *Config
	var err error
	isOk := true
	exists, err := fileUtils.IsExists(configPath)
	if err != nil {
		log.Errorln("find config file error: %v", err)
		return
	}
	if !exists {
		log.Warnln("cannot find config file, it will generate default to: %s", configPath)
		isOk := true
		// 检查目录是否存在
		var dirExists bool
		dirExists, err = fileUtils.IsExists(cConfig.DirPath)
		if err != nil {
			isOk = false
		} else {
			if !dirExists {
				// 不存在则创建目录
				err = os.MkdirAll(cConfig.DirPath, 0666)
				if err == nil {
					log.Infoln("[config] created default config directory: %s", cConfig.DirPath)
					isOk = true
				}
			}
		}
		if isOk {
			// 文件不存在，目录已准备，新建配置文件
			err = mapstructure.Decode(appConfig, &m)
			if err != nil {
				isOk = false
			} else {
				err = config.MergeConfigMap(*m)
				if err != nil {
					isOk = false
				} else {
					SaveConfig(appConfig)
					log.Infoln("[config] created default config file: %s", configPath)
				}
			}
		}
	} else {
		originalData := stringUtils.IgnoreErrorBytes(ioutil.ReadFile(configPath))
		err = yaml.Unmarshal(originalData, &m)
		if err != nil {
			isOk = false
		} else {
			var metadata mapstructure.Metadata
			err = mapstructure.DecodeMetadata(m, &originalConfig, &metadata)
			if err != nil {
				isOk = false
			} else if len(metadata.Unused) > 0 {
				isOk = false
				err = fmt.Errorf("found %d useless field(s)", len(metadata.Unused))
			}
		}
		if !isOk && len(originalData) > 0 {
			backupFile := time.Now().Format(configPath + "_20060102_150405.bak")
			log.Warnln("decode error: %s, it will backup the file to %s and regenerate a new", err.Error(), backupFile)
			err = ioutil.WriteFile(backupFile, originalData, 0644)
			if err != nil {
				log.Errorln("backup the file to %s failed: %s", backupFile, err.Error())
			}
		}
		err = mergo.Merge(appConfig, originalConfig, mergo.WithOverride)
		if err != nil {
			log.Errorln("merge config error: %s", err.Error())
			return
		}
		SaveConfig(appConfig)
	}
	//fmt.Println(appConfig)
	m = nil
	originalConfig = nil
	appConfig = nil

	//fmt.Println(config.AllSettings())
	err = config.ReadInConfig()
	if err != nil {
		log.Errorln("read config error: %s", err.Error())
		return
	}
	//fmt.Println(config.AllSettings())
}

// LoadConfig 加载配置
func LoadConfig() {
	//viper.NewWithOptions()
	config = viper.New()
	config.SetConfigName(cConfig.FileName)   // 文件名 (不含扩展名)
	config.SetConfigType(cConfig.FileFormat) // 扩展名
	config.AddConfigPath(cConfig.DirPath)    // 配置文件路径
	var err error
	// 查找并读取配置
	err = config.ReadInConfig()
	if err != nil && !errors.As(err, &viper.ConfigFileNotFoundError{}) {
		// 其他错误
		errMsg := fmt.Sprintf("[config] Fatal error config file: %v \n", err)
		notify.PushError("", errMsg)
		log.Fatalln(errMsg)
	}
	InitConfig()
	//SaveConfig(nil)
}

// SaveConfig 保存配置
func SaveConfig(data interface{}) {
	errPrefix := "[config] save config failed: "
	if config == nil {
		log.Errorln(errPrefix + "config is nil")
		return
	}
	if data == nil {
		data = config.AllSettings()
	}
	bs, err := yaml.Marshal(data)
	if err != nil {
		log.Errorln("unable to marshal config to YAML: %v", err)
	}
	buf := bytes.NewBufferString(fmt.Sprintf("# Clash.Mini\r\n# v%s\r\n# %s\r\n\r\n", app.Version, time.Now().Format("2006-01-02 15:04:05")))
	defer func() {
		buf.Reset()
		buf = nil
	}()
	buf.Write(bs)
	err = ioutil.WriteFile(configPath, buf.Bytes(), 0644)
	//err := config.WriteConfig()
	if err != nil {
		log.Errorln(errPrefix+"%s: %v", "write to file failed", err)
		return
	}
}

// Set 设置配置值
func Set(name string, value interface{}) {
	config.Set(name, value)
	//fmt.Println(config.AllSettings())
	// TODO: 1分钟内只保存一次？
	SaveConfig(nil)
}

// Get 获取配置值
func Get(name string) interface{} {
	name = stringUtils.ToLowerCamelCase(name)
	//config.IsSet()
	if config.InConfig(name) {
		return config.Get(name)
	}
	return nil
}

// GetOrDefault 获取配置值或默认值
func GetOrDefault(name string, defaultValue interface{}) interface{} {
	value := Get(name)
	if value == nil {
		return defaultValue
	}
	return value
}

// SetCmd 写入命令到配置
func SetCmd(value cmd.GeneralType) error {
	command := value.GetCommandType()
	if !command.IsValid(value) {
		return fmt.Errorf("command \"%s\" is not supported type \"%s\"", command.GetName(), value.String())
	}
	if !value.IsValid() {
		log.Infoln("被动新建键值: %s", command.GetName())
		value = value.GetDefault()
	}
	Set("cmd."+command.GetName(), value.String())
	return nil
}

// IsCmdPositive 判断命令是否为活动值，并更新配置
func IsCmdPositive(command cmd.CommandType) (b bool) {
	value := Get("cmd." + command.GetName())
	if value == nil {
		return false
	}

	cmdValue := parser.GetCmdOrDefaultValue(command, reflect.ValueOf(value).String())
	if SetCmd(cmdValue) != nil {
		return false
	}

	if cmdValue == cmd.Invalid {
		return false
	}
	return cmdValue.IsPositive()
}
