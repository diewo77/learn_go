package main

import "testing"

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startigBalance := Bitcoin(20)
		wallet := Wallet{startigBalance}
		got := wallet.Withdraw(Bitcoin(100))

		assertError(t, got, ErrInsufficientFunds)
		assertBalance(t, wallet, startigBalance)
	})
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Errorf("Wanted an error but didn't get one !")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("got an error but didn't want one !")
	}
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
