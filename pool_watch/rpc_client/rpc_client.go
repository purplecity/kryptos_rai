package rpc_client

import (
	"context"
	"errors"
	"sync"

	"kryptos_rai/pool_watch/config"
	"kryptos_rai/use/logger"
	"kryptos_rai/use/pb/message_notify"
	"time"

	"google.golang.org/grpc"
)

var (
	rpcPool = sync.Pool{
		New: func() interface{} {
			return InitRPCClient()
		},
	}
)

func InitClient(PoolSize int) {
	for i := 0; i < PoolSize; i++ {
		Client := InitRPCClient()
		rpcPool.Put(Client)
	}
}

func InitRPCClient() message_notify.MessageNotifyClient {
	addr := config.GetRPCAddr()
	for {
		mn_rpc, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithTimeout(time.Duration(config.GetRPCConnTimeout())*time.Second))
		if err != nil {
			logger.Error("InitRPCClient::init failed %v will try again", addr)
			continue
		}
		return message_notify.NewMessageNotifyClient(mn_rpc)
	}
}

func RPCPoolNotify(req *message_notify.PN_Request) error {
	client := rpcPool.Get().(message_notify.MessageNotifyClient)
	defer rpcPool.Put(client)

	r, err := client.PoolNotify(context.Background(), req)
	if err != nil {
		logger.Error("RPCPoolNotify:: PoolNotify failed:%+v", err)
		return err
	}
	if !r.Success {
		logger.Error("RPCPoolNotify:: PoolNotify not success:%+v", *req)
		return errors.New("result failed")
	}
	return nil
}

func RPCCurvePool3Notify(req *message_notify.C3P_Request) error {
	client := rpcPool.Get().(message_notify.MessageNotifyClient)
	defer rpcPool.Put(client)

	r, err := client.Curve3PoolNotify(context.Background(), req)
	if err != nil {
		logger.Error("RPCCurvePool3Notify:: PoolNotify failed:%+v", err)
		return err
	}
	if !r.Success {
		logger.Error("RPCCurvePool3Notify:: PoolNotify not success:%+v", *req)
		return errors.New("result failed")
	}
	return nil
}
