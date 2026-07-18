package account

import "fmt"

func ShowInfo(a Account) {
	switch v := a.(type) {
	case *SavingsAccount:
		fmt.Println(v.InterestRate)

	case *CreditAccount:
		fmt.Println(v.Debt, v.MounthlyPayment)

	case *BankAccount:
		fmt.Println(v.Balance)
	}
}
