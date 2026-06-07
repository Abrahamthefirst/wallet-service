package repository

import "gorm.io/gorm"

type DisbursmentRepository struct {
	db *gorm.DB
}

func NewDisbursmentRepository(db *gorm.DB) *DisbursmentRepository {
	return &DisbursmentRepository{
		db,
	}
}


func (r *DisbursmentRepository) Create(){}
func (r *DisbursmentRepository) GetByID(){}
func (r *DisbursmentRepository) GetByMerchantID(){}