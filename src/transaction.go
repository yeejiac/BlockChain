package src

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/yeejiac/BlockChain/models"
)

func GenerateNonce(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func GenerateHashValue(transaction models.Transaction) string {
	out, err := json.Marshal(transaction)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
	return string(out)
}

func GenerateTransactionHash(block models.Block) string {
	res := ""
	for _, value := range block.Transaction_ary {
		res += GenerateHashValue(value)
	}
	return res
}

func GenerateBlockHash(block models.Block, nonce string) string {
	blockStr := block.Previous_hash + block.Timestamp + GenerateTransactionHash(block) + block.Nonce
	h := sha1.New()
	h.Write([]byte(blockStr))
	return string(h.Sum([]byte{}))
}

func GenerateGenesisBlock() *models.Block {
	fmt.Println("Generate genesis block")
	var block models.Block
	return &block
}
