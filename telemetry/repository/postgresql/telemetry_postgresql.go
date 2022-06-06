package postgresql

import (
	"go-template/domain"

	"gorm.io/gorm"
)

type TelemetryRepository struct {
	db *gorm.DB
}

func NewTelemetryRepository(db *gorm.DB) domain.TelemetryRepository {
	return &TelemetryRepository{db}
}

func (r *TelemetryRepository) Create(telemetry domain.Telemetry) error {
	if err := r.db.Create(&telemetry).Error; err != nil {
		return err
	}
	return nil
}

func (r *TelemetryRepository) GetAll() ([]domain.Telemetry, error) {
	var telemetries []domain.Telemetry
	if err := r.db.Find(&telemetries).Error; err != nil {
		return telemetries, err
	}
	return telemetries, nil
}

func (r *TelemetryRepository) GetCount() ([]domain.CountResponse, error) {
	var responses []domain.CountResponse
	if err := r.db.Model(&domain.Telemetry{}).Select("endpoint", "method", "count(*) as count", "count(DISTINCT ip) as unique_user").Group("endpoint").Group("method").Find(&responses).Error; err != nil {
		return responses, err
	}
	return responses, nil
}
