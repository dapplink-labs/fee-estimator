package bitcoin

import "github.com/lightningnetwork/lnd/lnwallet/chainfee"

var _ BtcFeeEstimator = (*DynamicBtcFeeEstimator)(nil)

func (e *DynamicBtcFeeEstimator) Start() error {
	return e.estimator.Start()
}

func (e *DynamicBtcFeeEstimator) Stop() error {
	return e.estimator.Stop()
}

func (e *DynamicBtcFeeEstimator) EstimateFeePerKb() chainfee.SatPerKVByte {
	fee, err := e.estimator.EstimateFeePerKW(DefaultNumBlockForEstimation)

	if err != nil {
		return e.MaxFeeRate
	}

	estimatedFee := fee.FeePerKVByte()

	if estimatedFee < e.MinFeeRate {
		return e.MinFeeRate
	}

	if estimatedFee > e.MaxFeeRate {
		return e.MaxFeeRate
	}

	return estimatedFee
}
