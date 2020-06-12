package dbconf

import (
	"os"
	"strconv"
	"sync"
)

// DBConfig is database configuration
type DBConfig struct {
	DatabaseName                   string
	DatabaseHost                   string
	DatabasePort                   uint
	DatabaseUser                   string
	DatabasePassword               string
	DatabaseSSLMode                string
	DatabasePoolMaxIdleConnections int
	DatabasePoolMaxOpenConnections int
	DatabaseEnableLogging          bool
}

var configInstance *DBConfig
var configOnce sync.Once

// GetDBConfig reads the database config out of the environment
func GetDBConfig() *DBConfig {
	configOnce.Do(func() {
		databaseName := os.Getenv("DATABASE_NAME")
		if databaseName == "" {
			databaseName = "dbconf_test"
		}

		databaseHost := os.Getenv("DATABASE_HOST")
		if databaseHost == "" {
			databaseHost = "localhost"
		}

		databasePort, _ := strconv.ParseUint(os.Getenv("DATABASE_PORT"), 10, 64)
		if databasePort == 0 {
			databasePort = 5432
		}

		databaseUser := os.Getenv("DATABASE_USER")
		if databaseUser == "" {
			databaseUser = "root"
		}

		databasePassword := os.Getenv("DATABASE_PASSWORD")
		if databasePassword == "" {
			databasePassword = "password"
		}

		databaseSSLMode := os.Getenv("DATABASE_SSL_MODE")
		if databaseSSLMode == "" {
			databaseSSLMode = "disable"
		}

		databasePoolMaxIdleConnections, _ := strconv.ParseInt(os.Getenv("DATABASE_POOL_MAX_IDLE_CONNECTIONS"), 10, 8)
		if databasePoolMaxIdleConnections == 0 {
			databasePoolMaxIdleConnections = -1
		}

		databasePoolMaxOpenConnections, _ := strconv.ParseInt(os.Getenv("DATABASE_POOL_MAX_OPEN_CONNECTIONS"), 10, 8)

		configInstance = &DBConfig{
			DatabaseName:                   databaseName,
			DatabaseHost:                   databaseHost,
			DatabasePort:                   uint(databasePort),
			DatabaseUser:                   databaseUser,
			DatabasePassword:               databasePassword,
			DatabaseSSLMode:                databaseSSLMode,
			DatabasePoolMaxIdleConnections: int(databasePoolMaxIdleConnections),
			DatabasePoolMaxOpenConnections: int(databasePoolMaxOpenConnections),
			DatabaseEnableLogging:          os.Getenv("DATABASE_LOGGING") == "true",
		}
	})
	return configInstance
}

var dbConf = GetDBConfig()
