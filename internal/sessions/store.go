package sessions

import (
	"github.com/gorilla/sessions"
)

const (
	AdminSessionName    = "admin-session"
	TokenSessionName    = "auth_token"
	IsAuthenticatedName = "is_authenticated"
)

var Store *sessions.CookieStore

func NewStore() *sessions.CookieStore {
	Store = sessions.NewCookieStore()
	return Store
}
