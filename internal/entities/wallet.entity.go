package entities

import (
	"time"

	"github.com/Abrahamthefirst/finecore-practice/internal/enums"
)

type Wallet struct {
	ID           uint             `json:"id"`
	UserId       uint             `json:"user_id"`
	Balance      uint             `json:"wallet_balance"`
	Currency     enums.Currency   `json:"currency"`
	WalletType   enums.WalletType `json:"wallet_type"`
	Transactions *[]Transaction      `json:"transactions"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
