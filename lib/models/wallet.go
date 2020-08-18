package models

import (
	"github.com/go-pg/pg/v10"
	"github.com/weijinnx/boosty-test/lib/errors"
)

// Wallet model
type Wallet struct {
	tableName struct{} `pg:"wallets"`

	ID string `pg:"type:uuid,pk,default:gen_random_uuid()" json:"id"`
	Cur string `json:"cur"`
	Balance float64 `pg:",use_zero" json:"balance"`
}

// Transfer funds to another wallet
func (w *Wallet) Transfer(tx *pg.Tx, r *Wallet, amount float64) (*Transaction, error) {
	// sum that should be minused
	sum := amount + (amount*COMMISSION)

	// check that currencies are the same
	if w.Cur != r.Cur {
		return nil, errors.Prep(errors.ECShouldBeSameCurrency, []string{w.Cur, r.Cur})
	}
	// check that wallet has sufficient amount of funds
	if w.Balance < sum {
		return nil, errors.Prep(errors.ECInsufficientBalance, []float64{w.Balance, sum})
	}
	
	// update sender balance
	_, err := tx.Model(w).Set("balance = ?", w.Balance - sum).Where("id = ?id").Update()
	if err != nil {
		return nil, err
	}
	// update receiver balance
	_, err = tx.Model(r).Set("balance = ?", w.Balance + amount).Where("id = ?id").Update()
	if err != nil {
		return nil, err
	}

	transaction := NewTransaction(w.ID, r.ID, w.Cur, amount)
	_, err = tx.Model(transaction).Returning("*").Insert()
	if err != nil {
		return nil, err
	}

	return transaction, tx.Commit()
}