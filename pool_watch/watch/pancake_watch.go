package watch

import (
	"fmt"
	"kryptos_rai/pool_watch/config"
	"kryptos_rai/pool_watch/pool_info"
	"kryptos_rai/pool_watch/rpc_client"
	"kryptos_rai/support_pools/bsc/pancake/pancake_watch"
	"kryptos_rai/use/logger"
	"kryptos_rai/use/pb/message_notify"
)

func AddPancakePoolInfo(pool_address string, isVault bool) {
	poolInfo := &pool_info.PoolInfo{
		Address:  pool_address,
		WorkChan: make(chan int64),
	}

	pool_info.BscPoolList = append(pool_info.BscPoolList, poolInfo)
	go watchPancakePool(poolInfo, isVault)
}

func AddPancakePairInfo(pair_address string) {
	poolInfo := &pool_info.PoolInfo{
		Address:  pair_address,
		WorkChan: make(chan int64),
	}

	pool_info.BscPoolList = append(pool_info.BscPoolList, poolInfo)
	go watchPancakePair(poolInfo)
}

func AddPancakeFarmInfo(pair_address, farm_address string) {
	poolInfo := &pool_info.PoolInfo{
		Address:  farm_address,
		WorkChan: make(chan int64),
	}

	pool_info.BscPoolList = append(pool_info.BscPoolList, poolInfo)
	go watchPancakeFarm(poolInfo, pair_address)
}

func watchPancakePool(poolInfo *pool_info.PoolInfo, isVault bool) {

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
			curTVL, err := pancake_watch.GetPoolTVL(poolInfo.Address, nodeURL)
			if err != nil {
				logger.Error("GetPoolTVL failed check it :%+v", err)
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
				if isVault {
					//cake特殊处理
					req.MasterChelfAddr = "0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82"
				} else {
					req.MasterChelfAddr = ""
				}

				err := rpc_client.RPCPoolNotify(req)
				if err != nil {
					logger.Error("watchPancakePool::rpc to notify error:%+v", err)
				}

			}
		}
	}
}

func watchPancakePair(poolInfo *pool_info.PoolInfo) {

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
			curTVL, err := pancake_watch.GetPairTVL(poolInfo.Address, nodeURL)
			if err != nil {
				logger.Error("watchPancakePair::GetPairTVL failed check it :%+v", err)
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
					logger.Error("watchPancakePair::rpc to notify error:%+v", err)
				}

			}
		}
	}
}

func watchPancakeFarm(poolInfo *pool_info.PoolInfo, pair_address string) {

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
			curTVL, err := pancake_watch.GetFarmTVL(pair_address, poolInfo.Address, nodeURL)
			if err != nil {
				logger.Error("watchPancakeFarm::GetFarmTVL failed check it :%+v", err)
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
				req.MasterChelfAddr = pair_address
				err := rpc_client.RPCPoolNotify(req)
				if err != nil {
					logger.Error("watchPancakeFarm::rpc to notify error:%+v", err)
				}
			}
		}
	}
}
