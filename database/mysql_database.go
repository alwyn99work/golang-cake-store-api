package database

import (
	"context"
	"database/sql"
	"fmt"
	"golang-restful-api/helper"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewMysqlDB() *sql.DB {
	// db, err := sql.Open("mysql", "root:password@tcp(cake-store-mysql:3306)/cake?parseTime=true")
	db, err := sql.Open("mysql", "root:password@tcp(cake-store-mariadb:3306)/cake?parseTime=true")

	// connStr, err := loadConfig()
	// if err != nil {
	// 	return nil
	// }

	// db, err := sql.Open("mariadb", connStr)
	helper.PanicIfError(err)

	// See "Important settings" section.
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)

	return db
}

func Migrate(mysqlDB *sql.DB) {
	ctx := context.Background()
	var sql string = "CREATE TABLE IF NOT EXISTS cakes(id INT PRIMARY KEY AUTO_INCREMENT, title VARCHAR(255) NOT NULL, description TEXT NULL, rating FLOAT DEFAULT 0, image VARCHAR(191) NULL, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP)"
	_, err := mysqlDB.ExecContext(ctx, sql)
	helper.PanicIfError(err)
}

func loadConfig() (string, error) {
	if os.Getenv("DB_HOST") == "" {
		return "", fmt.Errorf("Environment variable DB_HOST must be set")
	}
	if os.Getenv("DB_PORT") == "" {
		return "", fmt.Errorf("Environment variable DB_PORT must be set")
	}
	if os.Getenv("DB_USER") == "" {
		return "", fmt.Errorf("Environment variable DB_USER must be set")
	}
	if os.Getenv("DB_DATABASE") == "" {
		return "", fmt.Errorf("Environment variable DB_DATABASE must be set")
	}
	if os.Getenv("DB_PASSWORD") == "" {
		return "", fmt.Errorf("Environment variable DB_PASSWORD must be set")
	}
	// db, err := sql.Open("mysql", "root:password@tcp(cake-store-mysql:3306)/golang?parseTime=true")
	connStr := fmt.Sprintf("mariadb://%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)
	return connStr, nil
}
