package walletservice

import (
	"context"
	"fmt"

	"github.com/Abrahamthefirst/finecore-practice/internal/db/repository"
	"github.com/Abrahamthefirst/finecore-practice/internal/dtos"
	"github.com/Abrahamthefirst/finecore-practice/internal/entities"
	"github.com/Abrahamthefirst/finecore-practice/internal/enums"
)

type WalletService struct {
	tx                    repository.Transactor
	ledgerRepository      repository.LedgerRepository
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
	// I would come back and handle the fee
	// Another question to ask is that would it be okay to throw http errors at the service layer
	if senderWallet.Balance < input.Amount {
		return nil, fmt.Errorf("Insufficient funds")
	}

	s.tx.WithTx(ctx, func(ctx context.Context) error {

		// Create a transaction

		inputTransaction := entities.Transaction{
			Amount:         input.Amount,
			OperationType:  enums.OperationTypeTransfer,
			Currency:       input.Currency,
			Description:    input.Description,
			IdempotencyKey: input.IdempotencyKey,
		}

		transaction, err := s.transactionRepository.Create(ctx, inputTransaction)

		if err != nil {
			return err
		}

		inputCreditLedger := entities.Ledger{
			TransactionID: transaction.ID,
			AccountID:     input.ReceiverWalletId,
			EntryType:     enums.EntryTypeCredit,
			Amount:        input.Amount,
			Currency:      input.Currency,
			Description:   input.Description,
		}

		_, err = s.ledgerRepository.Create(ctx, inputCreditLedger)

		if err != nil {
			return err
		}

		inputDebitLedger := entities.Ledger{
			TransactionID: transaction.ID,
			AccountID:     input.SenderWalletId,
			EntryType:     enums.EntryTypeDebit,
			Amount:        input.Amount,
			Currency:      input.Currency,
			Description:   input.Description,
		}

		_, err = s.ledgerRepository.Create(ctx, inputDebitLedger)

		if err != nil {
			return err
		}

		return nil
	})


	
}
