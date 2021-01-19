package exchanges

import (
	"fmt"

	. "github.com/evzpav/crex"
	"github.com/evzpav/crex/exchanges/binancefutures"
	"github.com/evzpav/crex/exchanges/bitmex"
	"github.com/evzpav/crex/exchanges/bybit"
	"github.com/evzpav/crex/exchanges/deribit"
	"github.com/evzpav/crex/exchanges/hbdm"
	"github.com/evzpav/crex/exchanges/hbdmswap"
	"github.com/evzpav/crex/exchanges/okexfutures"
	"github.com/evzpav/crex/exchanges/okexswap"
)

func NewExchange(name string, opts ...ApiOption) Exchange {
	params := &Parameters{}

	for _, opt := range opts {
		opt(params)
	}

	return NewExchangeFromParameters(name, params)
}

func NewExchangeFromParameters(name string, params *Parameters) Exchange {
	switch name {
	case BinanceFutures:
		return binancefutures.NewBinanceFutures(params)
	case BitMEX:
		return bitmex.NewBitMEX(params)
	case Deribit:
		return deribit.NewDeribit(params)
	case Bybit:
		return bybit.NewBybit(params)
	case Hbdm:
		return hbdm.NewHbdm(params)
	case HbdmSwap:
		return hbdmswap.NewHbdmSwap(params)
	case OkexFutures:
		return okexfutures.NewOkexFutures(params)
	case OkexSwap:
		return okexswap.NewOkexSwap(params)
	default:
		panic(fmt.Sprintf("new exchange error [%v]", name))
	}
}
