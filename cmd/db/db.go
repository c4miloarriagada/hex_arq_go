package db

import (
	"fmt"
	"sync"

	"github.com/c4miloarriagada/hexarq/cmd/internal/domain"
	log "github.com/sirupsen/logrus"
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
	DBName     = "go"
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

		log.Infof("Connecting to the database with DSN: %s", dsn)

		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		}), &gorm.Config{})

		if err != nil {
			log.Errorf("Error trying to connect: %v", err)
		}
	})

	return db, err
}

func ConnectDb() {
	db, err := GetDB()

	if err != nil {
		log.Fatalf("Error trying to connect: %v", err)
	}

	pg, err := db.DB()
	if err != nil {
		log.Fatalf("Error trying to get db connection: %v", err)
	}

	err = pg.Ping()
	if err != nil {
		log.Fatalf("Cannot establish ping to db: %v", err)
	} else {
		log.Info("Successfully connected to the database")
	}
}

func CreateTables() {
	err := db.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Info("Table users created successfully!")
}
