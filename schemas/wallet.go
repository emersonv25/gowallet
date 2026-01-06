package schemas

import (
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	Name     string
	Document string
	Email    string
	Balance  float64
}

type WalletResponse struct {
	ID        uint    `json:"id"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	Name      string  `json:"name"`
	Document  string  `json:"document"`
	Email     string  `json:"email"`
	Balance   float64 `json:"balance"`
}
