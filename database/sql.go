package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbGlobal *gorm.DB

func NewDB() (*gorm.DB, error) {
	if dbGlobal != nil {
		return dbGlobal, nil
	}

	// Ambil tipe database dari environment variable
	dbType := os.Getenv("DB_TYPE")

	var dsn string
	var dialector gorm.Dialector

	if dbType == "mysql" {
		// MySQL DSN format
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"))
		dialector = mysql.Open(dsn)
	} else if dbType == "postgresql" {
		// PostgreSQL DSN format
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"))
		dialector = postgres.Open(dsn)
	} else {
		return nil, fmt.Errorf("unsupported DB_TYPE: %s", dbType)
	}

	// Buka koneksi database
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	dbGlobal = db
	return dbGlobal, nil
}
