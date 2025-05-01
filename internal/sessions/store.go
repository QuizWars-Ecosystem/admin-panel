package sessions

import (
	"github.com/gorilla/sessions"
)

var Store *sessions.CookieStore

func NewStore() {
	Store = sessions.NewCookieStore()
}
