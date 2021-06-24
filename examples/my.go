package main

import (
	"fmt"
	"log"
	"time"

	. "github.com/evzpav/crex"
	"github.com/evzpav/crex/exchanges"
)

type BasicStrategy struct {
	StrategyBase
}

func (s *BasicStrategy) OnInit() error {
	symbols := []string{"BTCUSDT", "ETHUSDT", "DOTUSDT", "ADAUSDT"}
	candlePeriods := []string{"1m", "5m", "15m", "1h", "4h"}

	for _, ss := range symbols {
		for _, cp := range candlePeriods {
			_, err := s.Exchange.GetRecords(ss, cp, 0, 0, 1000)
			if err != nil {
				log.Printf(err.Error())
			}

			info, err := s.Exchange.IO("", "")
			if err != nil {
				log.Printf(err.Error())
			}
			fmt.Printf("\n\ninfo:\n %v\n", info)
		}

	}
	return nil
}

func (s *BasicStrategy) OnTick() error {
	// currency := "USDT"
	// symbol := "BTCUSDT"

	// balance, err := s.Exchange.GetBalance(currency)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("balance: %#v", balance)

	// positions, err := s.Exchange.GetPositions(symbol)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, p := range positions {
	// 	log.Printf("positions: %#v", *p)
	// }

	// openOrders, err := s.Exchange.GetOpenOrders(symbol)
	// if err != nil {
	// 	log.Printf(err.Error())
	// }

	// for _, oo := range openOrders {
	// 	log.Printf("openOrders: %#v", *oo)
	// }

	// externalID := "13915856387"
	// order, err := s.Exchange.GetOrder(symbol, externalID)
	// if err != nil {
	// 	log.Printf(err.Error())
	// }

	// log.Printf("order: %#v", *order)
	// log.Printf("order status: %#v", order.Status.String())

	// symbols := []string{"BTCUSDT", "ETHUSDT"}
	// candlePeriods := []string{"1m", "5m", "15m", "1h", "4h"}

	// for _, ss := range symbols {
	// 	for _, cp := range candlePeriods {
	// 		_, err := s.Exchange.GetRecords(ss, cp, 0, 0, 1000)
	// 		if err != nil {
	// 			log.Printf(err.Error())
	// 		}

	// 		info, err := s.Exchange.IO("", "")
	// 		if err != nil {
	// 			log.Printf(err.Error())
	// 		}
	// 		fmt.Printf("\n\ninfo:\n %v\n", info)
	// 	}

	// }

	// s.Exchange.GetOrderBook(symbol, 10)

	// s.Exchange.OpenLong(symbol, OrderTypeLimit, 5000, 10)
	// s.Exchange.CloseLong(symbol, OrderTypeLimit, 6000, 10)

	// s.Exchange.PlaceOrder(symbol,
	// 	Buy, OrderTypeLimit, 1000.0, 10, OrderPostOnlyOption(true))

	// s.Exchange.GetOpenOrders(symbol)
	// s.Exchange.GetPositions(symbol)
	return nil
}

func (s *BasicStrategy) Run() error {
	// run loop
	for {
		s.OnTick()
		time.Sleep(100 * time.Millisecond)
	}

}

func (s *BasicStrategy) OnExit() error {
	return nil
}

func main() {
	exchange := exchanges.NewExchange(exchanges.BinanceFutures,
		// ApiProxyURLOption("socks5://127.0.0.1:1080"), // 使用代理
		// ApiAccessKeyOption("rgUFP4t3neU2dmhaxqn3CtEcmmm0xvkbDNjp7KxRW3P0h39asWMEYazYC9NouNts"),
		// ApiSecretKeyOption("ObsY9FW7tXwe7nvrSoQxJgALLy6q7B2AMeGznmalLOmjqSoEfk9VVx9NTsrxuztY"),
		ApiAccessKeyOption(""),
		ApiSecretKeyOption(""),
		// ApiTestnetOption(true)
	)

	s := &BasicStrategy{}

	s.Setup(TradeModeLiveTrading, exchange)

	s.OnInit()
	s.Run()
	s.OnExit()
}
