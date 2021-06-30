package internal

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/yeejiac/BlockChain/models"
	"github.com/yeejiac/BlockChain/src"
)

func StartClient() {
	log.Println("Try connect to socket server")
	conn, err := net.Dial("tcp", ":1204")
	if err != nil {
		log.Println("Connect failed")
	}
	defer conn.Close()
	log.Println("Connect to server success")
	go sendInput(conn)
	sendTCP(conn)
}

func sendInput(conn net.Conn) {
	for {
		var tempStr string
		fmt.Println("Input string :\n")
		fmt.Scanln(&tempStr)
		conn.Write([]byte(tempStr))
	}
}

func sendTCP(conn net.Conn) {
	for {
		// conn.Write([]byte("<3"))
		bs := make([]byte, 1024)
		len, err := conn.Read(bs)
		if err != nil {
			log.Println(err)
		} else {
			log.Println(string(bs[:len]))
		}
	}
}

func HandleNewTransaction(rawStr string) {
	tempAry := strings.Split(rawStr, "|")
	var transaction models.Transaction
	transaction.Sender = tempAry[0]
	transaction.Receiver = tempAry[1]
	transaction.Amounts = tempAry[2]
	transaction.Fee = 10
	temp := src.InitNewBlock()
	temp.Transaction_ary = append(temp.Transaction_ary, transaction)
	src.AppendBlock(temp)
}
