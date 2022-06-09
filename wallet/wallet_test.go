package wallet

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	result := wallet.Balance()
	expect := Bitcoin(12)

	if result != expect {
		t.Errorf("result %s, expect %s", result, expect)
	}
}
