package middleware

import (
	"net/http"

	"github.com/aqi/AqiCloud-Agent/internal/model"
	"github.com/aqi/AqiCloud-Agent/internal/util"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			token = c.Query("token")
		}
		if token == "" {
			c.AbortWithStatusJSON(http.StatusOK, model.Error("Token不能为空", model.CodeNotLogin))
			return
		}

		claims, err := util.ParseLoginToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, model.Error(err.Error(), model.CodeNotLogin))
			return
		}

		c.Set("account_id", claims.AccountID)
		c.Set("username", claims.Username)
		c.Next()
	}
}

func ShareTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("share-token")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusOK, model.Error("分享凭证无效", model.CodeShareCodeIllegal))
			return
		}

		claims, err := util.ParseShareToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, model.Error(err.Error(), model.CodeShareCodeIllegal))
			return
		}

		c.Set("share_id", claims.ShareID)
		c.Next()
	}
}

var publicPaths = map[string]bool{
	"/api/account/v1/register":           true,
	"/api/account/v1/login":              true,
	"/api/account/v1/upload_avatar":      true,
	"/api/share/v1/check_share_code":     true,
	"/api/share/v1/visit":                true,
	"/api/share/v1/getShareSimpleDetail": true,
	"/api/share/v1/detail":               true,
	"/api/share/v1/getShareDetail":       true,
	"/api/share/v1/list_share_file":      true,
	"/api/share/v1/listShareFile":        true,
	"/api/share/v1/transfer":             true,
	"/api/file/v1/folder/tree":           true,
	"/v1/chat/completions":               true,
}

func ShouldSkipAuth(path string) bool {
	return publicPaths[path]
}

func AddPublicPath(path string) {
	publicPaths[path] = true
}
