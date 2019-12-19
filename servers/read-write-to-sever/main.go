package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal("Error listening for open connection", err.Error())
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error establishing tcp connection :", err.Error())
		}
		go handleConn(conn)
	}

}

// Use telnet or netstat to connect to the server
func handleConn(c net.Conn) {
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(c, "I heard you say %s\n", ln)
	}
	defer c.Close()
	fmt.Println("I guess we're done for now")
}
