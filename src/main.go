package main

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

var (
	apiURL    = "https://api.novadax.com"
	accesskey = ""
	secretkey = ""
)

func main() {
	godotenv.Load(".env")
	accesskey = os.Getenv("ACCESS_KEY")
	secretkey = os.Getenv("SECRET_KEY")

	ors := getOrders()
	ms := time.Now().Add(-time.Hour*4).UnixNano() / 1000 / 1000
	for _, o := range ors {
		if o.Timestamp < ms && o.Status == "PROCESSING" {
			cancel(cancelRequest{
				ID: o.ID,
			})
		}
	}

	p := os.Args[1]
	tick := getTick("BTC")
	buy(buyRequest{
		Symbol: "BTC_BRL",
		Type:   "LIMIT",
		Side:   "BUY",
		Price:  tick.Data.avg(), //btc average price of the day
		Amount: p,
	})

	getBalance("BRL")
	getBalance("BTC")

	func() {
		time.Sleep(2 * time.Hour)
		main()
	}()

}
