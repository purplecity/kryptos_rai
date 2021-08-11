package bsc_analysis

import (
	"kryptos_rai/address_analysis/config"
	"kryptos_rai/address_analysis/data"
	"kryptos_rai/support_pools/bsc/autofarm/autofarm_method"
	"kryptos_rai/use/const_def"
	"kryptos_rai/use/logger"
	analysis "kryptos_rai/use/pb/address_analysis"
	"math/big"
	"strings"
)

func AutoFarmAnalysis(req *analysis.HA_Single_Request, mr *data.WrapReplyList) {
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
	case "MasterChelf":
		for _, e := range req.MasterChelfList {
			tmpr := new(analysis.HA_Single_Reply)
			tmpr.PoolAddress = req.PoolAddress
			tmpr.MasterChelfAddress = e.Address

			s0, s1 := big.NewInt(0), big.NewInt(0)
			if strings.ToLower(e.Address) == strings.ToLower("0xa184088a740c695E156F91f5cC086a06bb78b827") {

				t0, err := autofarm_method.GetSingleAutoFarmAssets(req.QueryAdddress, 0, nodeURL)
				if err != nil {
					continue
				}
				s0 = t0
			} else {
				t0, t1, err := autofarm_method.GetAutoFarmAssets(req.QueryAdddress, e.Pid, nodeURL)
				if err != nil {
					continue
				}
				s0, s1 = t0, t1
			}

			logger.Debug("AutoFarmAnalysis::GetAutoFarmAssets:%+v,%+v", s0, s1)

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

	default:
		logger.Error("AutoFarmAnalysis::not support this pancake pool type:%+v", req.PoolType)
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
