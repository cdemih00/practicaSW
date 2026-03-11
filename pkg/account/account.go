// Package account provides a deliberately SHALLOW bank account module.
//
// This is the starting point for the coding dojo. Every field is exposed
// through getters and setters, and all business logic lives in free functions
// outside the struct. The caller must interrogate the account, make every
// decision, and mutate its state directly.
//
// Your job: refactor this through three rounds into a DEEP module.
package account

import (
	"errors"
	"fmt"
)

// Account is a shallow data holder — all fields are accessed through
// getters and setters, and all logic lives outside the struct.
type Account struct {
	owner          string
	balance        float64
	overdraftLimit float64
	frozen         bool
}

// New creates a new account with the given owner and overdraft limit.
func New(owner string, overdraftLimit float64) *Account {
	return &Account{
		owner:          owner,
		overdraftLimit: overdraftLimit,
	}
}

// Withdraw removes money from the account.
// The caller must deal with: frozen checks, amount validation, overdraft logic.

func (account *Account) Withdraw(amount float64) error {
	if account.frozen {
		return errors.New("account is frozen")
	}
	if amount <= 0 {
		return fmt.Errorf("invalid amount: %.2f", amount)
	}
	if account.balance-amount < account.overdraftLimit {
		return errors.New("insufficient funds")
	}

	account.balance = account.balance - amount

	return nil
}

// Deposit adds money to the account.
// The caller must deal with: frozen checks, amount validation.

func (account *Account) Deposit(amount float64) error {
	if account.frozen {
		return errors.New("account is frozen")
	}
	if amount <= 0 {
		return fmt.Errorf("invalid amount: %.2f", amount)
	}
	account.balance = account.balance + amount
	return nil
}

// Transfer moves money from one account to another.
// The caller must deal with: frozen checks on both accounts, amount validation,
// overdraft logic — reaching into the internals of two objects.

func (from *Account) Transfer(to *Account, amount float64) error {
	if from.frozen {
		return errors.New("source account is frozen")
	}
	if to.frozen {
		return errors.New("destination account is frozen")
	}
	if amount <= 0 {
		return fmt.Errorf("invalid amount: %.2f", amount)
	}
	if from.balance-amount < from.overdraftLimit {
		return errors.New("insufficient funds")
	}
	from.balance = from.balance - amount
	to.balance = to.balance + amount
	return nil
}

// Freeze prevents any further operations on the account.
func (account *Account) Freeze() error {
	if account.frozen {
		return errors.New("account is already frozen")
	}
	account.frozen = true
	return nil
}

// Unfreeze re-enables operations on the account.
func (account *Account) Unfreeze() error {
	if !account.frozen {
		return errors.New("account is not frozen")
	}
	account.frozen = false
	return nil
}
