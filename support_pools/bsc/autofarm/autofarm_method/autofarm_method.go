package autofarm_method

import (
	"kryptos_rai/use/logger"
	"math/big"
	"sync"

	"kryptos_rai/support_pools/bsc/autofarm/translate"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetSingleAutoFarmAssets(query_address string, pid int64, provider []string) (inputLP *big.Int, err error) {
	//logger.Debug("AutoFarm::GetAutoFarmAssets::faddr:%+v,lpaddr:%+v,qaddr:%+v,pid:%+v", farm_address, lpaddress, query_address, pid)
	farm_address := "0x763a05bdb9f8946d8C3FA72d1e0d3f5E68647e5C"
	for _, e := range provider {
		client, err := ethclient.Dial(e)
		if err != nil {
			logger.Error("AutoFarm::GetAutoFarmAssets:: dial err:%+v  farmAddr:%+v,query:%+v,%+v", err, farm_address, query_address, e)
			continue
		}

		addr := common.HexToAddress(farm_address)
		instance, err := translate.NewSingleAutoCaller(addr, client)
		if err != nil {
			logger.Error("AutoFarm::GetAutoFarmAssets:: NewAutoCaller err:%+v farmAddr:%+v,query:%+v,%+v", err, farm_address, query_address, e)
			continue
		}

		query_addr := common.HexToAddress(query_address)

		wg := new(sync.WaitGroup)
		var r2 *big.Int

		var err2 error
		wg.Add(1)

		go func(r *big.Int, e error) {
			tr2, err := instance.StakedWantTokens(nil, big.NewInt(pid), query_addr)
			if err != nil {
				logger.Warn("AutoFarm::GetAutoFarmAssets::  StakedWantTokens err:%+v farmAddr:%+v,query:%+v,%+v", err, farm_address, query_address, e)
			}
			logger.Debug("%+v,%+v,%+v,%+v", tr2, err, query_address, farm_address)
			r2 = tr2
			err2 = err
			wg.Done()
			return
		}(r2, err2)

		wg.Wait()
		if err2 != nil || r2 == nil {
			continue
		}

		inputLP = r2
		return inputLP, nil
	}
	return nil, err
}

func GetAutoFarmAssets(query_address string, pid int64, provider []string) (inputLP, reward *big.Int, err error) {
	//logger.Debug("AutoFarm::GetAutoFarmAssets::faddr:%+v,lpaddr:%+v,qaddr:%+v,pid:%+v", farm_address, lpaddress, query_address, pid)
	farm_address := "0x0895196562C7868C5Be92459FaE7f877ED450452"
	for _, e := range provider {
		client, err := ethclient.Dial(e)
		if err != nil {
			logger.Error("AutoFarm::GetAutoFarmAssets:: dial err:%+v  farmAddr:%+v,query:%+v,%+v", err, farm_address, query_address, e)
			continue
		}

		addr := common.HexToAddress(farm_address)
		instance, err := translate.NewAutoCaller(addr, client)
		if err != nil {
			logger.Error("AutoFarm::GetAutoFarmAssets:: NewAutoCaller err:%+v farmAddr:%+v,query:%+v,%+v", err, farm_address, query_address, e)
			continue
		}

		query_addr := common.HexToAddress(query_address)

		wg := new(sync.WaitGroup)
		var r1 *big.Int
		var r2 *big.Int
		var err1 error
		var err2 error
		wg.Add(2)
		go func(r *big.Int, e error) {

			tr1, err := instance.PendingAUTO(nil, big.NewInt(pid), query_addr)
			if err != nil {
				logger.Warn("AutoFarm::GetAutoFarmAssets::  PendingAUTO err:%+v farmAddr:%+v,query:%+v,%+v", err, farm_address, query_address, e)
			}
			logger.Debug("%+v,%+v,%+v,%+v", tr1, err, query_address, farm_address)
			err1 = err
			r1 = tr1
			wg.Done()
			return
		}(r1, err1)

		go func(r *big.Int, e error) {
			tr2, err := instance.StakedWantTokens(nil, big.NewInt(pid), query_addr)
			if err != nil {
				logger.Warn("AutoFarm::GetAutoFarmAssets::  StakedWantTokens err:%+v farmAddr:%+v,query:%+v,%+v", err, farm_address, query_address, e)
			}
			logger.Debug("%+v,%+v,%+v,%+v", tr2, err, query_address, farm_address)
			r2 = tr2
			err2 = err
			wg.Done()
			return
		}(r2, err2)

		wg.Wait()
		if err1 != nil || err2 != nil || r1 == nil || r2 == nil {
			continue
		}

		inputLP = r2
		reward = r1
		return inputLP, reward, nil
	}
	return nil, nil, err
}
