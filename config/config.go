package config

import (
	"encoding/hex"
	"io"
	"os"
	"time"

	"github.com/dapplink-labs/fee-estimator/estimator/types"
)

type Config struct {
	Btcd                 *Btcd
	Bitcoind             *Bitcoind
	BtcNodeBackendConfig *BtcNodeBackendConfig
}

type Btcd struct {
	RPCHost        string
	RPCUser        string
	RPCPass        string
	RPCCert        string
	RawRPCCert     string
	BlockCacheSize uint64
}

type Bitcoind struct {
	RPCHost              string
	RPCUser              string
	RPCPass              string
	ZMQPubRawBlock       string
	ZMQPubRawTx          string
	EstimateMode         string
	PrunedNodeMaxPeers   int
	RPCPolling           bool
	BlockPollingInterval time.Duration
	TxPollingInterval    time.Duration
	BlockCacheSize       uint64
}

type BtcNodeBackendConfig struct {
	Nodetype            string
	WalletType          string
	FeeMode             string
	MinFeeRate          int64
	MaxFeeRate          int64
	Btcd                *Btcd
	Bitcoind            *Bitcoind `group:"bitcoind" namespace:"bitcoind"`
	EstimationMode      types.FeeEstimationMode
	ActiveNodeBackend   types.SupportedNodeBackend
	ActiveWalletBackend types.SupportedWalletBackend
}

func ReadCertFile(rawCert string, certFilePath string) ([]byte, error) {
	if rawCert != "" {
		rpcCert, err := hex.DecodeString(rawCert)
		if err != nil {
			return nil, err
		}
		return rpcCert, nil
	}
	certFile, err := os.Open(certFilePath)
	if err != nil {
		return nil, err
	}
	defer func(certFile *os.File) {
		err := certFile.Close()
		if err != nil {
			return
		}
	}(certFile)

	rpcCert, err := io.ReadAll(certFile)
	if err != nil {
		return nil, err
	}

	return rpcCert, nil
}
