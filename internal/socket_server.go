package internal

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

var connectionMap map[int]net.Conn

var sequence = 1

func StartServer() {
	connectionMap = make(map[int]net.Conn)
	li, err := net.Listen("tcp", ":1203")
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
		go Handle(conn)
		go SendHeartbeat(conn)
	}
}

func Handle(conn net.Conn) {
	for {
		bs := make([]byte, 1024)
		len, err := conn.Read(bs)
		if err != nil {
			log.Println(err)
			break
		} else {
			log.Println(string(bs[:len]))
		}
	}
	fmt.Println("連線中斷.")
}

func SendHeartbeat(conn net.Conn) {
	for {
		conn.Write([]byte("<3\n"))
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
