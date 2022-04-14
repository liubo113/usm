package repo

//go:generate mockgen -destination=./mock/user.go -package=mock usm/internal/biz/repo UserRepo

import (
	"context"
	"time"
)

type User struct {
	ID         int
	Disabled   bool
	CreateTime time.Time
	UpdateTime time.Time

	// TODO: 补充自定义字段
	Username string
	Email    string
	Password string
}

type UserRepo interface {
	Create(ctx context.Context, m *User) (*User, error)
	Update(ctx context.Context, m *User) (*User, error)
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (*User, error)
	Enable(ctx context.Context, id int) error
	Disable(ctx context.Context, id int) error
	List(ctx context.Context, offset, limit int) ([]*User, error)

	// TODO: 补充自定义方法
	SetPassword(ctx context.Context, id int, password string) error
	GetByUsername(ctx context.Context, username string) (*User, error)
}
