package http

import (
	"go-template/domain"
	"go-template/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FarmHandler struct {
	FarmUsecase domain.FarmUsecase
}

func NewFarmHandler(r *gin.RouterGroup, farmUsecase domain.FarmUsecase) {
	handler := &FarmHandler{farmUsecase}
	api := r.Group("/farm")
	{
		api.POST("/", handler.Create)
		api.GET("/", handler.GetAll)
		api.GET("/:id", handler.GetById)
		api.PUT("/:id", handler.Update)
		api.DELETE("/:id", handler.Delete)
	}
}

func (h *FarmHandler) Create(c *gin.Context) {
	var input domain.CreateFarmDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	id, err := h.FarmUsecase.Create(input)
	if err != nil {
		c.JSON(utils.GetErrorCode(err), utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.NewSuccessResponse("successfully created new farm", gin.H{"id": id}))
}

func (h *FarmHandler) GetAll(c *gin.Context) {
	farms, err := h.FarmUsecase.GetAll()
	if err != nil {
		c.JSON(utils.GetErrorCode(err), utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully get all farms", farms))
}

func (h *FarmHandler) GetById(c *gin.Context) {
	var uri domain.FarmIdUriBinding
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadGateway, utils.NewFailResponse(err.Error()))
		return
	}
	farm, err := h.FarmUsecase.GetById(uri.Id)
	if err != nil {
		c.JSON(utils.GetErrorCode(err), utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully get farm", farm))
}

func (h *FarmHandler) Update(c *gin.Context) {
	var input domain.UpdateFarmDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	var uri domain.FarmIdUriBinding
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	input.Id = uri.Id
	if err := h.FarmUsecase.Update(input); err != nil {
		c.JSON(utils.GetErrorCode(err), utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully updated farm", nil))
}

func (h *FarmHandler) Delete(c *gin.Context) {
	var uri domain.FarmIdUriBinding
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	if err := h.FarmUsecase.Delete(uri.Id); err != nil {
		c.JSON(utils.GetErrorCode(err), utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully deleted farm", nil))
}
