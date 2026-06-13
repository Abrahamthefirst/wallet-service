package db

import (
	"time"

	"github.com/Abrahamthefirst/finecore-practice/internal/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPgDB(dsn string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	sqlDB, err := db.DB()

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	err = db.AutoMigrate(
		&models.DisbursmentModel{},
		&models.FeeModel{},
		&models.LedgerEntryModel{},
		&models.SystemAccountModel{},
		&models.TransactionModel{},
		&models.UserModel{},
		&models.WalletModel{},
	)

	if err != nil {
		panic("Database connection failed")
	}

	return db

}
