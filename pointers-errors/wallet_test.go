package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10.00))
		assertBalance(t, wallet, Bitcoin(10.00))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{10.0}
		err := wallet.Withdraw(Bitcoin(10.0))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(0.0))
	})

	t.Run("insufficient funds for withdrawal", func(t *testing.T) {
		wallet := Wallet{0.0}
		err := wallet.Withdraw(Bitcoin(10.0))
		assertError(t, err, ErrorInsufficientFunds)
	})
}

func assertError(t testing.TB, err, want error) {
	t.Helper()
	if err == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if want != err {
		t.Errorf("got %q, want %q", err, want)
	}
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("got an error %s but didn't want one", got.Error())
	}
}
