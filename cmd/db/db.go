package db

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

const (
	DBUser     = "admin"
	DBPassword = "admin"
	DBName     = "postgres"
	DBPort     = "5432"
	DBHost     = "localhost"
	TimeZone   = "Asia/Shanghai"
	SSLMode    = "disable"
)

func GetDB() (*gorm.DB, error) {
	var err error

	once.Do(func() {

		dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=%s TimeZone=%s",
			DBUser, DBPassword, DBName, DBPort, DBHost, SSLMode, TimeZone)

		fmt.Println("Connecting to the database with DSN:", dsn)

		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		}), &gorm.Config{})

		if err != nil {
			log.Printf("Error trying to connect: %v\n", err)
		}
	})

	return db, err
}
