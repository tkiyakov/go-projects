package main

type Account interface {
	Deposit(amount int) error
	Withdraw(amount int) error
	Print()
}

func PrintAccount(a Account) {
	a.Print()
}