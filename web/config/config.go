package config

import (
	"kryptos_rai/use/logger"
	"strings"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type MongoInfo struct {
	MongoHost      string
	Database       string
	Collection     string
	TxMongoHost    string
	TxDatabase     string
	TxCollection   string
	LogoMongoHost  string
	LogoDatabase   string
	LogoCollection string
	SubMongoHost   string
	SubDatabase    string
	SubCollection  string
	InMongoHost    string
	InDatabase     string
	InCollection   string
}
type Config struct {
	ListenAddr     string
	RPCAddr        string
	Mongo          MongoInfo
	RPCConnTimeout int64
	PlatformCheck  map[string]bool
	ChainCheck     map[string]bool
	PoolProperty   map[string]string
	LogoType       map[string]bool
	MasterChelf    map[string][]map[string]interface{}
}

var (
	configMutex   = sync.RWMutex{}
	config        = new(Config)
	configViper   = viper.New()
	configPath    = ""
	configFileAbs = ""
)

func watchConfig() error {
	configViper.SetConfigName("config")
	configViper.AddConfigPath(configPath)
	configViper.SetConfigType("toml")
	if err := configViper.ReadInConfig(); err != nil {
		logger.Info("read config err:%+v\n", err)
		return err
	}
	configViper.WatchConfig()
	configViper.OnConfigChange(func(e fsnotify.Event) {
		logger.Info("config is change :%+v, reload it \n", e.String())
		reloadConifg()
	})
	return nil
}

func LoadConf(configFilePath string) error {
	configPath = configFilePath
	configFileAbs = configFilePath + "config.toml"
	configMutex.Lock()
	defer configMutex.Unlock()
	if _, err := toml.DecodeFile(configFileAbs, &config); err != nil {
		logger.Error("decode config file failed %+v", err)
		return err
	}
	logger.Info("Config Load:%+v", config)
	if err := watchConfig(); err != nil {
		logger.Error("watch config file failed %+v", err)
		return err
	}
	return nil
}

func reloadConifg() error {
	configMutex.Lock()
	defer configMutex.Unlock()
	if _, err := toml.DecodeFile(configFileAbs, &config); err != nil {
		return err
	}
	logger.Info("Config ReLoad:%+v", config)
	return nil
}

func GetListenAddr() string {
	configMutex.RLock()
	defer configMutex.RUnlock()
	return config.ListenAddr
}

func GetRPCAddr() string {
	configMutex.RLock()
	defer configMutex.RUnlock()
	return config.RPCAddr
}

func GetRPCConnTimeout() int64 {
	configMutex.RLock()
	defer configMutex.RUnlock()
	return config.RPCConnTimeout
}

func GetMongoInfo() MongoInfo {
	configMutex.RLock()
	defer configMutex.RUnlock()
	return config.Mongo
}

func GetPlatformCheck() map[string]bool {
	configMutex.RLock()
	defer configMutex.RUnlock()
	return config.PlatformCheck
}
func GetChainCheck() map[string]bool {
	configMutex.RLock()
	defer configMutex.RUnlock()
	return config.ChainCheck

}

func GetPoolProperty() map[string]string {
	configMutex.RLock()
	defer configMutex.RUnlock()
	return config.PoolProperty
}

func GetLogoType() map[string]bool {
	configMutex.RLock()
	defer configMutex.RUnlock()
	return config.LogoType
}

func GetMasterChelf() map[string]map[string]map[string]interface{} {
	configMutex.RLock()
	defer configMutex.RUnlock()
	r := map[string]map[string]map[string]interface{}{}

	for k, v := range config.MasterChelf {
		itemMap := map[string]map[string]interface{}{}

		for _, i := range v {

			//约定配置文件有Address PID Symbol
			item := map[string]interface{}{}
			address := ""
			for m, n := range i {

				if m == "Address" {
					address = strings.ToLower(n.(string))
				}
				item[m] = n

			}
			itemMap[address] = item

		}
		r[k] = itemMap

	}
	return r
}
