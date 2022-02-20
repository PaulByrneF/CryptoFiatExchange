package main

import (
	"flag"
	"fmt"
	"github.com/paulbyrnef/CryptoExchange/client"
	"log"
)

func main() {

	fiatCurrency := flag.String(
		"fiat", "USD", "The name of the fiat currency you would like to know the price of your crypto in",
	)

	cryptoName := flag.String(
		"crypto", "BTC", "Input the name of the CryptoCurrency you would like to know the price of",
	)
	flag.Parse()

	crypto, err := client.FetchCryptoExchange(*fiatCurrency, *cryptoName)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(crypto)
}
