package services

import (
	"context"
	"fmt"
	"net"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/dapplink-labs/fee-estimator/config"
	"github.com/dapplink-labs/fee-estimator/proto/fee"
)

const MaxRecvMessageSize = 1024 * 1024 * 30000

type FeeService struct {
	RpcEndPoint string
	RpcPort     string

	fee.UnimplementedChainFeeServiceServer
	stopped atomic.Bool
}

func NewFeeRpcService(conf *config.Config) (*FeeService, error) {
	return &FeeService{}, nil
}

func (wbs *FeeService) Start(ctx context.Context) error {
	go func(wbs *FeeService) {
		rpcAddr := fmt.Sprintf("%s:%s", wbs.RpcEndPoint, wbs.RpcPort)
		log.Info("Rpc address", "rpcAddr", rpcAddr)
		listener, err := net.Listen("tcp", rpcAddr)
		if err != nil {
			log.Error("Could not start tcp listener. ")
		}

		opt := grpc.MaxRecvMsgSize(MaxRecvMessageSize)

		gs := grpc.NewServer(
			opt,
			grpc.ChainUnaryInterceptor(
				nil,
			),
		)

		reflection.Register(gs)

		fee.RegisterChainFeeServiceServer(gs, wbs)

		log.Info("grpc info", "addr", listener.Addr())

		if err := gs.Serve(listener); err != nil {
			log.Error("start rpc server fail", "err", err)
		}
	}(wbs)
	return nil
}

func (wbs *FeeService) Stop(ctx context.Context) error {
	wbs.stopped.Store(true)
	return nil
}

func (wbs *FeeService) Stopped() bool {
	return wbs.stopped.Load()
}
