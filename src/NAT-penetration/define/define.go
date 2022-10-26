package define

import (
	"net"
	"time"
)

const (
	NewConnection   = "New Connection\n"
	KeepAlive       = "keep alive\n"
	LocalHostAddr   = "192.168.101.119:9090"
	ControloerAddr  = "120.46.203.178:8081"
	TunnelAddr      = "120.46.203.178:8082"
	UserRequestAddr = "120.46.203.178:8083"
)

func CreateListener(addr string) (*net.TCPListener, error) {
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

func CreateConn(addr string) (*net.TCPConn, error) {
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

func KeepAliveService(conn *net.TCPConn) {
	for {
		_, err := conn.Write([]byte(KeepAlive))
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second * 3)
	}
}
func GetDataFromTcp(bufferSize int, tcpConn *net.TCPConn) ([]byte, error) {
	respFromServer := make([]byte, 0)
	for {
		buffer := make([]byte, bufferSize)
		n, err := tcpConn.Read(buffer)
		if err != nil {
			return nil, err
		}
		respFromServer = append(respFromServer, buffer[:n]...)
		if n < bufferSize {
			break
		}
	}
	return respFromServer, nil
}
