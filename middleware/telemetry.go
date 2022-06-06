package middleware

import (
	"go-template/domain"
	"go-template/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TelemetryMiddleware struct {
	telemetryUsecase domain.TelemetryUsecase
}

func NewTelemetryMiddleware(telemetryUsecase domain.TelemetryUsecase) TelemetryMiddleware {
	return TelemetryMiddleware{telemetryUsecase}
}

func (m *TelemetryMiddleware) Telemetry() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()

		req := c.Request.Header.Get("X-Forwarded-For")
		if len(req) == 0 {
			req = c.Request.Header.Get("X-Real-IP")
		}
		if len(req) == 0 {
			req = c.Request.RemoteAddr
		}

		data := domain.CreateTelemetryDto{
			Ip:       req,
			Method:   c.Request.Method,
			Endpoint: c.Request.RequestURI,
			Status:   c.Writer.Status(),
			Latency:  time.Since(t).Milliseconds(),
		}

		if err := m.telemetryUsecase.Create(data); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
			return
		}
	}
}
