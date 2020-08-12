package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type timestampResponse struct {
	Code    string `json:"code"`
	Data    int64  `json:"data"`
	Message string `json:"message"`
}

func getTimestamp() string {
	endpoint := "/v1/common/timestamp"
	r, err := http.NewRequest("GET", apiURL+endpoint, nil)

	// t := time.Now()
	// ts := t.UnixNano() / int64(time.Millisecond)
	// ms := strconv.FormatInt(ts, 10)

	// sign := getSha256(secretkey, "GET", endpoint, param, ms)
	// r.Header.Add("X-Nova-Access-Key", accesskey)
	// r.Header.Add("X-Nova-Signature", sign)
	// r.Header.Add("X-Nova-Timestamp", ms)
	c := &http.Client{}

	resp, err := c.Do(r)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	bs, _ := ioutil.ReadAll(resp.Body)
	t := &timestampResponse{}
	json.Unmarshal(bs, t)

	fmt.Println(string(bs))

	return strconv.FormatInt(t.Data, 10)
}
