package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// This clock is a TCP server that periodically writes the time
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal("Errror opening tcp connection :", err.Error())
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting incoming connection : %s", err.Error()) // e.g connection aborted
			continue
		}

		// handleConn(conn) // handle one connection at a time
		go handleConn(conn) // handle more than one connection at a time (concurrent)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			fmt.Println("Connection closed")
			return
		}
		time.Sleep(1 * time.Second)
	}

}
