package point

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	assertNoError := func(t *testing.T, got error) {
		if got != nil {
			t.Fatal("got an error but didn't want one")
		}
	}

	assertError := func(t *testing.T, got error, want error) {
		if got == nil {
			t.Error("wanted an error but didn't get one")
		}
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		fmt.Println("address of balance in test is", &wallet.balance)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(40)}
		err := wallet.Withdraw(Bitcoin(20))
		assertBalance(t, wallet, Bitcoin(20))
		assertNoError(t, err)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(40)}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, Bitcoin(40))
		assertError(t, err, InsufficientFundsError)
	})

}
