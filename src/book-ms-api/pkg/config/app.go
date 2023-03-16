package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {

	dsn := "host=localhost user=postgres password=postgres dbname=book-ms port=5432 sslmode=disable TimeZone=Africa/Lagos"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)

	}
	db = d

}

func GetDB() *gorm.DB {
	return db
}
