package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

    assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
        t.Helper()
        got := wallet.Balance()

        if got != want {
            t.Errorf("got %s want %s", got, want)
        }
    }

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

        assertBalance(t, wallet, startingBalance)

        if err == nil {
            t.Error("wanted error but didn't get one")
        }
    })
}