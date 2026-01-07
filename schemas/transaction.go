package schemas

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	WalletId uint
	Amount   float64
	Type     string
	Wallet   Wallet `gorm:"foreignKey:WalletId"`
}

type TransactionResponse struct {
	ID        uint    `json:"id"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
	DeletedAt string  `json:"deletedAt,omitempty"`
	WalletId  uint    `json:"walletId"`
	Amount    float64 `json:"amount"`
	Type      string  `json:"type"`
}

func NewTransaction(walletId uint, amount float64, ttype string) Transaction {
	return Transaction{
		WalletId: walletId,
		Amount:   amount,
		Type:     ttype,
	}
}
func NewTransactionResponse(t Transaction) TransactionResponse {
	return TransactionResponse{
		ID:        t.ID,
		CreatedAt: t.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: t.UpdatedAt.Format("2006-01-02 15:04:05"),
		WalletId:  t.WalletId,
		Amount:    t.Amount,
		Type:      t.Type,
	}
}
