package main

import (
	"encoding/json"
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
	c := &http.Client{}
	resp, err := c.Do(r)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	bs, _ := ioutil.ReadAll(resp.Body)
	t := &timestampResponse{}
	json.Unmarshal(bs, t)
	return strconv.FormatInt(t.Data, 10)
}
