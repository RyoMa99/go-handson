package repository

import (
	"context"
	"database/sql"
	"handson/internal/apierr"
	"handson/internal/model"

	"github.com/jmoiron/sqlx"
)

type User struct{}

func NewUser() *User {
	return &User{}
}

func (u *User) FindByEmail(ctx context.Context, queryer sqlx.QueryerContext, email string) (*model.User, error) {
	var user model.User
	if err := sqlx.GetContext(ctx, queryer, &user, "SELECT id from users where email = ?", email); err == sql.ErrNoRows {
		return nil, apierr.ErrUserNotExists
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) Create(ctx context.Context, execer sqlx.ExecerContext, user *model.User) error {
	rs, err := execer.ExecContext(ctx, "insert into users(first_name, last_name, email, password_hash) values (?, ?, ?, ?)",
		user.FirstName, user.LastName, user.Email, user.PasswordHash)
	if err != nil {
		return err
	}

	// IDはauto incrementで生成されるので、ResultSetからとってくる
	id, err := rs.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int(id)

	return nil
}
