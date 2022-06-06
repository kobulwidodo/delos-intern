package http

import (
	"go-template/domain"
	"go-template/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PondHandler struct {
	PondUsecase domain.PondUsecase
}

func NewPondHandler(r *gin.RouterGroup, PondUsecase domain.PondUsecase) {
	handler := &PondHandler{PondUsecase}
	api := r.Group("/pond")
	{
		api.POST("/:id", handler.Create)
		api.GET("/", handler.GetAll)
		api.GET("/:id", handler.GetById)
		api.GET("/farm/:id", handler.GetByFarmId)
		api.PUT("/:id", handler.Update)
		api.DELETE("/:id", handler.Delete)
	}
}

func (h *PondHandler) Create(c *gin.Context) {
	var input domain.CreatePondDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	var uri domain.FarmIdUriBinding
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	input.FarmId = uri.Id
	id, err := h.PondUsecase.Create(input)
	if err != nil {
		c.JSON(utils.GetErrorCode(err), utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.NewSuccessResponse("successfully created new pond", gin.H{"id": id}))
}

func (h *PondHandler) GetAll(c *gin.Context) {
	ponds, err := h.PondUsecase.GetAll()
	if err != nil {
		c.JSON(utils.GetErrorCode(err), utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully get all ponds", ponds))
}

func (h *PondHandler) GetById(c *gin.Context) {
	var uri domain.PondIdUriBinding
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	pond, err := h.PondUsecase.GetById(uri.Id)
	if err != nil {
		c.JSON(utils.GetErrorCode(err), utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully get pond", pond))
}

func (h *PondHandler) GetByFarmId(c *gin.Context) {
	var uri domain.FarmIdUriBinding
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	ponds, err := h.PondUsecase.GetByFarmId(uri.Id)
	if err != nil {
		c.JSON(utils.GetErrorCode(err), utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully get all ponds by farm id", ponds))
}

func (h *PondHandler) Update(c *gin.Context) {
	var input domain.UpdatePondDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	var uri domain.PondIdUriBinding
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	input.Id = uri.Id
	if err := h.PondUsecase.Update(input); err != nil {
		c.JSON(utils.GetErrorCode(err), utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully updated pond", nil))
}

func (h *PondHandler) Delete(c *gin.Context) {
	var uri domain.PondIdUriBinding
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	if err := h.PondUsecase.Delete(uri.Id); err != nil {
		c.JSON(utils.GetErrorCode(err), utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully deleted pond", nil))
}
