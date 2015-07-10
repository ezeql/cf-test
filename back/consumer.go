package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bitly/go-nsq"
	"github.com/gin-gonic/gin"
)

func main() {

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	config := nsq.NewConfig()
	prod, _ := nsq.NewProducer("127.0.0.1:4150", config)
	defer prod.Stop()

	r := gin.Default()
	r.POST("/api/v1/transaction", func(c *gin.Context) {
		//decode and validate
		content, err := ioutil.ReadAll(c.Request.Body)

		if err != nil {
			c.Writer.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())
			return
		}

		if _, err := NewTransactionFromJSONBuffer(content); err != nil {
			c.Writer.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())
			return
		}

		err = prod.Publish("transactions", content)
		if err != nil {
			log.Println(err.Error())
			c.Writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		c.Writer.WriteHeader(http.StatusOK)

	})

	go r.Run(":8080")
	<-termChan
}
