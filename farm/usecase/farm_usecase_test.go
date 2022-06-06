package usecase_test

import (
	"go-template/domain"
	"go-template/domain/mocks"
	"testing"

	_farmUsecase "go-template/farm/usecase"

	"github.com/stretchr/testify/suite"
)

type FarmUsecaseSuite struct {
	suite.Suite
	repository *mocks.FarmRepository
	usecase    domain.FarmUsecase
}

func TestFarmUsecase(t *testing.T) {
	suite.Run(t, new(FarmUsecaseSuite))
}

func (suite *FarmUsecaseSuite) SetupTest() {
	repository := new(mocks.FarmRepository)
	usecase := _farmUsecase.NewFarmUsecase(repository)
	suite.repository = repository
	suite.usecase = usecase
}

func (suite *FarmUsecaseSuite) TestCreateFarm_Positive() {
	farm := domain.Farm{
		Name: "Farm Testing",
	}
	suite.repository.On("Create", farm).Return(uint(1), nil)
	input := domain.CreateFarmDto{
		Name: "Farm Testing",
	}
	_, err := suite.usecase.Create(input)
	suite.Nil(err, "err is a nil pointer so no error in this process")
	suite.repository.AssertExpectations(suite.T())
}

func (suite *FarmUsecaseSuite) TestGetAll_Positive() {
	farms := []domain.Farm{
		{
			Name: "Farm test 1",
		},
		{
			Name: "Farm test 2",
		},
	}
	suite.repository.On("GetAll").Return(farms, nil)
	farmsRes, err := suite.usecase.GetAll()
	suite.NoError(err, "should not error while fetching data")
	suite.Equal(len(farmsRes), len(farms), "should containe same length")
	suite.Equal(farmsRes, farms, "should be equal")
}

func (suite *FarmUsecaseSuite) TestGetAllFarms_Nil_Negative() {
	emptyFarms := []domain.Farm(nil)
	suite.repository.On("GetAll").Return(emptyFarms, nil)
	farms, err := suite.usecase.GetAll()
	suite.Error(err, "error not found")
	suite.Equal(len(farms), 0, "farms must be empty")
}

func (suite *FarmUsecaseSuite) TestGetById_Positive() {
	id := 2
	farm := domain.Farm{
		Name: "Farm Testing",
	}
	farm.ID = uint(id)
	suite.repository.On("GetById", uint(id)).Return(farm, nil)
	farmRes, err := suite.usecase.GetById(uint(id))
	suite.NoError(err, "no error while fetching data")
	suite.Equal(farm, farmRes, "should be equal")
}

func (suite *FarmUsecaseSuite) TestUpdate_Positive() {
	id := 1
	inputFarm := domain.UpdateFarmDto{
		Name: "Updated Farm",
		Id:   uint(id),
	}
	farm := domain.Farm{
		Name: "Updated Farm",
	}
	farm.ID = uint(1)
	suite.repository.On("GetById", uint(id)).Return(farm, nil)
	suite.repository.On("Update", farm).Return(nil)
	err := suite.usecase.Update(inputFarm)
	suite.NoError(err, "no error while updating data")
}

func (suite *FarmUsecaseSuite) TestDelete_Positive() {
	id := 1
	farm := domain.Farm{
		Name: "Deleted Farm",
	}
	farm.ID = uint(1)
	suite.repository.On("GetById", uint(id)).Return(farm, nil)
	suite.repository.On("Delete", farm).Return(nil)
	err := suite.usecase.Delete(uint(id))
	suite.NoError(err, "no error while updating data")
}
