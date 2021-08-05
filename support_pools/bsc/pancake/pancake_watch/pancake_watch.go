package pancake_watch

import (
	"errors"
	"kryptos_rai/support_pools/bsc/pancake/translate"
	"kryptos_rai/use/logger"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetVaultTVL(provider []string) (*big.Int, error) {
	for _, x := range provider {
		client, err := ethclient.Dial(x)
		if err != nil {
			logger.Error("Pancake::GetVaultTVL:: dial err:%+v", err)
			continue
		}

		addr := common.HexToAddress("0xa80240Eb5d7E05d3F250cF000eEc0891d00b51CC")
		instance, err := translate.NewVaultCaller(addr, client)
		if err != nil {
			logger.Error("Pancake::GetVaultTVL:: NewVaultCaller err:%+v", err)
			continue
		}

		r, err := instance.BalanceOf(nil)
		if err != nil || r == nil {
			logger.Error("Pancake::GetVaultTVL:: BalanceOf err:%+v", err)
			continue
		} else {
			return r, nil
		}
	}
	return nil, errors.New("Pancake::GetVaultTVL: all provider failed")

}

func GetPairTVL(pair_address string, provider []string) (*big.Int, error) {
	for _, x := range provider {
		client, err := ethclient.Dial(x)
		if err != nil {
			logger.Error("Pancake::GetPairTVL:: dial err:%+v", err)
			continue
		}

		addr := common.HexToAddress(pair_address)
		instance, err := translate.NewPairCaller(addr, client)
		if err != nil {
			logger.Error("Pancake::GetPairTVL:: NewPairCaller err:%+v", err)
			continue
		}

		r, err := instance.TotalSupply(nil)
		if err != nil || r == nil {
			logger.Error("Pancake::GetPairTVL:: TotalSupply err:%+v", err)
			continue
		} else {
			return r, nil
		}
	}
	return nil, errors.New("Pancake::GetPairTVL: all provider failed")

}

//这里也可以作为masterchelf的
func GetPoolTVL(pool_address string, provider []string) (*big.Int, error) {
	for _, x := range provider {
		client, err := ethclient.Dial(x)
		if err != nil {
			logger.Error("Pancake::GetPoolTVL:: dial err:%+v", err)
			continue
		}

		cakeAddr := common.HexToAddress("0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82")
		pool_addr := common.HexToAddress(pool_address)
		instance, err := translate.NewCakeCaller(cakeAddr, client)
		if err != nil {
			logger.Error("Pancake::GetPoolTVL:: NewCakeCaller err:%+v", err)
			continue
		}

		r, err := instance.BalanceOf(nil, pool_addr)
		if err != nil || r == nil {
			logger.Error("Pancake::GetPoolTVL::  BalanceOf err:%+v", err)
			continue
		} else {
			return r, nil
		}
	}
	return nil, errors.New("Pancake::GetPoolTVL: all provider failed")

}

func GetFarmTVL(pair_address, farm_address string, provider []string) (*big.Int, error) {
	for _, x := range provider {
		client, err := ethclient.Dial(x)
		if err != nil {
			logger.Error("Pancake::GetFarmTVL:: dial err:%+v", err)
			continue
		}

		addr := common.HexToAddress(pair_address)
		instance, err := translate.NewPairCaller(addr, client)
		if err != nil {
			logger.Error("Pancake::GetFarmTVL:: NewPairCaller err:%+v", err)
			continue
		}

		farm_addr := common.HexToAddress(farm_address)
		r, err := instance.BalanceOf(nil, farm_addr)
		if err != nil || r == nil {
			logger.Error("Pancake::GetFarmTVL::  BalanceOf err:%+v", err)
			continue
		} else {
			return r, nil
		}
	}

	return nil, errors.New("Pancake::GetPairTVL: all provider failed")
}
