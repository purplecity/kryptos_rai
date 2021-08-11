package watch

import (
	"fmt"
	"kryptos_rai/pool_watch/config"
	"kryptos_rai/pool_watch/pool_info"
	"kryptos_rai/pool_watch/rpc_client"
	"kryptos_rai/support_pools/bsc/venus/venus_watch"
	"kryptos_rai/use/logger"
	"kryptos_rai/use/pb/message_notify"
)

func AddBNBVenusPoolInfo(pool_address string, isVault bool) {
	poolInfo := &pool_info.PoolInfo{
		Address:  pool_address,
		WorkChan: make(chan int64),
	}

	pool_info.BscPoolList = append(pool_info.BscPoolList, poolInfo)
	go watchBNBVenusPool(poolInfo, isVault)
}

func AddVenusPoolInfo(pool_address string, isVault bool) {
	poolInfo := &pool_info.PoolInfo{
		Address:  pool_address,
		WorkChan: make(chan int64),
	}

	pool_info.BscPoolList = append(pool_info.BscPoolList, poolInfo)
	go watchVenusPool(poolInfo, isVault)
}

func AddVaultPoolInfo(pool_address string, isVault bool) {
	poolInfo := &pool_info.PoolInfo{
		Address:  pool_address,
		WorkChan: make(chan int64),
	}

	pool_info.BscPoolList = append(pool_info.BscPoolList, poolInfo)
	go watchVaultPool(poolInfo, isVault)
}

func watchBNBVenusPool(poolInfo *pool_info.PoolInfo, isVault bool) {

	lastTVL := "0"
	lastBlock := int64(0)
	for block := range poolInfo.WorkChan {
		tmpNodeURL := config.GetBSCNodeURL()
		indexList := Perm(len(tmpNodeURL))
		nodeURL := []string{}
		for _, e := range indexList {
			nodeURL = append(nodeURL, tmpNodeURL[e])
		}
		if block > lastBlock {
			curTVL, err := venus_watch.GetBNBCashTVL(nodeURL)
			if err != nil {
				logger.Error("GetCashTVL failed check it :%+v", err)
				continue
			}
			tmpLastBlock := lastBlock
			tmpLastTVL := lastTVL
			lastTVL = curTVL.String()
			lastBlock = block
			if tmpLastBlock > 0 && config.GetSendNotify() {
				req := new(message_notify.PN_Request)
				req.CurCheckBlockHeight = fmt.Sprintf("%+v", lastBlock)
				req.CurCheckBlockHeightTvl = lastTVL
				req.LastCheckBlockHeight = fmt.Sprintf("%+v", tmpLastBlock)
				req.LastCheckBlockHeightTvl = tmpLastTVL
				req.PoolAddress = poolInfo.Address
				req.MasterChelfAddr = ""
				err := rpc_client.RPCPoolNotify(req)
				if err != nil {
					logger.Error("watchVenusPool::rpc to notify error:%+v", err)
				}
			}
		}
	}
}

func watchVenusPool(poolInfo *pool_info.PoolInfo, isVault bool) {

	lastTVL := "0"
	lastBlock := int64(0)
	for block := range poolInfo.WorkChan {
		tmpNodeURL := config.GetBSCNodeURL()
		indexList := Perm(len(tmpNodeURL))
		nodeURL := []string{}
		for _, e := range indexList {
			nodeURL = append(nodeURL, tmpNodeURL[e])
		}
		if block > lastBlock {
			curTVL, err := venus_watch.GetCashTVL(poolInfo.Address, nodeURL)
			if err != nil {
				logger.Error("GetCashTVL failed check it :%+v", err)
				continue
			}
			tmpLastBlock := lastBlock
			tmpLastTVL := lastTVL
			lastTVL = curTVL.String()
			lastBlock = block
			if tmpLastBlock > 0 && config.GetSendNotify() {
				req := new(message_notify.PN_Request)
				req.CurCheckBlockHeight = fmt.Sprintf("%+v", lastBlock)
				req.CurCheckBlockHeightTvl = lastTVL
				req.LastCheckBlockHeight = fmt.Sprintf("%+v", tmpLastBlock)
				req.LastCheckBlockHeightTvl = tmpLastTVL
				req.PoolAddress = poolInfo.Address
				req.MasterChelfAddr = ""
				err := rpc_client.RPCPoolNotify(req)
				if err != nil {
					logger.Error("watchVenusPool::rpc to notify error:%+v", err)
				}
			}
		}
	}
}

func watchVaultPool(poolInfo *pool_info.PoolInfo, isVault bool) {

	lastTVL := "0"
	lastBlock := int64(0)
	for block := range poolInfo.WorkChan {
		tmpNodeURL := config.GetBSCNodeURL()
		indexList := Perm(len(tmpNodeURL))
		nodeURL := []string{}
		for _, e := range indexList {
			nodeURL = append(nodeURL, tmpNodeURL[e])
		}
		if block > lastBlock {
			curTVL, err := venus_watch.GetVaultTVL(nodeURL)
			if err != nil {
				logger.Error("GetVaultTVL failed check it :%+v", err)
				continue
			}
			tmpLastBlock := lastBlock
			tmpLastTVL := lastTVL
			lastTVL = curTVL.String()
			lastBlock = block
			if tmpLastBlock > 0 && config.GetSendNotify() {
				req := new(message_notify.PN_Request)
				req.CurCheckBlockHeight = fmt.Sprintf("%+v", lastBlock)
				req.CurCheckBlockHeightTvl = lastTVL
				req.LastCheckBlockHeight = fmt.Sprintf("%+v", tmpLastBlock)
				req.LastCheckBlockHeightTvl = tmpLastTVL
				req.PoolAddress = poolInfo.Address
				req.MasterChelfAddr = ""
				err := rpc_client.RPCPoolNotify(req)
				if err != nil {
					logger.Error("watchVaultPool::rpc to notify error:%+v", err)
				}
			}
		}
	}
}
