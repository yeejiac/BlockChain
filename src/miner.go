package src

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/yeejiac/BlockChain/models"
)

func GenerateDifficultyTarget(times int) string {
	res := strings.Repeat("0", times)
	return res
}

func MineBlock(blockchain models.BlockChain) {
	fmt.Println("Mine Block")
	var block models.Block
	last_block := GenerateGenesisBlock()
	block.Previous_hash = last_block.Hash
	block.Hash = GenerateBlockHash(block, 0)

	temphash := block.Hash[0:blockchain.Difficulty]
	a := 1
	for temphash != GenerateDifficultyTarget(blockchain.Difficulty) {
		a++
		block.Hash = GenerateBlockHash(block, a)
		fmt.Println(block.Hash + " " + strconv.Itoa(a))
		temphash = block.Hash[0:blockchain.Difficulty]
		fmt.Println(temphash)
	}
	fmt.Println("Hash found " + block.Hash)
}
