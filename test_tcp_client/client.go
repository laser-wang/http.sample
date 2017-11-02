package main

import (
	"bufio"
	"fmt"
	"net"
	"time"

	"caton/xw.wang/utils"
)

var quitSemaphore chan bool

func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")

	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	fmt.Println("connected!")

	go onMessageRecived(conn)

	msg := ""
	for i := 0; i < 100000; i++ {
		msg = msg + "1234567890"
	}
	fmt.Println("msg's len:" + utils.Itoa(utils.Len(msg)))

	quitSemaphore = make(chan bool, 1)

	b := []byte(msg + "\n")
	conn.Write(b)
	fmt.Println("5 end!")
	<-quitSemaphore
	fmt.Println("end!")
}

func onMessageRecived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	for {
		fmt.Println("1!")
		msg, err := reader.ReadString('\n')
		fmt.Println("2!")

		fmt.Println(msg)
		if err != nil {
			fmt.Println("3!" + err.Error())
			quitSemaphore <- true
			break
		}
		fmt.Println("3 end!")
		time.Sleep(time.Second)
		//		b := []byte(msg)
		//		conn.Write(b)

		quitSemaphore <- true
		fmt.Println("4 end!")
		break
	}

	fmt.Println("2 end!")

}
