// test project main.go

package main

import (
	"fmt"
	"net"
)

//ip和port常量的定义
const (
	serverip   = "localhost"
	serverport = "54321"
)

var (
	maxRead  = 1100
	msgStop  = []byte("cmdStop")
	msgStart = []byte("cmdContinue")
)

func main() {
	hostAndPort := serverip + ":" + serverport
	println(hostAndPort)
	listener := initServer(hostAndPort)
	for {
		conn, err := listener.Accept()
		checkError(err, "Accept: ")
		//开一个goroutines处理客户端消息，这是golang的特色，实现并发就只go一下就好
		go connectionHandler(conn)
	}
}

//负责tcp服务端的初始化
func initServer(hostAndPort string) *net.TCPListener {
	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
	checkError(err, "Resolving address:port failed: '"+hostAndPort+"'")
	listener, err := net.ListenTCP("tcp", serverAddr)
	println("Listening to: ", listener.Addr().String())
	return listener
}

func connectionHandler(conn net.Conn) {
	connFrom := conn.RemoteAddr().String()
	println("Connection from: ", connFrom)
	talktoclients(conn)
	for {
		var ibuf []byte = make([]byte, maxRead+1)
		length, err := conn.Read(ibuf[0:maxRead])
		ibuf[maxRead] = 0 //to prevent overflow
		switch err {
		case nil:
			handleMsg(length, err, ibuf)
		default:
			goto DISCONNECT
		}
	}
DISCONNECT:
	err := conn.Close()
	println("Closed connection:", connFrom)
	checkError(err, "Close:")
}

func talktoclients(to net.Conn) {
	wrote, err := to.Write(msgStart)
	checkError(err, "Write: wrote "+string(wrote)+" bytes.")
}

func handleMsg(length int, err error, msg []byte) {
	if length > 0 {
		for i := 0; ; i++ {
			if msg[i] == 0 {
				break
			}
		}
		fmt.Printf("Received data: %v", string(msg[0:length]))
		fmt.Println("    length:", length)
	}
}

//错误检查，
func checkError(err error, info string) {
	if err != nil {
		panic("ERROR: " + info + " " + err.Error()) //terminate
	}
}
