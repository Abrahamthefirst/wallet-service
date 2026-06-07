package models

import "gorm.io/gorm"

type DisbursmentModel struct {
	gorm.Model `gorm:"uniqueIndex"`
}


func (*DisbursmentModel) TableName() string {
	return "disbursments"
}