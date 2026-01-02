package main

import (
	"fmt"
)

func main() {

	years := 5
	interestRate := float64(0.425)
	totalAmount := float64(1000.25)
	months := years * 12
	payment := float64(85)

	principal := totalAmount
	for i := 1; i <= months; i++ {

		//remaining balance * (1 + interestRate / 12) - payment
		principal = principal*(1+interestRate/12) - payment
		fmt.Printf("Month %d: Remaining Balance: %.2f\n", i, principal)
	}

}
