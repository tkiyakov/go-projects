package main 

// import (
// 	"fmt"
// )

func main() {
	tamer := BankAccount{
		Owner: "Tamer",
		Balance: 500,
	}

	adlet := BankAccount{
		Owner: "Adlet",
		Balance: 1000,
	}

	credit := CreditAccount{
		BankAccount: BankAccount{
			Owner: "Tamer",
			Balance: 10000,
		},
		Debt: 100000,
		MounthlyPayment: 20000,
	}

	save := SavingsAccount{
		BankAccount: BankAccount{
			Owner: "Tamer",
			Balance: 10000,
		},
		InterestRate: 18.5,
	}

	accounts := []Account{
		&tamer,
		&adlet,
		&credit,
		&save,
	}

	for _, acc := range Account {
		&tamer,
		&adlet,
		&credit,
		&save,
	}



	// err := tamer.Transfer(&adlet, -300)
	// if err != nil {
	// 	fmt.Println(err)
	// }


	
	// tamer.Print()
	// adlet.Print()
	// fmt.Println()

	// credit.Print()
	

}