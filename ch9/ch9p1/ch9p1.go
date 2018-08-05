// Exercise 9.1: Add a function Withdraw(amount int) bool to the gopl.io/ch9/bank1 program.
// The result should indicate whether the transaction succeeded or failed due to insufficient funds.
// The message sent to the monitor goroutine must contain both the amount to withdraw and a new
// channel over which the monitor goroutine can send the boolean result back to Withdraw.

package ch9p1

type withdraw struct {
	amount int
	result chan bool
}

var deposits = make(chan int)
var balances = make(chan int)
var withdraws = make(chan withdraw)

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func Withdraw(amount int) bool {
	result := make(chan bool)
	withdraws <- withdraw{amount, result}
	return <-result
}

// this is the monitor goroutine for variable balance
func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case w := <-withdraws:
			if w.amount > balance {
				w.result <- false
			} else {
				balance -= w.amount
				w.result <- true
			}
		}
	}
}

func init() {
	go teller()
}
