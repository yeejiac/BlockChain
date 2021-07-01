package internal

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/yeejiac/BlockChain/models"
	"github.com/yeejiac/BlockChain/src"
	"gopkg.in/ini.v1"
)

var socket_connection net.Conn

func StartClient() {
	cfg, err := ini.Load("./config/setting.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return
	}
	addr := ":"
	port := cfg.Section("socket").Key("port").String()
	fmt.Println("Try connect to socket server: " + addr + port)
	conn, err := net.Dial("tcp", addr+port)
	if err != nil {
		fmt.Println("Connect failed")
	}
	defer conn.Close()
	fmt.Println("Connect to server success")
	socket_connection = conn
	go sendInput()
	receive(conn)
}

func sendInput() {
	for {
		var tempStr string
		fmt.Println("Input string :")
		fmt.Scanln(&tempStr)
		socket_connection.Write([]byte(tempStr))
	}
}

func SendMsg(str string) {
	socket_connection.Write([]byte(str))
}

func receive(conn net.Conn) {
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
	case 40:
		log.Println("Get Nonce")
	default:
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
