package main

import (
	"bufio"
	"fmt"
	"net"
	"time"

	"caton/xw.wang/utils"
)

func main() {
	var tcpAddr *net.TCPAddr

	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")

	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)

	defer tcpListener.Close()

	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			continue
		}

		fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())
		go tcpPipe(tcpConn)
	}

}

func tcpPipe(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnected :" + ipStr)
		conn.Close()
	}()
	//	reader := bufio.NewReader(conn)
	reader := bufio.NewReaderSize(conn, 40960)

	for {
		fmt.Println("start:" + utils.GetNowUTC2())
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		fmt.Println("end:" + utils.GetNowUTC2())
		//		fmt.Println(string(message))
		msgbye := []byte(message)
		fmt.Println("msg's len:" + utils.Itoa(utils.Len(message)))
		fmt.Println("msg byte's len:" + utils.Itoa(len(msgbye)))

		msg := time.Now().String() + "\n"
		b := []byte(msg)
		conn.Write(b)
	}
}
