package web_socket

type UserSuccess struct {
	LoginTime string `json:"login_time"`
	MsgText   string `json:"msg_text"`
	MsgCode   uint   `json:"msg_code"`
}

type Heartbeat struct {
	Time uint64 `json:"time"`
	Icon string `json:"icon"`
}

type UserMessage struct {
	UserId  string `json:"user_id"`
	Message string `json:"message"`
	Time    uint   `json:"time"`
	Type    string `json:"type"`
}
