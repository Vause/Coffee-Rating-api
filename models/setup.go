package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"go.uber.org/zap"
)

var DB *gorm.DB

func (RatingSummary) TableName() string {
	return "rating_summary"
}

func ConnectDb(logger *zap.SugaredLogger) {
	db, err := gorm.Open("sqlite3", "./ratings.db")

	db.LogMode(true)

	if err != nil {
		logger.Panic("Failed to connect to DB",
			"err", err,
			"db", db)
	}

	db.AutoMigrate(&RatingSummary{})

	DB = db
}
