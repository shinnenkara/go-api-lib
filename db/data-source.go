package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataSource struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func GetDSN(dataSource DataSource) string {
	dsn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dataSource.Host,
		dataSource.Port,
		dataSource.Username,
		dataSource.Password,
		dataSource.Database,
	)

	return dsn
}

func Connect(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
