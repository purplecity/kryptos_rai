package mongo

import (
	"context"
	"kryptos_rai/use/logger"

	"github.com/qiniu/qmgo"
)

var MongoClient *qmgo.QmgoClient
var TxMongoClient *qmgo.QmgoClient
var LogoMongoClient *qmgo.QmgoClient
var SubMongoClient *qmgo.QmgoClient
var InternalTxClient *qmgo.QmgoClient

func InitQMgo(host, database, collection string) (*qmgo.QmgoClient, error) {
	ctx := context.Background()
	cli, err := qmgo.Open(ctx, &qmgo.Config{Uri: host, Database: database, Coll: collection})
	if err != nil {
		logger.Error("InitQMgo:connect mongo failed :%+v", err)
		return nil, err
	}

	err = cli.Ping(2)
	if err != nil {
		logger.Error("InitQMgo:connect mongo failed :%+v", err)
		return nil, err
	}

	return cli, nil
}
