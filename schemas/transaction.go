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
