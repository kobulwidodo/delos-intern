package utils

import (
	"fmt"
	"go-template/domain"

	"gorm.io/gorm"
)

type TestingHelper struct {
	db *gorm.DB
}

func NewTestingHelper(db *gorm.DB) TestingHelper {
	return TestingHelper{db}
}

func (e *TestingHelper) TruncateTable(tablename string) error {
	query := fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", tablename)
	if err := e.db.Exec(query).Error; err != nil {
		return err
	}
	return nil
}

func (e *TestingHelper) CreateFarm() error {
	farm := domain.Farm{
		Name: "Farm Test Get",
	}
	err := e.db.Create(&farm).Error
	if err != nil {
		return err
	}
	return nil
}
