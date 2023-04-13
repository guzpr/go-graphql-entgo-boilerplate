package gqlserver

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/sekalahita/epirus/internal/middleware"
)

func (s *Server) RegisterRouter() {
	r := func(method, path string, handler http.Handler) {
		s.Router.Method(method, path, handler)
	}

	r(http.MethodGet, "/", playground.Handler("GraphQL playground", "/query"))
	r(http.MethodPost, "/query", s.authMiddleware(s.gqlSrv))
	r(http.MethodGet, "/login/start", middleware.ErrorMiddleware(s.AuthHandler.LoginStartHandler))
	r(http.MethodGet, "/login/callback", middleware.ErrorMiddleware(s.AuthHandler.LoginCallbackHandler))
}
