package rpc_server

import (
	"context"
	"errors"
	"kryptos_rai/address_analysis/bsc_analysis"
	"kryptos_rai/address_analysis/config"
	"kryptos_rai/address_analysis/data"
	"kryptos_rai/use/logger"
	analysis "kryptos_rai/use/pb/address_analysis"
	"net"
	"strings"
	"sync"

	"google.golang.org/grpc"
)

type AnalysisService struct{}

func (s *AnalysisService) HandleAnalysis(ctx context.Context, in *analysis.HA_Request) (*analysis.HA_Reply, error) {
	//目前只有bsc
	logger.Info("HandleAnalysis:: receive request: %+v", *in)
	mr := new(data.WrapReplyList)
	replyList := []*analysis.HA_Single_Reply{}
	mr.ReplyList = replyList
	mr.Mu = new(sync.RWMutex)

	switch strings.ToLower(in.Chain) {
	case "bsc":
		wg := new(sync.WaitGroup)
		for _, e := range in.RequestInfo {
			wg.Add(1)
			go bsc_analysis.BSCAnalysisSingle(e, mr, wg)
		}
		wg.Wait()
	default:
		r := new(analysis.HA_Reply)
		r.ReplyList = []*analysis.HA_Single_Reply{}
		logger.Error("HandleAnalysis:: not support this chain: %+v", in.Chain)
		return r, errors.New("not support this chain")
	}

	r := new(analysis.HA_Reply)
	r.ReplyList = mr.ReplyList
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
