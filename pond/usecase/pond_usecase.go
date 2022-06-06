package usecase

import "go-template/domain"

type PondUsecase struct {
	pondRepository domain.PondRepository
}

func NewPondUsecase(pondRepository domain.PondRepository) domain.PondUsecase {
	return &PondUsecase{pondRepository}
}

func (u *PondUsecase) Create(input domain.CreatePondDto) (uint, error) {
	pond := domain.Pond{
		Name:   input.Name,
		FarmId: input.FarmId,
	}
	id, err := u.pondRepository.Create(pond)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *PondUsecase) GetAll() ([]domain.Pond, error) {
	ponds, err := u.pondRepository.GetAll()
	if err != nil {
		return ponds, err
	}
	if len(ponds) == 0 {
		return ponds, domain.ErrNotFound
	}
	return ponds, nil
}

func (u *PondUsecase) GetById(id uint) (domain.Pond, error) {
	pond, err := u.pondRepository.GetById(id)
	if err != nil {
		return pond, err
	}
	if pond.ID == 0 {
		return pond, domain.ErrNotFound
	}
	return pond, nil
}

func (u *PondUsecase) GetByFarmId(id uint) ([]domain.Pond, error) {
	ponds, err := u.pondRepository.GetByFarmId(id)
	if err != nil {
		return ponds, err
	}
	if len(ponds) == 0 {
		return ponds, domain.ErrNotFound
	}
	return ponds, nil
}

func (u *PondUsecase) Update(input domain.UpdatePondDto) error {
	pond, err := u.pondRepository.GetById(input.Id)
	if err != nil {
		return err
	}
	if pond.ID == 0 {
		return domain.ErrNotFound
	}
	pond.Name = input.Name
	if err := u.pondRepository.Update(pond); err != nil {
		return err
	}
	return nil
}

func (u *PondUsecase) Delete(id uint) error {
	pond, err := u.pondRepository.GetById(id)
	if err != nil {
		return err
	}
	if pond.ID == 0 {
		return domain.ErrNotFound
	}
	if err := u.pondRepository.Delete(pond); err != nil {
		return err
	}
	return nil
}
