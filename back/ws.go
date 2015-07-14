package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bitly/go-nsq"
	"github.com/garyburd/redigo/redis"
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

	redisPool := redis.NewPool(func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", "127.0.0.1:6379")

		if err != nil {
			return nil, err
		}

		return c, err
	}, 10)

	defer redisPool.Close()

	statsConsumer.AddConcurrentHandlers(nsq.HandlerFunc(func(message *nsq.Message) error {

		transaction, err := NewTransactionFromJSONBuffer(message.Body)
		if err != nil {
			log.Println("error decoding transaction from nsq msg")
			return err
		}
		conn := redisPool.Get()
		conn.Send("ZINCRBY", "transactions:bycountry", 1, transaction.OriginatingCountry)
		conn.Send("ZINCRBY", "transactions:bycurrtocurr", 1, transaction.CurrencyFrom+":"+transaction.CurrencyTo)
		conn.Flush()

		message.Finish()
		conn.Close()
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
		go h.run()
		http.HandleFunc("/ws", serveWs)
		go http.ListenAndServe(":8090", nil)
		<-termChan
	}
}
