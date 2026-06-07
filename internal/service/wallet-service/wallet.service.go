package walletservice

import (
	"context"

	"github.com/Abrahamthefirst/finecore-practice/internal/db/repository"
	"github.com/Abrahamthefirst/finecore-practice/internal/dtos"
)

type WalletService struct {
	tx               repository.Transactor
	walletRepository *repository.WalletRepository
}

func NewWalletService(tx repository.Transactor, walletRepository *repository.WalletRepository) *WalletService {
	return &WalletService{tx: tx, walletRepository: walletRepository}
}

func (s *WalletService) TransferBetweenUsers(ctx context.Context, senderId uint, input dtos.TransferBetweenUsersRequestBody) {

	if input.Amount <= 0 {
		// return errors.New("Amount must be positive")
	}

	if senderId == input.ReceiverId {
		// return errors.New("Sender and receiver must be of different accounts")
	}

	// s.tx.WithTx(ctx, func(ctx context.Context) error {
	// 	// run every transaction here
	// })

}
