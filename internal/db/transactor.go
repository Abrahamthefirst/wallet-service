package db

import (
	"context"

	"github.com/Abrahamthefirst/finecore-practice/internal/db/repository"
	"gorm.io/gorm"
)

type GormTransactor struct {
	db *gorm.DB
}

func NewGormTransactor(db *gorm.DB) *GormTransactor {
	return &GormTransactor{db: db}
}

func (t *GormTransactor) WithTx(ctx context.Context, fn repository.TxFunc) error {
	return t.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txCtx := WithTxDB(ctx, tx)
		return fn(txCtx)
	})
}
