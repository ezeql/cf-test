package main

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/shopspring/decimal"
)

type CustomTime time.Time

type Transaction struct {
	UserID,
	OriginatingCountry,
	CurrencyFrom, CurrencyTo string
	Rate, AmountBuy, AmountSell decimal.Decimal
	TimePlaced                  CustomTime
}

func NewTransactionFromJSONBuffer(bytes []byte) (*Transaction, error) {
	t := &Transaction{}

	if err := json.Unmarshal(bytes, t); err != nil {
		return nil, err
	}

	return t, t.Validate()
}

func (t *CustomTime) UnmarshalJSON(b []byte) error {
	//unquote.
	//better: https://github.com/shopspring/decimal/blob/master/decimal.go#L561
	s := string(b)
	s = s[1 : len(s)-1]

	auxTime, err := time.Parse("2-JAN-06 15:04:05", s)
	*t = CustomTime(auxTime)
	return err
}

func (t *Transaction) Validate() error {
	if t.AmountBuy.Cmp(decimal.Zero) <= 0 {
		return errors.New("AmountBuy must be greater than 0")
	}
	if t.AmountSell.Cmp(decimal.Zero) <= 0 {
		return errors.New("AmountSell must be greater than 0")
	}
	if !isValidCurrency(t.CurrencyFrom) {
		return errors.New("CurrencyFrom must be a valid currency")
	}
	if !isValidCurrency(t.CurrencyTo) {
		return errors.New("CurrencyTo must be a valid currency")
	}
	if !isValidCountry(t.OriginatingCountry) {
		return errors.New("OriginatingCountry must be a valid country")
	}
	if !t.AmountSell.Mul(t.Rate).Equals(t.AmountBuy) {
		log.Println(t.Rate, t.AmountSell, t.AmountBuy)
		return errors.New("AmountSell, AmountBuy and Rate do not match")
	}
	return nil
}

func isValidCurrency(currency string) bool {
	//FIXME
	return true
}

func isValidCountry(country string) bool {
	//FIXME
	return true
}
