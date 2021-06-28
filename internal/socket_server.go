package internal

import (
	"bufio"
	"fmt"
	"log"
	"net"
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
	}
}

func Handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()

	fmt.Println("連線中斷.")
}

func SendHeartbeat(conn net.Conn) {

}
