package main

import (
	"io"
	"log"
	"net"
	"os"
)

//TODO: Finish this off. Dial not yet working properly
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal("Error establishing a dial up connection :", err.Error())
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal("Something bad happened here", err.Error())
	}

}
