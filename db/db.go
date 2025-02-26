package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
    err := godotenv.Load()
    if err != nil {
        log.Println("Warning: Error loading .env file:", err)
    }
		
	if os.Getenv("GO_ENV") == "dev"{
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}
	fmt.Print(os.Getenv("GO_ENV"))
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"),os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),os.Getenv("POSTGRES_DB"))
	
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Database connected")
	return db
}

// DB close
// func CloseDB(db *gorm.DB) {
// 	fmt.Println("Database connection closed")

// 	sqlDB, _ := db.DB()
// 	if err := sqlDB.Close(); err != nil {
// 		log.Fatalln(err)
// 	}
// }