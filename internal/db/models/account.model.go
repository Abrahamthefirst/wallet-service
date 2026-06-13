package models

import "gorm.io/gorm"

type SystemAccountModel struct {
	gorm.Model 

}

func (*SystemAccountModel) TableName() string {
	return "system_accounts"
}
