package bootstrap

import (
	"fmt"
	"golang-rest-api/domain"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	LoadEnv()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		GetEnv("DB_HOST"),
		GetEnv("DB_USER"),
		GetEnv("DB_PASSWD"),
		GetEnv("DB_NAME"),
		GetEnv("DB_PORT"),
		GetEnv("DB_SSL_MODE"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	// database migration
	// if we don't have the tables gorm will create it for us

	log.Println("Running Database Migrations...")

	err = db.AutoMigrate(
		&domain.User{},
		&domain.Task{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database migration completed successfully.")
	// ----------------------

	return db
}
