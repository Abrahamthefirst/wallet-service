package enums

type WalletType string

const (
	WalletTypePersonal   WalletType = "personal"
	WalletTypeBusiness   WalletType = "business"
	WalletTypeSavings    WalletType = "savings"
	WalletTypeFloat      WalletType = "float"
	WalletTypeSettlement WalletType = "settlement"
	WalletTypeFee        WalletType = "fee"
	WalletTypeSuspense   WalletType = "suspense"
)

func (w WalletType) String() string {
	return string(w)
}
