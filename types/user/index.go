package user

// 接口返回的格式定义

type Message struct {
	Name      string `json:"name"`      // 用户名称
	Uuid      string `json:"uuid"`      // 用户uuid
	Introduce string `json:"introduce"` // 介绍
	Image     string `json:"image"`     // 头像
	UserID    string `json:"user_id"`   // 用户设置ID
	Phone     string `json:"phone"`     // 用户手机
	State     int    `json:"state"`     // 状态 0 离线 1 在线
	Sex       int    `json:"sex"`       // 性别
	Region    string `json:"region"`    // 地区
	Birthday  string `json:"birthday"`  // 生日
}

type MessageUpdate struct {
	Message Message `json:"message"`
	Res     string  `json:"res"`
}
