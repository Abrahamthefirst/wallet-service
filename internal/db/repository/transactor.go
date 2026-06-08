package repository

import (
	"context"

	"gorm.io/gorm"
)

type GormTransactor struct {
	db *gorm.DB
}

func NewGormTransactor(db *gorm.DB) *GormTransactor {
	return &GormTransactor{db: db}
}

func (t *GormTransactor) WithTx(ctx context.Context, fn TxFunc) error {
	return t.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txCtx := WithTxDB(ctx, tx)
		return fn(txCtx)
	})
}
