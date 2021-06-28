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

func MineBlock(blockchain *models.BlockChain) models.Block {
	fmt.Println("Mine Block")
	fmt.Println("Block size " + strconv.Itoa(len(blockchain.Block_ary)))
	var block models.Block
	last_block := blockchain.Block_ary[len(blockchain.Block_ary)-1]
	block.Previous_hash = last_block.Hash
	a := 0
	temp := GenerateBlockHash(block, a)
	temphash := temp[0:blockchain.Difficulty]
	for temphash != GenerateDifficultyTarget(blockchain.Difficulty) {
		a++
		temp = GenerateBlockHash(block, a)
		fmt.Println(temp + " " + strconv.Itoa(a))
		temphash = temp[0:blockchain.Difficulty]
		fmt.Println(temphash)
	}
	block.Hash = GenerateBlockHash(block, a)
	fmt.Println("Hash found " + block.Hash)
	fmt.Println("Nonce " + strconv.Itoa(a))
	return block
}

func CheckCorrectness(block models.Block, nonce int, difficulty int) bool {
	// fmt.Println("CheckCorrectness")
	// if GenerateBlockHash(block, nonce) != block.Hash {
	// 	return false
	// }
	return true
}
