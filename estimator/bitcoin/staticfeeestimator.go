package bitcoin

import "github.com/lightningnetwork/lnd/lnwallet/chainfee"

type StaticFeeEstimator struct {
	DefaultFee chainfee.SatPerKVByte
}

var _ BtcFeeEstimator = (*StaticFeeEstimator)(nil)

func NewStaticBtcFeeEstimator(defaultFee chainfee.SatPerKVByte) *StaticFeeEstimator {
	return &StaticFeeEstimator{
		DefaultFee: defaultFee,
	}
}

func (e *StaticFeeEstimator) Start() error {
	return nil
}

func (e *StaticFeeEstimator) Stop() error {
	return nil
}

func (e *StaticFeeEstimator) EstimateFeePerKb() chainfee.SatPerKVByte {
	return e.DefaultFee
}
