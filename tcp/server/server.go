package main

import (
	"log"
	"net"
)

func main() {

	port := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp", port)
	CheckError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	CheckError(err)
	log.Println("Server Started. Listening on ", port)

	receiveTCPConnection(listener)
}

func receiveTCPConnection(listener *net.TCPListener) {

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Println("Failed to accept client connection:", err)
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn *net.TCPConn) {

	defer conn.Close()
	buf := make([]byte, 1024)
	count, err := conn.Read(buf)
	if err != nil {
		log.Println("Failed to read client data", err)
		return
	}
	log.Println("Received data from ", conn.RemoteAddr().String(), ":", string(buf[:count]))

	message := "Hello Client:" + conn.RemoteAddr().String()
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Println("Failed ")
		return
	}
	log.Println("Writed data to ", conn.RemoteAddr().String())
}

func CheckError(err error) {

	if err != nil {
		log.Println("Error:", err)
	}
}
