package db

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPgDB(dsn string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	sqlDB, err := db.DB()

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	err = db.AutoMigrate()

	if err != nil {
		panic("Database connection failed")
	}

	return db

}
