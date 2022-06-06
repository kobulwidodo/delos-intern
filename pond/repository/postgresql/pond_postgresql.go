package postgresql

import (
	"go-template/domain"

	"gorm.io/gorm"
)

type PondRepository struct {
	db *gorm.DB
}

func NewPondRepository(db *gorm.DB) domain.PondRepository {
	return &PondRepository{db}
}

func (r *PondRepository) Create(pond domain.Pond) (uint, error) {
	if err := r.db.Create(&pond).Error; err != nil {
		return 0, err
	}
	return pond.ID, nil
}

func (r *PondRepository) GetAll() ([]domain.Pond, error) {
	var ponds []domain.Pond
	if err := r.db.Find(&ponds).Error; err != nil {
		return ponds, err
	}
	return ponds, nil
}

func (r *PondRepository) GetById(id uint) (domain.Pond, error) {
	var pond domain.Pond
	if err := r.db.Where("id = ?", id).Find(&pond).Error; err != nil {
		return pond, err
	}
	return pond, nil
}

func (r *PondRepository) GetByFarmId(id uint) ([]domain.Pond, error) {
	var ponds []domain.Pond
	if err := r.db.Where("farm_id = ?", id).Find(&ponds).Error; err != nil {
		return ponds, err
	}
	return ponds, nil
}

func (r *PondRepository) Update(pond domain.Pond) error {
	if err := r.db.Save(&pond).Error; err != nil {
		return err
	}
	return nil
}

func (r *PondRepository) Delete(pond domain.Pond) error {
	if err := r.db.Delete(&pond).Error; err != nil {
		return err
	}
	return nil
}
