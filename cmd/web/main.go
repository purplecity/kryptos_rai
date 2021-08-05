package main

import (
	"flag"
	"kryptos_rai/use/logger"
	"kryptos_rai/use/mongo"
	"kryptos_rai/web/app"
	"kryptos_rai/web/config"
	"kryptos_rai/web/rpc"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/cors"
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
	configPath := flag.String("config_path", "/Users/purplecity/test/kryptos_rai/web/", "config file")
	logFile := flag.String("log_file", "/Users/purplecity/test/kryptos_rai/web/web.log", "log file")
	logLevel := flag.Int("log_level", 6, "log level default -debug")
	flag.Parse()
	InitLogger(*logFile, *logLevel)

	err := config.LoadConf(*configPath)
	if err != nil {
		logger.Error("error load conf failed:%+v", err)
		return
	}

	dataCli, err1 := mongo.InitQMgo(config.GetMongoInfo().MongoHost, config.GetMongoInfo().Database, config.GetMongoInfo().Collection)
	txCli, err2 := mongo.InitQMgo(config.GetMongoInfo().TxMongoHost, config.GetMongoInfo().TxDatabase, config.GetMongoInfo().TxCollection)
	logocli, err3 := mongo.InitQMgo(config.GetMongoInfo().LogoMongoHost, config.GetMongoInfo().LogoDatabase, config.GetMongoInfo().LogoCollection)
	subcli, err4 := mongo.InitQMgo(config.GetMongoInfo().SubMongoHost, config.GetMongoInfo().SubDatabase, config.GetMongoInfo().SubCollection)
	internalcli, err5 := mongo.InitQMgo(config.GetMongoInfo().InMongoHost, config.GetMongoInfo().InDatabase, config.GetMongoInfo().InCollection)

	err6 := rpc.InitRPCClient()
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil {
		return
	}

	mongo.MongoClient = dataCli
	mongo.TxMongoClient = txCli
	mongo.LogoMongoClient = logocli
	mongo.SubMongoClient = subcli
	mongo.InternalTxClient = internalcli

	err7 := app.LoadDB()
	if err7 != nil {
		logger.Error("LoadDB failed:%+v", err)
		return
	}

	router := app.NewRouter() // router object has ServeHTTP method
	handler := cors.Default().Handler(router)
	go log.Fatal(http.ListenAndServe(config.GetListenAddr(), handler))

	go watchSignal()
	select {
	case _ = <-ch:
		logger.Info("---I'm done----")
		break
	}
}
