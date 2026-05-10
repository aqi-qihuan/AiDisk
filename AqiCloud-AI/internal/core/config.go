package core

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

type Config struct {
	// Redis
	RedisHost           string
	RedisPort           int
	RedisDB             int
	RedisPassword       string
	RedisMaxConnections int
	RedisChatHistoryTTL int // seconds

	// MySQL
	MySQLHost     string
	MySQLPort     int
	MySQLUser     string
	MySQLPassword string
	MySQLDatabase string

	// JWT
	JWTSecretKey           string
	JWTAlgorithm           string
	JWTAccessExpireMinutes int
	JWTLoginSubject        string

	// LLM
	LLMProvider      string // "ollama" or "openai_compatible"
	LLMModelName     string
	LLMOllamaBaseURL string
	LLMBaseURL       string
	LLMAPIKey        string
	LLMTemperature   float64
	LLMStreaming     bool

	// App
	AppName         string
	Debug           bool
	ListenAddr      string
	FrontendBaseURL string

	// Token Limits
	TokenDailyLimit         int
	TokenHourlyLimit        int
	MaxChatHistoryMessages  int
	MaxContextTokens        int // 单次请求最大 token 预算
	MaxSummaryLength        int
	SummaryTriggerThreshold int
}

var (
	cfg  *Config
	once sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		cfg = &Config{
			RedisHost:           getEnv("REDIS_HOST", ""),
			RedisPort:           getEnvInt("REDIS_PORT", 6379),
			RedisDB:             getEnvInt("REDIS_DB", 0),
			RedisPassword:       getEnv("REDIS_PASSWORD", ""),
			RedisMaxConnections: getEnvInt("REDIS_MAX_CONNECTIONS", 10),
			RedisChatHistoryTTL: getEnvInt("REDIS_CHAT_HISTORY_TTL", 86400),

			MySQLHost:     getEnv("MYSQL_HOST", ""),
			MySQLPort:     getEnvInt("MYSQL_PORT", 3306),
			MySQLUser:     getEnv("MYSQL_USER", ""),
			MySQLPassword: getEnv("MYSQL_PASSWORD", ""),
			MySQLDatabase: getEnv("MYSQL_DATABASE", "aqi-cloud-pan"),

			JWTSecretKey:           getEnv("JWT_SECRET_KEY", ""),
			JWTAlgorithm:           getEnv("JWT_ALGORITHM", "HS256"),
			JWTAccessExpireMinutes: getEnvInt("JWT_ACCESS_TOKEN_EXPIRE_MINUTES", 30),
			JWTLoginSubject:        getEnv("JWT_LOGIN_SUBJECT", "AQI"),

			LLMProvider:      getEnv("LLM_PROVIDER", "ollama"),
			LLMModelName:     getEnv("LLM_MODEL_NAME", "qwen3.5:9b"),
			LLMOllamaBaseURL: getEnv("LLM_OLLAMA_BASE_URL", ""),
			LLMBaseURL:       getEnv("LLM_BASE_URL", "https://dashscope.aliyuncs.com/compatible-mode/v1"),
			LLMAPIKey:        getEnv("LLM_API_KEY", ""),
			LLMTemperature:   getEnvFloat("LLM_TEMPERATURE", 0.7),
			LLMStreaming:     getEnvBool("LLM_STREAMING", true),

			AppName:         getEnv("APP_NAME", "AI智能体中心API服务"),
			Debug:           getEnvBool("DEBUG", false),
			ListenAddr:      getEnv("LISTEN_ADDR", ":8000"),
			FrontendBaseURL: getEnv("FRONTEND_BASE_URL", "http://127.0.0.1:8081"),

			TokenDailyLimit:         getEnvInt("TOKEN_DAILY_LIMIT", 100000),
			TokenHourlyLimit:        getEnvInt("TOKEN_HOURLY_LIMIT", 10000),
			MaxChatHistoryMessages:  getEnvInt("MAX_CHAT_HISTORY_MESSAGES", 20),
			MaxContextTokens:        getEnvInt("MAX_CONTEXT_TOKENS", 8000),
			MaxSummaryLength:        getEnvInt("MAX_SUMMARY_LENGTH", 500),
			SummaryTriggerThreshold: getEnvInt("SUMMARY_TRIGGER_THRESHOLD", 10),
		}
	})
	return cfg
}

// SetLLMProvider 运行时切换模型提供商（线程安全）
func (c *Config) SetLLMProvider(provider string) {
	cfg.LLMProvider = provider
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return fallback
}

func getEnvFloat(key string, fallback float64) float64 {
	if v := os.Getenv(key); v != "" {
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return f
		}
	}
	return fallback
}

func getEnvBool(key string, fallback bool) bool {
	if v := os.Getenv(key); v != "" {
		if b, err := strconv.ParseBool(v); err == nil {
			return b
		}
	}
	return fallback
}

// MySQLDSN 生成 MySQL 连接字符串
func (c *Config) MySQLDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.MySQLUser, c.MySQLPassword, c.MySQLHost, c.MySQLPort, c.MySQLDatabase)
}

// RedisAddr Redis 地址
func (c *Config) RedisAddr() string {
	return fmt.Sprintf("%s:%d", c.RedisHost, c.RedisPort)
}
