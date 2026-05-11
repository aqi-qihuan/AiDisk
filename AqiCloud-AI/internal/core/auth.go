package core

import (
	"errors"
	"log"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrTokenEmpty  = errors.New("Token不能为空")
	ErrTokenInvalid = errors.New("无效的认证凭据")
)

type UserClaims struct {
	AccountID int    `json:"accountId"`
	Username  string `json:"username"`
	jwt.RegisteredClaims
}

// ParseJWT 解析 JWT token，返回用户信息
func ParseJWT(tokenStr string) (*UserClaims, error) {
	if tokenStr == "" {
		return nil, ErrTokenEmpty
	}

	// 移除 AQI 前缀（Python 版逻辑）
	tokenStr = strings.TrimSpace(tokenStr)
	subject := GetConfig().JWTLoginSubject
	log.Printf("[ParseJWT] original=%q subject=%q", tokenStr, subject)
	tokenStr = strings.TrimPrefix(tokenStr, subject)
	tokenStr = strings.TrimSpace(tokenStr)
	log.Printf("[ParseJWT] after strip=%q", tokenStr)

	if tokenStr == "" {
		return nil, ErrTokenEmpty
	}

	// 获取 JWT 密钥（优先使用 JWT_SECRET，与 AqiCloud-Agent 统一）
	cfg := GetConfig()
	jwtSecret := cfg.JWTSecret
	if jwtSecret == "" {
		jwtSecret = cfg.JWTSecretKey // 兼容旧配置
	}
	log.Printf("[ParseJWT] using secret length: %d", len(jwtSecret))

	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		log.Printf("[ParseJWT] jwt.ParseWithClaims error: %v", err)
		return nil, ErrTokenInvalid
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		return nil, ErrTokenInvalid
	}

	if claims.AccountID == 0 {
		return nil, ErrTokenInvalid
	}

	return claims, nil
}
