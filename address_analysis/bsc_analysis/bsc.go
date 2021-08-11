package bsc_analysis

import (
	"fmt"
	"kryptos_rai/address_analysis/data"
	"kryptos_rai/use/logger"
	analysis "kryptos_rai/use/pb/address_analysis"
	"math/big"
	"strings"
	"sync"
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

func BSCAnalysisSingle(req *analysis.HA_Single_Request, mr *data.WrapReplyList, wg *sync.WaitGroup) {
	switch strings.ToLower(req.Platform) {
	case "pancakeswap":
		PancakeAnalysis(req, mr)
	case "venus":
		VenusAnalysis(req, mr)
	case "autofarm":
		AutoFarmAnalysis(req, mr)
	default:
		logger.Error("BSCAnalysisSingle::not support this platform: %+v", req.Platform)
	}
	wg.Done()
}
