package main

import (
	"fmt"

	"github.com/banking/main/accounts"
)

func main() {
	account := accounts.NewAccount("Danny")
	account.Deposit(10)
	fmt.Println(account.Balance())
	// err := account.WithDraw(20)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	fmt.Println(account)
}
