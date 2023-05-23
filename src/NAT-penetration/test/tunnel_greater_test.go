package test

import (
	"bufio"
	"io"
	"log"
	"net"
	"sync"
	"testing"
	"time"
)

const (
	controlerAddr     = "0.0.0.0:22222"
	requestServerAddr = "0.0.0.0:22223"
)

var wg sync.WaitGroup
var clientConn *net.TCPConn

func controler() {
	log.Println("contoler start!")
	tcpListener, err := createListener(controlerAddr)
	if err != nil {
		panic(err)
	}
	for {
		clientConn, err = tcpListener.AcceptTCP()
		if err != nil {
			panic(err)
		}
		//go keepAlive(clientConn)
	}
}

func requestServer() {
	log.Println("requestServer start!")
	tcpListener, err := createListener(requestServerAddr)
	if err != nil {
		panic(err)
	}
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			panic(err)
		}
		go func() {
			i, err := io.Copy(clientConn, tcpConn)
			if err != nil {
				log.Println(i)
				panic(err)
			}
			log.Println("copy tcpConn to clientConn")
		}()
		go func() {
			i, err := io.Copy(tcpConn, clientConn)
			if err != nil {
				log.Println(i)
				panic(err)
			}
			log.Println("copy clientConn to tcpConn")
		}()
	}
}

func createListener(addr string) (*net.TCPListener, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}
	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return nil, err
	}
	return tcpListener, nil
}

func createConn(addr string) (*net.TCPConn, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}
	tcpConn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, err
	}
	return tcpConn, nil
}

func keepAlive(conn *net.TCPConn) {
	for {
		_, err := conn.Write([]byte("keep alive\n"))
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second * 3)
	}
}

func TestControlCenter(t *testing.T) {
	wg.Add(1)
	go controler()
	go requestServer()
	wg.Wait()
}

func TestNAT(t *testing.T) {
	tcpConn, err := createConn(controlerAddr)
	if err != nil {
		panic(err)
	}
	for {
		readString, err := bufio.NewReader(tcpConn).ReadString('\n')
		if err != nil {
			panic(err)
		}
		log.Println(readString)
		if readString == "keep alive\n" {
			continue
		}
		_, err = tcpConn.Write([]byte(readString))
		if err != nil {
			panic(err)
		}
		break
	}
	err = tcpConn.Close()
	if err != nil {
		panic(err)
	}
}

func TestUserClient(t *testing.T) {
	tcpConn, err := createConn(requestServerAddr)
	if err != nil {
		panic(err)
	}
	t.Log("tcp conn create!")
	time.Sleep(time.Second * 100)
	_, err = tcpConn.Write([]byte("i am net package\n"))
	if err != nil {
		panic(err)
	}
	respFromServer := make([]byte, 0)
	for {
		var buffer [bufferSize]byte
		n, err := tcpConn.Read(buffer[:])
		if err != nil {
			t.Fatal(err)
		}
		respFromServer = append(respFromServer, buffer[:n]...)
		if n < bufferSize {
			break
		}
	}
	err = tcpConn.Close()
	if err != nil {
		panic(err)
	}
	log.Println(respFromServer)
}
