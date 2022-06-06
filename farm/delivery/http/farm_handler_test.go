package http_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-template/domain"
	"go-template/domain/mocks"
	"go-template/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	_farmHttpHandler "go-template/farm/delivery/http"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type FarmHandlerSuite struct {
	suite.Suite
	usecase       *mocks.FarmUsecase
	testingServer *httptest.Server
}

func TestFarmHandler(t *testing.T) {
	suite.Run(t, new(FarmHandlerSuite))
}

func (suite *FarmHandlerSuite) SetupSuite() {
	usecase := new(mocks.FarmUsecase)
	handler := _farmHttpHandler.FarmHandler{
		FarmUsecase: usecase,
	}
	r := gin.Default()
	r.POST("/api/farm", handler.Create)
	r.GET("/api/farm", handler.GetAll)
	r.GET("/api/farm/:id", handler.GetById)
	r.PUT("/api/farm/:id", handler.Update)
	r.DELETE("/api/farm/:id", handler.Delete)
	suite.testingServer = httptest.NewServer(r)
	suite.usecase = usecase
}

func (suite *FarmHandlerSuite) TearDownSuite() {
	defer suite.testingServer.Close()
}

func (suite *FarmHandlerSuite) TestCreate_Positive() {
	inputDto := domain.CreateFarmDto{
		Name: "Farm Testing",
	}
	suite.usecase.On("Create", inputDto).Return(uint(1), nil)
	reqBody, err := json.Marshal(&inputDto)
	suite.NoError(err, "no error while marshal struct to json")
	res, err := http.Post(fmt.Sprintf("%s/api/farm/", suite.testingServer.URL), "application/json", bytes.NewBuffer(reqBody))
	suite.NoError(err, "no error while calling endpoint")
	defer res.Body.Close()
	resBody := utils.Response{}
	json.NewDecoder(res.Body).Decode(&resBody)
	suite.Equal(http.StatusCreated, res.StatusCode)
	suite.Equal(resBody.IsSuccess, true)
	suite.usecase.AssertExpectations(suite.T())
}

func (suite *FarmHandlerSuite) TestGetAll_Positive() {
	farms := []domain.Farm{
		{
			Name: "Farm 1",
		},
		{
			Name: "Farm 2",
		},
		{
			Name: "Farm 3",
		},
	}
	suite.usecase.On("GetAll").Return(farms, nil)
	res, err := http.Get(fmt.Sprintf("%s/api/farm", suite.testingServer.URL))
	suite.NoError(err, "no error while calling endpoint")
	defer res.Body.Close()
	resBody := utils.Response{}
	json.NewDecoder(res.Body).Decode(&resBody)
	suite.Equal(http.StatusOK, res.StatusCode)
	suite.Equal(resBody.IsSuccess, true)
	suite.usecase.AssertExpectations(suite.T())
}

func (suite *FarmHandlerSuite) TestGetById_Positive() {
	id := uint(1)
	farm := domain.Farm{
		Name: "Farm Testing",
	}
	farm.ID = id
	suite.usecase.On("GetById", id).Return(farm, nil)
	res, err := http.Get(fmt.Sprintf("%s/api/farm/%d", suite.testingServer.URL, id))
	suite.NoError(err, "no error while calling endpoint")
	defer res.Body.Close()
	resBody := utils.Response{}
	json.NewDecoder(res.Body).Decode(&resBody)
	suite.Equal(http.StatusOK, res.StatusCode)
	suite.Equal(resBody.IsSuccess, true)
	suite.usecase.AssertExpectations(suite.T())
}

func (suite *FarmHandlerSuite) TestUpdate_Positive() {
	id := uint(1)
	inputDto := domain.UpdateFarmDto{
		Name: "Farm Testing",
		Id:   id,
	}
	suite.usecase.On("Update", inputDto).Return(nil)
	reqBody, err := json.Marshal(&inputDto)
	suite.NoError(err, "no error while marshal struct to json")
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/api/farm/%d", suite.testingServer.URL, id), bytes.NewBuffer(reqBody))
	suite.NoError(err, "no error while setup request")
	req.Header.Set("Content-Type", "application/json")
	cl := &http.Client{}
	res, err := cl.Do(req)
	suite.NoError(err, "no error while calling endpoint")
	defer res.Body.Close()
	resBody := utils.Response{}
	json.NewDecoder(res.Body).Decode(&resBody)
	suite.Equal(http.StatusOK, res.StatusCode)
	suite.Equal(resBody.IsSuccess, true)
	suite.usecase.AssertExpectations(suite.T())
}

func (suite *FarmHandlerSuite) TestDelete_Positive() {
	id := uint(1)
	suite.usecase.On("Delete", id).Return(nil)
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/farm/%d", suite.testingServer.URL, id), nil)
	suite.NoError(err, "no error while calling endpoint")
	cl := &http.Client{}
	res, err := cl.Do(req)
	suite.NoError(err, "no error while calling endpoint")
	defer res.Body.Close()
	resBody := utils.Response{}
	json.NewDecoder(res.Body).Decode(&resBody)
	suite.Equal(http.StatusOK, res.StatusCode)
	suite.Equal(resBody.IsSuccess, true)
	suite.usecase.AssertExpectations(suite.T())
}
