package bsc_analysis

import (
	"fmt"
	"kryptos_rai/address_analysis/config"
	"kryptos_rai/address_analysis/data"
	"kryptos_rai/support_pools/bsc/pancake/pancake_method"
	"kryptos_rai/use/const_def"
	"kryptos_rai/use/logger"
	analysis "kryptos_rai/use/pb/address_analysis"
	"math/big"
)

//保留8位
func GetTokenAmount(s *big.Int, precision int32) string {
	slen := len(s.String())
	ft, _ := new(big.Int).SetString(s.String(), 0)
	prFormat := "%0" + fmt.Sprintf("%+v", precision) + "d"
	prString := "1" + fmt.Sprintf(prFormat, 0)
	pr, _ := new(big.Int).SetString(prString, 0)
	mod := big.NewInt(0)
	d := big.NewInt(0)
	d.Div(ft, pr)
	mod.Mod(ft, pr)
	modLen := len(mod.String())
	logger.Debug("GetTokenAmount::slen:%+v,modlen:%+v,mod:%+v,d:%+v", slen, modLen, mod.String(), d.String())
	ec := slen - int(precision)
	if ec >= 0 {
		if mod.Cmp(big.NewInt(0)) == 0 || int(precision)-modLen >= 8 {
			return d.String()
		} else {
			format := "%0" + fmt.Sprintf("%+v", precision) + "d"
			m := fmt.Sprintf(format, mod.Int64())
			fs := m
			if len(m) > 8 {
				fs = m[:8]
			}

			if fs == "00000000" {
				return d.String()
			} else {
				return d.String() + "." + fs
			}

		}
	} else {
		format := "%0" + fmt.Sprintf("%+v", precision) + "d"
		m := fmt.Sprintf(format, mod.Int64())
		fs := m
		if len(m) > 8 {
			fs = m[:8]
		}
		if fs == "00000000" {
			return "0"
		} else {
			return "0." + fs
		}

	}
}

func PancakeAnalysis(req *analysis.HA_Single_Request, mr *data.WrapReplyList) {
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
		if len(req.InputTokenList) != 1 {
			logger.Error("PancakeAnalysis::input token num should be 1,req: %+v", *req)
			return
		}
		if len(req.RewardTokenList) != 1 {
			logger.Error("PancakeAnalysis::reward token num should be 1,req: %+v", *req)
			return
		}
		s0, s1, err := pancake_method.GetPoolUnderlyingAssets(req.PoolAddress, req.QueryAdddress, nodeURL)
		if err != nil {
			return
		}
		logger.Debug("PancakeAnalysis::GetPoolUnderlyingAssets:%+v,%+v", s0, s1)
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

	case "MasterChelf":
		for _, e := range req.MasterChelfList {
			tmpr := new(analysis.HA_Single_Reply)
			tmpr.PoolAddress = req.PoolAddress
			tmpr.MasterChelfAddress = e.Address
			s0, s1, err := pancake_method.GetFarmUnderlyingAssets(req.PoolAddress, e.Address, req.QueryAdddress, e.Pid, nodeURL)
			if err != nil {
				continue
			}
			logger.Debug("PancakeAnalysis::GetFarmUnderlyingAssets:%+v,%+v", s0, s1)

			if s0.Cmp(big.NewInt(0)) > 0 {
				out0 := new(analysis.OutTokenInfo)
				out0.TokenName = e.Name
				out0.TokenAddress = e.Address
				out0.TokenAmount = GetTokenAmount(s0, int32(e.Precision))
				out0.TokenType = const_def.STAKE
				tmpr.OutInfo = append(tmpr.OutInfo, out0)
			}

			if s1.Cmp(big.NewInt(0)) > 0 {
				out1 := new(analysis.OutTokenInfo)
				out1.TokenName = e.Name
				out1.TokenAddress = e.Address
				//cake固定18
				out1.TokenAmount = GetTokenAmount(s1, 18)
				out1.TokenType = const_def.REWARD
				tmpr.OutInfo = append(tmpr.OutInfo, out1)
			}
			fr = append(fr, tmpr)
		}
	case "DexPair":

		if len(req.InputTokenList) != 2 {
			logger.Error("PancakeAnalysis::input token num should is 2 req:%+v", *req)
			return
		}

		r0Info := req.InputTokenList[0]
		r1Info := req.InputTokenList[1]
		if req.InputTokenList[0].Address > req.InputTokenList[1].Address {
			r0Info = req.InputTokenList[1]
			r1Info = req.InputTokenList[0]
		}

		tr, r0, r1, err := pancake_method.GetPairCertificateBalance(req.CertificateAddress, req.QueryAdddress, nodeURL)
		if err != nil {
			return
		}
		logger.Debug("PancakeAnalysis::GetPairCertificateBalance:%+v,%+v,%+v", tr, r0, r1)
		if r0.Cmp(big.NewInt(0)) > 0 {
			out := new(analysis.OutTokenInfo)
			out.TokenName = r0Info.Name
			out.TokenAddress = r0Info.Address
			out.TokenType = const_def.TAKEOUT
			out.TokenAmount = GetTokenAmount(r0, r0Info.Precision)
			r.OutInfo = append(r.OutInfo, out)

		}

		if r1.Cmp(big.NewInt(0)) > 0 {
			out := new(analysis.OutTokenInfo)
			out.TokenName = r1Info.Name
			out.TokenAddress = r1Info.Address
			out.TokenType = const_def.TAKEOUT
			out.TokenAmount = GetTokenAmount(r1, r1Info.Precision)
			r.OutInfo = append(r.OutInfo, out)
		}

	case "Vault":
		if len(req.InputTokenList) != 1 {
			logger.Error("PancakeAnalysis::input token num should is 1 req:%+v", *req)
			return
		}
		if len(req.RewardTokenList) != 1 {
			logger.Error("PancakeAnalysis::reward token num should is 1 req:%+v", *req)
			return
		}
		s0, s1, err := pancake_method.GetVaultAssets(req.QueryAdddress, nodeURL)
		if err != nil {
			return
		}
		logger.Debug("PancakeAnalysis::GetVaultAssets:%+v,%+v", s0, s1)

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
	default:
		logger.Error("PancakeAnalysis::not support this pancake pool type:%+v", req.PoolType)
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
