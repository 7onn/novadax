package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Order{})

	brl := getBalance("BRL")
	fmt.Println("BRL balance", brl.Available)

	btc := getBalance("BTC")
	fmt.Println("BTC balance", btc.Available)

	t := getTick("BTC")
	fmt.Println("BTC highest", t.Data.High)
	fmt.Println("BTC lowest", t.Data.Low)

	b := buy(buyRequest{
		Symbol: "BTC_BRL",
		Type:   "LIMIT",
		Side:   "BUY",
		Price:  t.Data.Low,
		Amount: "0.001",
	})

	fmt.Println("\nbuyResp", b)

	b.Save()

}
