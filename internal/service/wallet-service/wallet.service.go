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

func (s *WalletService) TransferBetweenUsers(ctx context.Context, senderId uint, input dtos.TransferBetweenUsersRequestBody) (*dtos.Peer2PeerTransferServiceSuccess, error) {

	if input.Amount <= 0 {
		return nil, fmt.Errorf("Amount must be positive")
	}

	if input.SenderWalletID == input.ReceiverWalletID {
		return nil, fmt.Errorf("transfers must be between different wallets")
	}

	var senderBalance uint
	err := s.tx.WithTx(ctx, func(ctx context.Context) error {

		senderWallet, err := s.walletRepository.GetByIDForUpdate(ctx, input.SenderWalletID)

		if err != nil {
			// s.logger.Error("TransferBetweenUsers failed",
			// 	"senderId", input.SenderWalletId,
			// 	"receiverId", input.ReceiverWalletId,
			// 	"amount", input.Amount,
			// 	"error", err,
			// )
			return fmt.Errorf("Sender Wallet not found")
		}

		receiverWallet, err := s.walletRepository.GetByIDForUpdate(ctx, input.ReceiverWalletID)

		if err != nil {
			return fmt.Errorf("Receiver Wallet not found")
		}

		if senderWallet.Currency != receiverWallet.Currency {
			return fmt.Errorf("currency_mismatch: wallet currency is %s, request currency is %s", senderWallet.Currency, receiverWallet.Currency)
		}

		if senderWallet.Balance < input.Amount {
			return fmt.Errorf("Insufficient funds")
		}

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
			AccountID:     input.ReceiverWalletID,
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
			AccountID:     input.SenderWalletID,
			EntryType:     enums.EntryTypeDebit,
			Amount:        input.Amount,
			Currency:      input.Currency,
			Description:   input.Description,
		}

		_, err = s.ledgerRepository.Create(ctx, inputDebitLedger)

		if err != nil {
			return err
		}

		senderBalance = senderWallet.Balance - input.Amount
		_, err = s.walletRepository.UpdateBalance(ctx, input.SenderWalletID, senderBalance)

		if err != nil {
			return err
		}

		_, err = s.walletRepository.UpdateBalance(ctx, input.ReceiverWalletID, receiverWallet.Balance+input.Amount)

		if err != nil {
			return err
		}

		
		return nil
	})

	if err != nil {
		return nil, err
	}

	// I could probably build a notification service
	return &dtos.Peer2PeerTransferServiceSuccess{ReceiverWalletID: input.ReceiverWalletID, Description: input.Description, SenderBalance: senderBalance, EntryType: enums.EntryTypeDebit}, nil

}
