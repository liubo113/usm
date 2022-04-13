package data

import (
	"context"

	"usm/internal/biz"
	"usm/internal/biz/repo"
	"usm/internal/data/ent"
	"usm/internal/data/ent/user"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) repo.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, user *repo.User) (*repo.User, error) {
	u, err := r.data.db.User.
		Create().
		SetUsername(user.Username).
		SetEmail(user.Email).
		SetPassword(user.Password).
		Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			return nil, biz.ErrResourceAlreadyExists
		}
		return nil, err
	}
	return bizUserFromEntity(u), nil
}

func (r *userRepo) UpdateUser(ctx context.Context, user *repo.User) (*repo.User, error) {
	u, err := r.data.db.User.Get(ctx, int64(user.ID))
	if err != nil {
		return nil, err
	}
	u, err = u.Update().
		SetEmail(user.Email).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return bizUserFromEntity(u), nil
}

func (r *userRepo) SetUserPassword(ctx context.Context, id int64, password string) error {
	u, err := r.data.db.User.Get(ctx, id)
	if err != nil {
		return err
	}
	_, err = u.Update().
		SetPassword(password).
		Save(ctx)
	return err
}

func (r *userRepo) DeleteUser(ctx context.Context, id int64) error {
	return r.data.db.User.DeleteOneID(id).Exec(ctx)
}

func (r *userRepo) GetUser(ctx context.Context, id int64) (*repo.User, error) {
	u, err := r.data.db.User.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, biz.ErrResourceNotFound
		}
		return nil, err
	}
	return bizUserFromEntity(u), nil
}

func (r *userRepo) GetUserByUsername(ctx context.Context, username string) (*repo.User, error) {
	u, err := r.data.db.User.Query().Where(
		user.Username(username),
	).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, biz.ErrResourceNotFound
		}
		return nil, err
	}
	return bizUserFromEntity(u), nil
}

func (r *userRepo) EnableUser(ctx context.Context, id int64) error {
	u, err := r.data.db.User.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return biz.ErrResourceNotFound
		}
		return err
	}
	_, err = u.Update().SetDisabled(false).Save(ctx)
	return err
}

func (r *userRepo) DisableUser(ctx context.Context, id int64) error {
	u, err := r.data.db.User.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return biz.ErrResourceNotFound
		}
		return err
	}
	_, err = u.Update().SetDisabled(true).Save(ctx)
	return err
}

func (r *userRepo) ListUsers(ctx context.Context, offset, limit int) ([]*repo.User, error) {
	ents, err := r.data.db.User.Query().Offset(offset).Limit(limit).All(ctx)
	if err != nil {
		return nil, err
	}
	users := make([]*repo.User, 0, len(ents))
	for _, u := range ents {
		users = append(users, bizUserFromEntity(u))
	}
	return users, nil
}

func bizUserFromEntity(u *ent.User) *repo.User {
	return &repo.User{
		ID:         int(u.ID),
		Username:   u.Username,
		Email:      u.Email,
		Password:   u.Password,
		Disabled:   u.Disabled,
		CreateTime: u.CreateTime,
		UpdateTime: u.UpdateTime,
	}
}
