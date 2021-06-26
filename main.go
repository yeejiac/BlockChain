package main

import (
	"fmt"
	"os"
	"github.com/yeejiac/BlockChain/models"
	"github.com/yeejiac/BlockChain/internal"
	"github.com/yeejiac/BlockChain/src"
)

func main() {
	if os.Args[1] == "node" {
		fmt.Println("Node mode start")
		internal.StartClient()
	} else if os.Args[1] == "server" {
		var blockchain models.BlockChain
		blockchain.Difficulty = 1
		blockchain.Block_ary = append(blockchain.Block_ary, src.GenerateGenesisBlock())
		src.MineBlock(&blockchain)
		fmt.Println("Node mode start")
		// internal.StartServer()
		// var blockchain models.BlockChain
		// blockchain.Difficulty = 5
		// src.MineBlock(blockchain)
	} else {
		fmt.Println("Args error")
	}

}
