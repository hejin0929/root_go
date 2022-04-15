package web_socket

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net"
	"net/http"
	"time"
)

// socket 配置
var web = websocket.Upgrader{
	//HandshakeTimeout: 1024,
	//ReadBufferSize:   1024,
	//WriteBufferSize:  1024,
	//WriteBufferPool:  nil,
	//Subprotocols:     nil,
	//Error:            nil,
	CheckOrigin: func(r *http.Request) bool {
		if r.Method != "GET" {
			r.Response.StatusCode = 400
			r.Response.Status = "method is not GET"
			r.Header.Set("Content-Type", "application/json;charset=utf-8")
			res, _ := json.Marshal(r.Response)

			var w http.ResponseWriter

			_, err := w.Write(res)
			if err != nil {
				return false
			}

			return false
		}

		return true
	},
	//EnableCompression: false,
}

type Client struct {
	Port          string          // 客户端地址
	Upgrade       *websocket.Conn // 用户连接
	Send          []byte          // 待发送数据
	APPId         uint            // 登陆平台
	UserId        string          // 用户ID
	FirstTime     uint64          // 首次连接时间
	HeartbeatTime uint64          // 上次心跳
	LoginTime     uint64          // 登陆时间
	Read          []byte          // 读取数据
	Status        int             // 状态
	Addr          string          // 开启ws的端口
	Listener      net.Listener    // net 连接

}

func WebSocketNews(g *gin.Context) (*Client, error) {

	ws := new(Client)

	ws.Addr = "0.0.0.0:10215"

	var err error

	// 待写逻辑
	ws.Upgrade, err = web.Upgrade(g.Writer, g.Request, nil)

	if err != nil {

		err = ws.Upgrade.WriteMessage(http.StatusNotFound, []byte("服务端创建socket失败"))
		if err != nil {
			return nil, err
		}
		return nil, err
	}
	//
	//for {
	//	//读取ws中的数据
	//	mt, message, err := ws.Upgrade.ReadMessage()
	//
	//	fmt.Println("this is data ", mt, message)
	//
	//	if err != nil {
	//		break
	//	}
	//}

	fmt.Println("this is a ?? ")

	ws.SendHeartbeat()

	return ws, nil
}

// Start 启动websocket
func (_this *Client) Start() error {
	_this.FirstTime = uint64(time.Now().Unix())
	var err error = nil
	//_this.Listener, err = net.Listen("tcp", _this.Addr)

	//var test *WsServer

	//http.Serve(_this.Listener, _this)

	return err
}

func (_this *Client) SendHeartbeat() {
	select {
	case <-time.After(time.Second * 10):
		_this.FirstTime = uint64(time.Now().Unix())
		fmt.Println("this is a send")
		var res = struct {
			Heartbeat uint64 `json:"heartbeat"`
			Message   string `json:"message"`
		}{
			Heartbeat: _this.HeartbeatTime,
			Message:   "❤️",
		}

		bytes, _ := json.Marshal(res)

		_this.Upgrade.WriteMessage(websocket.TextMessage, bytes)
	}
}

// Sends 发送数据
func (_this *Client) Sends() {
	err := _this.Upgrade.WriteMessage(websocket.TextMessage, _this.Send)

	if err != nil {
		fmt.Println("this is sends ?? ", err)
		err = _this.Upgrade.Close()
		return
	}
}

// Reads 读取数据
func (_this *Client) Reads() {
	message, read, err := _this.Upgrade.ReadMessage()

	_this.Status = message

	_this.Read = read

	if err != nil {
		return
	}
}

// Close 关闭socket服务
func (_this *Client) Close() {
	err := _this.Upgrade.Close()
	if err != nil {
		return
	}
}
