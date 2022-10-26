package test

import (
	"io"
	"net"
	"strconv"
	"testing"
)

const (
	serverAddr = "0.0.0.0:22222"
	tunnelAddr = "0.0.0.0:22223"
)

func TestServer(t *testing.T) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		t.Fatal(err)
	}

	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		t.Fatal(err)
	}

	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			t.Fatal(err)
		}

		resp := make([]byte, 0)
		for {
			var buffer [bufferSize]byte
			n, err := tcpConn.Read(buffer[:])
			if err != nil {
				t.Fatal(err)
			}
			resp = append(resp, buffer[:n]...)
			if n < bufferSize {
				break
			}
		}
		t.Log(resp)
		i, err := strconv.Atoi(string(resp))
		if err != nil {
			t.Fatal(err)
		}
		_, err = tcpConn.Write([]byte(strconv.Itoa(i + 2)))
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestClient(t *testing.T) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", tunnelAddr)
	if err != nil {
		t.Fatal(err)
	}

	tcpConn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		t.Fatal(err)
	}

	//发送数据
	_, err = tcpConn.Write([]byte(strconv.Itoa(2300)))
	if err != nil {
		t.Fatal(err)
	}

	//接收数据
	resp := make([]byte, 0)
	for {
		var buffer [bufferSize]byte
		n, err := tcpConn.Read(buffer[:])
		if err != nil {
			t.Fatal(err)
		}

		resp = append(resp, buffer[:n]...)
		if n < bufferSize {
			break
		}
	}
	t.Log(string(resp))
	err = tcpConn.Close()
	if err != nil {
		t.Fatal(err)
	}
}

func TestTunnel(t *testing.T) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", tunnelAddr)
	if err != nil {
		t.Fatal(err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		t.Fatal(err)
	}

	for {
		tcpConn, err := listener.AcceptTCP()
		if err != nil {
			t.Fatal(err)
		}

		//respFromClient :=make([]byte,0)
		//for{
		//	var buffer [bufferSize]byte
		//	n, err := tcpConn.Read(buffer[:])
		//	if err != nil {
		//		t.Fatal(err)
		//	}
		//	respFromClient =append(respFromClient,buffer[:n]...)
		//	if n<bufferSize{
		//		break
		//	}
		//}
		//t.Log(respFromClient)

		serverTcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
		if err != nil {
			t.Fatal(err)
		}
		serverTcpConn, err := net.DialTCP("tcp", nil, serverTcpAddr)
		if err != nil {
			t.Fatal(err)
		}
		//_, err = serverTcpConn.Write(respFromClient)
		//if err != nil {
		//	t.Fatal(err)
		//}
		//respFromServer:=make([]byte,0)
		//for{
		//	var buffer [bufferSize]byte
		//	n, err := serverTcpConn.Read(buffer[:])
		//	if err != nil {
		//		t.Fatal(err)
		//	}
		//	respFromServer=append(respFromServer,buffer[:n]...)
		//	if n < bufferSize{
		//		break
		//	}
		//}
		//t.Log(respFromServer)
		//
		//_, err = tcpConn.Write(respFromServer)
		//if err != nil {
		//	t.Fatal(err)
		//}
		go func() {
			_, err := io.Copy(serverTcpConn, tcpConn)
			if err != nil {
				t.Error(err)
				return
			}
		}()

		go func() {
			_, err := io.Copy(tcpConn, serverTcpConn)
			if err != nil {
				t.Error(err)
				return
			}
		}()
	}
}
