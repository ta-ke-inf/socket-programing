package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: ex) 'main.go :8080'")
		os.Exit(1)
	}
	port := os.Args[1]

	localAdder, err := net.ResolveTCPAddr("tcp", port)
	CheckError(err)

	conn, err := net.DialTCP("tcp", nil, localAdder)
	CheckError(err)

	message := "Hello Server:" + port
	_, err = conn.Write([]byte(message))
	CheckError(err)
	log.Println("Writed data to ", conn.RemoteAddr().String())

	buf := make([]byte, 1024)
	count, err := conn.Read(buf)
	CheckError(err)
	log.Println("Received from ", conn.RemoteAddr().String(), ":", string(buf[:count]))

}

func CheckError(err error) {

	if err != nil {
		log.Println("Error:", err)
	}
}
