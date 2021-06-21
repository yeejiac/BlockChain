package src

import (
	"encoding/json"
	"fmt"

	"github.com/yeejiac/BlockChain/models"
)

func GenerateHashValue(transaction models.Transaction) string {
	out, err := json.Marshal(transaction)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
	return string(out)
}
