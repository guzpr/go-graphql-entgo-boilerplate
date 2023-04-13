package repo

import (
	"context"

	du "github.com/sekalahita/epirus/internal/domain/appuser"
	"github.com/sekalahita/epirus/internal/ent/gen"
	"github.com/sekalahita/epirus/internal/ent/gen/googleauth"
	"github.com/sekalahita/epirus/internal/ent/gen/user"
	"github.com/sekalahita/epirus/internal/errors"
	"github.com/sekalahita/epirus/internal/pagination"
)

type userRepo struct {
	EntClient *gen.Client
}

type UserRepoFactory func(*gen.Client) userRepo

func NewUserRepoFactory(entClient *gen.Client) UserRepoFactory {
	return func(txClient *gen.Client) userRepo {
		if txClient != nil {
			return userRepo{
				EntClient: txClient,
			}
		} else {
			return userRepo{
				EntClient: entClient,
			}
		}
	}
}

func (r *userRepo) Create(ctx context.Context, user du.User) (du.User, error) {
	usr, err := r.EntClient.User.
		Create().
		SetEmail(user.Email).
		SetOnboardingStatus(user.OnboardingStatus).
		Save(ctx)
	if err != nil {
		return du.User{}, errors.ErrorWithCurrentFuncName(err)
	}

	user.ID = usr.ID

	return user, nil
}

func (r *userRepo) UpdateOnboardingStatus(ctx context.Context, user *du.User) error {
	usr, err := r.EntClient.User.
		UpdateOneID(user.ID).
		SetOnboardingStatus(user.OnboardingStatus).
		Save(ctx)
	if err != nil {
		return errors.ErrorWithCurrentFuncName(err)
	}

	user.OnboardingStatus = du.OnboardingStatus(usr.OnboardingStatus)

	return nil
}

func (r *userRepo) CreateGoogleAuth(ctx context.Context, user du.User, auth du.GoogleAuth) (du.GoogleAuth, error) {
	gID, err := r.EntClient.GoogleAuth.Create().
		SetGoogleID(auth.GoogleID).
		SetUserID(user.ID).
		Save(ctx)
	if err != nil {
		return du.GoogleAuth{}, errors.ErrorWithCurrentFuncName(err)
	}

	auth.ID = gID.ID

	return auth, nil
}

func (r *userRepo) GetByGoogleAuthID(ctx context.Context, gid string) (du.User, error) {
	user, err := r.EntClient.User.
		Query().
		Where(
			user.HasGoogleAuthWith(
				googleauth.GoogleIDEQ(gid),
			),
		).
		WithGoogleAuth().
		Only(ctx)
	if err != nil {
		return du.User{}, errors.ErrorWithCurrentFuncName(err)
	}

	return du.User{
		ID: user.ID,
		GoogleAuth: du.GoogleAuth{
			ID:       user.Edges.GoogleAuth.ID,
			GoogleID: user.Edges.GoogleAuth.GoogleID,
		},
		Email: user.Email,
	}, nil
}

func (r *userRepo) GetByID(ctx context.Context, id string) (du.User, error) {
	user, err := r.EntClient.User.
		Query().
		Where(
			user.IDEQ(id),
		).
		WithGoogleAuth().
		Only(ctx)
	if err != nil {
		return du.User{}, errors.ErrorWithCurrentFuncName(err)
	}

	googleAuth, err := user.Edges.GoogleAuthOrErr()
	if err != nil {
		return du.User{}, errors.ErrorWithCurrentFuncName(err)
	}

	return du.User{
		ID: user.ID,
		GoogleAuth: du.GoogleAuth{
			ID:       googleAuth.ID,
			GoogleID: googleAuth.GoogleID,
		},
		OnboardingStatus: user.OnboardingStatus,
		Email:            user.Email,
	}, nil
}

type GetAllPaginationParam struct {
	pagination.CursorPagination
}

func (r *userRepo) GetAllPagination(ctx context.Context, param GetAllPaginationParam) (*gen.UserConnection, error) {
	user, err := r.EntClient.User.Query().
		Paginate(ctx, param.After.Cursor, param.First, param.Before.Cursor, param.Last)
	if err != nil {
		return nil, errors.ErrorWithCurrentFuncName(err)
	}

	return user, nil
}

func (r *userRepo) WithTx(client *gen.Client) {
	r.EntClient = client
}
