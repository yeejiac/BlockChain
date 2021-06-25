package main

import (
	"fmt"
	"os"

	"github.com/yeejiac/BlockChain/models"
	"github.com/yeejiac/BlockChain/src"
)

func main() {
	if os.Args[1] == "node" {
		fmt.Println("Node mode start")
	} else if os.Args[1] == "server" {
		fmt.Println("Node mode start")
		var blockchain models.BlockChain
		blockchain.Difficulty = 5
		src.MineBlock(blockchain)
	} else {
		fmt.Println("Args error")
	}

}
