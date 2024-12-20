package connection

import (
	"fmt"

	"github.com/abdur-rahman41/hall-management-backend/pkg/config"
	model "github.com/abdur-rahman41/hall-management-backend/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// database connection
func Connect() {
	

	d, err := gorm.Open(postgres.Open(config.LocalConfig.DBURL), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println("Error connecting to DB")
		panic(err)
	}

	fmt.Println("Database Connected")
	db = d
}

// create table
func migrate() {
	if err := db.Migrator().AutoMigrate(&model.User{}); err != nil {
		fmt.Println("Error migrating DB")
		panic(err)
	}

}

// function for getting db instance
func GetDB() *gorm.DB {
	if db == nil {
		Connect()
	}
	migrate()
	return db
}
