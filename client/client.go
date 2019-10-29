package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var (
	myAddr string
	input string
)
func MessageSend(conn net.Conn){
	reader:=bufio.NewReader(os.Stdin)

	// set login name
	fmt.Print("\nplease set your name : ")
	input,_ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	input = myAddr+"#setName#"+input
	conn.Write([]byte(input))

	// print usage
	fmt.Println("\nwelcome visitor, you can send message to anyone who is online too!\n")
	fmt.Println("send message 'list' to get online members\n")
	fmt.Println("send message 'quit' to exit\n")
	fmt.Println("message format is 'xxx #name' \n")
	//send messages
	for{
		input,_ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
		input = myAddr+"#"+input
		fmt.Println()
		_,err:=conn.Write([]byte(input))
		if err!=nil{
			fmt.Println("Send failed,err:",err)
			break
		}
	}
}
func main(){
	// make connection
	conn,err :=net.Dial("tcp","127.0.0.1:20000")
	if err!=nil{
		fmt.Println("Dial failed,err:",err)
	}
	defer conn.Close()

	go MessageSend(conn)

	//listen message
	buf :=make([]byte,1024)
	for{
		n,err:=conn.Read(buf)
		if err!=nil{
			fmt.Println("have a good day!")
			break
		}
		tmp :=string(buf[:n])
		if tmp[0]=='$'{
			myAddr = tmp[1:]
			continue
		}
		msg := strings.Split(tmp,"#")
		fmt.Println(msg[0],":",msg[1])
		fmt.Println()
	}
}