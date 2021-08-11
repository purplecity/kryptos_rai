package venus_watch

import (
	"errors"
	"kryptos_rai/support_pools/bsc/venus/translate"
	"kryptos_rai/use/logger"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetBNBCashTVL(provider []string) (*big.Int, error) {
	vtokenAddr := "0xA07c5b74C9B40447a954e1466938b865b6BBea36"
	for _, x := range provider {
		client, err := ethclient.Dial(x)
		if err != nil {
			logger.Error("Venus::GetBNBCashTVL:: dial err:%+v", err)
			continue
		}

		addr := common.HexToAddress(vtokenAddr)
		instance, err := translate.NewVbnbCaller(addr, client)
		if err != nil {
			logger.Error("Venus::GetBNBCashTVL:: NewVtokenCaller err:%+v", err)
			continue
		}

		r, err := instance.GetCash(nil)
		if err != nil || r == nil {
			logger.Error("Venus::GetBNBCashTVL:: TotalSupply err:%+v", err)
			continue
		} else {
			return r, nil
		}
	}
	return nil, errors.New("Venus::GetBNBCashTVL: all provider failed")
}

func GetCashTVL(vtokenAddr string, provider []string) (*big.Int, error) {
	for _, x := range provider {
		client, err := ethclient.Dial(x)
		if err != nil {
			logger.Error("Venus::GetCashTVL:: dial err:%+v", err)
			continue
		}

		addr := common.HexToAddress("vtokenAddr")
		instance, err := translate.NewVtokenCaller(addr, client)
		if err != nil {
			logger.Error("Venus::GetCashTVL:: NewVtokenCaller err:%+v", err)
			continue
		}

		r, err := instance.GetCash(nil)
		if err != nil || r == nil {
			logger.Error("Venus::GetCashTVL:: TotalSupply err:%+v", err)
			continue
		} else {
			return r, nil
		}
	}
	return nil, errors.New("Venus::GetCashTVL: all provider failed")
}

func GetVaultTVL(provider []string) (*big.Int, error) {
	for _, x := range provider {
		client, err := ethclient.Dial(x)
		if err != nil {
			logger.Error("Venus::GetVaultTVL:: dial err:%+v", err)
			continue
		}

		addr := common.HexToAddress("0x4BD17003473389A42DAF6a0a729f6Fdb328BbBd7")
		vault_addr := common.HexToAddress("0x0667Eed0a0aAb930af74a3dfeDD263A73994f216")
		instance, err := translate.NewVaiCaller(addr, client)
		if err != nil {
			logger.Error("Venus::GetVaultTVL:: NewVaiCaller err:%+v", err)
			continue
		}

		r, err := instance.BalanceOf(nil, vault_addr)
		if err != nil || r == nil {
			logger.Error("Venus::GetVaultTVL:: BalanceOf err:%+v", err)
			continue
		} else {
			return r, nil
		}
	}
	return nil, errors.New("Venus::GetVaultTVL: all provider failed")
}
