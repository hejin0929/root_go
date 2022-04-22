package web_socket

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	redis2 "modTest/service/redis"
	"modTest/utlis/token"
	"net"
	"net/http"
	"time"
)

// socket 配置
var web = websocket.Upgrader{
	//HandshakeTimeout: 1024,
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	WriteBufferPool: nil,
	Subprotocols:    nil,
	Error:           nil,
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
	Redis         *redis.Client   // Redis 数据交互
}

// MessageChannel 消息交互
type MessageChannel struct {
	ReceiveId string `json:"receive_id"`
	Message   string `json:"message"`
	SendId    string `json:"send_id"`
}

func WebSocketNews(g *gin.Context) {

	ws := new(Client)

	ws.Addr = "0.0.0.0:10215"

	var err error

	// 待写逻辑
	ws.Upgrade, err = web.Upgrade(g.Writer, g.Request, nil)

	if err != nil {

		err = ws.Upgrade.WriteMessage(http.StatusNotFound, []byte("服务端创建socket失败"))
		if err != nil {
			return
		}
		return
	}
	ws.FirstTime = uint64(time.Now().Unix())
	ws.Redis = redis2.CreateRedis(2)
	go ws.ReadMessage()
	go ws.SendHeartbeat()

	return
}

func (_this *Client) SendHeartbeat() {
	times := time.NewTicker(time.Second * 3)
	for {
		select {
		case <-times.C:
			{
				_this.HeartbeatTime = uint64(time.Now().Unix())

				var res = struct {
					Heartbeat uint64 `json:"heartbeat"`
					Message   string `json:"message"`
				}{
					Heartbeat: _this.HeartbeatTime,
					Message:   "❤️",
				}

				bytes, _ := json.Marshal(res)

				_ = _this.Upgrade.WriteMessage(websocket.TextMessage, bytes)
			}

		}

	}
}

func (_this *Client) ReadMessage() error {

	for {

		_, message, err := _this.Upgrade.ReadMessage()

		if err != nil {
			err := _this.Upgrade.WriteMessage(websocket.CloseMessage, []byte(err.Error()))
			if err != nil {
				return err
			}
		}

		if _this.UserId == "" {

			var data = struct {
				UserID string `json:"user_id"`
				Token  string `json:"token"`
			}{}

			_ = json.Unmarshal(message, &data)

			if data.UserID == "" {
				_this.Send = []byte("缺少user_id !")
				_this.Sends()
				_ = _this.Upgrade.Close()
				break
			}

			claim, err := token.VerifyAction(data.Token)

			if err != nil {
				_this.Send = []byte(err.Error())
				_this.Sends()
				_ = _this.Upgrade.Close()
				break
			}

			if time.Now().Unix() < claim.StandardClaims.IssuedAt { // IssuedAt token过期时间
				_this.Send = []byte("token时间过期 !")
				_this.Sends()
				_ = _this.Upgrade.Close()
				break
			}

			_this.UserId = data.UserID
			_this.Send = []byte("Success")
			_this.Sends()
		}

		var channel MessageChannel
		channel.SendId = _this.UserId

		_ = json.Unmarshal(message, &channel)

		if channel.ReceiveId != "" {
			var messageMap []MessageChannel

			data := _this.Redis.Get(context.Background(), channel.ReceiveId)
			if data != nil {
				json.Unmarshal([]byte(data.Val()), &messageMap)
			}

			messageMap = append(messageMap, channel)

			bytes, _ := json.Marshal(messageMap)

			_this.Redis.Set(context.Background(), channel.ReceiveId, string(bytes), 0)
		}
	}

	return nil
}

// Sends 发送数据
func (_this *Client) Sends() {

	err := _this.Upgrade.WriteMessage(websocket.TextMessage, _this.Send)

	if err != nil {
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
