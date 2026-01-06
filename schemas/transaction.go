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
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	WalletId  uint    `json:"wallet_id"`
	Amount    float64 `json:"amount"`
	Type      string  `json:"type"`
}
