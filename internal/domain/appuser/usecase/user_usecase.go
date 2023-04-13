package usecase

import (
	"context"

	"github.com/sekalahita/epirus/internal/db"
	du "github.com/sekalahita/epirus/internal/domain/appuser"
	ur "github.com/sekalahita/epirus/internal/domain/appuser/repo"
	"github.com/sekalahita/epirus/internal/ent/gen"
	"github.com/sekalahita/epirus/internal/errors"
	"github.com/sekalahita/epirus/internal/pagination"
)

type UserUseCase struct {
	newUserRepo ur.UserRepoFactory
	client      *gen.Client
}

func NewUserUseCase(entClient *gen.Client) UserUseCase {
	return UserUseCase{
		newUserRepo: ur.NewUserRepoFactory(entClient),
		client:      entClient,
	}
}

type LoginOrSignUpParam struct {
	GoogleID string
	Email    string
}

func (u UserUseCase) LoginOrSignUp(ctx context.Context, param LoginOrSignUpParam) (du.User, error) {
	r := u.newUserRepo(u.client)

	usr, err := r.GetByGoogleAuthID(ctx, param.GoogleID)
	if err != nil {
		if errors.IsNotFoundError(err) {
			err := db.WithTx(ctx, r.EntClient, func(tx *gen.Tx) error {
				r.WithTx(tx.Client())

				usr, err = r.Create(ctx, du.User{
					Email:            param.Email,
					OnboardingStatus: du.OnboardingStatusCreateBudgetPeriod,
				})
				if err != nil {
					return errors.ErrorWithCurrentFuncName(err)
				}

				_, err = r.CreateGoogleAuth(ctx, usr, du.GoogleAuth{
					GoogleID: param.GoogleID,
				})
				if err != nil {
					return errors.ErrorWithCurrentFuncName(err)
				}
				return nil
			})

			if err != nil {
				return du.User{}, errors.ErrorWithCurrentFuncName(err)
			}
		} else {
			return du.User{}, errors.ErrorWithCurrentFuncName(err)
		}
	}

	return usr, nil
}

type GetByIDParam struct {
	ID string
}

func (u UserUseCase) GetByID(ctx context.Context, param GetByIDParam) (du.User, error) {
	repo := u.newUserRepo(u.client)

	usr, err := repo.GetByID(ctx, param.ID)
	if err != nil {
		if errors.IsNotFoundError(err) {
			return du.User{}, du.NewErrorUserNotFound(err)
		}

		return du.User{}, errors.NewErrorInternalServer(err)
	}

	return usr, nil
}

type GetAllPaginationParam struct {
	pagination.CursorPagination
}

func (u UserUseCase) GetAllPagination(ctx context.Context, param GetAllPaginationParam) (*gen.UserConnection, error) {
	r := u.newUserRepo(u.client)

	return r.GetAllPagination(ctx, ur.GetAllPaginationParam(param))
}
