package dtos

import (
	"time"

	"github.com/Abrahamthefirst/finecore-practice/internal/enums"
)

type Peer2PeerTransferServiceSuccess struct {
	EntryType        enums.EntryType `json:"entry_type"`
	ReceiverWalletID uint            `json:"receiver_wallet_id"`
	Description      string          `json:"description"`
	SenderBalance    uint            `json:"sender_balance"`
	ReceiverName     string          `json:"receiver_name"`
	CreatedAt        time.Time       `json:"created_at"`
}
