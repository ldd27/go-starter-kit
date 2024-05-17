package router

import (
	"github.com/labstack/echo/v4"
	"github.com/ldd27/go-starter-kit/internal/router/api/v1"
	"github.com/ldd27/go-starter-kit/internal/wire"
)

// InitRouter
//
//	@title						go-start-kit
//	@version					1.0
//	@description				go-start-kit
//	@Accept						json
//	@Produce					json
//	@host						localhost:8080
//	@BasePath					/api/v1
//	@servers					[{"url":"http://localhost:8080/api/v1"},{"url":"https://api.example.com/api/v1"}]
//	@schemes					http https
//	@securityDefinitions.apikey	BearerToken
//	@in							header
//	@name						Authorization
//	@description				JWT授权 格式：Bearer {token} 即可，注意两者之间有空格
func InitRouter(engine *echo.Echo) {
	v1Group := engine.Group("api/v1")

	// health
	{
		health := v1.NewHealthController()
		v1Group.GET("/health", health.Health)
	}
	// example
	{
		example := wire.NewExampleController()
		v1Group.GET("/example", example.GetList)
		v1Group.GET("/example/page", example.GetPageList)
		v1Group.GET("/example/cursor", example.GetCursorPageList)
		v1Group.POST("/example", example.Create)
		v1Group.PUT("/example", example.Update)
		v1Group.DELETE("/example", example.Delete)
	}
}
