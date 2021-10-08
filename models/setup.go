package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func (RatingSummary) TableName() string {
	return "rating_summary"
}

func ConnectDb() {
	db, err := gorm.Open("sqlite3", "./ratings.db")

	db.LogMode(true)

	if err != nil {
		panic("Failed to connect to DB")
	}

	db.AutoMigrate(&RatingSummary{})

	DB = db
}
