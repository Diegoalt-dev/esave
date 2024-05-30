package database

import (
	"esave/database/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"sync"
	"time"
)

var (
	db   *gorm.DB
	once sync.Once
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func getDbConfig() *DbConfig {
	return &DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	}
}

func GetDb() *gorm.DB {
	return db
}

func InitializeDatabase() {
	once.Do(func() {
		config := getDbConfig()
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.Host, config.User, config.Password, config.Database, config.Port)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		log.Println("initializing database...")
		//dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.Database)
		//db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})

		if err != nil {
			log.Fatalf("failed to connect database: %v", err)
		}

		// Configurar el pool de conexiones
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("failed to get database: %v", err)
		}

		// Establecer el número máximo de conexiones abiertas
		sqlDB.SetMaxOpenConns(10)

		// Establecer el número máximo de conexiones inactivas
		sqlDB.SetMaxIdleConns(5)

		// Establecer el tiempo máximo de vida de una conexión
		sqlDB.SetConnMaxLifetime(time.Hour)

		err = db.AutoMigrate(&models.Spent{})
		if err != nil {
			log.Fatalf("failed to migrate database: %v", err)
		}
	})
}
