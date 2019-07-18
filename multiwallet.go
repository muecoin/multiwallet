package multiwallet

import (
	"errors"
	"github.com/muecoin/multiwallet/phore"
	"strings"
	"time"

	"github.com/muecoin/multiwallet/bitcoin"
	"github.com/muecoin/multiwallet/bitcoincash"
	"github.com/muecoin/multiwallet/client/blockbook"
	"github.com/muecoin/multiwallet/config"
	"github.com/muecoin/multiwallet/litecoin"
	"github.com/muecoin/multiwallet/service"
	"github.com/muecoin/multiwallet/zcash"
	"github.com/muecoin/multiwallet/util"
	"github.com/OpenBazaar/wallet-interface"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/op/go-logging"
	"github.com/tyler-smith/go-bip39"
)

var log = logging.MustGetLogger("multiwallet")

var UnsuppertedCoinError = errors.New("multiwallet does not contain an implementation for the given coin")

type MultiWallet map[util.ExtCoinType]wallet.Wallet

func NewMultiWallet(cfg *config.Config) (MultiWallet, error) {
	log.SetBackend(logging.AddModuleLevel(cfg.Logger))
	service.Log = log
	blockbook.Log = log

	if cfg.Mnemonic == "" {
		ent, err := bip39.NewEntropy(128)
		if err != nil {
			return nil, err
		}
		mnemonic, err := bip39.NewMnemonic(ent)
		if err != nil {
			return nil, err
		}
		cfg.Mnemonic = mnemonic
		cfg.CreationDate = time.Now()
	}

	multiwallet := make(MultiWallet)
	var err error
	for _, coin := range cfg.Coins {
		var w wallet.Wallet
		switch coin.CoinType {
		case util.CoinTypeMonetaryUnit:
			var params chaincfg.Params
			if cfg.Params.Name == monetaryunit.MonetaryUnitMainNetParams.Name {
				params = monetaryunit.MonetaryUnitMainNetParams
			} else {
				params = monetaryunit.MonetaryUnitTestNetParams
			}
			w, err = monetaryunit.NewMonetaryUnitWallet(coin, cfg.Mnemonic, &params, cfg.Proxy, cfg.Cache, cfg.DisableExchangeRates)
			if err != nil {
				return nil, err
			}
			if cfg.Params.Name == monetaryunit.MonetaryUnitMainNetParams.Name {
				multiwallet[util.CoinTypeMonetaryUnit] = w
			} else {
				multiwallet[util.CoinTypeMonetaryUnitTest] = w
			}
		case util.ExtendCoinType(wallet.Bitcoin):
			w, err = bitcoin.NewBitcoinWallet(coin, cfg.Mnemonic, cfg.Params, cfg.Proxy, cfg.Cache, cfg.DisableExchangeRates)
			if err != nil {
				return nil, err
			}
			if cfg.Params.Name == chaincfg.MainNetParams.Name {
				multiwallet[util.ExtendCoinType(wallet.Bitcoin)] = w
			} else {
				multiwallet[util.ExtendCoinType(wallet.TestnetBitcoin)] = w
			}
		case util.ExtendCoinType(wallet.BitcoinCash):
			w, err = bitcoincash.NewBitcoinCashWallet(coin, cfg.Mnemonic, cfg.Params, cfg.Proxy, cfg.Cache, cfg.DisableExchangeRates)
			if err != nil {
				return nil, err
			}
			if cfg.Params.Name == chaincfg.MainNetParams.Name {
				multiwallet[util.ExtendCoinType(wallet.BitcoinCash)] = w
			} else {
				multiwallet[util.ExtendCoinType(wallet.TestnetBitcoinCash)] = w
			}
		case util.ExtendCoinType(wallet.Zcash):
			w, err = zcash.NewZCashWallet(coin, cfg.Mnemonic, cfg.Params, cfg.Proxy, cfg.Cache, cfg.DisableExchangeRates)
			if err != nil {
				return nil, err
			}
			if cfg.Params.Name == chaincfg.MainNetParams.Name {
				multiwallet[util.ExtendCoinType(wallet.Zcash)] = w
			} else {
				multiwallet[util.ExtendCoinType(wallet.TestnetZcash)] = w
			}
		case util.ExtendCoinType(wallet.Litecoin):
			w, err = litecoin.NewLitecoinWallet(coin, cfg.Mnemonic, cfg.Params, cfg.Proxy, cfg.Cache, cfg.DisableExchangeRates)
			if err != nil {
				return nil, err
			}
			if cfg.Params.Name == chaincfg.MainNetParams.Name {
				multiwallet[util.ExtendCoinType(wallet.Litecoin)] = w
			} else {
				multiwallet[util.ExtendCoinType(wallet.TestnetLitecoin)] = w
			}
			//case wallet.Ethereum:
			//w, err = eth.NewEthereumWallet(coin, cfg.Mnemonic, cfg.Proxy)
			//if err != nil {
			//return nil, err
			//}
			//multiwallet[coin.CoinType] = w
		}
	}
	return multiwallet, nil
}

func (w *MultiWallet) Start() {
	for _, wallet := range *w {
		wallet.Start()
	}
}

func (w *MultiWallet) Close() {
	for _, wallet := range *w {
		wallet.Close()
	}
}

func (w *MultiWallet) WalletForCurrencyCode(currencyCode string) (wallet.Wallet, error) {
	for _, wl := range *w {
		if strings.ToUpper(wl.CurrencyCode()) == strings.ToUpper(currencyCode) || strings.ToUpper(wl.CurrencyCode()) == "T"+strings.ToUpper(currencyCode) {
			return wl, nil
		}
	}
	return nil, UnsuppertedCoinError
}
