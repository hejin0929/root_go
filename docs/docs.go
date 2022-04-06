// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
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
        "/api/home/message": {
            "get": {
                "description": "get string by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Home"
                ],
                "summary": "首页接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/module.HttpErrs"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/module.HttpErrs"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/module.HttpErrs"
                        }
                    }
                }
            }
        },
        "/api/login/user/login": {
            "post": {
                "tags": [
                    "Login"
                ],
                "summary": "用户登录操作",
                "operationId": "LoginInPasswordPaths",
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
        "/api/login/user/login_code": {
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
        "/api/login/user/sign": {
            "post": {
                "tags": [
                    "Login"
                ],
                "summary": "用户注册操作",
                "operationId": "LoginNewsUserPaths",
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
        "/api/phone_code/user/{phone}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Code"
                ],
                "summary": "用户验证码登录",
                "operationId": "GetPathsCode",
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
                            "$ref": "#/definitions/getCode.ResCode"
                        }
                    }
                }
            }
        },
        "/api/upload/deleteImg": {
            "get": {
                "tags": [
                    "Upload"
                ],
                "summary": "删除图片操作",
                "operationId": "UploadImagesDelete",
                "parameters": [
                    {
                        "type": "string",
                        "description": "图片名称",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JSON数据",
                        "schema": {
                            "$ref": "#/definitions/uploadFiles.FilesRes"
                        }
                    }
                }
            }
        },
        "/api/upload/deleteVideo": {
            "get": {
                "tags": [
                    "Upload"
                ],
                "summary": "删除图片操作",
                "operationId": "UploadVideoDelete",
                "parameters": [
                    {
                        "type": "string",
                        "description": "视频名称",
                        "name": "video",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JSON数据",
                        "schema": {
                            "$ref": "#/definitions/uploadFiles.FilesRes"
                        }
                    }
                }
            }
        },
        "/api/upload/images": {
            "post": {
                "tags": [
                    "Upload"
                ],
                "summary": "上传图片操作",
                "operationId": "UploadImages",
                "parameters": [
                    {
                        "type": "file",
                        "description": "file",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JSON数据",
                        "schema": {
                            "$ref": "#/definitions/uploadFiles.FilesRes"
                        }
                    }
                }
            }
        },
        "/api/upload/test": {
            "post": {
                "tags": [
                    "Upload"
                ],
                "summary": "删除图片操作",
                "operationId": "UploadTestPaths",
                "parameters": [
                    {
                        "description": "视频名称",
                        "name": "video",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/uploadFiles.FileTest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JSON数据",
                        "schema": {
                            "$ref": "#/definitions/uploadFiles.FilesRes"
                        }
                    }
                }
            }
        },
        "/api/upload/videos": {
            "post": {
                "tags": [
                    "Upload"
                ],
                "summary": "上传视频操作",
                "operationId": "UploadVideos",
                "parameters": [
                    {
                        "type": "file",
                        "description": "file",
                        "name": "videos",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JSON数据",
                        "schema": {
                            "$ref": "#/definitions/uploadFiles.FilesRes"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "getCode.GetPhoneCode": {
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
        "getCode.ResCode": {
            "type": "object",
            "required": [
                "mgsCode",
                "mgsText"
            ],
            "properties": {
                "body": {
                    "$ref": "#/definitions/getCode.GetPhoneCode"
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
        },
        "module.HttpErrs": {
            "type": "object",
            "required": [
                "mgsText",
                "msgCode"
            ],
            "properties": {
                "mgsText": {
                    "description": "报错信息",
                    "type": "string"
                },
                "msgCode": {
                    "description": "状态码",
                    "type": "integer"
                }
            }
        },
        "uploadFiles.FileTest": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                }
            }
        },
        "uploadFiles.FilesRes": {
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
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
