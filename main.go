package main

import (
	"fmt"
	"os"

	"github.com/yeejiac/BlockChain/internal"
	"github.com/yeejiac/BlockChain/models"
	"github.com/yeejiac/BlockChain/src"
)

func main() {
	if os.Args[1] == "node" {
		fmt.Println("Node mode start")
		internal.StartClient()
	} else if os.Args[1] == "server" {
		var blockchain models.BlockChain
		var tempblock models.Block
		blockchain.Difficulty = 1
		blockchain.Block_ary = append(blockchain.Block_ary, src.GenerateGenesisBlock())
		tempblock = src.MineBlock(&blockchain)
		if src.CheckCorrectness(tempblock, 5, blockchain.Difficulty) {
			fmt.Println("correct")
		} else {
			fmt.Println("false")
		}
		fmt.Println("Node mode start")

	} else {
		fmt.Println("Args error")
	}

}
