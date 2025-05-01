package middlewares

import (
	"github.com/QuizWars-Ecosystem/admin-panel/internal/sessions"
	"github.com/QuizWars-Ecosystem/go-common/pkg/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func NewAuthMiddleware(service *jwt.Service) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			session, err := sessions.Store.Get(c.Request(), sessions.AdminSessionName)
			if err != nil {
				return c.Redirect(http.StatusSeeOther, "/login")
			}

			token, ok := session.Values[sessions.TokenSessionName].(string)
			if !ok || token == "" || !isValidJWT(service, token) {
				return c.Redirect(http.StatusSeeOther, "/login")
			}

			return next(c)
		}
	}
}

func isValidJWT(service *jwt.Service, token string) bool {
	superErr := service.ValidateRoleToken(token, string(jwt.Super))
	adminErr := service.ValidateRoleToken(token, string(jwt.Admin))

	if superErr != nil && adminErr != nil {
		return false
	}

	return true
}
