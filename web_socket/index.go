package web_socket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net"
	"net/http"
	"time"
)

type WsServer struct {
	Listener      net.Listener
	Upgrade       *websocket.Upgrader
	Send          []byte // 待发送数据
	APPId         uint   // 登陆平台
	UserId        string // 用户ID
	FirstTime     uint64 // 首次连接时间
	HeartbeatTime uint64 // 上次心跳
	LoginTime     uint64 // 登陆时间
	Read          []byte // 读取数据
	Status        int    // 状态
	Addr          string // 开启ws的端口
	Port          string // 客户端地址
}

func NewWsServer() *WsServer {
	ws := new(WsServer)
	ws.Addr = "0.0.0.0:10215"
	ws.Upgrade = &websocket.Upgrader{
		HandshakeTimeout: 4096,
		ReadBufferSize:   4096,
		WriteBufferSize:  1024,
		WriteBufferPool:  nil,
		Subprotocols:     nil,
		Error:            nil,
		CheckOrigin: func(r *http.Request) bool {
			if r.Method != "GET" {
				fmt.Println("method is not GET")
				return false
			}
			if r.URL.Path != "/ws" {
				fmt.Println("path error")
				return false
			}
			return true
		},
		EnableCompression: false,
	}
	return ws
}

func (_this *WsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/ws" {
		httpCode := http.StatusInternalServerError
		errText := http.StatusText(httpCode)
		//fmt.Println("path error ", errText)
		http.Error(w, errText, httpCode)
		return
	}

	conn, err := _this.Upgrade.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("websocket error:", err)
		return
	}
	fmt.Println("client connect :", conn.RemoteAddr())

	go _this.connHandle(conn)

}

func (_this *WsServer) connHandle(conn *websocket.Conn) {
	defer func() {
		conn.Close()
	}()
	stopCh := make(chan int)
	go _this.send(conn, stopCh)
	for {
		conn.SetReadDeadline(time.Now().Add(time.Millisecond * time.Duration(5000)))
		_, msg, err := conn.ReadMessage()
		if err != nil {
			close(stopCh)
			// 判断是不是超时
			if netErr, ok := err.(net.Error); ok {
				if netErr.Timeout() {
					fmt.Printf("ReadMessage timeout remote: %v\n", conn.RemoteAddr())
					return
				}
			}
			// 其他错误，如果是 1001 和 1000 就不打印日志
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				fmt.Printf("ReadMessage other remote:%v error: %v \n", conn.RemoteAddr(), err)
			}
			return
		}
		fmt.Println("收到消息：", string(msg))
	}
}

//测试一次性发送 10万条数据给 client, 如果不使用 time.Sleep browser 过了超时时间会断开
func (_this *WsServer) send10(conn *websocket.Conn) {
	for i := 0; i < 100000; i++ {
		data := fmt.Sprintf("hello websocket test from server %v", time.Now().UnixNano())
		err := conn.WriteMessage(1, []byte(data))
		if err != nil {
			fmt.Println("send msg faild ", err)
			return
		}
		// time.Sleep(time.Millisecond * 1)
	}
}

func (_this *WsServer) send(conn *websocket.Conn, stopCh chan int) {
	_this.send10(conn)
	for {
		select {
		case <-stopCh:
			fmt.Println("connect closed")
			return
		case <-time.After(time.Second * 1):
			data := fmt.Sprintf("hello websocket test from server %v", time.Now().UnixNano())
			err := conn.WriteMessage(1, []byte(data))
			fmt.Println("sending....")
			if err != nil {
				fmt.Println("send msg faild ", err)
				return
			}
		}
	}
}

func (_this *WsServer) Start() (err error) {
	_this.Listener, err = net.Listen("tcp", _this.Addr)
	if err != nil {
		fmt.Println("net listen error:", err)
		return
	}
	err = http.Serve(_this.Listener, _this)
	if err != nil {
		fmt.Println("http serve error:", err)
		return
	}
	return nil
}
