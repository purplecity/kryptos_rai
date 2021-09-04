package watch

import (
	"kryptos_rai/pool_watch/config"
	"kryptos_rai/pool_watch/pool_info"
	"kryptos_rai/pool_watch/rpc_client"
	"kryptos_rai/use/pool"
)

var bscBlockHeight int64

func Init() {
	bscBlockHeight = 0
	pool_info.BscPoolList = []*pool_info.PoolInfo{}
	pool_info.EthPoolList = []*pool_info.PoolInfo{}
	rpc_client.InitClient(int(config.GetRPCClientSize()))
	pool.InitClient(int(config.GetHTTPPoolSize()))
}
