package main

import (
	"bufio"
	"go_project/src/NAT-penetration/define"
	"io"
	"log"
)

func main() {
	tcpConn, err := define.CreateConn(define.ControloerAddr)
	if err != nil {
		panic(err)
	}
	log.Printf("连接成功: %v\n", tcpConn.RemoteAddr().String())
	for {
		readString, err := bufio.NewReader(tcpConn).ReadString('\n')
		if err != nil {
			panic(err)
		}
		if readString == define.NewConnection {
			go messageFoward()
		}
	}
}

func messageFoward() {
	tunnelTcpConn, err := define.CreateConn(define.TunnelAddr)
	if err != nil {
		panic(err)
	}

	serverTcpConn, err := define.CreateConn(define.LocalHostAddr)
	if err != nil {
		panic(err)
	}

	go func() {
		_, err := io.Copy(serverTcpConn, tunnelTcpConn)
		if err != nil {
			panic(err)
		}
	}()
	go func() {
		_, err := io.Copy(tunnelTcpConn, serverTcpConn)
		if err != nil {
			panic(err)
		}
	}()

}
