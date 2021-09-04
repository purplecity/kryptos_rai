package pool_info

type PoolInfo struct {
	Address  string
	WorkChan chan int64
}

var BscPoolList []*PoolInfo
var EthPoolList []*PoolInfo
