package walletservice

import (
	"context"
	"fmt"

	"github.com/Abrahamthefirst/finecore-practice/internal/db/repository"
	"github.com/Abrahamthefirst/finecore-practice/internal/dtos"
	"github.com/Abrahamthefirst/finecore-practice/internal/entities"
)

type WalletService struct {
	tx                    repository.Transactor
	transactionRepository *repository.TransactionRepository
	walletRepository      *repository.WalletRepository
}

func NewWalletService(tx repository.Transactor, walletRepository *repository.WalletRepository) *WalletService {
	return &WalletService{tx: tx, walletRepository: walletRepository}
}

func (s *WalletService) TransferBetweenUsers(ctx context.Context, senderId uint, input dtos.TransferBetweenUsersRequestBody) (*entities.Transaction, error) {

	if input.Amount <= 0 {
		return nil, fmt.Errorf("Amount must be positive")
	}

	if input.SenderWalletId == input.ReceiverWalletId {
		return nil, fmt.Errorf("Transfers could only be done between different wallets")
	}

	senderWallet, err := s.walletRepository.GetByID(ctx, input.SenderWalletId)

	if err != nil {
		return nil, fmt.Errorf("Sender Wallet not found")
	}

	receiverWallet, err := s.walletRepository.GetByID(ctx, input.ReceiverWalletId)

	if err != nil {
		return nil, fmt.Errorf("Receiver Wallet not found")
	}

	if senderWallet.Currency != receiverWallet.Currency {
		return nil, fmt.Errorf("currency_mismatch: wallet currency is %s, request currency is %s", senderWallet.Currency, receiverWallet.Currency)
	}

	s.tx.WithTx(ctx, func(ctx context.Context) error {
		// run every transaction here

		

		s.transactionRepository.Create(ctx, input)
		return
	})

}
