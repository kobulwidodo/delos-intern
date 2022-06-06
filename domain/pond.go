package domain

import "gorm.io/gorm"

type Pond struct {
	gorm.Model
	Name   string `json:"name"`
	FarmId uint   `json:"farm_id"`
}

type PondRepository interface {
	Create(pond Pond) (uint, error)
	GetAll() ([]Pond, error)
	GetById(id uint) (Pond, error)
	GetByFarmId(id uint) ([]Pond, error)
	Update(pond Pond) error
	Delete(pond Pond) error
}

type PondUsecase interface {
	Create(input CreatePondDto) (uint, error)
	GetAll() ([]Pond, error)
	GetById(id uint) (Pond, error)
	GetByFarmId(id uint) ([]Pond, error)
	Update(input UpdatePondDto) error
	Delete(id uint) error
}

type CreatePondDto struct {
	Name   string `json:"name" binding:"required"`
	FarmId uint
}

type UpdatePondDto struct {
	Name string `json:"name" binding:"required"`
	Id   uint
}

type PondIdUriBinding struct {
	Id uint `uri:"id" binding:"required"`
}
