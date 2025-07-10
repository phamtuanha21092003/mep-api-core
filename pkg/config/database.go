package config

import (
	"os"
	"strconv"
	"time"
)

type DB struct {
	DBUri           string
	MaxOpenConn     int
	MaxIdleConn     int
	MaxConnLifetime time.Duration
}

var db = &DB{}

// GetDBConfig returns the default DB configuration
func GetDBConfig() *DB {
	return db
}

// LoadDBConfig loads DB configuration
func loadDbConfig() {
	db.DBUri = os.Getenv("DB_URI")
	db.MaxOpenConn, _ = strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTIONS"))
	db.MaxIdleConn, _ = strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	lifeTime, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))
	db.MaxConnLifetime = time.Duration(lifeTime) * time.Minute
}
