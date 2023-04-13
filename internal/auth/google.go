package auth

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GoogleOauthConfig(r *http.Request) *oauth2.Config {
	scheme := "http"
	env := os.Getenv("APP_ENV")
	if r.TLS != nil ||
		r.URL.Scheme == "https" ||
		r.Header.Get("X-Forwarded-Proto") == "https" ||
		env == "production" ||
		env == "staging" {
		scheme = "https"
	}

	redirectHost := fmt.Sprintf("%s://%s", scheme, r.Host)

	return &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  redirectHost + "/login/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
	}
}
