package auth

import (
	"LinkTer-api/config"
	"net/http"
	"time"
)

type Auth struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Auth {
	return &Auth{
		cfg: cfg,
	}
}

func (a *Auth) GenCookie(name string, value string, exp string) *http.Cookie {
	t, _ := time.ParseDuration(exp)
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.HttpOnly = true
	cookie.Expires = time.Now().Add(t)
	cookie.Secure = true
	return cookie
}
