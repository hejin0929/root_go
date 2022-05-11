package chum

import (
	"modTest/module"
	"time"
)

const (
	SourceSearch = iota // 搜索ID
	SourceMobile        // 手机号
	SourceShare         // 分享
	SourceGroup         // 群组

)

const (
	LOADING = iota // 未处理
	DELETE         // 拒绝
	CONSENT        // 同意
	OVERDUE        // 过期
)

type ApplyChumSql struct {
	module.Model
	Uuid        string     `json:"uuid"`        // 用户ID
	FriendID    string     `json:"friend_id"`   // 好友ID
	Note        string     `json:"note"`        // 备注
	Top         *time.Time `json:"top"`         // 该用户是否置顶
	Blacklist   *time.Time `json:"blacklist"`   // 黑名单
	Star        string     `json:"star"`        // 星标
	Source      int        `json:"source"`      // 来源
	Permissions string     `json:"permissions"` // 权限
	Hello       string     `json:"hello"`       // 招呼语
	Status      int        `json:"status"`      // 状态
}

type UserChumSql struct {
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

type FocusChumSql struct {
	module.Model
	Uuid     string `json:"uuid"`      // 用户ID
	FriendID string `json:"friend_id"` // 好友ID
	Star     string `json:"star"`      // 星标
}
