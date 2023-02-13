package usecase

import (
	"context"
	"handson/internal/apierr"
	"handson/internal/model"
	"testing"

	"github.com/jmoiron/sqlx"
)

type userRepositoryMock struct {
	findByEmailFn func(ctx context.Context, queryer sqlx.QueryerContext, email string) (*model.User, error)
	createFn      func(ctx context.Context, execer sqlx.ExecerContext, m *model.User) error
}

func (s *userRepositoryMock) FindByEmail(ctx context.Context, queryer sqlx.QueryerContext, email string) (*model.User, error) {
	return s.findByEmailFn(ctx, queryer, email)
}

func (s *userRepositoryMock) Create(ctx context.Context, execer sqlx.ExecerContext, m *model.User) error {
	return s.createFn(ctx, execer, m)
}

func TestUser_Create(t *testing.T) {
	var passedMail string

	mock := &userRepositoryMock{
		findByEmailFn: func(ctx context.Context, queryer sqlx.QueryerContext, email string) (user *model.User, err error) {
			passedMail = email
			return nil, apierr.ErrUserNotExists
		},
		createFn: func(ctx context.Context, execer sqlx.ExecerContext, m *model.User) error {
			return nil
		},
	}
	userUsecase := NewUser(mock, nil)

	if err := userUsecase.Create(context.Background(), &model.User{
		FirstName:    "test_first_name",
		LastName:     "test_last_name",
		Email:        "test@example.com",
		PasswordHash: "aaa",
	}); err != nil {
		t.Fatal(err)
	}

	if passedMail != "test@example.com" {
		t.Errorf("email must be test@example.com but %s", passedMail)
	}
}
