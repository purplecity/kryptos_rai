package autofarm_watch

import (
	"errors"
	"kryptos_rai/support_pools/bsc/autofarm/translate"
	"kryptos_rai/use/logger"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

//因为对于venus  autofarm官方暂时已经不支持了 2021/08/12
func GetStratTVL(stratAddr string, provider []string) (*big.Int, error) {
	for _, x := range provider {
		client, err := ethclient.Dial(x)
		if err != nil {
			logger.Error("AutoFarm::GetStratTVL:: dial err:%+v", err)
			continue
		}

		addr := common.HexToAddress(stratAddr)
		instance, err := translate.NewStratCaller(addr, client)
		if err != nil {
			logger.Error("AutoFarm::GetStratTVL:: NewStratCaller err:%+v", err)
			continue
		}

		r, err := instance.WantLockedTotal(nil)
		if err != nil || r == nil {
			logger.Error("AutoFarm::GetStratTVL:: WantLockedTotal err:%+v", err)
			continue
		} else {
			return r, nil
		}
	}
	return nil, errors.New("AutoFarm::GetStratTVL:: all provider failed")
}
