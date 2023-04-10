package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	go read(conn)
	input(conn)
}

func read(conn net.Conn) {
	r := bufio.NewReader(conn)
	for {
		message, err := r.ReadString('\n')
		fmt.Println("Message from server : " + message)
		fmt.Println("Text to send: ")
		if err != nil {
			fmt.Println(err)
			if err != io.EOF {
				return
			}
			break
		}
	}
	fmt.Print("Read is closing")
}

func input(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Text to send: ")
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Fprintf(conn, text)
	}
}
