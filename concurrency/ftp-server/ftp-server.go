package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {
	port := *flag.String("port", "8000", "port used to connect to program")
	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal("Error listening to open connections :", err.Error())
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Error establishing tcp connection :", err.Error())
	}

	handleConn(conn)
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)

	for input.Scan() {
		fmt.Fprintln(c, "\t", input.Text())
	}
	c.Close()
}
