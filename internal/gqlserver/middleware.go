package gqlserver

import (
	"encoding/json"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/sekalahita/epirus/internal/auth"
	"github.com/sekalahita/epirus/internal/errors"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (s *Server) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authKey := r.Header["Authentication"]
		if len(authKey) == 0 {
			err := auth.NewErrorAuthTokenNotFound(
				errors.ErrorWithCurrentFuncName(
					errors.New("auth header is empty"),
				))

			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(graphql.Response{
				Errors: gqlerror.List{
					&gqlerror.Error{
						Message: err.Message,
					},
				},
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}
