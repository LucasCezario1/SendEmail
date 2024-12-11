package database

import (
	"SendEmail/internal/campaign"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {
	dsn := "host=localhost user=postgres password=lucas dbname=postgres port=5432 sslmode=disable search_path=sistema"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect database")
	}

	// auto migracao para criar tabela igual o jdbc
	db.AutoMigrate(&campaign.Campaign{}, &campaign.Contact{})

	return db
}
