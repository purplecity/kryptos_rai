package curve_watch

import (
	"errors"
	"kryptos_rai/support_pools/eth/curve/translate"
	"kryptos_rai/use/logger"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

//因为对于venus  autofarm官方暂时已经不支持了 2021/08/12
func Get3poolInfo(poolAddress string, provider []string) (*big.Int, *big.Int, *big.Int, error) {
	for _, x := range provider {
		client, err := ethclient.Dial(x)
		if err != nil {
			logger.Error("Curve::Get3poolInfo:: dial err:%+v", err)
			continue
		}

		addr := common.HexToAddress(poolAddress)
		instance, err := translate.NewPool3Caller(addr, client)
		if err != nil {
			logger.Error("Curve::Get3poolInfo:: NewPool3Caller err:%+v", err)
			continue
		}

		wg := new(sync.WaitGroup)
		var r1 *big.Int
		var r2 *big.Int
		var r3 *big.Int

		var err1 error
		var err2 error
		var err3 error
		wg.Add(3)

		go func(r *big.Int, e error) {
			tr, err := instance.Balances(nil, big.NewInt(0))
			if err != nil {
				logger.Warn("Curve::Balances:: failed:%+v", err)
			}
			logger.Debug("tr1:%+v", tr)
			r1 = tr
			err1 = err
			wg.Done()
			return
		}(r1, err1)
		go func(r *big.Int, e error) {
			tr, err := instance.Balances(nil, big.NewInt(1))
			if err != nil {
				logger.Warn("Curve::Balances:: failed:%+v", err)
			}
			logger.Debug("tr2:%+v", tr)
			r2 = tr
			err2 = err
			wg.Done()
			return
		}(r2, err2)
		go func(r *big.Int, e error) {
			tr, err := instance.Balances(nil, big.NewInt(2))
			if err != nil {
				logger.Warn("Curve::Balances:: failed:%+v", err)
			}
			logger.Debug("tr3:%+v", tr)
			r3 = tr
			err3 = err
			wg.Done()
			return
		}(r3, err3)

		wg.Wait()
		if err2 != nil || r2 == nil || err3 != nil || r3 == nil || err1 != nil || r1 == nil {
			continue
		}
		return r1, r2, r3, nil

	}
	return nil, nil, nil, errors.New("curve::Get3poolInfo:: all provider failed")
}
