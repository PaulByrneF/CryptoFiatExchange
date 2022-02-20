package model

import "fmt"

// Response to store the json response
type CryptoResponse []struct {
	Name              string `json:"name"`
	Price             string `json:"price"`
	Rank              string `json:"rank"`
	High              string `json:"high"`
	CirculatingSupply string `json:"circulating_supply"`
}

//Format data to string
func (c CryptoResponse) TextOutput() string {
	p := fmt.Sprintf(
		"Name: %s\nPrice: %s\nRank: %s\nHigh: %s\nCirculatingSupply: %s",
		c[0].Name, c[0].Price, c[0].Rank, c[0].High, c[0].CirculatingSupply)
	return p
}
