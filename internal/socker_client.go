package internal

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/yeejiac/BlockChain/models"
	"github.com/yeejiac/BlockChain/src"
	"gopkg.in/ini.v1"
)

var socket_connection net.Conn
var mode models.Mode

func SetMode(modeType models.Mode) {
	mode = modeType
}

func StartClient() {
	cfg, err := ini.Load("./config/setting.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return
	}
	// addr := cfg.Section("socket").Key("addr").String() + ":"
	port := cfg.Section("socket").Key("port").String()
	addr := ":"
	if mode == models.TEST {
		addr = "172.28.0.3:"
	}
	fmt.Println("Try connect to socket server: " + addr + port)
	conn, err := net.Dial("tcp", addr+port)
	if err != nil {
		fmt.Println("Connect failed")
	}
	defer conn.Close()
	fmt.Println("Connect to server success")
	socket_connection = conn
	go TestServerHandle(conn)
	go sendHeartbeat(conn)
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
			break
		} else {
			log.Println(string(bs[:len]))
			HandleMsg(string(bs[:len]))
		}
	}
}

func HandelServerMsg(rawStr string) {
	if rawStr == "<3" {
		return
	}
	num, err := strconv.Atoi(rawStr[0:2])
	if err != nil {
		// handle error
		fmt.Println("Invalid message stucture: " + rawStr)
		return
	}

	switch num {
	case 10:
		log.Println("Get New Transaction")
		HandleNewTransaction(rawStr)
	case 20:
		log.Println("Get hash block str")
		res := HandleHashBlock(rawStr)
		SendMsg(res)
	default:
		fmt.Println("Invalid message: " + rawStr)
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

func TestServerHandle(conn net.Conn) {
	if mode == models.TEST {
		cfg, err := ini.Load("./test/test.ini")
		if err != nil {
			fmt.Printf("Fail to read file: %v", err)
			return
		}
		testList := cfg.Section("ServerTest").KeysHash()
		for key := range testList {
			conn.Write([]byte(testList[key] + "\n"))
			time.Sleep(1 * time.Second)
		}
	}
	// os.Exit(0)
}
