package center

import "modTest/module"

// Message sql 数据表结构
type Message struct {
	module.Model
	Name      string `json:"name"`      // 用户名称
	Uuid      string `json:"uuid"`      // 用户uuid
	Introduce string `json:"introduce"` // 介绍
	Image     string `json:"image"`     // 头像
	UserID    string `json:"user_id"`   // 用户设置ID
	Phone     string `json:"phone"`     // 用户手机
	State     int    `json:"state"`     // 状态 0 离线 1 在线
}
