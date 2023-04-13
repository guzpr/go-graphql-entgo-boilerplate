package gql

import (
	"context"

	"github.com/sekalahita/epirus/internal/domain/appuser/usecase"
	"github.com/sekalahita/epirus/internal/ent/gen"
	"github.com/sekalahita/epirus/internal/pagination"
)

type Resolver struct {
	useCase usecase.UserUseCase
}

func NewResolver(entClient *gen.Client) Resolver {
	return Resolver{
		useCase: usecase.NewUserUseCase(entClient),
	}
}

type GetAllPaginationParam struct {
	pagination.CursorPagination
}

func (r Resolver) UserQuery(ctx context.Context, param GetAllPaginationParam) (*gen.UserConnection, error) {
	return r.useCase.GetAllPagination(ctx, usecase.GetAllPaginationParam(param))
}
