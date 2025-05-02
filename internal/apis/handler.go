package apis

import (
	"context"
	"github.com/QuizWars-Ecosystem/admin-panel/internal/render"
	"github.com/QuizWars-Ecosystem/admin-panel/internal/sessions"
	"github.com/QuizWars-Ecosystem/admin-panel/ui/pages"
	gorilla "github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Handler struct {
	r       *render.Render
	store   *gorilla.CookieStore
	service *Service
	logger  *zap.Logger
}

func NewHandler(service *Service, store *gorilla.CookieStore, logger *zap.Logger) *Handler {
	return &Handler{
		r:       &render.Render{},
		service: service,
		store:   store,
		logger:  logger,
	}
}

func (h *Handler) MainPage(c echo.Context) error {
	session, _ := h.store.Get(c.Request(), sessions.AdminSessionName)
	auth, _ := session.Values[sessions.IsAuthenticatedName].(bool)

	return h.r.Render(c, http.StatusOK, pages.MainPage(auth))
}

func (h *Handler) LoginPage(c echo.Context) error {
	return h.r.Render(c, http.StatusOK, pages.LoginPage())
}

func (h *Handler) LoginSubmitForm(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	requestCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	token, err := h.service.Login(requestCtx, email, password)
	if err != nil {
		h.logger.Error("Login failed", zap.Error(err))
		return c.Redirect(http.StatusFound, "/login")
	}

	session, _ := h.store.Get(c.Request(), sessions.AdminSessionName)
	session.Values[sessions.TokenSessionName] = token

	return c.String(http.StatusOK, "Login received")
}
