package main

import (
	"fmt"
	"errors"
)


type BankAccount struct {
	Owner string
	Balance int
}

func (b *BankAccount) Deposit(amount int) error {
	if amount <= 0 {
		return errors.New("сумма должна быть больше нуля")
	} 

	b.Balance += amount
	return nil
}

func (b *BankAccount) Withdraw(amount int) error  {
	if amount <= 0 {
		return errors.New("Сумма должна быть больше нуля")
	}

	if amount > b.Balance {
		return errors.New("недостаточно средств")
	}

	b.Balance -= amount
	return nil
}


func (b *BankAccount) Transfer(to *BankAccount, amount int) error {
	if b == to {
		return errors.New("ты не можешь отправить самому себе")
	}

	err := b.Withdraw(amount) 

	if err != nil {
		return err	
	}
	err = to.Deposit(amount)

	if err != nil {
		return err
	}
	return nil

}


func (b *BankAccount) Print() {
	fmt.Println(b.Owner)
	fmt.Println(b.Balance)
}

func (b *BankAccount) IsRich() bool {
	return b.Balance > 0
}
