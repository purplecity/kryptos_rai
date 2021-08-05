package rpc_server

import (
	"context"
	"errors"
	"kryptos_rai/notify/config"
	"kryptos_rai/notify/email"
	"kryptos_rai/use/logger"
	"kryptos_rai/use/mongo"
	"kryptos_rai/use/pb/message_notify"
	"kryptos_rai/web/app"
	"math/big"
	"net"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
)

type NotifyService struct{}

func (s *NotifyService) FlashEventNotify(ctx context.Context, in *message_notify.FEN_Request) (*message_notify.FEN_Reply, error) {
	logger.Notic("FlashEventNotify:: receive flashloan event notify, pooladdress:%+v,hash:%+v", in.PoolAddress, in.Hash)
	toList := []string{}
	poolInfo, ok := config.GetChangeMap()[in.PoolAddress]
	if !ok {
		logger.Error("FlashEventNotify:: not set the pool info of this address:%+v", in.PoolAddress)
		return nil, errors.New("not set poolInfo ")
	}

	projectName, ok1 := poolInfo["project_name"]
	poolName, ok2 := poolInfo["pool_name"]
	if !ok1 || !ok2 {
		logger.Error("FlashEventNotify:: not set the pool info of this address:%+v", in.PoolAddress)
		return nil, errors.New("not set poolInfo")
	}

	subData := []app.SubscribeInfo{}
	err := mongo.SubMongoClient.Find(context.Background(), bson.M{"pool_address": strings.ToLower(in.PoolAddress), "watch_fl": "true"}).All(&subData)
	if err != nil {
		logger.Error("FlashEventNotify:: get subcribe info failed pool_address:%+v,err:%+v", in.PoolAddress, err)
	} else {
		subMap := map[string]bool{}
		for _, e := range subData {
			_, ok := subMap[e.Email]
			if !ok {
				subMap[e.Email] = true
			}
		}

		for k, _ := range subMap {
			toList = append(toList, k)
		}
	}

	subject := projectName.(string) + " " + poolName.(string) + " " + "疑似闪电贷通知"
	body := "您订阅的合约: (" + in.PoolAddress + ") 疑似发生一笔闪电贷,交易哈希为: " + in.Hash + " ,请留意相关风险"
	err = email.SendEmail(toList, subject, body)
	r := new(message_notify.FEN_Reply)
	r.Success = true
	if err != nil {
		r.Success = false
		logger.Error("FlashEventNotify:: send email failed:%+v,data is: pooladdress:%+v,hash:%+v", err, in.PoolAddress, in.Hash)
	}
	return r, err
}

func (s *NotifyService) PoolNotify(ctx context.Context, in *message_notify.PN_Request) (*message_notify.SuccessReply, error) {
	logger.Notic("PoolNotify:: receive pool notify:pooladdress:%+v, lastCheckHeight:%+v,curHeight:%+v,lastCheckHeightTVL:%+v,curHeightTVL:%+v", in.PoolAddress, in.LastCheckBlockHeight, in.CurCheckBlockHeight, in.LastCheckBlockHeightTvl, in.CurCheckBlockHeightTvl)
	r := new(message_notify.SuccessReply)
	r.Success = true

	toList := []string{}

	lastTVL := new(big.Int)
	lastTVL, ok := lastTVL.SetString(in.LastCheckBlockHeightTvl, 10)
	if !ok {
		logger.Error("PoolNotify:: lastTVL: %+v,big.Int SetString failed !", in.LastCheckBlockHeightTvl)
		return r, nil
	}

	curTVL := new(big.Int)
	curTVL, ok = curTVL.SetString(in.CurCheckBlockHeightTvl, 10)
	if !ok {
		logger.Error("PoolNotify:: curTVL: %+v,big.Int SetString failed !", in.CurCheckBlockHeightTvl)
	}

	poolInfo, ok := config.GetChangeMap()[strings.ToLower(in.PoolAddress)]
	if !ok {
		logger.Error("PoolNotify:: not set the pool info of this address:%+v", in.PoolAddress)
		return r, nil
	}

	projectName, ok1 := poolInfo["project_name"]
	poolName, ok2 := poolInfo["pool_name"]
	if !ok1 || !ok2 {
		logger.Error("PoolNotify:: not set the pool info of this address:%+v", in.PoolAddress)
		return r, nil
	}

	hundred := big.NewInt(int64(100))

	zero := big.NewInt(0)
	exceed := (curTVL.Cmp(lastTVL) != 0)
	increase := (curTVL.Cmp(lastTVL) >= 0)

	if exceed && lastTVL.Cmp(zero) > 0 {
		change := "下降"
		delta := big.NewInt(int64(0))
		delta.Sub(lastTVL, curTVL).Mul(delta, hundred).Div(delta, lastTVL)
		if increase {
			change = "增加"
			delta.Sub(curTVL, lastTVL).Mul(delta, hundred).Div(delta, lastTVL)
		}
		subData := []app.SubscribeInfo{}
		isInt64 := delta.IsInt64() //大部分数字应该都可以 否则应该算是大到超标 因为超过了int64的范围
		var changeNum int64
		if isInt64 {
			changeNum = delta.Int64()
		} else {
			changeNum = 100
		}

		if changeNum < 1 {
			logger.Notic("PoolNotify:: changeis:%+v,pooladdress:%+v,projectName:%+v poolName:%+v,lastCheckHeight:%+v,curHeight:%+v,lastCheckHeightTVL:%+v,curHeightTVL:%+v", changeNum, in.PoolAddress, projectName, poolName, in.LastCheckBlockHeight, in.CurCheckBlockHeight, in.LastCheckBlockHeightTvl, in.CurCheckBlockHeightTvl)
			return r, nil
		}
		if increase {
			err := mongo.SubMongoClient.Find(context.Background(), bson.M{"pool_address": strings.ToLower(in.PoolAddress), "master_chelf_addr": in.MasterChelfAddr, "tvl_inc": bson.M{"$lte": changeNum, "$gt": 0}}).All(&subData)
			if err != nil {
				logger.Error("PoolNotify::get sub info failed %+v", in.PoolAddress)
				return r, nil
			}
		} else {
			err := mongo.SubMongoClient.Find(context.Background(), bson.M{"pool_address": strings.ToLower(in.PoolAddress), "master_chelf_addr": in.MasterChelfAddr, "tvl_dec": bson.M{"$lte": changeNum, "$gt": 0}}).All(&subData)
			if err != nil {
				logger.Error("PoolNotify::get sub info failed %+v", in.PoolAddress)
				return r, nil
			}
		}

		//logger.Debug("PoolNotify::isInt64:%+v,change:%+v,changeNum:%+v,subData:%+v", isInt64, change, changeNum, subData)
		subMap := map[string]bool{}
		for _, e := range subData {
			_, ok := subMap[e.Email]
			if !ok {
				subMap[e.Email] = true
			}
		}
		for k, _ := range subMap {
			toList = append(toList, k)
		}

		//logger.Debug("toList:%+v", toList)
		subject := projectName.(string) + " " + poolName.(string) + " " + "监控提醒"
		body := "当前区块高度: " + in.CurCheckBlockHeight + ",TVL为: " + in.CurCheckBlockHeightTvl + "\n" + "上一个检测区块高度: " + in.LastCheckBlockHeight + ",TVL为: " + in.LastCheckBlockHeightTvl + "\n" + "同比" + change + "百分之" + delta.String()
		err := email.SendEmail(toList, subject, body)
		if err != nil {
			logger.Error("PoolNotify:: send email failed:%+v,pooladdress:%+v,projectName:%+v poolName:%+v,lastCheckHeight:%+v,curHeight:%+v,lastCheckHeightTVL:%+v,curHeightTVL:%+v", err, in.PoolAddress, projectName, poolName, in.LastCheckBlockHeight, in.CurCheckBlockHeight, in.LastCheckBlockHeightTvl, in.CurCheckBlockHeightTvl)
			return r, nil
		}
		logger.Notic("send success!")
	}
	return r, nil
}

func InitRPCService() error {
	logger.Notic("InitRPCService::start analysis rpc server")
	addr := config.GetListenAddr()
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Notic("InitRPCService::listen addr failed: %+v", err)
		return err
	}
	s := grpc.NewServer()

	message_notify.RegisterMessageNotifyServer(s, &NotifyService{})
	if err = s.Serve(listen); err != nil {
		logger.Info("InitRPCService::failed to start rpc service at: %v error: %v", addr, err)
		return err
	}
	return nil
}
