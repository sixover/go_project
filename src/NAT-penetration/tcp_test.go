package NAT_penetration

import (
	"net"
	"testing"
)

const (
	addr       = "0.0.0.0:22222"
	bufferSize = 10
)

//监听
func TestTcpListen(t *testing.T) {
	//拿到一个真实的地址
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tcpAddr)

	//监听端口
	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		t.Fatal(err)
	}

	for {
		//从端口监听器获取到一个tcp连接
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			t.Fatal(err)
		}

		/*
			下面这段读数据的逻辑是：
			从一个tcp连接中读取数据
			一次性读取bufferSize字节的byte
			如果此时read<bufferSize，证明发送的数据读取完了，就break循环
			如果此时read==bufferSize，证明发送的数据还没有被读取完，那就继续循环读取
		*/
		for {
			//读取数据
			var buffer [bufferSize]byte
			read, err := tcpConn.Read(buffer[:])
			if err != nil {
				t.Fatal(err)
			}
			//read是读取到的数据的长度
			t.Log(string(buffer[:read]))

			if read < bufferSize {
				break
			}
		}

		//发送数据
		_, err = tcpConn.Write([]byte("hello world!"))
		if err != nil {
			t.Fatal(err)
		}
	}
}

//创建连接
func TestCreateTcp(t *testing.T) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}

	//创建一个tcp连接
	tcpConn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		t.Fatal(err)
	}

	//写数据
	_, err = tcpConn.Write([]byte("client ==> hello world!"))
	if err != nil {
		t.Fatal(err)
	}

	for {
		//读数据
		var buffer [bufferSize]byte
		read, err := tcpConn.Read(buffer[:])
		if err != nil {
			t.Fatal(err)
		}
		t.Log(string(buffer[:read]))

		if read < bufferSize {
			break
		}
	}
}
