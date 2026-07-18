package account

import (
	"fmt"
	// "errors"
)

type SavingsAccount struct {
	BankAccount

	InterestRate float64
}

func (s *SavingsAccount) Print() {
	s.BankAccount.Print()
	fmt.Println(s.InterestRate)
}

func (s *SavingsAccount) ApplyInterest() error {
	interest := float64(s.Balance) * s.InterestRate / 100

	err := s.Deposit(int(interest))

	if err != nil {
		return err
	}
	return nil

	// но можно написать короче
	// -- return s.Deposit(int(interest)) --
	// это значит если будет ошибка или успех, все вернется в любом случае

}
