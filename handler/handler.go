package handler

import (
	"net/http"

	"github.com/kave08/url-shortener/config"
	"github.com/kave08/url-shortener/models"
	"github.com/kave08/url-shortener/service"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	lgc *service.Logics
}

func NewHandlers(lgc *service.Logics) *Handlers {
	return &Handlers{
		lgc: lgc,
	}
}

func (h *Handlers) ShortenUrlHandler() func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		var newBody = new(models.Request)
		ctx.Bind(newBody)

		if newBody.ApiKey != config.Cfg.ApiToken.Token {
			return ctx.JSON(http.StatusForbidden, map[string]string{"error": "Token is not valid"})
		}

		shortUrl, err := h.lgc.ShortenUrl(ctx.Request().Context(), newBody.LongUrl)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		newResponse := &models.Response{
			ShortUrl: shortUrl,
		}

		return ctx.JSON(http.StatusOK, newResponse)
	}
}

func (h *Handlers) ResolvUrlHandler() func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		shortUrl := ctx.Param("shortUrl")

		longUrl, err := h.lgc.ResolveUrl(ctx.Request().Context(), shortUrl)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return ctx.Redirect(302, longUrl)
	}
}
