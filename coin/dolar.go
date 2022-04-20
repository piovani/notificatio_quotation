package coin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
)

const URL = "https://economia.awesomeapi.com.br/last/USD-BRL"

type ResponseDolar struct {
	Usdbrl struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func GetQuotationDolar() float64 {
	responseDolar, err := makeRequest()
	if err != nil {
		return returnError(err)
	}

	dolar, err := getDolarFormted(responseDolar)
	if err != nil {
		return returnError(err)
	}

	return dolar
}

func returnError(err error) float64 {
	fmt.Println("Error: ", err)
	return 0
}

func makeRequest() (ResponseDolar, error) {
	var responseDolar ResponseDolar

	response, err := http.Get(URL)
	if err != nil {
		return responseDolar, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return responseDolar, err
	}

	if err := json.Unmarshal(body, &responseDolar); err != nil {
		return responseDolar, err
	}

	return responseDolar, nil
}

func getDolarFormted(responseDolar ResponseDolar) (float64, error) {
	var dolarFloat float64 = 0

	dolarFloat, err := strconv.ParseFloat(responseDolar.Usdbrl.Ask, 64)
	if err != nil {
		return dolarFloat, err
	}

	return math.Round(dolarFloat*100) / 100, nil
}
