package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthController struct {
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

// Health
//
//	@Summary		检查服务健康状态
//	@Description	返回服务是否健康的状态信息
//	@Tags			健康检查
//	@Success		200	{object}	string	"ok"
//	@Security		BearerToken
//	@Router			/health [GET]
func (r *HealthController) Health(c echo.Context) (err error) {
	return c.String(http.StatusOK, "ok")
}
