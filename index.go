package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"modTest/module"
	"modTest/utlis/key"
	"net/http"
	"path/filepath"
	"strings"
)

// UserVerify
// 统一处理用户验证
func UserVerify() gin.HandlerFunc {
	return func(g *gin.Context) {
		var url = g.Request.URL.Path

		// 过滤静态资源
		if strings.Index(url, "oss") != -1 || strings.Index(url, "ws") != -1 || strings.Index(url, "swagger") != -1 || strings.Index(url, "login") != -1 || strings.Index(url, "phone_code") != -1 || strings.Index(url, "upload") != -1 {
			return
		}

		// token验证
		//_, err := token.VerifyToken(g)
		//if err != nil {
		//	g.JSON(http.StatusOK, module.Resp{MgsCode: 500, MgsText: "token失败"})
		//	g.Abort()
		//}

		if g.Request.Header.Get("key") != "" {
			baseUrl, _ := filepath.Abs("./static/keys/")

			text := key.RSA_Decrypt([]byte(g.Request.Header.Get("key")), baseUrl+"/private.pem")

			post, _ := ioutil.ReadAll(g.Request.Body)

			var PostData = struct {
				Data interface{} `json:"data"`
			}{}

			var keyData interface{}

			var bytesPost []byte

			_ = json.Unmarshal([]byte(text), &keyData)

			bytesPost, _ = json.Marshal(keyData)

			g.Request.Body = ioutil.NopCloser(bytes.NewBuffer(post))

			_ = json.Unmarshal(post, &PostData)

			post, _ = json.Marshal(PostData.Data)

			if string(post) == "null" {
				return
			}

			fmt.Println(string(bytesPost), string(post))

			if string(bytesPost) != string(post) {
				g.JSON(http.StatusOK, module.Resp{MgsCode: 500, MgsText: "休想修改数据"})
				g.Abort()
			}
		}

	}
}
