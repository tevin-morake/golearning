package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	portNum := flag.String("port", "8080", "It's the listening port for this tcp program")
	flag.Parse()
	// This clock is a TCP server that periodically writes the time
	port := "localhost:" + *portNum
	listener, err := net.Listen("tcp", port)
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
