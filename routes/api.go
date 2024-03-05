package routes

import (
	handlers "github.com/kave08/url-shortener/handler"
	"github.com/labstack/echo/v4"
)

func InitializeGroup(e *echo.Echo, handler *handlers.Handlers) {
	api := e.Group("/api/v1")

	api.GET("/:shortUrl", handler.ResolvUrlHandler())
	api.POST("/shorten", handler.ShortenUrlHandler())
}
