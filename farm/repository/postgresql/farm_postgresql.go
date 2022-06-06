package postgresql

import (
	"go-template/domain"

	"gorm.io/gorm"
)

type FarmRepository struct {
	db *gorm.DB
}

func NewFarmRepository(db *gorm.DB) domain.FarmRepository {
	return &FarmRepository{db}
}

func (r *FarmRepository) Create(farm domain.Farm) (uint, error) {
	if err := r.db.Create(&farm).Error; err != nil {
		return 0, err
	}
	if farm.ID == 0 {
		return 0, nil
	}
	return farm.ID, nil
}

func (r *FarmRepository) GetAll() ([]domain.Farm, error) {
	var farms []domain.Farm
	if err := r.db.Find(&farms).Error; err != nil {
		return farms, err
	}
	return farms, nil
}

func (r *FarmRepository) GetById(id uint) (domain.Farm, error) {
	var farm domain.Farm
	if err := r.db.Where("id = ?", id).Preload("Ponds").Find(&farm).Error; err != nil {
		return farm, err
	}
	return farm, nil
}

func (r *FarmRepository) Update(farm domain.Farm) error {
	if err := r.db.Save(&farm).Error; err != nil {
		return err
	}
	return nil
}

func (r *FarmRepository) Delete(farm domain.Farm) error {
	if err := r.db.Delete(&farm).Error; err != nil {
		return err
	}
	return nil
}
