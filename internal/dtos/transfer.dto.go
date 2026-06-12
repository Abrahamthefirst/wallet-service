

import (
	"time"

	"github.com/Abrahamthefirst/finecore-practice/internal/enums"
)

type Peer2PeerTransferServiceSuccess struct {
	EntryType        enums.EntryType
	ReceiverWalletID uint
	Description      string
	SenderBalance    uint
	SenderName       string
	CreatedAt        time.Time
}
