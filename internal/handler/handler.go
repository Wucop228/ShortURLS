package handler

import (
	"database/sql"
	"github.com/Wucop228/ShortURLS/internal/cache"
	"github.com/Wucop228/ShortURLS/internal/service"
	"net/http"

	"github.com/Wucop228/ShortURLS/internal/model"
	"github.com/Wucop228/ShortURLS/pkg/shortener"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	db  *sql.DB
	rdb *cache.RedisCache
}

func NewHandlers(db *sql.DB, rdb *cache.RedisCache) *Handlers {
	return &Handlers{db: db, rdb: rdb}
}

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (h *Handlers) Redirect(c echo.Context) error {
	key := c.Param("key")
	if key == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Missing key"})
	}

	baseURL, err := service.GetOriginalURL(h.db, h.rdb, key)
	if err != nil {
		return c.JSON(http.StatusNotFound, ErrorResponse{Error: "URL not found"})
	}

	return c.Redirect(http.StatusFound, baseURL)
}

func (h *Handlers) Shorten(c echo.Context) error {
	var req ShortenRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
	}

	id, err := model.GetLastID(h.db)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "DB error"})
	}

	shortURL := shortener.Generate(id + 1)
	url := model.URL{BaseUrl: req.URL, ShortUrl: shortURL}

	if err := service.SetOriginalURL(h.db, h.rdb, url); err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "DB error"})
	}

	return c.JSON(http.StatusCreated, ShortenResponse{
		OriginalURL: url.BaseUrl,
		ShortURL:    url.ShortUrl,
	})
}
