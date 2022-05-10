package chum

import (
	"modTest/module"
	"modTest/module/user"
	user2 "modTest/types/user"
	"time"
)

const (
	SourceSearch = iota // 搜索ID
	SourceMobile        // 手机号
	SourceShare         // 分享
	SourceGroup         // 群组

)

type AddReq struct {
	Uuid        string `json:"uuid"`        // 本人ID
	FriendID    string `json:"friend_id"`   // 朋友ID
	Source      int    `json:"source"`      // 来源
	Permissions string `json:"permissions"` // 权限
}

type AddRes struct {
	module.Resp
	Body string `json:"body"`
}

type ResChum struct {
	module.Resp
	Body struct {
		User   user.Message `json:"user"`   // 用户信息
		Source int          `json:"source"` // 搜索方式
	} `json:"body"`
}

type UserChum struct {
	module.Model
	Uuid        string     `json:"uuid"`        // 用户ID
	FriendID    string     `json:"friend_id"`   // 好友ID
	Note        string     `json:"note"`        // 备注
	Top         *time.Time `json:"top"`         // 该用户是否置顶
	Blacklist   *time.Time `json:"blacklist"`   // 黑名单
	Star        string     `json:"star"`        // 星标
	Source      int        `json:"source"`      // 来源
	Permissions string     `json:"permissions"` // 权限
}

type SearchUserRes struct {
	module.Resp
	Body user2.Message `json:"body"`
}
