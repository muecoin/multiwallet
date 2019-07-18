package util

 import "github.com/OpenBazaar/wallet-interface"

 type ExtCoinType wallet.CoinType

 func ExtendCoinType(coinType wallet.CoinType) ExtCoinType {
	return ExtCoinType(uint32(coinType))
}

 const (
	CoinTypeMonetaryUnit     ExtCoinType = 31
	CoinTypeMonetaryUnitTest             = 100031
)

 func (c *ExtCoinType) String() string {
	ct := wallet.CoinType(uint32(*c))
	str := ct.String()
	if str != "" {
		return str
	}

 	switch *c {
	case CoinTypeMonetaryUnit:
		return "MonetaryUnit"
	case CoinTypeMonetaryUnitTest:
		return "Testnet MonetaryUnit"
	default:
		return ""
	}
}

 func (c *ExtCoinType) CurrencyCode() string {
	ct := wallet.CoinType(uint32(*c))
	str := ct.CurrencyCode()
	if str != "" {
		return str
	}

 	switch *c {
	case CoinTypeMonetaryUnit:
		return "MUE"
	case CoinTypeMonetaryUnitTest:
		return "TMUE"
	default:
		return ""
	}
}

 func (c ExtCoinType) ToCoinType() wallet.CoinType {
	return wallet.CoinType(uint32(c))
}
