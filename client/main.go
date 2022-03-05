package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		// handle error
		fmt.Println("Error:", err)
	}
	defer conn.Close()
	for {
		// read from connection, read until \n
		if _, err := conn.Write([]byte("hello server\n")); err != nil {
			// handle error
			fmt.Println(err)
		}

		result, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			// handle error
			fmt.Println(err)
		}
		fmt.Println(result)
		time.Sleep(time.Second * 1)
	}
	// status := make([]byte, 1024)
	// _, err = conn.Read(status)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(status)
}
