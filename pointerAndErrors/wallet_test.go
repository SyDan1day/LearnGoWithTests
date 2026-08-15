package pointeranderrors

import (
	"errors"
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, w wallet, want Bitcoin) {
		got := w.Balance()
		if got != want {
			t.Errorf("want %s but got %s", want, got)
		}
	}

	assertError := func(t *testing.T, got error, want error) {
		if got == nil {
			t.Fatal("want a error but didn't get one")
		}

		if !errors.Is(got, want) {
			t.Errorf("want error %s but got error %s", want, got)
		}
	}

	assertNoError := func(t *testing.T, got error) {
		if got != nil {
			t.Error("didn't want a error but get one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := wallet{}
		wallet.Deposit(10)
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient Bitcoin", func(t *testing.T) {
		wallet := wallet{20}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, Bitcoin(20))
		assertError(t, err, InsufficientFundsError)
	})
}
