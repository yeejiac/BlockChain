package internal

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/yeejiac/BlockChain/models"
	"github.com/yeejiac/BlockChain/src"
)

// var socket_connection net.Conn

func StartClient() {
	fmt.Println("Try connect to socket server")
	conn, err := net.Dial("tcp", ":1203")
	if err != nil {
		fmt.Println("Connect failed")
	}
	defer conn.Close()
	fmt.Println("Connect to server success")
	// socket_connection = conn
	go sendInput(conn)
	sendTCP(conn)
}

func sendInput(conn net.Conn) {
	for {
		var tempStr string
		fmt.Println("Input string :")
		fmt.Scanln(&tempStr)
		conn.Write([]byte(tempStr))
	}
}

// func SendMsg(str string) {
// 	socket_connection.Write([]byte(str))
// }

func sendTCP(conn net.Conn) {
	for {
		bs := make([]byte, 1024)
		len, err := conn.Read(bs)
		if err != nil {
			log.Println(err)
		} else {
			log.Println(string(bs[:len]))
		}
	}
}

func HandleMsg(rawStr string) {
	num, err := strconv.Atoi(rawStr[0:2])
	if err != nil {
		// handle error
		fmt.Println("Invalid message stucture")
		return
	}

	switch num {
	case 10:
		log.Println("Get New Transaction")
		HandleNewTransaction(rawStr)
	case 20:
		log.Println("Get hash block str")
		// res := HandleHashBlock(rawStr)
		// SendMsg(res)
	case 2:
		log.Println("C")
	case 3:
		log.Println("D")
	default: //default:當前面條件都沒有滿足時將會執行此處內包含的方法
		fmt.Println("Invalid message")
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

func HandleHashBlock(hashStr string) string {
	nonce := src.NonceCalculate(hashStr)
	return strconv.Itoa(nonce)
}
