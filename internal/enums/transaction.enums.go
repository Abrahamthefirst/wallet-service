package enums

// On the ledger entry
type EntryType string

const (
	EntryTypeDebit  EntryType = "debit"
	EntryTypeCredit EntryType = "credit"
)

func (e EntryType) String() string{
	return string(e)
}
