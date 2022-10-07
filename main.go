package main

import (
	"fmt"

	"github.com/Eric-lab-star/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("kim")
	fmt.Println(account)
}
