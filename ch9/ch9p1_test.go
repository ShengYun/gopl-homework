package ch9

import (
	"fmt"
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		Deposit(100)
		if Withdraw(400) {
			fmt.Println("alice can withdraw 400")
		} else {
			fmt.Println("alice cannot withdraw 400")
		}
		done <- struct{}{}
	}()

	// Bob
	go func() {
		Deposit(100)
		if Withdraw(100) {
			fmt.Println("bob can withdraw 100")
		} else {
			fmt.Println("bob cannot withdraw 100")
		}
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := Balance(), 100; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
