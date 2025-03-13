package services

import (
	"context"
	"github.com/dapplink-labs/fee-estimator/proto/fee"
)

func (fs *FeeService) GetSupportChains(ctx context.Context, req *fee.SupportChainsRequest) (*fee.SupportChainsResponse, error) {
	return &fee.SupportChainsResponse{
		Code:    fee.ReturnCode_SUCCESS,
		Msg:     "Support this chain",
		Support: true,
	}, nil
}

func (fs *FeeService) GetFeeByChain(ctx context.Context, req *fee.ChainFeeRequest) (*fee.ChainFeeResponse, error) {
	return &fee.ChainFeeResponse{
		Code:      fee.ReturnCode_SUCCESS,
		Msg:       "Support this chain",
		LowFee:    "0",
		NormalFee: "0",
		FastFee:   "0",
		OtherFee:  "0",
	}, nil
}
