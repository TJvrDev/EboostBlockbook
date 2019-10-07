package eboost

import (
	"blockbook/bchain"
	"blockbook/bchain/coins/btc"
	"encoding/json"

	"github.com/golang/glog"
)

// EboostRPC is an interface to JSON-RPC bitcoind service.
type EboostRPC struct {
	*btc.BitcoinRPC
}

// NewEboostRPC returns new EboostRPC instance.
func NewEboostRPC(config json.RawMessage, pushHandler func(bchain.NotificationType)) (bchain.BlockChain, error) {
	b, err := btc.NewBitcoinRPC(config, pushHandler)
	if err != nil {
		return nil, err
	}

	s := &EboostRPC{
		b.(*btc.BitcoinRPC),
	}
	s.RPCMarshaler = btc.JSONMarshalerV2{}
	s.ChainConfig.SupportsEstimateFee = false

	return s, nil
}

// Initialize initializes EboostRPC instance.
func (b *EboostRPC) Initialize() error {
	ci, err := b.GetChainInfo()
	if err != nil {
		return err
	}

	chainName := ci.Chain

	params := GetChainParams(chainName)

	// always create parser
	b.Parser = NewEboostParser(params, b.ChainConfig)

	


	// parameters for getInfo request
	if params.Net == MainnetMagic {
		b.Testnet = false
		b.Network = "livenet"
	} else {
		b.Testnet = true
		b.Network = "testnet"
	}

	glog.Info("rpc: block chain ", params.Name)

	return nil
}
