package chum

import (
	"modTest/module"
	"time"
)

type AddReq struct {
	Uuid        string `json:"uuid"`        // 本人ID
	FriendID    string `json:"friend_id"`   // 朋友ID
	Source      string `json:"source"`      // 来源
	Permissions string `json:"permissions"` // 权限
}

type AddRes struct {
	module.Resp
	Body string `json:"body"`
}

type ResChum struct {
	Name string `json:"name" binding:"required"`
}

type UserChum struct {
	module.Model
	Uuid        string     `json:"uuid"`        // 用户ID
	FriendID    string     `json:"friend_id"`   // 好友ID
	Note        string     `json:"note"`        // 备注
	Top         *time.Time `json:"top"`         // 该用户是否置顶
	Blacklist   *time.Time `json:"blacklist"`   // 黑名单
	Star        string     `json:"star"`        // 星标
	Source      string     `json:"source"`      // 来源
	Permissions string     `json:"permissions"` // 权限
}

type SearchUserRes struct {
}