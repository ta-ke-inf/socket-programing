package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: ex) 'main.go :8080'")
		os.Exit(1)
	}
	port := os.Args[1]
	//bind IP/port
	localAdder, err := net.ResolveTCPAddr("tcp", port)
	CheckError(err)
	//接続
	conn, err := net.DialTCP("tcp", nil, localAdder)
	CheckError(err)

	//サーバーに対して書き込み
	message := "Hello Server:" + port
	_, err = conn.Write([]byte(message))
	CheckError(err)
	log.Println("Writed data to ", conn.RemoteAddr().String())

	//サーバーからのデータを読み込み
	buf := make([]byte, 1024)
	count, err := conn.Read(buf)
	CheckError(err)
	log.Println("Received from ", conn.RemoteAddr().String(), ":", string(buf[:count]))

}
