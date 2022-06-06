package infrastructure

import (
	"fmt"
	"go-template/domain"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConfig struct {
	DbHost string
	DbUser string
	DbPass string
	DbName string
	DbPort string
}

func NewDbConfig() DbConfig {
	return DbConfig{DbHost: os.Getenv("DB_HOST"), DbUser: os.Getenv("DB_USER"), DbPass: os.Getenv("DB_PASS"), DbName: os.Getenv("DB_NAME"), DbPort: os.Getenv("DB_PORT")}
}

func (c *DbConfig) InitDb() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		c.DbHost,
		c.DbUser,
		c.DbPass,
		c.DbName,
		c.DbPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}

	if err := db.AutoMigrate(&domain.Farm{}, &domain.Pond{}, &domain.Telemetry{}); err != nil {
		return nil, err
	}

	return db, nil
}
