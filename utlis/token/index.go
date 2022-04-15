package token

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"modTest/module"
	"net/http"
	"time"
)

var (
	Secret     = "HeJin" // 加盐
	ExpireTime = 3600    // token有效期

)

func init() {

}

// CreateToken
// 创建token函数
func CreateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// VerifyToken
// token的验证
func VerifyToken(c *gin.Context) (bool, error) {
	strToken := c.Request.Header.Get("Authorization")
	claim, err := VerifyAction(strToken)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return false, err
	}

	nowTime := time.Now().Unix()

	if nowTime < claim.StandardClaims.IssuedAt { // IssuedAt token过期时间
		return false, errors.New("token已过期")
	}
	return true, nil
}

// refresh
// token的刷新
func refresh(c *gin.Context) {
	strToken := c.Param("token")
	claims, err := VerifyAction(strToken)
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

// VerifyAction
// token的验证
func VerifyAction(strToken string) (*module.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(strToken, &module.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, errors.New(err.Error())
	}
	claims, ok := token.Claims.(*module.JWTClaims)
	if !ok {
		return nil, errors.New("ErrorReason_ReLogin")
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New(err.Error())
	}

	//fmt.Println("verify")
	return claims, nil
}
