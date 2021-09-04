package watch

import (
	"fmt"
	"kryptos_rai/pool_watch/config"
	"kryptos_rai/pool_watch/pool_info"
	"kryptos_rai/pool_watch/rpc_client"
	"kryptos_rai/support_pools/eth/curve/curve_watch"
	"kryptos_rai/use/logger"
	"kryptos_rai/use/pb/message_notify"
)

func AddCurvePool3(pool_address string) {
	poolInfo := &pool_info.PoolInfo{
		Address:  pool_address,
		WorkChan: make(chan int64),
	}

	pool_info.EthPoolList = append(pool_info.EthPoolList, poolInfo)
	go watchCurve3Pool(poolInfo)
}

func watchCurve3Pool(poolInfo *pool_info.PoolInfo) {

	lastTVL1 := "0"

	lastTVL2 := "0"

	lastTVL3 := "0"
	lastBlock := int64(0)
	for block := range poolInfo.WorkChan {
		tmpNodeURL := config.GetETHNodeURL()
		indexList := Perm(len(tmpNodeURL))
		nodeURL := []string{}
		for _, e := range indexList {
			nodeURL = append(nodeURL, tmpNodeURL[e])
		}
		if block > lastBlock {
			t1, t2, t3, err := curve_watch.Get3poolInfo(poolInfo.Address, nodeURL)
			if err != nil {
				logger.Error("Get3poolInfo failed check it :%+v", err)
				continue
			}
			tmpLastBlock := lastBlock
			tmpLastTVL1 := lastTVL1
			lastTVL1 = t1.String()

			tmpLastTVL2 := lastTVL2
			lastTVL2 = t2.String()

			tmpLastTVL3 := lastTVL3
			lastTVL3 = t3.String()
			lastBlock = block
			if tmpLastBlock > 0 && config.GetSendNotify() {
				req := new(message_notify.C3P_Request)
				req.CurCheckBlockHeight = fmt.Sprintf("%+v", lastBlock)
				req.CurCheckBlockHeightTvl1 = lastTVL1
				req.CurCheckBlockHeightTvl2 = lastTVL2
				req.CurCheckBlockHeightTvl3 = lastTVL3
				req.LastCheckBlockHeight = fmt.Sprintf("%+v", tmpLastBlock)
				req.LastCheckBlockHeightTvl1 = tmpLastTVL1
				req.LastCheckBlockHeightTvl2 = tmpLastTVL2
				req.LastCheckBlockHeightTvl3 = tmpLastTVL3
				req.PoolAddress = poolInfo.Address

				err := rpc_client.RPCPoolNotify(req)
				if err != nil {
					logger.Error("watchPancakePool::rpc to notify error:%+v", err)
				}

			}
		}
	}
}
