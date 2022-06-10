package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var apiUrl = "https://openapi.twse.com.tw/v1/exchangeReport/BWIBBU_ALL"

func GetStocks() []Stock {
	bs, err := getAllStockInfo()
	if err != nil {
		log.Fatal(err)
	}
	var stocks []Stock
	err = json.Unmarshal(bs, &stocks)
	if err != nil {
		log.Fatal(err)
	}

	return stocks
}

func getAllStockInfo() ([]byte, error) {
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
