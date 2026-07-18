package main

import (
	"fmt"
)

type CreditAccount struct {
	BankAccount 

	Debt int
	MounthlyPayment int
}

func (c *CreditAccount) Print() {
	c.BankAccount.Print()
	fmt.Println(c.Debt)
	fmt.Println(c.MounthlyPayment)
}

