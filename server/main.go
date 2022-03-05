package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			fmt.Println(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	for {
		// read from connection, read until \n
		result, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			// handle error
			fmt.Println(err)
		}
		fmt.Println(result)
		if _, err := conn.Write([]byte("hello client\n")); err != nil {
			// handle error
			fmt.Println(err)

		}
		time.Sleep(time.Second * 1)
	}

}
