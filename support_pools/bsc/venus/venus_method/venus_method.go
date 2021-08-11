package venus_method

import (
	"errors"
	"kryptos_rai/support_pools/bsc/venus/translate"
	"kryptos_rai/use/logger"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 获取底层最大token(暂不考虑借贷的影响) 主要目的是预警  包括了精度

func GetBNBSupplyMaxWithdrawToken(query_address string, provider []string) (*big.Int, error) {
	for _, e := range provider {
		client, err := ethclient.Dial(e)
		if err != nil {
			logger.Error("Venus::GetBNBSupplyMaxWithdrawToken:: dial err:%+v,%+v", err, e)
			continue
		}

		vtokenAddr := "0xA07c5b74C9B40447a954e1466938b865b6BBea36"
		addr := common.HexToAddress(vtokenAddr)
		instance, err := translate.NewVtokenCaller(addr, client)
		if err != nil {
			logger.Error("Venus::GetBNBSupplyMaxWithdrawToken::  NewPairCaller err:%+v,%+v", err, e)
			continue
		}

		query_addr := common.HexToAddress(query_address)
		r, err := instance.BalanceOf(nil, query_addr)
		if err != nil || r == nil {
			logger.Error("Venus::GetBNBSupplyMaxWithdrawToken::  BalanceOf err:%+v,%+v", err, e)
			continue
		}

		if r.Cmp(big.NewInt(0)) <= 0 {
			logger.Debug("Venus::GetBNBSupplyMaxWithdrawToken:: balance of is :%+v", r)
			return big.NewInt(0), nil
		}

		//然后估算能取出来多少
		wg := new(sync.WaitGroup)
		var (
			getCash        *big.Int
			totalBorrows   *big.Int
			totalReservers *big.Int
			totalSupply    *big.Int
			err1           error
			err2           error
			err3           error
			err4           error
		)

		wg.Add(4)

		go func(z0 *big.Int, e error) {
			tr, err := instance.GetCash(nil)
			if err != nil {
				logger.Warn("Venus::GetBNBSupplyMaxWithdrawToken::  GetCash err:%+v,url:%+v", err, e)
			}
			logger.Debug("Venus::GetBNBSupplyMaxWithdrawToken:: %+v,%+v,%+v", tr, query_address, vtokenAddr)
			err1 = err
			getCash = tr
			wg.Done()
			return
		}(getCash, err1)

		go func(z0 *big.Int, e error) {
			tr, err := instance.TotalBorrows(nil)
			if err != nil {
				logger.Warn("Venus::GetBNBSupplyMaxWithdrawToken::  TotalBorrows err:%+v,url:%+v", err, e)
			}
			logger.Debug("Venus::GetBNBSupplyMaxWithdrawToken:: %+v,%+v,%+v", tr, query_address, vtokenAddr)
			err2 = err
			totalBorrows = tr
			wg.Done()
			return
		}(totalBorrows, err2)

		go func(z0 *big.Int, e error) {
			tr, err := instance.TotalReserves(nil)
			if err != nil {
				logger.Warn("Venus::GetBNBSupplyMaxWithdrawToken::  totalReservers err:%+v,url:%+v", err, e)
			}
			logger.Debug("Venus::GetBNBSupplyMaxWithdrawToken:: %+v,%+v,%+v", tr, query_address, vtokenAddr)
			err3 = err
			totalReservers = tr
			wg.Done()
			return
		}(totalReservers, err3)

		go func(z0 *big.Int, e error) {
			tr, err := instance.TotalSupply(nil)
			if err != nil {
				logger.Warn("Venus::GetBNBSupplyMaxWithdrawToken::  totalSupply err:%+v,url:%+v", err, e)
			}
			logger.Debug("Venus::GetBNBSupplyMaxWithdrawToken:: %+v,%+v,%+v", tr, query_address, vtokenAddr)
			err4 = err
			totalSupply = tr
			wg.Done()
			return
		}(totalSupply, err4)

		wg.Wait()

		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			logger.Debug("err1:%+v,err2:%+v,err3:%+v,err4:%+v,err5:%+v", err1, err2, err3, err4)
			continue
		}

		logger.Debug("Venus::GetBNBSupplyMaxWithdrawToken:: %+v,%+v,%+v,%+v", getCash, totalBorrows, totalReservers, totalSupply)
		//得先乘 防止出现第一步为0
		out := big.NewInt(0)
		out.Add(getCash, totalBorrows).Sub(out, totalReservers).Mul(out, r).Div(out, totalSupply)
		return out, nil
	}
	return nil, errors.New("GetBNBSupplyMaxWithdrawToken: all provider failed")
}

func GetSupplyMaxWithdrawToken(vtokenAddr, query_address string, provider []string) (*big.Int, error) {
	for _, e := range provider {
		client, err := ethclient.Dial(e)
		if err != nil {
			logger.Error("Venus::GetSupplyMaxWithdrawToken:: dial err:%+v,%+v", err, e)
			continue
		}

		addr := common.HexToAddress(vtokenAddr)
		instance, err := translate.NewVtokenCaller(addr, client)
		if err != nil {
			logger.Error("Venus::GetSupplyMaxWithdrawToken::  NewPairCaller err:%+v,%+v", err, e)
			continue
		}

		query_addr := common.HexToAddress(query_address)
		r, err := instance.BalanceOf(nil, query_addr)
		if err != nil || r == nil {
			logger.Error("Venus::GetSupplyMaxWithdrawToken::  BalanceOf err:%+v,%+v", err, e)
			continue
		}

		if r.Cmp(big.NewInt(0)) <= 0 {
			logger.Debug("Venus::GetSupplyMaxWithdrawToken:: balance of is :%+v", r)
			return big.NewInt(0), nil
		}

		//然后估算能取出来多少
		wg := new(sync.WaitGroup)
		var (
			getCash        *big.Int
			totalBorrows   *big.Int
			totalReservers *big.Int
			totalSupply    *big.Int
			err1           error
			err2           error
			err3           error
			err4           error
		)

		wg.Add(4)

		go func(z0 *big.Int, e error) {
			tr, err := instance.GetCash(nil)
			if err != nil {
				logger.Warn("Venus::GetSupplyMaxWithdrawToken::  GetCash err:%+v,url:%+v", err, e)
			}
			logger.Debug("Venus::GetSupplyMaxWithdrawToken:: %+v,%+v,%+v", tr, query_address, vtokenAddr)
			err1 = err
			getCash = tr
			wg.Done()
			return
		}(getCash, err1)

		go func(z0 *big.Int, e error) {
			tr, err := instance.TotalBorrows(nil)
			if err != nil {
				logger.Warn("Venus::GetSupplyMaxWithdrawToken::  TotalBorrows err:%+v,url:%+v", err, e)
			}
			logger.Debug("Venus::GetSupplyMaxWithdrawToken:: %+v,%+v,%+v", tr, query_address, vtokenAddr)
			err2 = err
			totalBorrows = tr
			wg.Done()
			return
		}(totalBorrows, err2)

		go func(z0 *big.Int, e error) {
			tr, err := instance.TotalReserves(nil)
			if err != nil {
				logger.Warn("Venus::GetSupplyMaxWithdrawToken::  totalReservers err:%+v,url:%+v", err, e)
			}
			logger.Debug("Venus::GetSupplyMaxWithdrawToken:: %+v,%+v,%+v", tr, query_address, vtokenAddr)
			err3 = err
			totalReservers = tr
			wg.Done()
			return
		}(totalReservers, err3)

		go func(z0 *big.Int, e error) {
			tr, err := instance.TotalSupply(nil)
			if err != nil {
				logger.Warn("Venus::GetSupplyMaxWithdrawToken::  totalSupply err:%+v,url:%+v", err, e)
			}
			logger.Debug("Venus::GetSupplyMaxWithdrawToken:: %+v,%+v,%+v", tr, query_address, vtokenAddr)
			err4 = err
			totalSupply = tr
			wg.Done()
			return
		}(totalSupply, err4)

		wg.Wait()

		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			logger.Debug("err1:%+v,err2:%+v,err3:%+v,err4:%+v,err5:%+v", err1, err2, err3, err4)
			continue
		}

		logger.Debug("Venus::GetSupplyMaxWithdrawToken:: %+v,%+v,%+v,%+v", getCash, totalBorrows, totalReservers, totalSupply)
		//得先乘 防止出现第一步为0
		out := big.NewInt(0)
		out.Add(getCash, totalBorrows).Sub(out, totalReservers).Mul(out, r).Div(out, totalSupply)
		return out, nil
	}
	return nil, errors.New("GetSupplyMaxWithdrawToken: all provider failed")
}

//获取vai mint/suppley/borrow的xvs奖励

//有余额的操作 这个已经写完测试正确 暂不放在这里 先统一前端界面返回样式
func Get_MSB_XVS() {

}

//获取vai挖矿的奖励
func Get_VAI_XVS(query_address string, provider []string) (*big.Int, *big.Int, error) {
	for _, e := range provider {
		client, err := ethclient.Dial(e)
		if err != nil {
			logger.Error("Venus::Get_VAI_XVS:: dial err:%+v,%+v", err, e)
			continue
		}

		addr := common.HexToAddress("0x0667Eed0a0aAb930af74a3dfeDD263A73994f216")
		instance, err := translate.NewVaultCaller(addr, client)
		if err != nil {
			logger.Error("Venus::Get_VAI_XVS::  NewVaultCaller err:%+v,%+v", err, e)
			continue
		}

		query_addr := common.HexToAddress(query_address)

		//然后估算能取出来多少
		wg := new(sync.WaitGroup)
		var (
			pendingXVS    *big.Int
			depositAmount *big.Int
			err1          error
			err2          error
		)

		wg.Add(2)

		go func(z0 *big.Int, e error) {
			tr, err := instance.PendingXVS(nil, query_addr)
			if err != nil {
				logger.Warn("Venus::Get_VAI_XVS::  GetCash err:%+v,url:%+v", err, e)
			}
			logger.Debug("Venus::Get_VAI_XVS:: %+v,%+v", tr, query_address)
			err1 = err
			pendingXVS = tr
			wg.Done()
			return
		}(pendingXVS, err1)

		go func(z0 *big.Int, e error) {
			tr, err := instance.UserInfo(nil, query_addr)
			if err != nil {
				logger.Warn("Venus::Get_VAI_XVS::  UserInfo err:%+v,url:%+v", err, e)
			}
			logger.Debug("Venus::Get_VAI_XVS:: %+v,%+v", tr, query_address)
			err2 = err
			depositAmount = tr.Amount
			wg.Done()
			return
		}(depositAmount, err2)

		wg.Wait()

		if err1 != nil || err2 != nil {
			logger.Debug("err1:%+v,err2:%+v", err1, err2)
			continue
		}
		return depositAmount, pendingXVS, nil
	}
	return nil, nil, errors.New("Get_VAI_XVS: all provider failed")
}
