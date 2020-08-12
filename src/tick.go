package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type quote struct {
	Time time.Time `json:"time"`
	High string    `json:"high24h"`
	Low  string    `json:"low24h"`
}

type tick struct {
	Code    string `json:"code"`
	Data    quote  `json:"data"`
	Message string `json:"message"`
}

func getTick(currency string) *tick {
	endpoint := "/v1/market/ticker"
	param := fmt.Sprintf("/?symbol=%s_BRL", currency)
	r, err := http.NewRequest("GET", apiURL+endpoint+param, nil)

	t := time.Now()
	ts := t.UnixNano() / int64(time.Millisecond)
	ms := strconv.FormatInt(ts, 10)

	sign := getSha256(secretkey, "GET", endpoint, param, ms)
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
	b := &tick{}
	json.Unmarshal(bs, b)
	b.Data.Time = t
	return b
}
