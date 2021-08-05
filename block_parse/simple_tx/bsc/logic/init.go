package logic

import (
	"kryptos_rai/block_parse/simple_tx/bsc/config"
	"kryptos_rai/use/pool"
)

func BSCInit() {
	curBlock = config.GetStartBlock()
	pool.InitClient(int(config.GetPoolSize()))
}
