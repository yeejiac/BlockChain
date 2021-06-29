package internal

import (
	"log"
	"net"
)

func StartClient() {
	log.Println("Try connect to socket server")
	conn, err := net.Dial("tcp", ":1204")
	if err != nil {
		log.Println("Connect failed")
	}
	defer conn.Close()
	log.Println("Connect to server success")
	sendTCP(conn)

}

func sendTCP(conn net.Conn) {
	for {
		conn.Write([]byte("<3"))
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

}
