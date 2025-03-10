package types

import "fmt"

type FeeEstimationMode int
type SupportedNodeBackend int
type SupportedWalletBackend int

const (
	StaticFeeEstimation FeeEstimationMode = iota
	DynamicFeeEstimation
)

const (
	BitcoindNodeBackend SupportedNodeBackend = iota
	BtcdNodeBackend
)

const (
	BitcoindWalletBackend SupportedWalletBackend = iota
	BtcwalletWalletBackend
)

func NewNodeBackend(backend string) (SupportedNodeBackend, error) {
	switch backend {
	case "btcd":
		return BtcdNodeBackend, nil
	case "bitcoind":
		return BitcoindNodeBackend, nil
	default:
		return BtcdNodeBackend, fmt.Errorf("invalid node type: %s", backend)
	}
}

func NewWalletBackend(backend string) (SupportedWalletBackend, error) {
	switch backend {
	case "btcwallet":
		return BtcwalletWalletBackend, nil
	case "bitcoind":
		return BitcoindWalletBackend, nil
	default:
		return BtcwalletWalletBackend, fmt.Errorf("invalid wallet type: %s", backend)
	}
}
