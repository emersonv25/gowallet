package config

import (
	"os"

	"github.com/emersonv25/gowallet/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

func InitializeSqlLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/sqlite.db"
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.New(sqlite.Config{
		DriverName: "sqlite", // ← força modernc.org/sqlite
		DSN:        dbPath,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Wallet{}, &schemas.Transaction{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
