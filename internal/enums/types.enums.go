package enums

type WalletType string

const (
	UserWallet     WalletType = "user_wallet"
	MerchantWallet WalletType = "merchant_wallet"
)

func (w WalletType) String() string {
	return string(w)
}
