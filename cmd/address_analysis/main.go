package main

import (
	"flag"
	"kryptos_rai/address_analysis/config"
	"kryptos_rai/address_analysis/rpc_server"
	"kryptos_rai/use/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var ch chan int = make(chan int)

func InitLogger(file string, lvl int) {
	logger.New(file, lvl, logger.Rotate{Size: logger.GB, Expired: time.Hour * 24 * 30, Interval: time.Hour * 24})
}

func watchSignal() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	s := <-c
	logger.Info("received signal:%v", s)
	ch <- 1
}

func main() {
	configPath := flag.String("config_path", "/Users/purplecity/test/kryptos_rai/address_analysis/", "config file")
	logFile := flag.String("log_file", "/Users/purplecity/test/kryptos_rai/address_analysis/address_analysis.log", "log file")
	logLevel := flag.Int("log_level", 6, "log level default -debug")
	flag.Parse()
	InitLogger(*logFile, *logLevel)

	err := config.LoadConf(*configPath)
	if err != nil {
		logger.Error("error load conf failed:%+v", err)
		return
	}

	err1 := rpc_server.InitRPCService()
	if err1 != nil {
		return
	}

	go watchSignal()
	select {
	case _ = <-ch:
		logger.Info("---I'm done----")
		break
	}
}
