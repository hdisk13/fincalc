package main

import (
	"fmt"
)

/*

 Monthly Payment =  principal * [ interest rate / 12 ]
				  ---------------------------------
				  1 - (1 + interest rate / 12)^(-number of months)


*/

func main() {

	years := 5
	interestRate := float64(0.0425)
	totalAmount := float64(1000.25)
	months := years * 12
	payment := float64(39)

	principal := totalAmount
	for i := 1; i <= months; i++ {
		fmt.Printf("Mnth %d: Starting Balance: %.2f ", i, principal)
		//remaining balance * (1 + interestRate / 12) - payment
		principal = principal*(1+interestRate/12) - payment
		fmt.Printf(": Remaining Balance: %.2f\n", principal)
	}

}
