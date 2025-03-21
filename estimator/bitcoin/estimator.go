package bitcoin

import (
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"

	"github.com/dapplink-labs/fee-estimator/estimator/types"
	"github.com/lightningnetwork/lnd/lnwallet/chainfee"

	"github.com/dapplink-labs/fee-estimator/config"
)

const (
	DefaultNumBlockForEstimation = 1
)

type BtcFeeEstimator interface {
	Start() error
	Stop() error
	EstimateFeePerKb() chainfee.SatPerKVByte
}

type DynamicBtcFeeEstimator struct {
	estimator  chainfee.Estimator
	MinFeeRate chainfee.SatPerKVByte
	MaxFeeRate chainfee.SatPerKVByte
}

func NewDynamicBtcFeeEstimator(cfg *config.BtcNodeBackendConfig, _ chaincfg.Params) (*DynamicBtcFeeEstimator, error) {
	minFeeRate := chainfee.SatPerKVByte(cfg.MinFeeRate)
	maxFeeRate := chainfee.SatPerKVByte(cfg.MaxFeeRate)
	switch cfg.ActiveNodeBackend {
	case types.BitcoindNodeBackend:
		rpcConfig := rpcclient.ConnConfig{
			Host:                 cfg.Bitcoind.RPCHost,
			User:                 cfg.Bitcoind.RPCUser,
			Pass:                 cfg.Bitcoind.RPCPass,
			DisableConnectOnNew:  true,
			DisableAutoReconnect: false,
			DisableTLS:           true,
			HTTPPostMode:         true,
		}
		est, err := chainfee.NewBitcoindEstimator(
			rpcConfig, cfg.Bitcoind.EstimateMode, maxFeeRate.FeePerKWeight(),
		)
		if err != nil {
			return nil, err
		}
		return &DynamicBtcFeeEstimator{
			estimator:  est,
			MinFeeRate: minFeeRate,
			MaxFeeRate: maxFeeRate,
		}, nil

	case types.BtcdNodeBackend:
		fmt.Println("BitcoindNodeBackend")
		cert, err := config.ReadCertFile(cfg.Btcd.RawRPCCert, cfg.Btcd.RPCCert)

		if err != nil {
			return nil, err
		}

		rpcConfig := rpcclient.ConnConfig{
			Host:                 cfg.Btcd.RPCHost,
			Endpoint:             "ws",
			User:                 cfg.Btcd.RPCUser,
			Pass:                 cfg.Btcd.RPCPass,
			Certificates:         cert,
			DisableTLS:           false,
			DisableConnectOnNew:  true,
			DisableAutoReconnect: false,
		}
		est, err := chainfee.NewBtcdEstimator(
			rpcConfig, maxFeeRate.FeePerKWeight(),
		)
		if err != nil {
			return nil, err
		}
		return &DynamicBtcFeeEstimator{
			estimator:  est,
			MinFeeRate: minFeeRate,
			MaxFeeRate: maxFeeRate,
		}, nil
	default:
		return nil, fmt.Errorf("unknown node backend: %v", cfg.ActiveNodeBackend)
	}
}
