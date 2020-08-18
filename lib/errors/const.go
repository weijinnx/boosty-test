package errors

import "net/http"

const (
	ECBindError = 1000
	ECInvalidCurrency = 1001
	ECSameWallet = 1002
	ECShouldBeSameCurrency = 1003
	ECInsufficientBalance = 1004
	ECValidationError = 1005
	ECFailBeginTx = 1006
	ECFailCommitTx = 1007
)

var (
	CodeMap = map[int]int{
		ECInvalidCurrency: http.StatusBadRequest,
		ECSameWallet: http.StatusBadRequest,
		ECShouldBeSameCurrency: http.StatusBadRequest,
		ECInsufficientBalance: http.StatusBadRequest,
		ECValidationError: http.StatusBadRequest,
	}

	CodeText = map[int]string{
		ECInvalidCurrency: "invalid/unsupported currency",
		ECSameWallet: "wallet to receive shouldn't be the same as sender",
		ECShouldBeSameCurrency: "currency of receiver wallet should be the same as sender",
		ECInsufficientBalance: "insufficient funds on the wallet",
		ECValidationError: "input validation errors",
	}
)