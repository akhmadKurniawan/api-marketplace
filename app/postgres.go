package app

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dB *gorm.DB

// DBInit create connection to database
func DBInitP() *gorm.DB {
	fmt.Println("ok")
	// e := godotenv.Load() //Load .env file
	// if e != nil {
	// 	fmt.Print(e)
	// }

	// host := os.Getenv("POSTGRES_HOST")
	// port := os.Getenv("POSTGRES_PORT")
	// user := os.Getenv("POSTGRES_USER_DB")
	// password := os.Getenv("POSTGRES_PASSWORD")
	// dbName := os.Getenv("POSTGRES_NAME_DB")
	// dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
	// 	host, port, dbName, user, password)
	// fmt.Println(host)
	// fmt.Println(port)
	// fmt.Println(user)
	// fmt.Println(dbName)
	// fmt.Println(password)

	dsn, err := pq.ParseURL(os.Getenv("DATABASE_URL"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	fmt.Println(db)

	if err != nil {
		log.Panic("failed to connect to database")
	}
	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxLifetime(time.Minute * 5)
	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetMaxOpenConns(5)
	// defer sqlDB.Close()

	fmt.Println("ok")
	dB = db
	return dB
}

// GetDB getdb
func GetDBP() *gorm.DB {
	return dB
}
