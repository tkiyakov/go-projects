package main

import "fmt"

type BankAccount struct {
	Owner string
	Balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.Balance += amount
}

func (b *BankAccount) Withdraw(amount int) bool  {
	if amount > b.Balance {
		return false
	}
	b.Balance -= amount
	return true
}


func (b *BankAccount) Transfer(to *BankAccount, amount int) {
	ok := b.Withdraw(amount)

	if ok {
		to.Deposit(amount)
	}

}


func (b *BankAccount) Print() {
	fmt.Println(b.Owner)
	fmt.Println(b.Balance)
}

func (b *BankAccount) IsRich() bool {
	return b.Balance > 0
}




func main() {
	tamer := BankAccount{
		Owner: "Tamer",
		Balance: 1000,
	}

	adlet := BankAccount{
		Owner: "Adlet",
		Balance: 1000,
	}

	tamer.Transfer(&adlet, 600)
	tamer.Print()
	adlet.Print()
	fmt.Println()

	fmt.Println(tamer.IsRich())
	fmt.Println(tamer.Owner)
}