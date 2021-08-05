package pancake_method

import (
	"errors"
	"kryptos_rai/support_pools/bsc/pancake/translate"
	"kryptos_rai/use/logger"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetPairCertificateBalance(pair_address, query_address string, provider []string) (r, r0, r1 *big.Int, e error) {
	for _, e := range provider {
		client, err := ethclient.Dial(e)
		if err != nil {
			logger.Error("Pancake::GetPairCertificateBalance:: dial err:%+v,%+v", err, e)
			continue
		}

		addr := common.HexToAddress(pair_address)
		instance, err := translate.NewPairCaller(addr, client)
		if err != nil {
			logger.Error("Pancake::GetPairCertificateBalance:: NewPairCaller err:%+v,%+v", err, e)
			continue
		}

		query_addr := common.HexToAddress(query_address)
		r, err := instance.BalanceOf(nil, query_addr)
		if err != nil || r == nil {
			logger.Error("Pancake::GetPairCertificateBalance::  BalanceOf err:%+v,%+v", err, e)
			continue
		}

		//然后估算能取出来多少
		wg := new(sync.WaitGroup)
		var r0 *big.Int
		var r1 *big.Int
		var r2 *big.Int
		var err1 error
		var err2 error
		wg.Add(2)

		go func(z0, z1 *big.Int, e error) {
			tr1, err := instance.GetReserves(nil)
			if err != nil {
				logger.Warn("Pancake::GetPairCertificateBalance::  GetReserves err:%+v,url:%+v", err, e)
			}
			logger.Debug("Pancake::GetPairCertificateBalance:: %+v,%+v,%+v,%+v", tr1.Reserve0, tr1.Reserve1, query_address, pair_address)
			err1 = err
			r0 = tr1.Reserve0
			r1 = tr1.Reserve1
			wg.Done()
			return
		}(r0, r1, err1)

		go func(r *big.Int, e error) {
			tr2, err := instance.TotalSupply(nil)
			if err != nil {
				logger.Warn("Pancake::GetPairCertificateBalance::  TotalSupply err:%+v,url:%+v", err, e)
			}
			logger.Debug("Pancake::GetPairCertificateBalance:: totalsupply:%+v,%+v,%+v,%+v", tr2, err, query_address, pair_address)
			r2 = tr2
			err2 = err
			wg.Done()
			return
		}(r2, err2)

		wg.Wait()

		if err1 != nil || err2 != nil || r1 == nil || r2 == nil || r0 == nil || r == nil {
			logger.Debug("err1:%+v,err2:%+v,r0:%+v,r1:%+v,r2:%+v.r:%+v", err1, err2, r0, r1, r2, r)
			continue
		}

		//得先乘 防止出现第一步为0
		out0 := new(big.Int)
		out0.Mul(r, r0).Div(out0, r2)
		out1 := new(big.Int)
		out1.Mul(r, r1).Div(out1, r2)
		logger.Debug("##### r:%+v out0%+v,out1:%+v", r, out0, out1)
		return r, out0, out1, nil
	}
	return nil, nil, nil, errors.New("GetPairCertificateBalance: all provider failed")

}

func GetPoolUnderlyingAssets(pool_address, query_address string, provider []string) (input, reward *big.Int, err error) {
	for _, e := range provider {
		client, err := ethclient.Dial(e)
		if err != nil {
			logger.Error("Pancake::GetPoolUnderlyingAssets:: dial err:%+v,%+v", err, e)
			continue
		}

		addr := common.HexToAddress(pool_address)
		instance, err := translate.NewPoolCaller(addr, client)
		if err != nil {
			logger.Error("Pancake::GetPoolUnderlyingAssets:: NewPoolCaller err:%+v,%+v", err, e)
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
			tr1, err := instance.UserInfo(nil, query_addr)
			if err != nil {
				logger.Warn("Pancake::GetPoolUnderlyingAssets::  UserInfo err:%+v,%+v", err, e)
			}
			logger.Debug("%+v,%+v,%+v,%+v", tr1, err, query_address, pool_address)
			err1 = err
			r1 = tr1.Amount
			wg.Done()
			return
		}(r1, err1)

		go func(r *big.Int, e error) {
			tr2, err := instance.PendingReward(nil, query_addr)
			if err != nil {
				logger.Warn("Pancake::GetPoolUnderlyingAssets::  PendingReward err:%+v", err)
			}
			logger.Debug("%+v,%+v,%+v,%+v", tr2, err, query_address, pool_address)
			err1 = err
			r2 = tr2
			err2 = err
			wg.Done()
			return
		}(r2, err2)

		wg.Wait()
		if err1 != nil || err2 != nil || r1 == nil || r2 == nil {
			continue
		}
		input = r1
		reward = r2
		return input, reward, nil
	}
	return nil, nil, errors.New("GetPairCertificateBalance: all provider failed")

}

func GetFarmUnderlyingAssets(farm_address, lpaddress, query_address string, pid int64, provider []string) (inputLP, reward *big.Int, err error) {
	//logger.Debug("Pancake::GetFarmUnderlyingAssets::faddr:%+v,lpaddr:%+v,qaddr:%+v,pid:%+v", farm_address, lpaddress, query_address, pid)
	for _, e := range provider {
		client, err := ethclient.Dial(e)
		if err != nil {
			logger.Error("Pancake::GetFarmUnderlyingAssets:: dial err:%+v  farmAddr:%+v,query:%+v,%+v", err, farm_address, query_address, e)
			continue
		}

		addr := common.HexToAddress(farm_address)
		instance, err := translate.NewFarmCaller(addr, client)
		if err != nil {
			logger.Error("Pancake::GetFarmUnderlyingAssets:: NewFarmCaller err:%+v farmAddr:%+v,query:%+v,%+v", err, farm_address, query_address, e)
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

			tr1, err := instance.UserInfo(nil, big.NewInt(pid), query_addr)
			if err != nil {
				logger.Warn("Pancake::GetFarmUnderlyingAssets::  UserInfo err:%+v farmAddr:%+v,query:%+v,%+v", err, farm_address, query_address, e)
			}
			logger.Debug("%+v,%+v,%+v,%+v", tr1, err, query_address, farm_address)
			err1 = err
			r1 = tr1.Amount
			wg.Done()
			return
		}(r1, err1)

		go func(r *big.Int, e error) {
			tr2, err := instance.PendingCake(nil, big.NewInt(pid), query_addr)
			if err != nil {
				logger.Warn("Pancake::GetFarmUnderlyingAssets::  PendingReward err:%+v farmAddr:%+v,query:%+v,%+v", err, farm_address, query_address, e)
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

		inputLP = r1
		reward = r2
		return inputLP, reward, nil
	}
	return nil, nil, err
}

func GetVaultAssets(query_address string, provider []string) (inputLP, reward *big.Int, err error) {
	//logger.Debug("Pancake::GetVaultAssets::qaddr:%+v", query_address)
	for _, e := range provider {
		client, err := ethclient.Dial(e)
		if err != nil {
			logger.Error("Pancake::GetVaultAssets:: dial err:%+v ,query:%+v,url:%+v", err, query_address, e)
			continue
		}

		addr := common.HexToAddress("0xa80240Eb5d7E05d3F250cF000eEc0891d00b51CC")
		instance, err := translate.NewVaultCaller(addr, client)
		if err != nil {
			logger.Error("Pancake::GetVaultAssets:: NewVaultCaller err:%+v ,query:%+v,url:%+v", err, query_address, e)
			continue
		}

		query_addr := common.HexToAddress(query_address)

		wg := new(sync.WaitGroup)
		var r1 *big.Int
		var r2 *big.Int
		var r3 *big.Int
		var r4 *big.Int
		var err1 error
		var err2 error
		var err3 error
		wg.Add(3)
		go func(z0, z1 *big.Int, e error) {

			tr1, err := instance.UserInfo(nil, query_addr)
			if err != nil {
				logger.Warn("Pancake::GetVaultAssets::  UserInfo err:%+v ,query:%+v,url:%+v", err, query_address, e)
			}
			logger.Debug("%+v,%+v,%+v", tr1, err, query_address)
			err1 = err
			r1 = tr1.Shares
			r2 = tr1.CakeAtLastUserAction
			wg.Done()
			return
		}(r1, r2, err1)

		go func(r *big.Int, e error) {
			tr2, err := instance.BalanceOf(nil)
			if err != nil {
				logger.Warn("Pancake::GetVaultAssets::  BalanceOf err:%+v,query:%+v,url:%+v", err, query_address, e)
			}
			logger.Debug("%+v,%+v,%+v", tr2, err, query_address)
			r3 = tr2
			err2 = err
			wg.Done()
			return
		}(r3, err2)

		go func(r *big.Int, e error) {
			tr2, err := instance.TotalShares(nil)
			if err != nil {
				logger.Warn("Pancake::GetVaultAssets::  TotalShares err:%+v,query:%+v,url:%+v", err, query_address, e)
			}
			logger.Debug("%+v,%+v,%+v", tr2, err, query_address)
			r4 = tr2
			err3 = err
			wg.Done()
			return
		}(r4, err3)

		wg.Wait()
		if err1 != nil || err2 != nil || err3 != nil || r1 == nil || r2 == nil || r3 == nil || r4 == nil {
			continue
		}

		inputLP = r2

		//得先乘 防止出现第一步为0
		reward := new(big.Int)
		reward.Mul(r1, r3).Div(reward, r4)
		freward := new(big.Int)
		freward.Sub(reward, r2)
		if freward.Cmp(big.NewInt(0)) < 0 {
			logger.Error("Pancake::GetVaultAssets::what happendreward is neg check this reward:%+v,input:%+v", reward, inputLP)
			freward = big.NewInt(0)
		}
		return inputLP, freward, nil
	}
	return nil, nil, err
}
