package config

import (
	"kryptos_rai/use/logger"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type MongoInfo struct {
	Database   string
	Collection string
	MongoHost  string
}
type Config struct {
	PoolSize      int64
	BSCNodeURL    []string
	Mongo         MongoInfo
	StartBlock    int64
	CheckInterval int64
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

func GetPoolSize() int64 {
	configMutex.RLock()
	defer configMutex.RUnlock()
	return config.PoolSize
}

func GetBSCNodeURL() []string {
	configMutex.RLock()
	defer configMutex.RUnlock()
	return config.BSCNodeURL
}

func GetMongo() MongoInfo {
	configMutex.RLock()
	defer configMutex.RUnlock()
	return config.Mongo
}

func GetStartBlock() int64 {
	configMutex.RLock()
	defer configMutex.RUnlock()
	return config.StartBlock
}

func GetCheckInterval() int64 {
	configMutex.RLock()
	defer configMutex.RUnlock()
	return config.CheckInterval
}
