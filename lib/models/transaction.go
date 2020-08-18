package models

import "time"

// COMMISSION for every transaction
const COMMISSION = 0.015

// Transaction model
type Transaction struct {
	tableName struct{} `pg:"transactions"`

	ID string `pg:"type:uuid,pk,default:gen_random_uuid()" json:"id"`
	Sender string `pg:"type:uuid,nopk" json:"sender"`
	Receiver string `pg:"type:uuid,nopk" json:"receiver"`
	Currency string `json:"currency"`
	Amount float64 `json:"amount"`
	Commission float64 `json:"commission"`
	CreatedAt time.Time `json:"created_at"`
}

// NewTransaction instance
func NewTransaction(from, to, currency string, amount float64) *Transaction {
	return &Transaction{
		Sender: from,
		Receiver: to,
		Currency: currency,
		Amount: amount,
		Commission: amount*COMMISSION,
	}
}