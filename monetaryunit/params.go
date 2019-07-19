package monetaryunit

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/wire"
	"math/big"
	"time"
)

// These variables are the chain proof-of-work limit parameters for each default
// network.
var (
	// bigOne is 1 represented as a big.Int.  It is defined here to avoid
	// the overhead of creating it multiple times.
	bigOne = big.NewInt(1)

	// mainPowLimit is the highest proof of work value a Bitcoin block can
	// have for the main network.  It is the value 2^224 - 1.
	mainPowLimit = new(big.Int).Sub(new(big.Int).Lsh(bigOne, 255), bigOne)
)

const (
	// MainNet represents the main bitcoin network.
	MainMonetaryUnitNet wire.BitcoinNet = 0x504852 // MUE
	MainMonetaryUnitTestNet wire.BitcoinNet = 0x54504852 // TMUE
)

var MonetaryUnitMainNetParams = chaincfg.Params{
	Name:        "mainMonetaryUnit",
	Net:         MainMonetaryUnitNet,
	DefaultPort: "19687",
	DNSSeeds: []chaincfg.DNSSeed{
		{"dns1.monetaryunit.org", true},
		{"dns2.monetaryunit.org", true},
	},

	// Chain parameters
	GenesisBlock:             nil, // unused
	GenesisHash:              nil, // unused
	PowLimit:                 mainPowLimit,
	PowLimitBits:             0x207fffff,
	BIP0034Height:            0, // unused
	BIP0065Height:            0, // unused
	BIP0066Height:            0, // unused
	CoinbaseMaturity:         50,
	TargetTimespan:           time.Second * 400, // 400 seconds
	TargetTimePerBlock:       time.Second * 40, // 40 seconds
	RetargetAdjustmentFactor: 4,
	ReduceMinDifficulty:      false,
	MinDiffReductionTime:     0,
	GenerateSupported:        true,

	// Checkpoints ordered from oldest to newest.
	Checkpoints: []chaincfg.Checkpoint{},

	// Mempool parameters
	RelayNonStdTxs: false,

	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173.
	Bech32HRPSegwit: "mue", // always bc for main net

	// Address encoding magics
	PubKeyHashAddrID: 0x37, // starts with 1
	ScriptHashAddrID: 0x0d, // starts with 3
	PrivateKeyID:     0xd4, // starts with 5 (uncompressed) or K (compressed)

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x02, 0x2d, 0x25, 0x33}, // starts with xprv
	HDPublicKeyID:  [4]byte{0x02, 0x21, 0x31, 0x2b}, // starts with xpub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 0x8000001f,
}

var MonetaryUnitTestNetParams = MonetaryUnitMainNetParams
