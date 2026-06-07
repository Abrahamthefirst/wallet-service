package repository

import "gorm.io/gorm"

type SystemAccountRepository struct {
	db *gorm.DB
}

func NewSystemAccountRepository(db *gorm.DB) *SystemAccountRepository {
	return &SystemAccountRepository{
		db,
	}
}

// this is when a person pays to a merchant, which is basically the platform, then we use this to pay out
func (r SystemAccountRepository) Create() {}
