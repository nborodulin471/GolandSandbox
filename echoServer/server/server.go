package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	tcp, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer tcp.Close()
	fmt.Println("Launching server...")
	for {
		conn, err := tcp.Accept()
		if err != nil {
			log.Println(err)
		}
		fmt.Println("Conn is running")
		go newConnect(conn)
	}
}

func newConnect(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		conn.Read(buf)
		fmt.Printf("Message Received: %s \n", string(buf))
		conn.Write(buf)
	}
}
