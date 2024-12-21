package utils

import (
	"fmt"
	"github.com/saleh-ghazimoradi/ShopSphere/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func DBURI() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.AppConfig.DBConfig.DbHost, config.AppConfig.DBConfig.DbPort, config.AppConfig.DBConfig.DbUser, config.AppConfig.DBConfig.DbPassword, config.AppConfig.DBConfig.DbName, config.AppConfig.DBConfig.DbSslMode)
}

func DBConnection(DBMigrator func(db *gorm.DB) error) (*gorm.DB, error) {
	fmt.Println(DBURI())
	uri := DBURI()

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	log.Println("Successfully connected to database")

	if err = DBMigrator(db); err != nil {
		log.Fatalf("Unable to migrate database: %v\n", err)
	}

	return db, nil
}
