package models

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/set"
	"github.com/gorilla/websocket"
	"net"
	"net/http"
	"strconv"
	"sync"
)

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

var clientMap map[uint]*Node = make(map[uint]*Node, 0)

var rwLocker sync.RWMutex

func Chat(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()

	//获取参数
	//token:=query.Get("token")
	senderId := query.Get("senderId") //发送者
	//msgType := query.Get("type")      //发送类型
	//receverId := query.Get("receverId")
	//content := query.Get("content") //消息内容
	//media := query.Get("media")     //消息类型

	//校验合法性
	isvalid := true //校验

	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return isvalid
		},
	}).Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	//获取连接
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}

	//用户关系

	//绑定senderId与node，并加锁
	rwLocker.Lock()
	senderIdInt, _ := strconv.Atoi(senderId)
	clientMap[uint(senderIdInt)] = node
	rwLocker.Unlock()

	//发送逻辑
	go sendProc(node)
	//接收逻辑
	go recvProc(node)

	sendMesg(uint(senderIdInt), []byte("welcome to chat room!"))
}

func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func recvProc(node *Node) {
	for {
		mesgType, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		broadMesg(data)
		fmt.Println("message type: ", mesgType)
		fmt.Println("[ws] <<<<<", data)
	}
}

var udpSenderChan chan []byte = make(chan []byte, 1024)

func broadMesg(data []byte) {
	udpSenderChan <- data
}

func init() {
	go udpSendProc()
	go udpRecvProc()
}

func udpSendProc() {
	udpConn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(192, 168, 211, 255), //!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
		Port: 3000,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(udpConn *net.UDPConn) {
		err := udpConn.Close()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}(udpConn)

	for {
		select {
		case data := <-udpSenderChan:
			_, err := udpConn.Write(data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func udpRecvProc() {
	udpConn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3000,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(udpConn *net.UDPConn) {
		err := udpConn.Close()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}(udpConn)

	for {
		var buf [512]byte
		n, err := udpConn.Read(buf[:])
		if err != nil {
			fmt.Println(err)
			return
		}
		dispatch(buf[:n])
	}
}

func dispatch(data []byte) {
	mesg := Message{}
	err := json.Unmarshal(data, &mesg)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch mesg.Type {
	case 1: //私聊
		sendMesg(mesg.ReceverId, data)
		//case 2: //群发
		//
		//case 3: //广播
		//
		//case 4: //其他

	}
}

func sendMesg(receverId uint, mesg []byte) {
	rwLocker.Lock()
	node, ok := clientMap[receverId]
	rwLocker.Unlock()
	if ok {
		node.DataQueue <- mesg
	}
}
