package main

import (
	"fmt"
	"log"

	"github.com/apisit/binance-go"
)

// MarketPrices get prices from binance
func MarketPrices() string {
	p, err := binance.Market().Prices()
	if err != nil {
		log.Printf("err = %v", err)
		return ""
	}
	return fmt.Sprintf("el precio de bitcoin en UESEDETE es de %s", p["BTCUSDT"].Price)
}
