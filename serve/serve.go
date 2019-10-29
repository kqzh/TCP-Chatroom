package main

import (
	"fmt"
	"net"
	"strings"
)

var (
	onlineConns  = make(map[string]net.Conn)
	onlineNames  = make(map[string]string)
	onlineAddrs  = make(map[string]string)
	messageQueue = make(chan string, 1000)
)

func ProcessInfo(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf[:])

		if err != nil {
			fmt.Printf("Connection Over:%v\n", err)
			delete(onlineConns, conn.RemoteAddr().String())
			break
		}
		if n != 0 {
			message := string(buf[:n])
			if string(message) == "quit" {
				conn.Close()
			} else {
				messageQueue <- message
			}
		}
	}
}

func ConsumeMessage() {
	for {
		select {
		case message := <-messageQueue:
			ProcessMessage(message)
		}
	}
}

func ProcessMessage(s string) {
	contents := strings.Split(s, "#")

	if contents[1] == "list" {
		msg := ""
		for i, _ := range onlineConns {
			msg += onlineAddrs[i] + " "
		}
		_, err := onlineConns[contents[0]].Write([]byte("online members#" + msg))
		if err != nil {
			fmt.Println("Error:", err)
		}
	} else if contents[1] == "setName" {
		onlineNames[contents[2]] = contents[0]
		onlineAddrs[contents[0]] = contents[2]
	} else if contents[1] == "quit" {
		onlineConns[contents[0]].Close()
	} else if len(contents) > 2 {
		fromAddr := onlineAddrs[contents[0]]
		sendMsg := contents[1]
		toAddr := onlineNames[contents[2]]

		if conn, ok := onlineConns[toAddr]; ok {
			_, err := conn.Write([]byte(fromAddr + "#" + sendMsg))
			if err != nil {
				fmt.Println("Error:", err)
			}
		}
	}
}

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:20000")

	if err != nil {
		fmt.Printf("Listen failed,err:%v\n", err)
	}
	defer listen.Close()

	fmt.Println("Serve is waiting...")

	// distribute message
	go ConsumeMessage()

	// listen connection
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("Accept failed,err:%v\n", err)
			continue
		}
		addr := conn.RemoteAddr().String()
		onlineConns[addr] = conn
		fmt.Println("visitor", addr, "connected")

		_, err = conn.Write([]byte("$" + addr))
		if err != nil {
			fmt.Println("Err,", err)
		}
		go ProcessInfo(conn)
	}
}