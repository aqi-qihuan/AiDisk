package config

import (
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Config struct {
	// Server
	AppName    string
	Debug      bool
	ListenAddr string

	// MySQL
	MySQLHost     string
	MySQLPort     string
	MySQLUser     string
	MySQLPassword string
	MySQLDatabase string

	// MinIO / S3
	MinIOEndpoint         string
	MinIOExternalEndpoint string // 外部访问端点（用于生成预签名 URL）
	MinIOAccessKeyID      string
	MinIOSecretAccessKey  string
	MinIOBucketName       string
	MinIOAvatarBucket     string

	// JWT
	JWTSecret              string
	JWTLoginSubject        string
	JWTShareSubject        string
	JWTAlgorithm           string // JWT 签名算法（HS256/HS384/HS512）
	JWTExpireDays          int    // Token 过期天数
	JWTAccessExpireMinutes int    // Token 过期分钟数（访问令牌）

	// AI
	DashScopeAPIKey string
	DashScopeBase   string
	OllamaBaseURL   string

	// AI Agent Stream (external SSE proxy)
	StreamBaseURL  string
	StreamChatPath string

	// App
	FrontendBaseURL string

	// Defaults
	DefaultStorageSize int64
	RootFolderName     string
	RootParentID       int64
	MaxUploadSize      int64
}

var (
	once   sync.Once
	config *Config
)

func GetConfig() *Config {
	once.Do(func() {
		config = &Config{
			AppName:                getEnv("APP_NAME", "AqiCloud-AgentPan API"),
			Debug:                  getEnv("DEBUG", "false") == "true",
			ListenAddr:             getEnv("LISTEN_ADDR", ":8080"),
			MySQLHost:              getEnv("MYSQL_HOST", ""),
			MySQLPort:              getEnv("MYSQL_PORT", "3306"),
			MySQLUser:              getEnv("MYSQL_USER", ""),
			MySQLPassword:          getEnv("MYSQL_PASSWORD", ""),
			MySQLDatabase:          getEnv("MYSQL_DATABASE", "aqi-cloud-pan"),
			MinIOEndpoint:          getEnv("MINIO_ENDPOINT", ""),
			MinIOExternalEndpoint:  getEnv("MINIO_EXTERNAL_ENDPOINT", ""), // 外部访问端点（用于生成预签名 URL）
			MinIOAccessKeyID:       getEnv("MINIO_ACCESS_KEY", ""),
			MinIOSecretAccessKey:   getEnv("MINIO_SECRET_KEY", ""),
			MinIOBucketName:        getEnv("MINIO_BUCKET", "ai-pan"),
			MinIOAvatarBucket:      getEnv("MINIO_AVATAR_BUCKET", "avatar"),
			JWTSecret:              getEnv("JWT_SECRET", ""),
			JWTLoginSubject:        getEnv("JWT_LOGIN_SUBJECT", "AQI"),
			JWTShareSubject:        getEnv("JWT_SHARE_SUBJECT", "AQI_SHARE"),
			JWTAlgorithm:           getEnv("JWT_ALGORITHM", "HS256"),
			JWTExpireDays:          getEnvInt("JWT_EXPIRE_DAYS", 7),
			JWTAccessExpireMinutes: getEnvInt("JWT_ACCESS_TOKEN_EXPIRE_MINUTES", 30),
			DashScopeAPIKey:        getEnv("DASHSCOPE_API_KEY", ""),
			DashScopeBase:          getEnv("DASHSCOPE_BASE", "https://dashscope.aliyuncs.com/compatible-mode/v1"),
			OllamaBaseURL:          getEnv("OLLAMA_BASE_URL", ""),
			StreamBaseURL:          getEnv("STREAM_BASE_URL", ""),
			StreamChatPath:         getEnv("STREAM_CHAT_PATH", "/api/chat/stream"),
			FrontendBaseURL:        getEnv("FRONTEND_BASE_URL", "http://127.0.0.1:8080"),
			DefaultStorageSize:     getEnvInt64("DEFAULT_STORAGE_SIZE", 10*1024*1024*1024),
			RootFolderName:         getEnv("ROOT_FOLDER_NAME", "全部文件夹"),
			RootParentID:           0,
			MaxUploadSize:          getEnvInt64("MAX_UPLOAD_SIZE", 100*1024*1024),
		}
	})
	return config
}

func getEnv(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}

func getEnvInt(key string, defaultVal int) int {
	if v := os.Getenv(key); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return defaultVal
}

// GetSigningMethod 根据配置返回 JWT 签名方法
func GetSigningMethod() jwt.SigningMethod {
	cfg := GetConfig()
	switch cfg.JWTAlgorithm {
	case "HS384":
		return jwt.SigningMethodHS384
	case "HS512":
		return jwt.SigningMethodHS512
	case "HS256":
		fallthrough
	default:
		return jwt.SigningMethodHS256
	}
}

// GetTokenExpireDuration 返回 Token 过期时长
func GetTokenExpireDuration() time.Duration {
	cfg := GetConfig()
	if cfg.JWTAccessExpireMinutes > 0 {
		return time.Duration(cfg.JWTAccessExpireMinutes) * time.Minute
	}
	return time.Duration(cfg.JWTExpireDays) * 24 * time.Hour
}

func getEnvInt64(key string, defaultVal int64) int64 {
	if v := os.Getenv(key); v != "" {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil {
			return n
		}
	}
	return defaultVal
}
