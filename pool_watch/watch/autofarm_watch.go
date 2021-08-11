package watch

import (
	"fmt"
	"kryptos_rai/pool_watch/config"
	"kryptos_rai/pool_watch/pool_info"
	"kryptos_rai/pool_watch/rpc_client"
	"kryptos_rai/support_pools/bsc/autofarm/autofarm_watch"
	"kryptos_rai/use/logger"
	"kryptos_rai/use/pb/message_notify"
)

func AddAutoFarmPoolInfo(farm_address, pool_address, strat_address string, isVault bool) {
	poolInfo := &pool_info.PoolInfo{
		Address:  farm_address,
		WorkChan: make(chan int64),
	}

	pool_info.BscPoolList = append(pool_info.BscPoolList, poolInfo)
	go watchAutoFarmPool(poolInfo, isVault, pool_address, strat_address)
}

func watchAutoFarmPool(poolInfo *pool_info.PoolInfo, isVault bool, pool_address, strat_address string) {

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
			curTVL, err := autofarm_watch.GetStratTVL(strat_address, nodeURL)
			if err != nil {
				logger.Error("watchAutoFarmPool failed check it :%+v", err)
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
				req.MasterChelfAddr = pool_address
				err := rpc_client.RPCPoolNotify(req)
				if err != nil {
					logger.Error("watchAutoFarmPool::rpc to notify error:%+v", err)
				}
			}
		}
	}
}
