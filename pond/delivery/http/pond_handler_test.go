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

	_pondHttpHandler "go-template/pond/delivery/http"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type PondHandlerSuite struct {
	suite.Suite
	usecase       *mocks.PondUsecase
	testingServer *httptest.Server
}

func TestPondHandler(t *testing.T) {
	suite.Run(t, new(PondHandlerSuite))
}

func (suite *PondHandlerSuite) SetupSuite() {
	usecase := new(mocks.PondUsecase)
	handler := _pondHttpHandler.PondHandler{
		PondUsecase: usecase,
	}
	r := gin.Default()
	r.POST("/api/pond/:id", handler.Create)
	r.GET("/api/pond", handler.GetAll)
	r.GET("/api/pond/:id", handler.GetById)
	r.GET("/api/pond/farm/:id", handler.GetByFarmId)
	r.PUT("/api/pond/:id", handler.Update)
	r.DELETE("/api/pond/:id", handler.Delete)
	suite.testingServer = httptest.NewServer(r)
	suite.usecase = usecase
}

func (suite *PondHandlerSuite) TearDownSuite() {
	defer suite.testingServer.Close()
}

func (suite *PondHandlerSuite) TestCreate_Positive() {
	id := uint(1)
	inputDto := domain.CreatePondDto{
		Name:   "Pond Testing",
		FarmId: id,
	}
	suite.usecase.On("Create", inputDto).Return(uint(1), nil)
	reqBody, err := json.Marshal(&inputDto)
	suite.NoError(err, "no error while marshal struct to json")
	res, err := http.Post(fmt.Sprintf("%s/api/pond/%d", suite.testingServer.URL, id), "application/json", bytes.NewBuffer(reqBody))
	suite.NoError(err, "no error while calling endpoint")
	defer res.Body.Close()
	resBody := utils.Response{}
	json.NewDecoder(res.Body).Decode(&resBody)
	suite.Equal(http.StatusCreated, res.StatusCode)
	suite.Equal(resBody.IsSuccess, true)
	suite.usecase.AssertExpectations(suite.T())
}

func (suite *PondHandlerSuite) TestGetAll_Positive() {
	ponds := []domain.Pond{
		{
			Name:   "Pond 1",
			FarmId: 1,
		},
		{
			Name:   "Pond 2",
			FarmId: 1,
		},
		{
			Name:   "Pond 3",
			FarmId: 1,
		},
	}
	suite.usecase.On("GetAll").Return(ponds, nil)
	res, err := http.Get(fmt.Sprintf("%s/api/pond", suite.testingServer.URL))
	suite.NoError(err, "no error while calling endpoint")
	defer res.Body.Close()
	resBody := utils.Response{}
	json.NewDecoder(res.Body).Decode(&resBody)
	suite.Equal(http.StatusOK, res.StatusCode)
	suite.Equal(resBody.IsSuccess, true)
	suite.usecase.AssertExpectations(suite.T())
}

func (suite *PondHandlerSuite) TestGetById_Positive() {
	id := uint(1)
	pond := domain.Pond{
		Name:   "Pond Testing",
		FarmId: 1,
	}
	pond.ID = id
	suite.usecase.On("GetById", id).Return(pond, nil)
	res, err := http.Get(fmt.Sprintf("%s/api/pond/%d", suite.testingServer.URL, id))
	suite.NoError(err, "no error while calling endpoint")
	defer res.Body.Close()
	resBody := utils.Response{}
	json.NewDecoder(res.Body).Decode(&resBody)
	suite.Equal(http.StatusOK, res.StatusCode)
	suite.Equal(resBody.IsSuccess, true)
	suite.usecase.AssertExpectations(suite.T())
}

func (suite *PondHandlerSuite) TestGetByFarmId_Positive() {
	farmId := uint(1)
	ponds := []domain.Pond{
		{
			Name:   "Pond Testing",
			FarmId: farmId,
		},
	}
	suite.usecase.On("GetByFarmId", farmId).Return(ponds, nil)
	res, err := http.Get(fmt.Sprintf("%s/api/pond/farm/%d", suite.testingServer.URL, farmId))
	suite.NoError(err, "no error while calling endpoint")
	defer res.Body.Close()
	resBody := utils.Response{}
	json.NewDecoder(res.Body).Decode(&resBody)
	suite.Equal(http.StatusOK, res.StatusCode)
	suite.Equal(resBody.IsSuccess, true)
	suite.usecase.AssertExpectations(suite.T())
}

func (suite *PondHandlerSuite) TestUpdate_Positive() {
	id := uint(1)
	inputDto := domain.UpdatePondDto{
		Name: "Pond Testing",
		Id:   id,
	}
	suite.usecase.On("Update", inputDto).Return(nil)
	reqBody, err := json.Marshal(&inputDto)
	suite.NoError(err, "no error while marshal struct to json")
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/api/pond/%d", suite.testingServer.URL, id), bytes.NewBuffer(reqBody))
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

func (suite *PondHandlerSuite) TestDelete_Positive() {
	id := uint(1)
	suite.usecase.On("Delete", id).Return(nil)
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/pond/%d", suite.testingServer.URL, id), nil)
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
