package apis

import (
	"github.com/QuizWars-Ecosystem/admin-panel/internal/render"
	"github.com/QuizWars-Ecosystem/admin-panel/internal/sessions"
	"github.com/QuizWars-Ecosystem/admin-panel/ui/pages"
	gorilla "github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

type Handler struct {
	r      *render.Render
	store  *gorilla.CookieStore
	logger *zap.Logger
}

func NewHandler(store *gorilla.CookieStore, logger *zap.Logger) *Handler {
	return &Handler{
		r:      &render.Render{},
		store:  store,
		logger: logger,
	}
}

func (h *Handler) MainPage(ctx echo.Context) error {
	session, _ := h.store.Get(ctx.Request(), sessions.AdminSessionName)
	auth, _ := session.Values[sessions.IsAuthenticatedName].(bool)

	return h.r.Render(ctx, http.StatusOK, pages.MainPage(auth))
}

func (h *Handler) LoginPage(ctx echo.Context) error {
	return h.r.Render(ctx, http.StatusOK, pages.LoginPage())
}
