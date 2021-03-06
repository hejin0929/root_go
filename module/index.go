package module

import (
	"net/http"
	"time"
)

// A Resp 项目统一返回的结构
// swagger:response Resp
type Resp struct {
	// 返回体状态码
	MgsCode int `json:"mgsCode" binding:"required"`
	// 返回体信息
	MgsText string `json:"mgsText" binding:"required"`
}

type ResponsesData struct {
	MgsCode int `json:"mgs_code" binding:"required"` //
}

type HttpErrs struct {
	MsgCode int    `json:"msgCode" binding:"required"` // 状态码
	MgsText string `json:"mgsText" binding:"required"` // 报错信息
}

// Model gorm.Model 定义
type Model struct {
	ID        int        `gorm:"primary_key"` // 字段名为 ID 的字段默认作为表的主键
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Response struct {
	// 返回体状态码
	MgsCode int `json:"mgsCode" binding:"required"`
	// 返回体信息
	MgsText string `json:"mgsText" binding:"required"`
	// body
	Body interface{} `json:"body"`
}

type ResponseBodyInString struct {
	Resp
	Body string `json:"body"`
}

// ResponseSuccess 统一返回类型
func ResponseSuccess(data interface{}) Response {
	return Response{
		Body:    data,
		MgsCode: http.StatusOK,
		MgsText: "Success",
	}
}

// ResponseErrorParams 统一参数报错返回
func ResponseErrorParams(err interface{}) Response {
	return Response{
		Body:    err,
		MgsCode: http.StatusBadRequest,
		MgsText: "Params Not",
	}
}

func ResponseServerError(err interface{}) Response {
	return Response{
		Body:    err,
		MgsCode: http.StatusInternalServerError,
		MgsText: "Fail",
	}
}
