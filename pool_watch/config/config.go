package config

import (
	"kryptos_rai/use/logger"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	BSCNodeURL       []string
	SendNotify       bool
	BSC              map[string]bool
	HTTPPoolSize     int64
	RPCClientSize    int64
	RPCAddr          string
	RPCConnTimeout   int64
	BSCCheckInterval int64
	ETHNodeURL       []string
	ETHCheckInterval int64
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

func GetSendNotify() bool {
	configMutex.RLock()
	defer configMutex.RUnlock()
	return config.SendNotify
}

func GetRPCClientSize() int64 {
	configMutex.RLock()
	defer configMutex.RUnlock()
	return config.RPCClientSize
}

func GetBSC() map[string]bool {
	configMutex.RLock()
	defer configMutex.RUnlock()
	return config.BSC
}

func GetBSCNodeURL() []string {
	configMutex.RLock()
	defer configMutex.RUnlock()
	return config.BSCNodeURL
}

func GetHTTPPoolSize() int64 {
	configMutex.RLock()
	defer configMutex.RUnlock()
	return config.HTTPPoolSize
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

func GetBSCCheckInterval() int64 {
	configMutex.RLock()
	defer configMutex.RUnlock()
	return config.BSCCheckInterval
}

func GetETHNodeURL() []string {
	configMutex.RLock()
	defer configMutex.RUnlock()
	return config.ETHNodeURL
}

func GetETHCheckInterval() int64 {
	configMutex.RLock()
	defer configMutex.RUnlock()
	return config.ETHCheckInterval
}
