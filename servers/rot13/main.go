package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	var port = ":" + *flag.String("port", "8000", "The port that the program will be listening on to run the program")
	/*
	 * It is said that Julius Caesar used an encryption to send messages to his soldiers in the battlefield.
	 * This message would have the words he'd written reversed by a certain number of characters...
	 */

	li, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Error listenting on port", port, err.Error())
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal("Connection closed", err.Error())
		}

		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		ln := strings.ToLower(scanner.Text())
		bs := []byte(ln)
		reversedBs := rot13(bs)
		fmt.Fprintf(c, "%s - %s \n\n", ln, reversedBs)
	}
}

func rot13(bs []byte) []byte {
	_bs := make([]byte, len(bs))

	for i, v := range bs {
		if v <= 109 {
			_bs[i] = v + 13
		} else {
			_bs[i] = v - 13
		}
	}
	return _bs
}
