package rpc_server

import (
	"context"
	"kryptos_rai/dispatching_center/config"
	"kryptos_rai/dispatching_center/rpc_client"
	"kryptos_rai/use/logger"
	analysis "kryptos_rai/use/pb/address_analysis"
	"math/rand"
	"net"

	"google.golang.org/grpc"
)

type AnalysisService struct{}

func (s *AnalysisService) HandleAnalysis(ctx context.Context, in *analysis.HA_Request) (*analysis.HA_Reply, error) {
	logger.Info("HandleAnalysis:: receive request:%+v", *in)
	ipMap := config.GetIPMap()
	openIPList := []string{}
	for k, v := range ipMap {
		if v {
			openIPList = append(openIPList, k)
		}
	}
	r := new(analysis.HA_Reply)
	if len(openIPList) > 0 {
		index := rand.Intn(len(openIPList))
		ip := openIPList[index]
		client := rpc_client.Aa_client_map[ip]
		tr, err := client.HandleAnalysis(context.Background(), in)
		if err != nil || tr == nil {
			r.ReplyList = []*analysis.HA_Single_Reply{}
		} else {
			r = tr
		}

	} else {
		r.ReplyList = []*analysis.HA_Single_Reply{}
	}
	return r, nil
}

func InitRPCService() error {
	logger.Notic("InitRPCService::start analysis rpc server")
	addr := config.GetListenAddr()
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Notic("InitRPCService::listen addr failed: %+v", err)
		return err
	}
	s := grpc.NewServer()

	analysis.RegisterAnalysisServer(s, &AnalysisService{})
	if err = s.Serve(listen); err != nil {
		logger.Info("InitRPCService::failed to start rpc service at: %v error: %v", addr, err)
		return err
	}
	return nil
}
