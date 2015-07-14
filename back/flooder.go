package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ezeql/go-randomdata"
	"github.com/shopspring/decimal"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	cant := 0
	log.Println("Flooding with transactions... =D")
	initialDate, _ := time.Parse("2-JAN-06 15:04:05", "1-JAN-14 10:00:00")

	for i := 0; i < 5; i++ {
		go func() {
			for {
				id := fmt.Sprintf("%v", rand.Intn(10000))
				from := randomdata.Currency()
				to := randomdata.Currency()
				rate := decimal.NewFromFloat(randomdata.Decimal(1, 2, 2))
				sell := decimal.NewFromFloat(randomdata.Decimal(50, 1000, 2))
				buy := rate.Mul(sell)
				tp := initialDate.Add(time.Hour * 24 * time.Duration(cant)).
					Format("2-JAN-06 15:04:05")

				oc := randomdata.Country(randomdata.TwoCharCountry)

				m := map[string]interface{}{
					"userId":             id,
					"currencyFrom":       from,
					"currencyTo":         to,
					"amountSell":         sell,
					"amountBuy":          buy,
					"rate":               rate,
					"timePlaced":         tp,
					"originatingCountry": oc,
				}
				mJson, _ := json.Marshal(m)
				contentReader := bytes.NewReader(mJson)
				req, _ := http.NewRequest("POST", "http://localhost:8080/api/v1/transaction", contentReader)
				req.Header.Set("Content-Type", "application/json")
				client := &http.Client{}
				client.Do(req)
				cant++

				time.Sleep(time.Millisecond * time.Duration(3000))
			}
		}()
	}

	<-termChan
	log.Println("total messages:", cant)
}
