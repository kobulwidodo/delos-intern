package postgresql_test

import (
	"fmt"
	"go-template/domain"
	"go-template/infrastructure"
	"go-template/utils"
	"log"
	"testing"

	_pondRepository "go-template/pond/repository/postgresql"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
)

type PondRepositorySuite struct {
	suite.Suite
	repository domain.PondRepository
	helper     utils.TestingHelper
}

func TestPondRepository(t *testing.T) {
	if err := godotenv.Load("../../../.env"); err != nil {
		fmt.Println("failed to load env")
	}
	suite.Run(t, new(PondRepositorySuite))
}

func (suite *PondRepositorySuite) SetupSuite() {
	dbConfig := infrastructure.NewDbConfig()
	db, err := dbConfig.InitDb()
	if err != nil {
		log.Fatal("failed to connect with database")
		panic(err)
	}
	repository := _pondRepository.NewPondRepository(db)
	suite.NoError(err, "no error while creating new farm")
	suite.repository = repository
	suite.helper = utils.NewTestingHelper(db)
}

func (suite *PondRepositorySuite) TearDownTest() {
	defer suite.helper.TruncateTable("ponds")
	defer suite.helper.TruncateTable("farms")
}

func (suite *PondRepositorySuite) TestCreatePond_Positive() {
	err := suite.helper.CreateFarm()
	suite.NoError(err, "no error while creating new farm")
	pond := domain.Pond{
		Name:   "Pond Test",
		FarmId: 1,
	}
	_, err = suite.repository.Create(pond)
	suite.NoError(err, "no error while creating new pond")
}

func (suite *PondRepositorySuite) TestGetAllPond_Positive() {
	err := suite.helper.CreateFarm()
	suite.NoError(err, "no error while creating new farm")
	pond := domain.Pond{
		Name:   "pond Test Get",
		FarmId: 1,
	}
	_, err = suite.repository.Create(pond)
	suite.NoError(err, "no error while creating new pond")
	ponds, err := suite.repository.GetAll()
	suite.NoError(err, "no error while get all ponds")
	suite.NotEmpty(ponds, "should not empty data")
}

func (suite *PondRepositorySuite) TestGetPondById_Positive() {
	err := suite.helper.CreateFarm()
	suite.NoError(err, "no error while creating new farm")
	newPond := domain.Pond{
		Name:   "Pond for Get By Id",
		FarmId: 1,
	}
	id, err := suite.repository.Create(newPond)
	suite.NoError(err, "no error while creating pond")
	pond, err := suite.repository.GetById(id)
	suite.NoError(err, "no error while get farm by id")
	suite.Equal(pond.Name, newPond.Name, "farm name should equal")
}

func (suite *PondRepositorySuite) TestGetPondByFarmId_Positive() {
	err := suite.helper.CreateFarm()
	suite.NoError(err, "no error while creating new farm")
	newPond := domain.Pond{
		Name:   "Pond for Get By Id",
		FarmId: 1,
	}
	_, err = suite.repository.Create(newPond)
	suite.NoError(err, "no error while creating pond")
	ponds, err := suite.repository.GetByFarmId(newPond.FarmId)
	suite.NoError(err, "no error while get farm by id")
	suite.NotEmpty(ponds, "should not empty data")
}

func (suite *PondRepositorySuite) TestUpdatePond_Positive() {
	err := suite.helper.CreateFarm()
	suite.NoError(err, "no error while creating new farm")
	newPond := domain.Pond{
		Name:   "Pond for Get By Id",
		FarmId: 1,
	}
	id, err := suite.repository.Create(newPond)
	suite.NoError(err, "no error while creating pond")
	pond, err := suite.repository.GetById(id)
	suite.NoError(err, "no error while get pond by id")
	pond.Name = "Pond for update, updated"
	err = suite.repository.Update(pond)
	suite.NoError(err, "no error while updating data")
}

func (suite *PondRepositorySuite) TestDeletePond_Positive() {
	err := suite.helper.CreateFarm()
	suite.NoError(err, "no error while creating new farm")
	newPond := domain.Pond{
		Name:   "Pond for Update",
		FarmId: 1,
	}
	id, err := suite.repository.Create(newPond)
	suite.NoError(err, "no error while creating pond")
	pond, err := suite.repository.GetById(id)
	suite.NoError(err, "no error while get pond by id")
	err = suite.repository.Delete(pond)
	suite.NoError(err, "no error while deleting data")
}
