package database

import (
	"fmt"
	"time"

	"github.com/ArdyJunata/grpc-go/internal/config"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func ConnectPostgres(dbConfig config.DBConfig) (db *sqlx.DB, err error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.SSLMode)

	db, err = sqlx.Open("postgres", dsn)
	if err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		return
	}

	// set connection pool
	db.SetMaxOpenConns(dbConfig.ConnectionPool.MaxOpenConnection)
	db.SetMaxIdleConns(dbConfig.ConnectionPool.MaxIdleConnection)
	db.SetConnMaxLifetime(time.Duration(dbConfig.ConnectionPool.MaxLifeTime) * time.Second)
	db.SetConnMaxIdleTime(time.Duration(dbConfig.ConnectionPool.MaxIdleTime) * time.Second)

	return
}
