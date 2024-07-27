package main

import (
	"fmt"
)

// Define dados uteis para o erro.

// Por possuir um metodo `Error() string`, essa struct implementa a interface `error`,
// assim ela se torna do type `error`
type InsufficientFundsError struct {
	msg     string
	balance float64
	amount  float64
}

// Faz com que a struct implemente a interface `error`
func (e *InsufficientFundsError) Error() string {
	return fmt.Sprintf("%s: attempted to withdraw %.2f with balance %.2f", e.msg, e.amount, e.balance)
}

func withdraw(b, a float64) (float64, error) {
	if a > b {
		return 0, &InsufficientFundsError{"Insufficient funds", b, a} // &InsufficientFundsError eh do tipo `error`
	} else {
		return b - a, nil
	}
}

func main() {
	amount := 200.0
	balance := 100.5
	if _, err := withdraw(balance, amount); err != nil {
		fmt.Printf("Error while withdrawing: %v\n", err)
	} else {
		fmt.Println("Withdrawal of successful")
	}
}
