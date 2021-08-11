package bsc_analysis

import (
	"kryptos_rai/address_analysis/config"
	"kryptos_rai/address_analysis/data"
	"kryptos_rai/support_pools/bsc/venus/venus_method"
	"kryptos_rai/use/const_def"
	"kryptos_rai/use/logger"
	analysis "kryptos_rai/use/pb/address_analysis"
	"math/big"
	"strings"
)

func VenusAnalysis(req *analysis.HA_Single_Request, mr *data.WrapReplyList) {
	tmpNodeURL := config.GetBSCNodeURL()
	indexList := data.Perm(len(tmpNodeURL))
	nodeURL := []string{}
	for _, e := range indexList {
		nodeURL = append(nodeURL, tmpNodeURL[e])
	}

	fr := []*analysis.HA_Single_Reply{}
	r := new(analysis.HA_Single_Reply)
	r.PoolAddress = req.PoolAddress
	r.MasterChelfAddress = ""
	r.OutInfo = []*analysis.OutTokenInfo{}

	switch req.PoolType {
	case "SingleTokenStakePool":
		//vai挖xvs
		if len(req.InputTokenList) != 1 {
			logger.Error("VenusAnalysis::input token num should be 1,req: %+v", *req)
			return
		}
		if len(req.RewardTokenList) != 1 {
			logger.Error("VenusAnalysis::reward token num should be 1,req: %+v", *req)
			return
		}
		s0, s1, err := venus_method.Get_VAI_XVS(req.QueryAdddress, nodeURL)
		if err != nil {
			return
		}
		logger.Debug("VenusAnalysis::Get_VAI_XVS:%+v,%+v", s0, s1)
		if s0.Cmp(big.NewInt(0)) > 0 {
			out0 := new(analysis.OutTokenInfo)
			out0.TokenName = req.InputTokenList[0].Name
			out0.TokenAddress = req.InputTokenList[0].Address
			out0.TokenAmount = GetTokenAmount(s0, req.InputTokenList[0].Precision)
			out0.TokenType = const_def.STAKE
			r.OutInfo = append(r.OutInfo, out0)
		}

		if s1.Cmp(big.NewInt(0)) > 0 {
			out1 := new(analysis.OutTokenInfo)
			out1.TokenName = req.RewardTokenList[0].Name
			out1.TokenAddress = req.RewardTokenList[0].Address
			out1.TokenAmount = GetTokenAmount(s1, req.RewardTokenList[0].Precision)
			out1.TokenType = const_def.REWARD
			r.OutInfo = append(r.OutInfo, out1)

		}

	case "Lending":
		//借贷池
		if len(req.InputTokenList) != 1 {
			logger.Error("VenusAnalysis::input token num should be 1,req: %+v", *req)
			return
		}

		s0 := big.NewInt(0)
		if strings.ToLower(req.PoolAddress) == strings.ToLower("0xA07c5b74C9B40447a954e1466938b865b6BBea36") {
			t0, err := venus_method.GetBNBSupplyMaxWithdrawToken(req.QueryAdddress, nodeURL)
			if err != nil {
				return
			}
			s0 = t0
		} else {
			t0, err := venus_method.GetSupplyMaxWithdrawToken(req.PoolAddress, req.QueryAdddress, nodeURL)
			if err != nil {
				return
			}
			s0 = t0
		}

		logger.Debug("VenusAnalysis::Get_VAI_XVS:%+v", s0)
		if s0.Cmp(big.NewInt(0)) > 0 {
			out0 := new(analysis.OutTokenInfo)
			out0.TokenName = req.InputTokenList[0].Name
			out0.TokenAddress = req.InputTokenList[0].Address
			out0.TokenAmount = GetTokenAmount(s0, req.InputTokenList[0].Precision)
			out0.TokenType = const_def.STAKE
			r.OutInfo = append(r.OutInfo, out0)
		}

	default:
		logger.Error("VenusAnalysis::not support this pancake pool type:%+v", req.PoolType)
		return
	}

	fr = append(fr, r)

	//logger.Debug("queryaddress:%+v,pool_address:%+v,fr:%+v", req.QueryAdddress, req.PoolAddress, fr)

	mr.Mu.Lock()
	for _, x := range fr {
		if len(x.OutInfo) > 0 {
			mr.ReplyList = append(mr.ReplyList, x)
		}
	}
	mr.Mu.Unlock()
	return
}
