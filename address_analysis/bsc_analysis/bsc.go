package bsc_analysis

import (
	"kryptos_rai/address_analysis/data"
	"kryptos_rai/use/logger"
	analysis "kryptos_rai/use/pb/address_analysis"
	"strings"
	"sync"
)

func BSCAnalysisSingle(req *analysis.HA_Single_Request, mr *data.WrapReplyList, wg *sync.WaitGroup) {
	switch strings.ToLower(req.Platform) {
	case "pancakeswap":
		PancakeAnalysis(req, mr)
	default:
		logger.Error("BSCAnalysisSingle::not support this platform: %+v", req.Platform)
	}
	wg.Done()
}
