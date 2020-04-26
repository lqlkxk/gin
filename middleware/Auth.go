package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lqlkxk/gin/utils"
	"net/http"
	"strings"
)

func CookieAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 读取cookie
		cookie, e := context.Request.Cookie("user_cookie")
		if e == nil {
			context.SetCookie(cookie.Name, cookie.Value, 1000, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
			context.Next()
		} else {
			context.Abort()
			context.HTML(http.StatusUnauthorized, "401.tmpl", nil)
		}
	}
}

func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		result := utils.ResponseMsg{
			Code: http.StatusUnauthorized,
			Msg:  "无法认证，重新登录",
			Data: "",
		}
		auth := context.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			context.Abort()
			context.JSON(http.StatusUnauthorized, gin.H{
				"result": result,
			})
		}
		auth = strings.Fields(auth)[1]
		// 校验token
		_, err := parseToken(auth)
		if err != nil {
			context.Abort()
			result.Msg = "token 过期"
			context.JSON(http.StatusUnauthorized, gin.H{
				"result": result,
			})
		} else {
			context.Next()
		}

	}
}

func parseToken(token string) (*jwt.StandardClaims, error) {
	_, secret, _ := utils.GetJWT()
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(secret), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}
