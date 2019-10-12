package main

import (
	"learning-golang/util"
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet util.Wallet, want util.Bitcoin) {
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t *testing.T, got error, want error) {
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}

	t.Run("Deposit:", func(t *testing.T) {
		wallet := util.Wallet{}
		wallet.Deposit(util.Bitcoin(10))
		want := util.Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := util.Wallet{}
		wallet.Deposit(util.Bitcoin(20))
		wallet.Withdraw(util.Bitcoin(10))
		assertBalance(t, wallet, util.Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := util.Wallet{}
		wallet.Deposit(util.Bitcoin(20))
		err := wallet.Withdraw(util.Bitcoin(100))
		assertBalance(t, wallet, util.Bitcoin(20))
		assertError(t, err, util.InsufficientFundsError)
	})
}
