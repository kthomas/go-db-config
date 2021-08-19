package dbconf

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres
)

var dbInstance *gorm.DB
var dbOnce sync.Once

// DatabaseConnection returns a leased database connection from the underlying
// pool configured from the environment-configured database connection
func DatabaseConnection() *gorm.DB {
	dbOnce.Do(func() {
		db, err := DatabaseConnectionFactory(dbConf)
		if err != nil {
			msg := fmt.Sprintf("Database connection failed; %s", err.Error())
			panic(msg)
		}

		dbInstance = db
	})
	return dbInstance
}

// DatabaseConnectionFactory returns a leased database connection from the underlying
// pool configured from the given database configuration
func DatabaseConnectionFactory(cfg *DBConfig) (*gorm.DB, error) {
         args := fmt.Sprintf(
                "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
                cfg.DatabaseHost,
                cfg.DatabasePort,
                cfg.DatabaseUser,
                cfg.DatabasePassword,
                cfg.DatabaseName,
                cfg.DatabaseSSLMode,
        )

	db, err := gorm.Open("postgres", args)

	if err != nil {
		return nil, err
	}

	db.LogMode(cfg.DatabaseEnableLogging)
	db.DB().SetMaxOpenConns(cfg.DatabasePoolMaxOpenConnections)

	if cfg.DatabasePoolMaxIdleConnections >= 0 {
		db.DB().SetMaxIdleConns(cfg.DatabasePoolMaxIdleConnections)
	}

	return db, nil
}
