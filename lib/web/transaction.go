package web

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/labstack/echo/v4"
	"github.com/weijinnx/boosty-test/lib/db"
	"github.com/weijinnx/boosty-test/lib/errors"
	"github.com/weijinnx/boosty-test/lib/util"
)

// TransferFundsInput contains request data
// to make transaction from wallet to wallet
type TransferFundsInput struct {
	From string `json:"from"`
	To string `json:"to"`
	Amount float64 `json:"amount"`
}

// Validate input data
func (i TransferFundsInput) Validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.From, validation.Required, is.UUIDv4),
		validation.Field(&i.To, validation.Required, is.UUIDv4),
		validation.Field(&i.Amount, validation.Required),
	)
}

// TransferFunds route handler
func TransferFunds(ctx echo.Context) error {
	cctx := ctx.Get("cctx").(*util.AppContext)

	// bind request data to struct
	i := new(TransferFundsInput)
	err := ctx.Bind(&i)
	if err != nil {
		return errors.Prep(errors.ECBindError, nil)
	}

	// validate data input
	if err := i.Validate(); err != nil {
		return errors.Prep(errors.ECValidationError, err)
	}

	// check that wallets are not the same
	if i.From == i.To {
		return errors.Prep(errors.ECSameWallet, i.From)
	}

	// get sender wallet
	sender, err := db.GetWallet(cctx.DB, i.From)
	if err != nil {
		return errors.Prep(http.StatusInternalServerError, []string{i.From})
	}
	// get receiver wallet
	receiver, err := db.GetWallet(cctx.DB, i.To)
	if err != nil {
		return errors.Prep(http.StatusInternalServerError, []string{i.To})
	}

	// begin transaction
	tx, err := cctx.DB.Begin()
	if err != nil {
		return errors.Prep(errors.ECFailBeginTx, err)
	}
	defer tx.Rollback()

	// make a transfer
	transaction, err := sender.Transfer(tx, receiver, i.Amount)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"transaction": transaction,
	})
}