package token

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"modTest/module"
	"net/http"
	"time"
)

var (
	Secret     = "dong_tech" // 加盐
	ExpireTime = 3600        // token有效期

)

func init() {

}

// CreateToken
// 创建token函数
func CreateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signedToken, err := token.SignedString([]byte(Secret))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// verify
// token的验证
func verify(c *gin.Context) {
	strToken := c.Param("token")
	claim, err := verifyAction(strToken)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, "verify,", claim.Username)
}

// refresh
// token的刷新
func refresh(c *gin.Context) {
	strToken := c.Param("token")
	claims, err := verifyAction(strToken)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	claims.ExpiresAt = time.Now().Unix() + (claims.ExpiresAt - claims.IssuedAt)

	signedToken, err := CreateToken(claims)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, signedToken)
}

// verifyAction
// token的验证
func verifyAction(strToken string) (*module.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(strToken, &module.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, errors.New("ErrorReason_ServerBusy")
	}
	claims, ok := token.Claims.(*module.JWTClaims)
	if !ok {
		return nil, errors.New("ErrorReason_ReLogin")
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New("ErrorReason_ReLogin")
	}
	fmt.Println("verify")
	return claims, nil
}
