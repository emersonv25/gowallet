package schemas

import (
	"time"

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
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Document  string    `json:"document"`
	Email     string    `json:"email"`
	Balance   float64   `json:"balance"`
}

func NewWallet(name string, document string, email string) Wallet {
	return Wallet{
		Name:     name,
		Document: document,
		Email:    email,
		Balance:  0.0,
	}
}

func NewWalletResponse(wallet Wallet) WalletResponse {
	return WalletResponse{
		ID:        wallet.ID,
		CreatedAt: wallet.CreatedAt,
		UpdatedAt: wallet.UpdatedAt,
		Name:      wallet.Name,
		Document:  wallet.Document,
		Email:     wallet.Email,
		Balance:   wallet.Balance,
	}
}
