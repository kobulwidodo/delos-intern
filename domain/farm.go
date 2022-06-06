package domain

import "gorm.io/gorm"

type Farm struct {
	gorm.Model
	Name  string `json:"name" gorm:"not null"`
	Ponds []Pond `json:"ponds"`
}

type FarmRepository interface {
	Create(farm Farm) (uint, error)
	GetAll() ([]Farm, error)
	GetById(id uint) (Farm, error)
	Update(farm Farm) error
	Delete(farm Farm) error
}

type FarmUsecase interface {
	Create(input CreateFarmDto) (uint, error)
	GetAll() ([]Farm, error)
	GetById(id uint) (Farm, error)
	Update(input UpdateFarmDto) error
	Delete(id uint) error
}

type CreateFarmDto struct {
	Name string `json:"name" binding:"required"`
}

type UpdateFarmDto struct {
	Name string `json:"name" binding:"required"`
	Id   uint
}

type FarmIdUriBinding struct {
	Id uint `uri:"id" binding:"required"`
}
