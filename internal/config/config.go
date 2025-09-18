package config

import (
	"fmt"
	"maps"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

// AppConfig holds all configuration parsed from environment variables.
type AppConfig struct {
	// DB
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     int
	DBName     string
	DBSSLMode  string // disable, require, verify-ca, verify-full

	// DB Pool
	DBMaxOpenConns    int           // default 10
	DBMaxIdleConns    int           // default 5
	DBConnMaxLifetime time.Duration // default 30m
	DBConnMaxIdleTime time.Duration // default 10m

	// App
	AppPort   int    // default 8080
	AppSecure bool   // session cookie secure flag
	AppCSP    string // optional override for Content-Security-Policy
}

// Load reads configuration from .env files and environment variables.
// Order of precedence: env vars > .env.local > .env
func Load() (*AppConfig, error) {
	// Load .env and .env.local if present. Behavior:
	//  - Real environment variables always win (they are not overwritten)
	//  - .env.local overrides values from .env
	// We read both files into maps and set environment variables only when
	// the key is not already present in the actual environment.
	envMap := map[string]string{}
	if m, err := godotenv.Read(".env"); err == nil {
		maps.Copy(envMap, m)
	}
	if m, err := godotenv.Read(".env.local"); err == nil {
		// .env.local overrides .env
		maps.Copy(envMap, m)
	}
	for k, v := range envMap {
		if _, ok := os.LookupEnv(k); !ok {
			_ = os.Setenv(k, v)
		}
	}

	cfg := &AppConfig{
		DBUser:            getenv("DB_USER", "postgres"),
		DBPassword:        getenv("DB_PASSWORD", ""),
		DBHost:            getenv("DB_HOST", "localhost"),
		DBPort:            getint("DB_PORT", 5432),
		DBName:            getenv("DB_NAME", "postgres"),
		DBSSLMode:         getenv("DB_SSLMODE", "disable"),
		DBMaxOpenConns:    getint("DB_MAX_OPEN_CONNS", 10),
		DBMaxIdleConns:    getint("DB_MAX_IDLE_CONNS", 5),
		DBConnMaxLifetime: getduration("DB_CONN_MAX_LIFETIME", 30*time.Minute),
		DBConnMaxIdleTime: getduration("DB_CONN_MAX_IDLE_TIME", 10*time.Minute),
		AppPort:           getint("APP_PORT", 8080),
		AppSecure:         getbool("APP_SECURE", false),
		AppCSP:            getenv("APP_CSP", ""),
	}
	return cfg, nil
}

// BuildPostgresDSN returns a lib/pq compatible DSN.
// Example: postgres://user:pass@host:port/dbname?sslmode=disable
func (c *AppConfig) BuildPostgresDSN() string {
	user := urlEscape(c.DBUser)
	pass := urlEscape(c.DBPassword)
	host := c.DBHost
	if c.DBPort > 0 {
		host = fmt.Sprintf("%s:%d", c.DBHost, c.DBPort)
	}
	// Note: if password is empty, omit the colon segment
	auth := user
	if pass != "" {
		auth = fmt.Sprintf("%s:%s", user, pass)
	}
	return fmt.Sprintf("postgres://%s@%s/%s?sslmode=%s", auth, host, c.DBName, c.DBSSLMode)
}

// Helpers
func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func getint(key string, def int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return def
}

func getbool(key string, def bool) bool {
	if v := os.Getenv(key); v != "" {
		switch strings.ToLower(v) {
		case "1", "t", "true", "y", "yes", "on":
			return true
		case "0", "f", "false", "n", "no", "off":
			return false
		}
	}
	return def
}

func getduration(key string, def time.Duration) time.Duration {
	if v := os.Getenv(key); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			return d
		}
	}
	return def
}

// urlEscape performs a minimal percent-encoding for username/password in DSN.
func urlEscape(s string) string {
	r := strings.NewReplacer(
		" ", "%20",
		"#", "%23",
		"%", "%25",
		"?", "%3F",
		"/", "%2F",
		":", "%3A",
		"@", "%40",
	)
	return r.Replace(s)
}
