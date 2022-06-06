package http

import (
	"go-template/domain"
	"go-template/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TelemetryHandler struct {
	telemetryUsecase domain.TelemetryUsecase
}

func NewTelemetryHandler(r *gin.RouterGroup, telemetryUsecase domain.TelemetryUsecase) {
	handler := &TelemetryHandler{telemetryUsecase}
	api := r.Group("/telemetry")
	{
		api.GET("/", handler.GetTelemetry)
	}
}

func (h *TelemetryHandler) GetTelemetry(c *gin.Context) {
	telemetries, err := h.telemetryUsecase.GetCount()
	if err != nil {
		c.JSON(utils.GetErrorCode(err), utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully get telemetries", telemetries))
}
