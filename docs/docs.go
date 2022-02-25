// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login/user/login": {
            "post": {
                "tags": [
                    "Login"
                ],
                "summary": "用户登录操作",
                "operationId": "LoginsUserPassword",
                "parameters": [
                    {
                        "description": "JSON数据",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/login.UserName"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JSON数据",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/login.UserBody"
                            }
                        }
                    }
                }
            }
        },
        "/login/user/login_code": {
            "post": {
                "tags": [
                    "Login"
                ],
                "summary": "用户登录操作",
                "operationId": "LoginsUserCode",
                "parameters": [
                    {
                        "description": "JSON数据",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/login.UserCode"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JSON数据",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/login.UserBody"
                            }
                        }
                    }
                }
            }
        },
        "/login/user/sign": {
            "post": {
                "tags": [
                    "Login"
                ],
                "summary": "用户登录操作",
                "operationId": "SignUser",
                "parameters": [
                    {
                        "description": "JSON数据",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/login.UserSignType"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JSON数据",
                        "schema": {
                            "$ref": "#/definitions/login.UserSignReps"
                        }
                    }
                }
            }
        },
        "/login/user/{phone}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Login"
                ],
                "summary": "用户验证码登录",
                "operationId": "GetSingCode",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户手机号",
                        "name": "phone",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JSON数据",
                        "schema": {
                            "$ref": "#/definitions/login.RepsGetCode"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "login.GetPhoneCode": {
            "type": "object",
            "required": [
                "code"
            ],
            "properties": {
                "code": {
                    "type": "string"
                }
            }
        },
        "login.RepsGetCode": {
            "type": "object",
            "required": [
                "mgsCode",
                "mgsText"
            ],
            "properties": {
                "body": {
                    "description": "in Body",
                    "$ref": "#/definitions/login.GetPhoneCode"
                },
                "mgsCode": {
                    "description": "返回体状态码",
                    "type": "integer"
                },
                "mgsText": {
                    "description": "返回体信息",
                    "type": "string"
                }
            }
        },
        "login.User": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "login.UserBody": {
            "type": "object",
            "required": [
                "mgsCode",
                "mgsText"
            ],
            "properties": {
                "body": {
                    "$ref": "#/definitions/login.User"
                },
                "mgsCode": {
                    "description": "返回体状态码",
                    "type": "integer"
                },
                "mgsText": {
                    "description": "返回体信息",
                    "type": "string"
                }
            }
        },
        "login.UserCode": {
            "type": "object",
            "required": [
                "code",
                "phone"
            ],
            "properties": {
                "code": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "login.UserName": {
            "type": "object",
            "required": [
                "password",
                "phone"
            ],
            "properties": {
                "password": {
                    "description": "用户密码",
                    "type": "string"
                },
                "phone": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "login.UserSignReps": {
            "type": "object",
            "required": [
                "mgsCode",
                "mgsText"
            ],
            "properties": {
                "body": {
                    "type": "string"
                },
                "mgsCode": {
                    "description": "返回体状态码",
                    "type": "integer"
                },
                "mgsText": {
                    "description": "返回体信息",
                    "type": "string"
                }
            }
        },
        "login.UserSignType": {
            "type": "object",
            "required": [
                "code",
                "password",
                "phone"
            ],
            "properties": {
                "code": {
                    "description": "手机验证码",
                    "type": "string"
                },
                "password": {
                    "description": "用户密码",
                    "type": "string"
                },
                "phone": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
