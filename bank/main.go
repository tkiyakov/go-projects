package main

import (
	"fmt"
	"tamer/bank/account"
)

// import (
// 	"fmt"
// )

func main() {
	tamer := account.BankAccount{
		Owner:   "Tamer",
		Balance: 500,
	}

	adlet := account.BankAccount{
		Owner:   "Adlet",
		Balance: 1000,
	}

	credit := account.CreditAccount{
		BankAccount: account.BankAccount{
			Owner:   "Tamer",
			Balance: 20000,
		},
		Debt:            100000,
		MounthlyPayment: 30000,
	}

	save := account.SavingsAccount{
		BankAccount: account.BankAccount{
			Owner:   "Tamer",
			Balance: 5000,
		},
		InterestRate: 18.5,
	}

	// accounts := []Account{
	// 	&tamer,
	// 	&adlet,
	// 	&credit,
	// 	&save,
	// }

	// for _, acc := range accounts {
	// 	PrintAccount(acc)
	// }

	fmt.Println(tamer)
	fmt.Println(credit)
	fmt.Println(adlet)
	fmt.Println(save)

	// err := tamer.Transfer(&adlet, -300)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// tamer.Print()
	// adlet.Print()
	// fmt.Println()

	// credit.Print()

}
