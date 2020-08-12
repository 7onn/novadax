package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jinzhu/gorm"
)

type buyRequest struct {
	Symbol string `json:"symbol"`
	Type   string `json:"type"`
	Side   string `json:"side"`
	Amount string `json:"amount"`
	Price  string `json:"price"`
}

//Order is for meu ovo
type Order struct {
	gorm.Model
	Amount    string `json:"amount"`
	ID        string `json:"id"`
	Price     string `json:"price"`
	Side      string `json:"side"`
	Status    string `json:"status"`
	Symbol    string `json:"symbol"`
	Timestamp string `json:"timestamp"`
	Type      string `json:"type"`
	Value     string `json:"value"`
}

type orderResponse struct {
	Code    string `json:"code"`
	Data    Order  `json:"data"`
	Message string `json:"message"`
}

func buy(b buyRequest) Order {
	j, _ := json.Marshal(b)
	hash := getMd5(j)
	body := bytes.NewBuffer(j)
	endpoint := "/v1/orders/create"
	r, err := http.NewRequest("POST", apiURL+endpoint, body)
	ms := getTimestamp()
	sign := getSha256(secretkey, "POST", endpoint, hash, ms)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Nova-Access-Key", accesskey)
	r.Header.Add("X-Nova-Signature", sign)
	r.Header.Add("X-Nova-Timestamp", ms)
	c := &http.Client{}
	resp, err := c.Do(r)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	bs, _ := ioutil.ReadAll(resp.Body)

	o := &orderResponse{}
	json.Unmarshal(bs, o)

	return o.Data
}
