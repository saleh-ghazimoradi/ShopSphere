package utils

import (
	"github.com/saleh-ghazimoradi/ShopSphere/internal/service/serviceModels"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(&serviceModels.User{}, &serviceModels.BankAccount{})
}
