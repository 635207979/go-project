package money

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t *testing.T, got error, want string) {
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}

		if got.Error() != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{money: Bitcoin(10)}
		wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(0)
		assertBalance(t, wallet, want)
	})
	t.Run("withdraw more than balance", func(t *testing.T) {
		wallet := Wallet{money: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(20))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
		assertError(t, err, "cannot withdraw, insufficient funds")
	})

}
