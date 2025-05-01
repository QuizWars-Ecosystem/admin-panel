package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/DavidMovas/gopherbox/pkg/closer"
	"github.com/QuizWars-Ecosystem/admin-panel/assets"
	"github.com/QuizWars-Ecosystem/admin-panel/internal/apis"
	"github.com/QuizWars-Ecosystem/admin-panel/internal/config"
	"github.com/QuizWars-Ecosystem/admin-panel/internal/sessions"
	"github.com/QuizWars-Ecosystem/go-common/pkg/abstractions"
	"github.com/QuizWars-Ecosystem/go-common/pkg/jwt"
	"github.com/QuizWars-Ecosystem/go-common/pkg/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"net/http"
)

var _ abstractions.Server = (*Server)(nil)

type Server struct {
	e      *echo.Echo
	logger *log.Logger
	cfg    *config.Config
	closer *closer.Closer
}

func NewServer(_ context.Context, cfg *config.Config) (*Server, error) {
	cl := closer.NewCloser()

	logger := log.NewLogger(cfg.Local, cfg.LogLevel)
	cl.PushIO(logger)

	sessions.NewStore()

	authService := jwt.NewService(&jwt.Config{
		Secret:            cfg.JWT.Secret,
		AccessExpiration:  cfg.JWT.AccessTimeout,
		RefreshExpiration: cfg.JWT.RefreshTimeout,
	})

	_ = authService

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.HTTPErrorHandler = func(err error, c echo.Context) {}

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	h := apis.NewHandler(logger.Zap())

	e.StaticFS("/assets", assets.Assets)

	e.GET("/login", nil)
	e.POST("/login", nil)

	authGroup := e.Group("")
	//authGroup.Use(middlewares.NewAuthMiddleware(authService))
	authGroup.GET("/", h.MainPage)

	return &Server{
		e:      e,
		logger: logger,
		cfg:    cfg,
		closer: cl,
	}, nil
}

func (s *Server) Start() error {
	z := s.logger.Zap()

	z.Info("Starting server", zap.Int("port", s.cfg.Port))

	if err := s.e.Start(fmt.Sprintf(":%d", s.cfg.Port)); err != nil {
		z.Error("Failed to start server", zap.Error(err))
		return err
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	z := s.logger.Zap()

	z.Info("Shutting down server")

	if err := s.e.Shutdown(ctx); !errors.Is(err, http.ErrServerClosed) {
		z.Error("Failed to shutdown server", zap.Error(err))
		return err
	}

	return s.closer.Close(ctx)
}
