package schemas

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	Name     string `json:"name"`
	Document string `json:"document"`
	Email    string `json:"email"`
}
