package config

import (
	"kryptos_rai/use/logger"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	ListenAddr string
	BSCNodeURL []string
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

func GetBSCNodeURL() []string {
	configMutex.RLock()
	defer configMutex.RUnlock()
	return config.BSCNodeURL
}
