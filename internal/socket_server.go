package internal

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/yeejiac/BlockChain/models"
	"gopkg.in/ini.v1"
)

var connectionMap map[int]net.Conn

var sequence = 1

func StartServer() {
	cfg, err := ini.Load("./config/setting.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return
	}
	addr := ":"
	port := cfg.Section("socket").Key("port").String()

	connectionMap = make(map[int]net.Conn)
	li, err := net.Listen("tcp", addr+port)
	if err != nil {
		log.Fatalln(err)
	}
	// defer li.Close()

	log.Println("Start tcp server")
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		connectionMap[sequence] = conn
		sequence++
		go waitClientMsg(conn)
		go sendHeartbeat(conn)
		go TestClientHandle(conn)
	}
}

func waitClientMsg(conn net.Conn) {
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
	fmt.Println("連線中斷.")
}

func sendHeartbeat(conn net.Conn) {
	for {
		conn.Write([]byte("<3"))
		time.Sleep(5 * time.Second)
	}
}

func BroadCast(msg string) {
	if len(connectionMap) == 0 {
		return
	}
	for key, val := range connectionMap {
		fmt.Println("Send to Connection : " + strconv.Itoa(key))
		val.Write([]byte(msg))
	}
	fmt.Println("BroadCast finished")
}

func HandleMsg(rawStr string) {
	if rawStr == "<3" {
		return
	}
	num, err := strconv.Atoi(rawStr[0:2])
	if err != nil {
		// handle error
		fmt.Println("Invalid message stucture")
		return
	}
	switch num {
	case 40:
		log.Println("Get Nonce")
	case 44:
		if mode == models.TEST {
			os.Exit(0)
		}
	default:
		fmt.Println("Invalid message")
	}
}

func TestClientHandle(conn net.Conn) {
	if mode == models.TEST {
		cfg, err := ini.Load("./test/test.ini")
		if err != nil {
			fmt.Printf("Fail to read file: %v", err)
			return
		}
		testList := cfg.Section("ClientTest").KeysHash()
		for key := range testList {
			conn.Write([]byte(testList[key] + "\n"))
			time.Sleep(1 * time.Second)
		}
	}
	// os.Exit(0)
}
