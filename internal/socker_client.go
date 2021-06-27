package internal

import (
	"fmt"
	"log"
	"net"
)

func StartClient() {
	res, err := sendTCP("127.0.0.1:1203", "hi")
	if err != nil {
		fmt.Println("Lose connection")
	} else {
		fmt.Println(res)
	}
}

func sendTCP(addr, msg string) (string, error) {
	log.Println("Try connect to socket server")
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println("Connect failed")
		return "", err
	}
	defer conn.Close()
	log.Println("Connect to server success")
	// send to socket
	conn.Write([]byte(msg))

	// listen for reply
	bs := make([]byte, 1024)
	len, err := conn.Read(bs)
	if err != nil {
		return "", err
	} else {
		return string(bs[:len]), err
	}
}
