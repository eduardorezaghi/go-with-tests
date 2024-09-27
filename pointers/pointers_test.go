package pointers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		fmt.Printf("Address of wallet balance in test is %p\n", &wallet.balance)
		wallet.Deposit(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)

		assert.Equal(t, want, got)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		assert.NoError(t, err)
		assert.Equal(t, want, got)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assert.Error(t, err, ErrInsufficientFunds)
		assert.Equal(t, startingBalance, wallet.Balance(), "Balance should not have changed")
	})

	t.Run("Test Bitcoin stringer", func(t *testing.T) {
		bitcoin := Bitcoin(10)
		got := bitcoin.String()
		want := "10 BTC"

		assert.Equal(t, want, got)
	})
}
