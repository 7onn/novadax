package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type buyRequest struct {
	Symbol string  `json:"symbol"`
	Type   string  `json:"type"`
	Side   string  `json:"side"`
	Amount string  `json:"amount"`
	Price  float64 `json:"price"`
}
type cancelRequest struct {
	ID string `json:"id"`
}

type cancelResponse struct {
	Code    string                `json:"code"`
	Data    struct{ Result bool } `json:"data"`
	Message string                `json:"message"`
}

//Order is for meu ovo
type Order struct {
	Amount    string `json:"amount"`
	ID        string `json:"id"`
	Price     string `json:"price"`
	Side      string `json:"side"`
	Status    string `json:"status"`
	Symbol    string `json:"symbol"`
	Timestamp int64  `json:"timestamp"`
	Type      string `json:"type"`
	Value     string `json:"value"`
}

type orderResponse struct {
	Code    string `json:"code"`
	Data    Order  `json:"data"`
	Message string `json:"message"`
}
type ordersResponse struct {
	Code    string  `json:"code"`
	Data    []Order `json:"data"`
	Message string  `json:"message"`
}

func buy(b buyRequest) *Order {

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

	or := &orderResponse{}
	json.Unmarshal(bs, or)

	o := &Order{
		ID:        or.Data.ID,
		Symbol:    or.Data.Symbol,
		Amount:    or.Data.Amount,
		Type:      or.Data.Type,
		Value:     or.Data.Value,
		Timestamp: or.Data.Timestamp,
		Price:     or.Data.Price,
		Side:      or.Data.Side,
	}

	log.Println(*o, "|| orders/create")

	return o
}

func cancel(cr cancelRequest) error {

	j, _ := json.Marshal(cr)
	hash := getMd5(j)
	body := bytes.NewBuffer(j)
	endpoint := "/v1/orders/cancel"
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

	cres := &cancelResponse{}
	json.Unmarshal(bs, cres)

	log.Printf("%+v %s %s", *cres, cr.ID, "|| orders/cancel/")

	if !cres.Data.Result {
		return fmt.Errorf(fmt.Sprintf("order not cancelled, %x", cr))
	}

	return nil
}

func getOrders() []Order {
	endpoint := "/v1/orders/list"
	r, err := http.NewRequest("GET", apiURL+endpoint, nil)
	ms := getTimestamp()
	sign := getSha256(secretkey, "GET", endpoint, "", ms)
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

	or := &ordersResponse{}
	json.Unmarshal(bs, or)

	return or.Data
}
