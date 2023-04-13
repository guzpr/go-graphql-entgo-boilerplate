package gqlserver

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi/v5"
	"github.com/sekalahita/epirus/internal/auth"
	"github.com/sekalahita/epirus/internal/db"
	usr "github.com/sekalahita/epirus/internal/domain/appuser/usecase"
	"github.com/sekalahita/epirus/internal/ent/gen"
	"github.com/sekalahita/epirus/internal/env"
	"github.com/sekalahita/epirus/internal/gql"
)

type Server struct {
	ec          *gen.Client
	gqlSrv      *handler.Server
	AuthHandler auth.AuthHandler
	Router      *chi.Mux
}

func NewServer() *Server {
	env.InitDotenv()

	ec := db.CreateEntClient()

	// TODO move to separate process
	if err := ec.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	authHandler := auth.AuthHandler{
		Service: usr.NewUserUseCase(ec),
	}

	gqlSrv := handler.NewDefaultServer(gql.NewSchema(ec))

	server := &Server{
		ec:          ec,
		AuthHandler: authHandler,
		gqlSrv:      gqlSrv,
		Router:      chi.NewMux(),
	}

	server.RegisterRouter()

	return server
}

func (s *Server) Run(ctx context.Context, port int) {
	httpS := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: s.Router,
	}

	log.Printf("gqlserver run on port %s", httpS.Addr)

	go func() {
		if err := httpS.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen")
		}
	}()

	<-ctx.Done()

	log.Printf("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		cancel()
	}()

	err := httpS.Shutdown(ctxShutDown)
	if err != nil {
		log.Fatal("server Shutdown Failed")
	}

	log.Printf("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	err = s.ec.Close()
	if err != nil {
		log.Fatal("unable to close db")
	}
}
