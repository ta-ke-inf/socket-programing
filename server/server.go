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
	//ソケットの作成、TCP の IP:port
	tcpAddr, err := net.ResolveTCPAddr("tcp", port)
	CheckError(err)

	//リスナーを作成して接続を待機
	listener, err := net.ListenTCP("tcp", tcpAddr)
	CheckError(err)
	log.Println("Server Started. Listening on ", port)

	receiveTCPConnection(listener)

}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// acceptで受け取る
func receiveTCPConnection(listener *net.TCPListener) {

	//クライアントから受け取り続ける
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Println("Failed to accept client connection:", err)
			continue
		}
		//AccpetTCPは接続が確立されるまでブロックされるため並行処理させる
		go handleClient(conn)
	}
}

func handleClient(conn *net.TCPConn) {
	defer conn.Close()

	//クライアントからのデータの読み込み処理
	//1024byteのスライス型バッファを作成
	buf := make([]byte, 1024)
	count, err := conn.Read(buf)
	if err != nil {
		log.Println("Failed to read client data", err)
		return
	}
	log.Println("Received data from ", conn.RemoteAddr().String(), ":", string(buf[:count]))

	//書き込み
	message := "Hello Client:" + conn.RemoteAddr().String()
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Println("Failed ")
		return
	}
	log.Println("Writed data to ", conn.RemoteAddr().String())
}
