package eboost

import (
	"blockbook/bchain/coins/btc"

	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
)

const (
	MainnetMagic wire.BitcoinNet = 0xfbcfccd4
	TestnetMagic wire.BitcoinNet = 0xfbcdccd3
	RegtestMagic wire.BitcoinNet = 0xfabfb5da
)

var (
	MainNetParams chaincfg.Params
	TestNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic
	MainNetParams.PubKeyHashAddrID = []byte{23}
	MainNetParams.ScriptHashAddrID = []byte{30}
	MainNetParams.Bech32HRPSegwit = "ebst"

	TestNetParams = chaincfg.TestNet3Params
	TestNetParams.Net = TestnetMagic
	TestNetParams.PubKeyHashAddrID = []byte{33}
	TestNetParams.ScriptHashAddrID = []byte{55}
	TestNetParams.Bech32HRPSegwit = "tala"
}

// EboostParser handle
type EboostParser struct {
	*btc.BitcoinParser
}

// NewEboostParser returns new EboostParser instance
func NewEboostParser(params *chaincfg.Params, c *btc.Configuration) *EboostParser {
	return &EboostParser{BitcoinParser: btc.NewBitcoinParser(params, c)}
}

// GetChainParams contains network parameters for the main Eboost network,
// and the test Eboost network
func GetChainParams(chain string) *chaincfg.Params {
	// register bitcoin parameters in addition to eboost parameters
	// eboost has dual standard of addresses and we want to be able to
	// parse both standards
	if !chaincfg.IsRegistered(&chaincfg.MainNetParams) {
		chaincfg.RegisterBitcoinParams()
	}
	if !chaincfg.IsRegistered(&MainNetParams) {
		err := chaincfg.Register(&MainNetParams)
		if err == nil {
			err = chaincfg.Register(&TestNetParams)
		}
		if err != nil {
			panic(err)
		}
	}
	switch chain {
	case "test":
		return &TestNetParams
	default:
		return &MainNetParams
	}
}
