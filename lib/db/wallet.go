package db

import (
	"github.com/go-pg/pg/v10"
	"github.com/weijinnx/boosty-test/lib/models"
)

// GetWallet by ID
func GetWallet(db *pg.DB, id string) (*models.Wallet, error) {
	wallet := new(models.Wallet)
	err := db.Model(wallet).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

// GetWallets get all wallets
func GetWallets(db *pg.DB) ([]models.Wallet, error) {
	var wallets []models.Wallet

	err := db.Model(&wallets).Select()
	if err != nil {
		return nil, err
	}

	return wallets, nil
}