package render

import (
	"github.com/QuizWars-Ecosystem/admin-panel/internal/sessions"
	"github.com/QuizWars-Ecosystem/admin-panel/ui/layouts"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type Render struct{}

func (r *Render) Render(ctx echo.Context, statusCode int, contents ...templ.Component) error {
	session, _ := sessions.Store.Get(ctx.Request(), "admin-session")
	auth, _ := session.Values["is_authenticated"].(bool)

	page := layouts.BaseLayout(auth, contents...)

	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := page.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}
