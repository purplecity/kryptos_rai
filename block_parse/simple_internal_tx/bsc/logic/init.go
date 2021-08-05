package logic

import (
	"kryptos_rai/block_parse/simple_internal_tx/bsc/config"
	"kryptos_rai/use/pool"
)

func BSCInit() {
	pool.InitClient(int(config.GetPoolSize()))
}
