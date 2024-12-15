package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host         = "localhost"
	port         = 5432
	databaseName = "mydb"
	username     = "admin"
	password     = "admin123"
)

type Db struct {
	Query gorm.DB
}

func New() Db {
	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		})

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, databaseName)
	newDb, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("connect success", newDb)
	newDb.AutoMigrate(&Book{}, &User{})

	return Db{Query: *newDb}

}
