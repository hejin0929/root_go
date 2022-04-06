package web_socket

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

// socket 配置
var web = websocket.Upgrader{
	HandshakeTimeout: 1024,
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	WriteBufferPool:  nil,
	Subprotocols:     nil,
	Error:            nil,
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
	EnableCompression: false,
}

type Client struct {
	Port          string          // 客户端地址
	Socket        *websocket.Conn // 用户连接
	Send          []byte          // 待发送数据
	APPId         uint            // 登陆平台
	UserId        string          // 用户ID
	FirstTime     uint64          // 首次连接时间
	HeartbeatTime uint64          // 上次心跳
	LoginTime     uint64          // 登陆时间
	Read          []byte          // 读取数据
}

func WebSocketNews(g *gin.Context) (*Client, error) {

	var ws Client
	var err error
	ws.Socket, err = web.Upgrade(g.Writer, g.Request, nil)

	if err != nil {
		err = ws.Socket.WriteMessage(200, []byte("服务端创建socket失败"))
		err = ws.Socket.Close()
		if err != nil {
			return nil, err
		}
		return nil, err
	}
	return &ws, nil
}

// Start 启动websocket
func (_this *Client) Start() {
	_this.FirstTime = uint64(time.Now().Unix())

}

// Sends 发送数据
func (_this *Client) Sends() {
	err := _this.Socket.WriteMessage(http.StatusOK, _this.Send)
	if err != nil {
		err = _this.Socket.Close()
		return
	}
}

// Reads 读取数据
func (_this *Client) Reads() {
	message, read, err := _this.Socket.ReadMessage()

	_this.Read = read

	fmt.Println("this is a ?? ", message, read)

	if err != nil {
		return
	}
}

