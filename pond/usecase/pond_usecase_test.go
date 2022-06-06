package usecase_test

import (
	"go-template/domain"
	"go-template/domain/mocks"
	"testing"

	_pondUsecase "go-template/pond/usecase"

	"github.com/stretchr/testify/suite"
)

type PondUsecaseSuite struct {
	suite.Suite
	repository *mocks.PondRepository
	usecase    domain.PondUsecase
}

func TestPondUsecase(t *testing.T) {
	suite.Run(t, new(PondUsecaseSuite))
}

func (suite *PondUsecaseSuite) SetupTest() {
	repository := new(mocks.PondRepository)
	usecase := _pondUsecase.NewPondUsecase(repository)
	suite.repository = repository
	suite.usecase = usecase
}

func (suite *PondUsecaseSuite) TestCreatePond_Positive() {
	pond := domain.Pond{
		Name:   "Pond Testing",
		FarmId: 1,
	}
	suite.repository.On("Create", pond).Return(uint(1), nil)
	input := domain.CreatePondDto{
		Name:   "Pond Testing",
		FarmId: 1,
	}
	_, err := suite.usecase.Create(input)
	suite.Nil(err, "err is a nil pointer so no error in this process")
	suite.repository.AssertExpectations(suite.T())
}

func (suite *PondUsecaseSuite) TestGetAll_Positive() {
	ponds := []domain.Pond{
		{
			Name:   "Pond test 1",
			FarmId: 1,
		},
		{
			Name:   "Pond test 2",
			FarmId: 1,
		},
	}
	suite.repository.On("GetAll").Return(ponds, nil)
	pondRes, err := suite.usecase.GetAll()
	suite.NoError(err, "should not error while fetching data")
	suite.Equal(len(pondRes), len(ponds), "should containe same length")
	suite.Equal(pondRes, ponds, "should be equal")
}

func (suite *PondUsecaseSuite) TestGetAllPonds_Nil_Negative() {
	emptyPonds := []domain.Pond(nil)
	suite.repository.On("GetAll").Return(emptyPonds, nil)
	ponds, err := suite.usecase.GetAll()
	suite.Error(err, "error not found")
	suite.Equal(len(ponds), 0, "ponds must be empty")
}

func (suite *PondUsecaseSuite) TestGetById_Positive() {
	id := 2
	pond := domain.Pond{
		Name: "Pond Testing",
	}
	pond.ID = uint(id)
	suite.repository.On("GetById", uint(id)).Return(pond, nil)
	pondRes, err := suite.usecase.GetById(uint(id))
	suite.NoError(err, "no error while fetching data")
	suite.Equal(pond, pondRes, "should be equal")
}

func (suite *PondUsecaseSuite) TestGetByFarmId_Positive() {
	id := 1
	ponds := []domain.Pond{
		{
			Name:   "Pond Testing",
			FarmId: 1,
		},
	}
	suite.repository.On("GetByFarmId", uint(id)).Return(ponds, nil)
	pondRes, err := suite.usecase.GetByFarmId(uint(id))
	suite.NoError(err, "no error while fetching data")
	suite.Equal(len(pondRes), len(ponds), "should containe same length")
	suite.Equal(pondRes, ponds, "should be equal")
}

func (suite *PondUsecaseSuite) TestUpdate_Positive() {
	id := 1
	inputPond := domain.UpdatePondDto{
		Name: "Updated Pond",
		Id:   uint(id),
	}
	pond := domain.Pond{
		Name:   "Updated Pond",
		FarmId: 1,
	}
	pond.ID = uint(1)
	suite.repository.On("GetById", uint(id)).Return(pond, nil)
	suite.repository.On("Update", pond).Return(nil)
	err := suite.usecase.Update(inputPond)
	suite.NoError(err, "no error while updating data")
}

func (suite *PondUsecaseSuite) TestDelete_Positive() {
	id := 1
	pond := domain.Pond{
		Name: "Deleted Pond",
	}
	pond.ID = uint(1)
	suite.repository.On("GetById", uint(id)).Return(pond, nil)
	suite.repository.On("Delete", pond).Return(nil)
	err := suite.usecase.Delete(uint(id))
	suite.NoError(err, "no error while updating data")
}
