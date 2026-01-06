package schemas

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	Id       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name     string    `json:"name"`
	Document string    `json:"document"`
	Email    string    `json:"email"`
	Balance  float64   `json:"balance"`
}
