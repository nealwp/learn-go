package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

    t.Run("deposit", func(t *testing.T){
        wallet := Wallet{}
        wallet.Deposit(Bitcoin(10))
        want := Bitcoin(10)

        assertBalance(t, wallet, want)
    })

    t.Run("withdraw", func(t *testing.T){
        wallet := Wallet{balance: Bitcoin(20)}
        wallet.Withdraw(10)
        want := Bitcoin(10)

        assertBalance(t, wallet, want)
    })

    t.Run("withdaw insufficient funds", func(t *testing.T){
        startingBalance := Bitcoin(20)
        wallet := Wallet{startingBalance}
        err := wallet.Withdraw(100)

        assertError(t, err, ErrInsufficientFunds)
        assertBalance(t, wallet, startingBalance)
    })
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
    t.Helper()
    got := wallet.Balance()

    if got != want {
        t.Errorf("got %s want %s", got, want)
    }
}

func assertError(t testing.TB, got, want error) {
    t.Helper()

    if got == nil {
        t.Fatal("wanted error but didn't get one")
    }

    if got != want {
        t.Errorf("got %q wanted %q", got, want)
    }
}
