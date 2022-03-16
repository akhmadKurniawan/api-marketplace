package app

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql" //for import mysql
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dB *gorm.DB

// DBInit create connection to database
func DBInit() *gorm.DB {
	fmt.Println("ok")
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	username := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=%s", username, password, host, port, dbName, "Asia%2FJakarta")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Panic("failed to connect to database")
	}
	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxLifetime(time.Minute * 5)
	sqlDB.SetMaxIdleConns(0)
	sqlDB.SetMaxOpenConns(5)
	// defer sqlDB.Close()

	dB = db
	return dB
}

// GetDB getdb
func GetDB() *gorm.DB {
	return dB
}
