package main

import (
	"fmt"
	"os"

	"github.com/yeejiac/BlockChain/internal"
	"github.com/yeejiac/BlockChain/models"
	"github.com/yeejiac/BlockChain/src"
)

func main() {
	// f, err := os.OpenFile("./log/testlogfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalf("error opening file: %v", err)
	// }
	// defer f.Close()
	// log.SetOutput(f)
	if os.Args[2] == "test" {
		fmt.Println("Test mode start")
		internal.SetMode(models.TEST)
	}
	if os.Args[1] == "node" {
		fmt.Println("Node mode start")
		internal.StartClient()
	} else if os.Args[1] == "server" {
		src.InitBlockChain()
		// tempblock = src.MineBlock(&blockchain)
		// if src.CheckCorrectness(tempblock, 5, blockchain.Difficulty) {
		// 	fmt.Println("correct")
		// } else {
		// 	fmt.Println("false")
		// }
		fmt.Println("Server mode start")
		internal.StartServer()
	} else {
		fmt.Println("Args error")
	}

}
