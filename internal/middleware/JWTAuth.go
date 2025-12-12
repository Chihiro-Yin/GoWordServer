package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"imxy.top/wordserver/internal/tool" // 你的响应工具包
)

var jwtSecret = []byte("wtf??cpx,yzp,rw,gss,wnh,zxz,zh,xzm,zyn,fhl,yy,ymx,wt")

type CustomClaims struct {
	UserID uint   `json:"user_id"`
	Nick   string `json:"nick"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, nick string) (string, error) {
	claims := CustomClaims{
		UserID: userID,
		Nick:   nick,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(365 * 24 * time.Hour)), // 365天过期
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "yzp", // 签发者
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
	return token, err
}

// JWTAuth Chi框架中间件：验证Token并将用户信息存入上下文
func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 从请求头提取Token（格式：Bearer <token>）
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			tool.JSONResponse(w, map[string]string{"msg": "请先登录"}, http.StatusUnauthorized)
			return
		}

		// 2. 解析Bearer Token格式（分割出纯Token字符串）
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			tool.JSONResponse(w, map[string]string{"msg": "Token格式错误"}, http.StatusUnauthorized)
			return
		}
		tokenStr := parts[1]

		// 3. 解析并验证Token
		claims := &CustomClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			// 验证签名算法（防止算法篡改）
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})

		// 4. Token验证失败处理（过期、签名错误、无效等）
		if err != nil || !token.Valid {
			tool.JSONResponse(w, map[string]string{"msg": "Token无效或已过期"}, http.StatusUnauthorized)
			return
		}

		// 5. 将解析出的用户信息存入上下文（供后续接口使用）
		ctx := context.WithValue(r.Context(), "user_id", claims.UserID) // 用户ID
		ctx = context.WithValue(ctx, "nick", claims.Nick)               // 用户昵称
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
