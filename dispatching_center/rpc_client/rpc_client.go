package rpc_client

import (
	"kryptos_rai/dispatching_center/config"
	"kryptos_rai/use/logger"
	analysis "kryptos_rai/use/pb/address_analysis"
	"time"

	"google.golang.org/grpc"
)

//也可以搞成pool那种
var Aa_client_map map[string]analysis.AnalysisClient

func InitRPCClient() error {
	Aa_client_map = make(map[string]analysis.AnalysisClient)
	ipMap := config.GetIPMap()
	openIPList := []string{}
	for k, v := range ipMap {
		if v {
			openIPList = append(openIPList, k)
		}
	}
	for _, e := range openIPList {
		aa_rpc, err := grpc.Dial(e, grpc.WithInsecure(), grpc.WithTimeout(time.Duration(config.GetRPCConnTimeout())*time.Second))
		if err != nil {
			logger.Error("init aa rpc client with addr:%+v failed %v", e, err)
			return err
		}
		Aa_client_map[e] = analysis.NewAnalysisClient(aa_rpc)
	}
	return nil
}
