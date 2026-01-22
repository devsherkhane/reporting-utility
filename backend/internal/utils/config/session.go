package config

import "github.com/gorilla/sessions"

var Store *sessions.CookieStore

func InitSession() {
	// Using the secret key from AppConfig
	Store = sessions.NewCookieStore([]byte(AppConfig.Session.SecretKey))
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600000,
		HttpOnly: true,
	}
}