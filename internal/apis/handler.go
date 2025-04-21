package apis

import (
	"github.com/QuizWars-Ecosystem/admin-panel/internal/render"
	"github.com/QuizWars-Ecosystem/admin-panel/ui/layouts"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

type Handler struct {
	r      *render.Render
	logger *zap.Logger
}

func NewHandler(logger *zap.Logger) *Handler {
	return &Handler{
		r:      &render.Render{},
		logger: logger,
	}
}

func (h *Handler) MainPage(ctx echo.Context) error {
	return h.r.Render(ctx, http.StatusOK, layouts.BaseLayout())
}
