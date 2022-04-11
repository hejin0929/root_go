package web_socket

import (
	"encoding/json"
	"errors"
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
	UserId        string `json:"user_id"` // 用户ID
	FirstTime     uint64 // 首次连接时间
	HeartbeatTime uint64 // 上次心跳
	LoginTime     uint64 // 登陆时间
	Read          []byte // 读取数据
	Status        int    // 状态
	Addr          string // 开启ws的端口
	Port          string // 客户端地址
	Token         string `json:"token"` // token
}

func NewWsServer() *WsServer {
	ws := new(WsServer)
	ws.Addr = "0.0.0.0:10215"
	ws.Status = 1
	ws.Send = []byte("一次握手")

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
		http.Error(w, errText, httpCode)
		return
	}

	conn, err := _this.Upgrade.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("websocket error:", err)
		return
	}

	go _this.connHandle(conn)

}

func (_this *WsServer) connHandle(conn *websocket.Conn) {

	defer func() {
		err := conn.Close()
		if err != nil {
			return
		}
	}()

	stopCh := make(chan int)

	go _this.send(conn, stopCh)

	for {
		err := conn.SetReadDeadline(time.Now().Add(time.Millisecond * time.Duration(5000)))
		if err != nil {
			return
		}
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

		if _this.UserId == "" {
			err := _this.login(msg)

			if err != nil {
				_this.Send = []byte(err.Error())
				_this.Status = 3

			} else {

				var res = struct {
					Success string `json:"success"`
				}{
					"登陆成功",
				}

				_this.Send, _ = json.Marshal(res)
			}
		}
	}
}

func (_this *WsServer) login(data []byte) error {
	err := json.Unmarshal(data, _this)
	_this.LoginTime = uint64(time.Now().Unix())

	if _this.UserId == "" {
		return errors.New("userID 缺少")
	}

	if _this.Token == "" {
		return errors.New("token 缺少")
	}

	if err != nil {
		return err
	}

	return nil
}

//测试一次性发送 10万条数据给 client, 如果不使用 time.Sleep browser 过了超时时间会断开
func (_this *WsServer) send10(conn *websocket.Conn) {
	for i := 0; i < 100000; i++ {
		//data := fmt.Sprintf("hello websocket test from server %v", time.Now().UnixNano())
		err := conn.WriteMessage(_this.Status, _this.Send)
		if err != nil {
			fmt.Println("send msg ", err)
			return
		}
		// time.Sleep(time.Millisecond * 1)
	}
}

func (_this *WsServer) send(conn *websocket.Conn, stopCh chan int) {
	//_this.send10(conn)
	for {
		select {
		case <-stopCh:
			//fmt.Println("connect closed", stopCh)
			err := conn.WriteMessage(_this.Status, []byte("close socket ！！！"))
			_ = conn.Close()
			if err != nil {
				fmt.Println("send msg ", err)
				return
			}
			return
		case <-time.After(time.Second * 1):
			if _this.Status == 3 {
				_ = conn.WriteMessage(1, _this.Send)
				_ = conn.Close()
				return
			}
			err := conn.WriteMessage(_this.Status, _this.Send)

			if err != nil {
				fmt.Println("send msg ", err)
				return
			}
		}
	}
}

func (_this *WsServer) Start() (err error) {

	//fmt.Println("this is a value ?? ", _this.Listener)
	//
	//_this.Listener, err = net.Listen("tcp", _this.Addr)
	//
	//if err != nil {
	//	return
	//}

	err = http.ListenAndServe(_this.Addr, _this)

	_this.FirstTime = uint64(time.Now().Unix())

	if err != nil {
		fmt.Println("http serve error:", err)
		return
	}

	return nil
}
