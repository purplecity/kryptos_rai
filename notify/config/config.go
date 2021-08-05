package config

import (
	"kryptos_rai/use/logger"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type MongoInfo struct {
	SubMongoHost  string
	SubDatabase   string
	SubCollection string
}

type Config struct {
	Mongo                 MongoInfo
	ListenAddr            string
	TestEmailRecvAddrList []string
	ChangeMap             map[string]map[string]interface{}
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

func GetTestEmailRecvAddrList() []string {
	configMutex.RLock()
	defer configMutex.RUnlock()
	return config.TestEmailRecvAddrList
}

func GetChangeMap() map[string]map[string]interface{} {
	configMutex.RLock()
	defer configMutex.RUnlock()
	return config.ChangeMap
}

func GetMongoInfo() MongoInfo {
	configMutex.RLock()
	defer configMutex.RUnlock()
	return config.Mongo
}
