package client

import (
	"encoding/json"
	"github.com/paulbyrnef/CryptoExchange/model"
	"io"
	"log"
	"net/http"
)

func FetchCryptoExchange(fiat, crypto string) (string, error) {
	URL := "https://api.nomics.com/v1/currencies/ticker?key=3990ec554a414b59dd85d29b2286dd85&interval=1d&ids=" + crypto + "&convert=" + fiat

	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal("ooopsss an error occurred, please try again")
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	//Create a variable of the same type as our model
	var cResp model.CryptoResponse

	//Decode the data
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		log.Fatal(err)
	}

	//Invoke the text output function & return it with nil as the error value
	return cResp.TextOutput(), nil
}
