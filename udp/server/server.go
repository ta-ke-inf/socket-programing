package main

import (
	"log"
	"net"
)

func main() {

	port := ":7778"
	udpAddr, err := net.ResolveUDPAddr("udp", port)
	CheckError(err)

	conn, err := net.ListenUDP("udp", udpAddr)
	CheckError(err)
	log.Println("Server Started listening on", port)

	for {
		handleClient(conn)
	}

}

func handleClient(conn *net.UDPConn) {

	buf := make([]byte, 1024)
	count, addr, err := conn.ReadFromUDP(buf)
	if err != nil {
		log.Println("Failed to reading:", err)
		return
	}
	log.Println("Received data from", addr.IP, addr.Port, ":", string(buf[:count]))
	message := "Hello Client !!"
	send_msg := []byte(message)
	_, err = conn.WriteToUDP(send_msg, addr)
	if err != nil {
		log.Println("Failed to write:", err)
		return
	}
	log.Println("Wrote data to", addr.IP, addr.Port, ":", message)
}

func CheckError(err error) {
	if err != nil {
		log.Println("Error:", err)
	}
}
