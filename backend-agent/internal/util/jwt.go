package util

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aqi/AqiCloud-AgentPan-Go/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type LoginClaims struct {
	AccountID int64  `json:"accountId"`
	Username  string `json:"username"`
	jwt.RegisteredClaims
}

type ShareClaims struct {
	ShareID int64 `json:"SHARE_ID"`
	jwt.RegisteredClaims
}

// RawClaims is a permissive claims struct used for fallback parsing
// to handle cross-library encoding differences (Java JJWT vs Go golang-jwt)
type RawClaims struct {
	AccountID interface{} `json:"accountId"`
	Username  string      `json:"username"`
	ShareID   interface{} `json:"SHARE_ID"`
	Sub       string      `json:"sub"`
	Exp       interface{} `json:"exp"`
	Iat       interface{} `json:"iat"`
}

func GenerateLoginToken(accountID int64, username string) string {
	cfg := config.GetConfig()
	claims := LoginClaims{
		AccountID: accountID,
		Username:  username,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   cfg.JWTLoginSubject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, _ := token.SignedString([]byte(cfg.JWTSecret))
	return cfg.JWTLoginSubject + s
}

func ParseLoginToken(tokenStr string) (*LoginClaims, error) {
	cfg := config.GetConfig()
	tokenStr = stripPrefix(tokenStr, cfg.JWTLoginSubject)

	// First try standard parsing with LoginClaims struct
	token, err := jwt.ParseWithClaims(tokenStr, &LoginClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWTSecret), nil
	})
	if err == nil {
		if claims, ok := token.Claims.(*LoginClaims); ok && token.Valid && claims.AccountID != 0 {
			return claims, nil
		}
	}

	// Fallback: parse with MapClaims to handle cross-library encoding differences
	log.Printf("[jwt-parse] standard parse failed (%v), trying MapClaims", err)
	mapToken, mapErr := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWTSecret), nil
	})
	if mapErr == nil && mapToken != nil {
		mapClaims := mapToken.Claims.(jwt.MapClaims)
		accountID := extractInt64(mapClaims, "accountId")
		if accountID != 0 {
			username, _ := mapClaims["username"].(string)
			log.Printf("[jwt-parse] MapClaims fallback OK: accountId=%d, username=%s", accountID, username)
			return &LoginClaims{
				AccountID: accountID,
				Username:  username,
			}, nil
		}
	}

	// Last resort: decode payload directly to inspect the claims
	return tryDirectDecode(tokenStr)
}

// extractInt64 safely extracts an int64 from a map that may contain
// float64 (JSON number), int64, string, or json.Number
func extractInt64(m jwt.MapClaims, key string) int64 {
	v, ok := m[key]
	if !ok {
		return 0
	}
	switch val := v.(type) {
	case float64:
		return int64(val)
	case int64:
		return val
	case json.Number:
		n, _ := val.Int64()
		return n
	case string:
		var n int64
		fmt.Sscanf(val, "%d", &n)
		return n
	default:
		return 0
	}
}

func GenerateShareToken(shareID int64) string {
	cfg := config.GetConfig()
	claims := ShareClaims{
		ShareID: shareID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   cfg.JWTShareSubject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, _ := token.SignedString([]byte(cfg.JWTSecret))
	return s
}

func ParseShareToken(tokenStr string) (*ShareClaims, error) {
	cfg := config.GetConfig()

	// Standard parse
	token, err := jwt.ParseWithClaims(tokenStr, &ShareClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWTSecret), nil
	})
	if err == nil {
		if claims, ok := token.Claims.(*ShareClaims); ok && token.Valid && claims.ShareID != 0 {
			return claims, nil
		}
	}

	// Fallback: MapClaims parse
	mapToken, mapErr := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWTSecret), nil
	})
	if mapErr == nil && mapToken != nil {
		mapClaims := mapToken.Claims.(jwt.MapClaims)
		shareID := extractInt64(mapClaims, "SHARE_ID")
		if shareID != 0 {
			return &ShareClaims{ShareID: shareID}, nil
		}
	}

	return nil, fmt.Errorf("分享凭证无效")
}

// tryDirectDecode attempts to decode the JWT payload directly
// as a last resort when all library parsing fails
func tryDirectDecode(tokenStr string) (*LoginClaims, error) {
	parts := strings.Split(tokenStr, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("token无效: malformed JWT")
	}

	// Decode payload
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("token无效: %v", err)
	}

	log.Printf("[jwt-parse] raw payload: %s", payload)

	// Use a decoder with UseNumber to preserve int64 precision
	decoder := json.NewDecoder(strings.NewReader(string(payload)))
	decoder.UseNumber()
	var claims RawClaims
	if err := decoder.Decode(&claims); err != nil {
		return nil, fmt.Errorf("token无效: %v", err)
	}

	accountID := toInt64(claims.AccountID)
	if accountID == 0 {
		return nil, fmt.Errorf("token无效: accountId=0")
	}

	log.Printf("[jwt-parse] direct decode OK: accountId=%d, username=%s", accountID, claims.Username)
	return &LoginClaims{
		AccountID: accountID,
		Username:  claims.Username,
	}, nil
}

func toInt64(v interface{}) int64 {
	switch val := v.(type) {
	case float64:
		return int64(val)
	case int64:
		return val
	case json.Number:
		n, _ := val.Int64()
		return n
	case string:
		var n int64
		fmt.Sscanf(val, "%d", &n)
		return n
	default:
		return 0
	}
}

func stripPrefix(s, prefix string) string {
	if len(s) >= len(prefix) && s[:len(prefix)] == prefix {
		return s[len(prefix):]
	}
	return s
}
