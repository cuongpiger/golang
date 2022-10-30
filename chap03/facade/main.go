package main

import (
	"fmt"
)

type (
	// Account is part of the subsystem.
	Account struct {
		name string
	}

	// Wallet is part of the subsystem.
	Wallet struct {
		balance int
	}

	// SecurityCode is part of the subsystem.
	SecurityCode struct {
		code int
	}

	// Notification is part of the subsystem.
	Notification struct{}

	// Ledger is part of the subsystem.
	Ledger struct{}

	WalletFacade struct {
		account      *Account
		wallet       *Wallet
		securityCode *SecurityCode
		notification *Notification
		ledger       *Ledger
	}
)

// Account's collection of methods
func newAccount(accountName string) *Account {
	return &Account{name: accountName}
}

func (s *Account) checkAccount(accountName string) error {
	if s.name != accountName {
		return fmt.Errorf("account name is not correct")
	}

	fmt.Println("Account is correct")
	return nil
}

// Wallet's collection of methods
func newWallet() *Wallet {
	return &Wallet{balance: 0}
}

func (s *Wallet) creditBalance(amount int) {
	s.balance += amount
	fmt.Printf("Wallet balance is %d", s.balance)
}

func (s *Wallet) debitBalance(amount int) error {
	if s.balance < amount {
		return fmt.Errorf("not enough balance")
	}

	s.balance -= amount
	fmt.Printf("Wallet balance is %d", s.balance)
	return nil
}

// SecurityCode's collection of methods
func newSecurityCode(code int) *SecurityCode {
	return &SecurityCode{code: code}
}

func (s *SecurityCode) checkCode(code int) error {
	if s.code != code {
		return fmt.Errorf("security code is not correct")
	}

	fmt.Println("Security code is correct")
	return nil
}

// Notification's collection of methods
func (s *Notification) sendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
}

func (s *Notification) sendWalletDebitNotification() {
	fmt.Println("Sending wallet debit notification")
}

// Ledger's collection of methods
func (s *Ledger) makeEntry(accountID, txnType string, amount int) {
	fmt.Printf("Make ledger entry for accountId %s with txnType %s for amount %d\n", accountID, txnType, amount)
}

// WalletFacade's collection of methods
func newWalletFacade(accountID string, code int) *WalletFacade {
	fmt.Println("Starting create account")
	walletFacade := &WalletFacade{
		account:      newAccount(accountID),
		securityCode: newSecurityCode(code),
		wallet:       newWallet(),
		notification: new(Notification),
		ledger:       new(Ledger),
	}
	fmt.Println("Account created")
	return walletFacade
}

func (w *WalletFacade) addMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	w.wallet.creditBalance(amount)
	w.notification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}

func (w *WalletFacade) deductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	err = w.wallet.debitBalance(amount)
	if err != nil {
		return err
	}
	w.notification.sendWalletDebitNotification()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}

// main function
func main() {
	walletFacade := newWalletFacade("abc", 1234)
	fmt.Println()

	err := walletFacade.addMoneyToWallet("abc", 1234, 10)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.deductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}
