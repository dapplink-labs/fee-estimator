package config

import (
	"encoding/hex"
	"io"
	"os"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/dapplink-labs/fee-estimator/estimator/types"
	"github.com/dapplink-labs/fee-estimator/flags"
)

type Config struct {
	Migrations string
	RpcServer  ServerConfig
	RestServer ServerConfig
	Metrics    ServerConfig
	MasterDB   DBConfig
	SlaveDB    DBConfig
	BtcConfig  BtcConfig
}

type DBConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

type ServerConfig struct {
	Host string
	Port int
}

type BtcConfig struct {
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

func NewConfig(ctx *cli.Context) Config {
	return Config{
		Migrations: ctx.String(flags.MigrationsFlag.Name),
		RpcServer: ServerConfig{
			Host: ctx.String(flags.RpcHostFlag.Name),
			Port: ctx.Int(flags.RpcPortFlag.Name),
		},
		RestServer: ServerConfig{
			Host: ctx.String(flags.HttpHostFlag.Name),
			Port: ctx.Int(flags.HttpPortFlag.Name),
		},
		Metrics: ServerConfig{
			Host: ctx.String(flags.MetricsHostFlag.Name),
			Port: ctx.Int(flags.MetricsPortFlag.Name),
		},
		MasterDB: DBConfig{
			Host:     ctx.String(flags.MasterDbHostFlag.Name),
			Port:     ctx.Int(flags.MasterDbPortFlag.Name),
			Name:     ctx.String(flags.MasterDbNameFlag.Name),
			User:     ctx.String(flags.MasterDbUserFlag.Name),
			Password: ctx.String(flags.MasterDbPasswordFlag.Name),
		},
		SlaveDB: DBConfig{
			Host:     ctx.String(flags.SlaveDbHostFlag.Name),
			Port:     ctx.Int(flags.SlaveDbPortFlag.Name),
			Name:     ctx.String(flags.SlaveDbNameFlag.Name),
			User:     ctx.String(flags.SlaveDbUserFlag.Name),
			Password: ctx.String(flags.SlaveDbPasswordFlag.Name),
		},
	}
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
