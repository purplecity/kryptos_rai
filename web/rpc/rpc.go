package rpc

import (
	"context"
	"kryptos_rai/use/logger"
	analysis "kryptos_rai/use/pb/address_analysis"
	"kryptos_rai/web/config"
	"time"

	"google.golang.org/grpc"
)

var aa_client analysis.AnalysisClient

func InitRPCClient() error {
	addr := config.GetRPCAddr()
	aa_rpc, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithTimeout(time.Duration(config.GetRPCConnTimeout())*time.Second))
	if err != nil {
		logger.Error("init aa rpc client failed %v", addr)
		return err
	}
	aa_client = analysis.NewAnalysisClient(aa_rpc)
	return nil
}

func RPCAddressAnalysis(req *analysis.HA_Request) (*analysis.HA_Reply, error) {

	r, err := aa_client.HandleAnalysis(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return r, nil
}
