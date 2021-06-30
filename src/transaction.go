package src

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"math/rand"

	"github.com/yeejiac/BlockChain/models"
)

var blockchain models.BlockChain
var block models.Block

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

func GenerateBlockHash(block models.Block, nonce int) string {
	blockStr := block.Previous_hash + block.Timestamp + GenerateTransactionHash(block) + GenerateNonce(nonce)
	h := fnv.New32a()
	h.Write([]byte(blockStr))

	strval := fmt.Sprintf("%0*d", 10, h.Sum32())
	str := fmt.Sprint(strval)
	return str
}

func GenerateBlockHashStr(blockstr string, nonce int) string {
	h := fnv.New32a()
	h.Write([]byte(blockstr))

	strval := fmt.Sprintf("%0*d", 10, h.Sum32())
	str := fmt.Sprint(strval)
	return str
}

func GenerateGenesisBlock() models.Block {
	fmt.Println("Generate genesis block")
	var block models.Block
	block.Previous_hash = "Generate genesis block"
	block.Difficulty = 1
	block.Miner = "Test123"
	return block
}

func AppendBlock(block models.Block) {
	blockchain.Block_ary = append(blockchain.Block_ary, block)
}

func InitBlockChain() {
	blockchain.Difficulty = 1
	blockchain.Block_ary = append(blockchain.Block_ary, GenerateGenesisBlock())

}

func InitNewBlock() models.Block {
	block.Difficulty = 3
	block.Nonce = 0
	return block
}
