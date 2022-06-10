package wallet

import "testing"

func TestWallet(t *testing.T) {
	confirmBalance := func(t *testing.T, wallet Wallet, expect Bitcoin) {
		t.Helper()
		result := wallet.Balance()

		if result != expect {
			t.Errorf("result %v, expect %v", result, expect)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		confirmBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))

		confirmBalance(t, wallet, Bitcoin(10))
	})
}
