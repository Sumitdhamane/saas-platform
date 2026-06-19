package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sumitdhamane/saas-platform/configs"
)

var DB *sql.DB

func Connect(cfg *configs.Config) error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// Connection Pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)

	err = db.Ping()
	if err != nil {
		return err
	}

	DB = db

	log.Println("✅ MySQL Connected")

	return nil
}
