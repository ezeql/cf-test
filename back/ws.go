package main

import (
	"log"
	"math"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/VividCortex/ewma"
	"github.com/bitly/go-nsq"
)

func main() {
	config := nsq.NewConfig()

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	transactionConsumer, _ := nsq.NewConsumer("transactions", "ch", config)
	statsConsumer, _ := nsq.NewConsumer("transactions", "stats", config)

	transactionConsumer.AddConcurrentHandlers(nsq.HandlerFunc(func(message *nsq.Message) error {
		h.broadcast <- message.Body
		message.Finish()
		return nil
	}), 2)

	//PRICE ANALYSYS FOR EUR/USD
	ewmaEURUSD := ewma.NewMovingAverage()
	var delta float64 = 0

	statsConsumer.AddConcurrentHandlers(nsq.HandlerFunc(func(message *nsq.Message) error {

		transaction, err := NewTransactionFromJSONBuffer(message.Body)
		if err != nil {
			log.Println("error decoding transaction from nsq msg")
			return err
		}

		if transaction.CurrencyFrom == "EUR" && transaction.CurrencyTo == "USD" {
			current := ewmaEURUSD.Value()
			newValue, _ := transaction.Rate.Float64()
			ewmaEURUSD.Add(newValue)
			if current != 0 && math.Abs(newValue-current) > delta {
				if newValue > current {
					//					h.broadcast <- []byte("{""message"":""trend"",""data"":{""from"":""EUR"",""to"":""USD"",""trend:""rising""}}"""
					log.Println("price is rising")
				} else {
					// h.broadcast <- []byte("{'message':'trend','data':{'from':'EUR','to':'USD','trend:'falling'}}")
					log.Println("price is falling")
				}
			} else {
				// h.broadcast <- []byte("{'message':'trend','data':{'from':'EUR','to':'USD','trend:'stable'}}")
				log.Println("price is stable")
			}

		}
		message.Finish()
		return nil
	}), 2)

	connect := func(consumers ...*nsq.Consumer) error {
		for _, consumer := range consumers {
			if err := consumer.ConnectToNSQLookupd("127.0.0.1:4161"); err != nil {
				return err
			}
		}
		return nil
	}

	if err := connect(transactionConsumer, statsConsumer); err != nil {
		log.Println("Could not connect")
		log.Println(err)
	} else {
		http.HandleFunc("/ws", serveWs)
		go h.run()
		go http.ListenAndServe(":8090", nil)
		//TODO: add grafecul shutdown lib for http server
		<-termChan

	}
}
