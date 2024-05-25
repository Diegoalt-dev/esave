package database

import (
	"esave/database/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
	"time"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetDb() *gorm.DB {
	return db
}

func InitializeDatabase() {
	once.Do(func() {
		log.Println("initializing database...")
		var err error
		dns := "root:rootroot@tcp(localhost:3308)/esave_db?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
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
