package accounts

import (
	"errors"
	"fmt"
)

// main function 없으면 compile 하지 않음

type bankAccount struct {
	// 구조체, 함수 뿐만 아니라 구조체 내부의 속성도 대문자로 만들어야 패키지 외부에서 사용가능(public)
	owner   string
	balance int
}

var errNoMoney = errors.New("can not withdraw, you don't have enough money")

// function을 이용해 construct를 대신하는 패턴을 이용
func NewAccount(owner string) *bankAccount {
	account := bankAccount{owner: owner, balance: 0}
	return &account
}

// struct Method
func (account *bankAccount) Deposit(amount int) {
	// 포인터 사용하지 않고 그냥 bankAccount로 받으면 account는 복사본을 받음
	// method를 보호할 때 사용(내부 변경을 원하지 않을 때)
	account.balance += amount
}

func (account bankAccount) Balance() int {
	return account.balance
}

func (account *bankAccount) WithDraw(amount int) error {
	if account.balance < amount {
		// Go lang에는  try except나 try catch 구문이 존재하지 않음
		return errNoMoney
	}
	account.balance -= amount
	return nil
}

func (account *bankAccount) ChangeOwner(newOnwer string) {
	account.owner = newOnwer
}

func (account bankAccount) Owner() string {
	return account.owner
}

// python class에서 __str__ 과 동일한 역할 print시 이 함수를 호출하게 된다
func (account bankAccount) String() string {
	return fmt.Sprint(account.owner, "'s account.\nhas: ", account.balance)
}
