package usecase

import "go-template/domain"

type FarmUsecase struct {
	farmRepository domain.FarmRepository
}

func NewFarmUsecase(farmRepository domain.FarmRepository) domain.FarmUsecase {
	return &FarmUsecase{farmRepository}
}

func (u *FarmUsecase) Create(input domain.CreateFarmDto) (uint, error) {
	farm := domain.Farm{
		Name: input.Name,
	}
	id, err := u.farmRepository.Create(farm)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *FarmUsecase) GetAll() ([]domain.Farm, error) {
	farms, err := u.farmRepository.GetAll()
	if err != nil {
		return farms, err
	}
	if len(farms) == 0 {
		return farms, domain.ErrNotFound
	}
	return farms, nil
}

func (u *FarmUsecase) GetById(id uint) (domain.Farm, error) {
	farm, err := u.farmRepository.GetById(id)
	if err != nil {
		return farm, err
	}
	if farm.ID == 0 {
		return farm, domain.ErrNotFound
	}
	return farm, nil
}

func (u *FarmUsecase) Update(input domain.UpdateFarmDto) error {
	farm, err := u.farmRepository.GetById(input.Id)
	if err != nil {
		return err
	}
	if farm.ID == 0 {
		return domain.ErrNotFound
	}
	farm.Name = input.Name
	if err := u.farmRepository.Update(farm); err != nil {
		return err
	}
	return nil
}

func (u *FarmUsecase) Delete(id uint) error {
	farm, err := u.farmRepository.GetById(id)
	if err != nil {
		return err
	}
	if farm.ID == 0 {
		return domain.ErrNotFound
	}
	if err := u.farmRepository.Delete(farm); err != nil {
		return err
	}
	return nil
}
