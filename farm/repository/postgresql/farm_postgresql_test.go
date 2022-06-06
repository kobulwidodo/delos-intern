package postgresql_test

import (
	"fmt"
	"go-template/domain"
	"go-template/infrastructure"
	"go-template/utils"
	"log"
	"testing"

	_farmRepository "go-template/farm/repository/postgresql"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
)

type FarmRepositorySuite struct {
	suite.Suite
	repository domain.FarmRepository
	helper     utils.TestingHelper
}

func TestFarmRepository(t *testing.T) {
	if err := godotenv.Load("../../../.env"); err != nil {
		fmt.Println("failed to load env")
	}
	suite.Run(t, new(FarmRepositorySuite))
}

func (suite *FarmRepositorySuite) SetupSuite() {
	dbConfig := infrastructure.NewDbConfig()
	db, err := dbConfig.InitDb()
	if err != nil {
		log.Fatal("failed to connect with database")
		panic(err)
	}
	repository := _farmRepository.NewFarmRepository(db)
	suite.repository = repository
	suite.helper = utils.NewTestingHelper(db)
}

func (suite *FarmRepositorySuite) TearDownTest() {
	defer suite.helper.TruncateTable("farms")
}

func (suite *FarmRepositorySuite) TestCreateFarm_Positive() {
	farm := domain.Farm{
		Name: "Farm Test",
	}
	_, err := suite.repository.Create(farm)
	suite.NoError(err, "no error while creating new farm")
}

func (suite *FarmRepositorySuite) TestGetAllFarm_Positive() {
	farm := domain.Farm{
		Name: "Farm Test Get",
	}
	_, err := suite.repository.Create(farm)
	suite.NoError(err, "no error while creating new farm")
	farms, err := suite.repository.GetAll()
	suite.NoError(err, "no error while get all farms")
	suite.NotEmpty(farms, "should not empty data")
}

func (suite *FarmRepositorySuite) TestGetFarmById_Positive() {
	newFarm := domain.Farm{
		Name: "Farm for Get By Id",
	}
	id, err := suite.repository.Create(newFarm)
	suite.NoError(err, "no error while creating farm")
	farm, err := suite.repository.GetById(id)
	suite.NoError(err, "no error while get farm by id")
	suite.Equal(farm.Name, newFarm.Name, "farm name should equal")
}

func (suite *FarmRepositorySuite) TestUpdateFarm_Positive() {
	newFarm := domain.Farm{
		Name: "Farm for Update",
	}
	id, err := suite.repository.Create(newFarm)
	suite.NoError(err, "no error while creating farm")
	farm, err := suite.repository.GetById(id)
	suite.NoError(err, "no error while get farm by id")
	farm.Name = "Farm for update, updated"
	err = suite.repository.Update(farm)
	suite.NoError(err, "no error while updating data")
}

func (suite *FarmRepositorySuite) TestDeleteFarm_Positive() {
	newFarm := domain.Farm{
		Name: "Farm for Update",
	}
	id, err := suite.repository.Create(newFarm)
	suite.NoError(err, "no error while creating farm")
	farm, err := suite.repository.GetById(id)
	suite.NoError(err, "no error while get farm by id")
	err = suite.repository.Delete(farm)
	suite.NoError(err, "no error while deleting data")
}
