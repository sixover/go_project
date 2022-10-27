package main

import (
	"go_project/src/NAT-penetration/define"
	"io"
	"log"
	"net"
	"sync"
)

var controlTcpConn *net.TCPConn
var userTcpConn *net.TCPConn
var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go controlListen()

	go userRequestListen()

	go tunnelListen()
	wg.Wait()
}

func tunnelListen() {
	listener, err := define.CreateListener(define.TunnelAddr)
	if err != nil {
		panic(err)
	}

	log.Printf("隧道监听中: %v\n", listener.Addr().String())
	for {
		TcpConn, err := listener.AcceptTCP()
		if err != nil {
			panic(err)
		}
		go func() {
			_, err := io.Copy(userTcpConn, TcpConn)
			if err != nil {
				panic(err)
			}
		}()
		go func() {
			defer TcpConn.Close()
			_, err := io.Copy(TcpConn, userTcpConn)
			if err != nil {
				panic(err)
			}
		}()
	}
}

func userRequestListen() {
	listener, err := define.CreateListener(define.UserRequestAddr)
	if err != nil {
		panic(err)
	}

	log.Printf("用户请求监听中: %v\n", listener.Addr().String())
	for {
		userTcpConn, err = listener.AcceptTCP()
		if err != nil {
			panic(err)
		}
		_, err = controlTcpConn.Write([]byte(define.NewConnection))
		if err != nil {
			panic(err)
		}
	}
}

func controlListen() {
	listener, err := define.CreateListener(define.ControloerAddr)
	if err != nil {
		panic(err)
	}

	log.Printf("控制中心监听中: %v\n", listener.Addr().String())
	for {
		controlTcpConn, err = listener.AcceptTCP()
		if err != nil {
			panic(err)
		}
		go define.KeepAliveService(controlTcpConn)
	}
}
