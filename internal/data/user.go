package data

import (
	"context"

	"usm/internal/biz"
	"usm/internal/biz/repo"
	"usm/internal/data/ent"
	"usm/internal/data/ent/user"
)

type userRepo struct {
	data *Data
}

func NewUserRepo(data *Data) repo.UserRepo {
	return &userRepo{
		data: data,
	}
}

func (r *userRepo) Create(ctx context.Context, user *repo.User) (*repo.User, error) {
	u, err := r.data.DB(ctx).User.
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
	return r.userFromEntity(u), nil
}

func (r *userRepo) Update(ctx context.Context, user *repo.User) (*repo.User, error) {
	u, err := r.data.DB(ctx).User.Get(ctx, int64(user.ID))
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, biz.ErrResourceNotFound
		}
		return nil, err
	}
	u, err = u.Update().
		SetEmail(user.Email).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.userFromEntity(u), nil
}

func (r *userRepo) Delete(ctx context.Context, id int) error {
	return r.data.DB(ctx).User.DeleteOneID(int64(id)).Exec(ctx)
}

func (r *userRepo) Get(ctx context.Context, id int) (*repo.User, error) {
	u, err := r.data.DB(ctx).User.Get(ctx, int64(id))
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, biz.ErrResourceNotFound
		}
		return nil, err
	}
	return r.userFromEntity(u), nil
}

func (r *userRepo) Enable(ctx context.Context, id int) error {
	u, err := r.data.DB(ctx).User.Get(ctx, int64(id))
	if err != nil {
		if ent.IsNotFound(err) {
			return biz.ErrResourceNotFound
		}
		return err
	}
	_, err = u.Update().SetDisabled(false).Save(ctx)
	return err
}

func (r *userRepo) Disable(ctx context.Context, id int) error {
	u, err := r.data.DB(ctx).User.Get(ctx, int64(id))
	if err != nil {
		if ent.IsNotFound(err) {
			return biz.ErrResourceNotFound
		}
		return err
	}
	_, err = u.Update().SetDisabled(true).Save(ctx)
	return err
}

func (r *userRepo) List(ctx context.Context, offset, limit int) ([]*repo.User, error) {
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 20
	}
	ents, err := r.data.DB(ctx).User.Query().Offset(offset).Limit(limit).All(ctx)
	if err != nil {
		return nil, err
	}
	users := make([]*repo.User, 0, len(ents))
	for _, u := range ents {
		users = append(users, r.userFromEntity(u))
	}
	return users, nil
}

func (r *userRepo) SetPassword(ctx context.Context, id int, password string) error {
	u, err := r.data.DB(ctx).User.Get(ctx, int64(id))
	if err != nil {
		if ent.IsNotFound(err) {
			return biz.ErrResourceNotFound
		}
		return err
	}
	_, err = u.Update().
		SetPassword(password).
		Save(ctx)
	return err
}

func (r *userRepo) GetByUsername(ctx context.Context, username string) (*repo.User, error) {
	u, err := r.data.DB(ctx).User.Query().Where(
		user.Username(username),
	).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, biz.ErrResourceNotFound
		}
		return nil, err
	}
	return r.userFromEntity(u), nil
}

func (r *userRepo) userFromEntity(u *ent.User) *repo.User {
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
