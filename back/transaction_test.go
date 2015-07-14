package main

import (
	"encoding/json"
	_ "fmt"
	"testing"
)

func TestJsonMarshaling(t *testing.T) {

	//http: //play.golang.org/p/IqWLPtqhEf

	transaction := Transaction{}

	var jsonTransaction = []byte(`{"userId": "134256", "currencyFrom": "EUR", 
						"currencyTo": "GBP", "amountSell": 1000, 
						"amountBuy": 747.10, "rate": 0.7471, 
						"timePlaced" : "24-JAN-15 10:27:44", 
						"originatingCountry" : "FR"}`)

	err := json.Unmarshal(jsonTransaction, &transaction)
	if err != nil {
		t.Error(err)
	}
	err = transaction.Validate()
	if err != nil {
		t.Error(err)
	}
}
