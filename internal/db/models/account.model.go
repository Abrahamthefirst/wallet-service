package models

import "gorm.io/gorm"

type SystemAccountModel struct {
	gorm.Model `gorm:"uniqueIndex"`

}

func (*SystemAccountModel) TableName() string {
	return "system_accounts"
}
