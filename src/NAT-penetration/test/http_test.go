package test

import (
	"context"
	"encoding/json"
	"go_project/src/NAT-penetration/define"
	"io"
	"log"
	"net"
	"net/http"
	"sync"
	"testing"
)

func TestHttpRet(t *testing.T) {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		q := request.URL.Query()
		marshal, err := json.Marshal(q)
		if err != nil {
			panic(err)
		}
		_, err = writer.Write(marshal)
		if err != nil {
			panic(err)
		}
	})
	err := http.ListenAndServe("0.0.0.0:9090", nil)
	if err != nil {
		panic(err)
	}
}

func Nat(ch chan int) {
	for {
		select {
		case <-ch:
			tunnelTcpConn, err := define.CreateConn("0.0.0.0:9091")
			if err != nil {
				panic(err)
			}
			log.Println("create tcpconn 0.0.0.0:9091")

			serverTcpConn, err := define.CreateConn("0.0.0.0:9090")
			if err != nil {
				panic(err)
			}
			log.Println("create tcpconn 0.0.0.0:9090")

			endCh := make(chan int)
			go func(endch chan int) {
				_, err := io.Copy(serverTcpConn, tunnelTcpConn)
				if err != nil {
					log.Println("tunnelNAT tcp conn done!")
				}
				log.Println("TUNNELNAT tcp conn done!")
				endch <- 1
				log.Println("cancel DDDDDDDDDDDDDDone!")
			}(endCh)
			//go func() {
			//	_, err := io.Copy(tunnelTcpConn, serverTcpConn)
			//	if err != nil {
			//		log.Println("server tcp conn done!")
			//	}
			//	log.Println("SERVER tcp conn done!")
			//}()
			go func(endch chan int) {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()
				for {
					select {
					case <-ctx.Done():
						//将serverTcpConn放到连接池里
						return
					default:
						resp, err := define.GetDataFromTcp(1024, serverTcpConn)
						if err != nil {
							panic(err)
						}
						_, err = tunnelTcpConn.Write(resp)
						if err != nil {
							panic(err)
						}
					}
				}
			}(endCh)
		}
	}
}

func NewSer(ch chan int) {
	listenerUser, err := define.CreateListener("0.0.0.0:9092")
	if err != nil {
		panic(err)
	}

	log.Printf("用户请求监听中: %v\n", listenerUser.Addr().String())

	listenerTun, err := define.CreateListener("0.0.0.0:9091")
	if err != nil {
		panic(err)
	}

	log.Printf("隧道监听中: %v\n", listenerTun.Addr().String())

	for {
		userTcpConn, err := listenerUser.AcceptTCP()
		if err != nil {
			log.Println(err)
			panic(err)
		}
		ch <- 1

		tunConn, err := listenerTun.AcceptTCP()
		if err != nil {
			panic(err)
		}
		go func() {
			_, err := io.Copy(userTcpConn, tunConn)
			if err != nil {
				log.Println("tunnel tcp conn done!")
			}
			log.Println("TUNNEL tcp conn done!")
		}()
		go func() {
			defer func(tunConn *net.TCPConn) {
				err := tunConn.Close()
				if err != nil {
					log.Println("b")
				}
			}(tunConn)
			_, err := io.Copy(tunConn, userTcpConn)
			if err != nil {
				log.Println("user tcp conn done!")
			}
			log.Println("USER tcp conn done!")
		}()
	}
}

func TestAll(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan int)
	go Nat(ch)
	go NewSer(ch)
	wg.Wait()
}
