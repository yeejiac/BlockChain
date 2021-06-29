package internal

import (
	"log"
	"net"
)

func StartClient() {
	log.Println("Try connect to socket server")
	conn, err := net.Dial("tcp", "127.0.0.1:1203")
	if err != nil {
		log.Println("Connect failed")
	}
	defer conn.Close()
	log.Println("Connect to server success")
	sendTCP(conn)

}

func sendTCP(conn net.Conn) {
	for {
		conn.Write([]byte("HI"))
		bs := make([]byte, 1024)
		len, err := conn.Read(bs)
		if err != nil {
			log.Println(err)
		} else {
			log.Println(string(bs[:len]))
		}
	}
}
