package config

import (
	"os"
	"strconv"
	"time"
)

// App holds the App configuration
type App struct {
	Name        string
	Version     string
	Host        string
	Port        int
	Environment string
	Debug       bool
	ReadTimeout time.Duration
	// JWT Conf
	JWTSecretKey             string
	JWTRefreshTokenSecretKey string
	JWTSecretExpire          int
	JWTSecretRefreshExpire   int
	MessageAPIKey            string
	Observability            ObservabilityConfig
}

type ObservabilityConfig struct {
	Enable bool   // Indicates if observability is enabled.
	Mode   string // Specifies the observability mode.
}

var app = &App{}

// AppCfg returns the default App configuration
func AppCfg() *App {
	return app
}

// loadApp loads App configuration
func loadApp() {
	app.Name = "Init project GIN"
	app.Version = "1.0"
	app.Host = os.Getenv("APP_HOST")
	app.Environment = os.Getenv("APP_ENV")
	app.Port, _ = strconv.Atoi(os.Getenv("APP_PORT"))
	app.Debug, _ = strconv.ParseBool(os.Getenv("APP_DEBUG"))
	timeOut, _ := strconv.Atoi(os.Getenv("APP_READ_TIMEOUT"))
	app.ReadTimeout = time.Duration(timeOut) * time.Second

	app.JWTSecretKey = os.Getenv("JWT_SECRET_KEY")
	app.JWTRefreshTokenSecretKey = os.Getenv("JWT_REFRESH_TOKEN_SECRET_KEY")
	app.JWTSecretExpire, _ = strconv.Atoi(os.Getenv("AUTH_JWT_EXPIRY"))
	app.JWTSecretRefreshExpire, _ = strconv.Atoi(os.Getenv("AUTH_JWT_REFRESH_EXPIRY"))
}
