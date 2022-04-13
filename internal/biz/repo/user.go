package repo

import (
	"context"
	"time"
)

type User struct {
	ID         int
	Username   string
	Email      string
	Password   string
	Disabled   bool
	CreateTime time.Time
	UpdateTime time.Time
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	UpdateUser(ctx context.Context, user *User) (*User, error)
	SetUserPassword(ctx context.Context, id int64, password string) error
	DeleteUser(ctx context.Context, id int64) error
	GetUser(ctx context.Context, id int64) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	EnableUser(ctx context.Context, id int64) error
	DisableUser(ctx context.Context, id int64) error
	ListUsers(ctx context.Context, offset, limit int) ([]*User, error)
}
