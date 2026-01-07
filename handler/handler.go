package handler

import (
	"github.com/emersonv25/gowallet/config"
	"gorm.io/gorm"
)

var logger *config.Logger
var db *gorm.DB

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
