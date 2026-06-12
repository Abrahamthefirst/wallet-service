package enums

type OperationType string

const (
	OperationTypeFund         OperationType = "fund"
	OperationTypeTransfer     OperationType = "transfer"
	OperationTypeWithdraw     OperationType = "withdraw"
	OperationTypeDisbursement OperationType = "disbursement"
	OperationTypeReversal     OperationType = "reversal"
	OperationTypeRefund       OperationType = "refund"
	OperationTypeFeeCharge    OperationType = "fee_charge"
	OperationTypeFeeReversal  OperationType = "fee_reversal"
	OperationTypeCreditAdjust OperationType = "credit_adjustment"
	OperationTypeDebitAdjust  OperationType = "debit_adjustment"
)

func (o OperationType) String() string {
	return string(o)
}
