package schemas

import (
	"gorm.io/gorm"
)

type TransactionType string

const (
	TransactionTypeCredit TransactionType = "credit"
	TransactionTypeDebit  TransactionType = "debit"
)

type Transaction struct {
	gorm.Model
	WalletId    uint
	Amount      float64
	Type        TransactionType
	Description string
	Wallet      Wallet `gorm:"foreignKey:WalletId"`
}

type TransactionResponse struct {
	ID          uint            `json:"id"`
	CreatedAt   string          `json:"createdAt"`
	UpdatedAt   string          `json:"updatedAt"`
	DeletedAt   string          `json:"deletedAt,omitempty"`
	WalletId    uint            `json:"walletId"`
	Amount      float64         `json:"amount"`
	Type        TransactionType `json:"type"`
	Description string          `json:"description"`
}

func NewTransaction(walletId uint, amount float64, ttype TransactionType, description string) Transaction {
	return Transaction{
		WalletId:    walletId,
		Amount:      amount,
		Type:        ttype,
		Description: description,
	}
}
func NewTransactionResponse(t Transaction) TransactionResponse {
	return TransactionResponse{
		ID:          t.ID,
		CreatedAt:   t.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   t.UpdatedAt.Format("2006-01-02 15:04:05"),
		WalletId:    t.WalletId,
		Amount:      t.Amount,
		Description: t.Description,
		Type:        t.Type,
	}
}
