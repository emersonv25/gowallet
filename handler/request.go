package handler

import (
	"fmt"

	"github.com/emersonv25/gowallet/schemas"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param %s (type: %s) is required", name, typ)
}

type CreateWalletRequest struct {
	Name     string `json:"name"`
	Document string `json:"document"`
	Email    string `json:"email"`
}

func (r *CreateWalletRequest) Validate() error {
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Document == "" {
		return errParamIsRequired("document", "string")
	}
	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}
	return nil
}

type CreateTransactionRequest struct {
	WalletID    int                     `json:"walletId"`
	Amount      float64                 `json:"amount"`
	Description string                  `json:"description"`
	Type        schemas.TransactionType `json:"type"`
}

func (r *CreateTransactionRequest) Validate() error {
	if r.WalletID == 0 {
		return errParamIsRequired("walletId", "int")
	}
	if r.Amount == 0 {
		return errParamIsRequired("amount", "float64")
	}
	if r.Description == "" {
		return errParamIsRequired("description", "string")
	}
	return nil
}
