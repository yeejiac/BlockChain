package main

import (
	"fmt"

	"github.com/yeejiac/BlockChain/models"
	"github.com/yeejiac/BlockChain/src"
)

func main() {
	a := models.Transaction{Sender: "Yee", Receiver: "Arisa", Amounts: "56", Fee: "50", Message: "test"}
	src.GenerateHashValue(a)
	fmt.Println(src.GenerateNonce(5))
}
