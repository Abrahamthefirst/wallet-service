package repository

import "context"

type TxFunc func(ctx context.Context) error

type Transactor interface {
	WithTx(ctx context.Context, fn TxFunc) error
}

