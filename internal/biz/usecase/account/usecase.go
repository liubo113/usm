package account

import (
	"context"

	"usm/internal/biz/repo"
)

type Usecase struct {
	userRepo repo.UserRepo
}

func NewUsecase(userRepo repo.UserRepo) *Usecase {
	return &Usecase{userRepo}
}

func (uc *Usecase) CreateUser(ctx context.Context, user *repo.User) (*repo.User, error) {
	return uc.userRepo.CreateUser(ctx, user)
}

func (uc *Usecase) UpdateUser(ctx context.Context, user *repo.User) (*repo.User, error) {
	return uc.userRepo.UpdateUser(ctx, user)
}

func (uc *Usecase) SetUserPassword(ctx context.Context, id int64, password string) error {
	return uc.userRepo.SetUserPassword(ctx, id, password)
}

func (uc *Usecase) DeleteUser(ctx context.Context, id int64) error {
	return uc.userRepo.DeleteUser(ctx, id)
}

func (uc *Usecase) GetUser(ctx context.Context, id int64) (*repo.User, error) {
	return uc.userRepo.GetUser(ctx, id)
}

func (uc *Usecase) GetUserByUsername(ctx context.Context, username string) (*repo.User, error) {
	return uc.userRepo.GetUserByUsername(ctx, username)
}

func (uc *Usecase) ListUsers(ctx context.Context, offset, limit int) ([]*repo.User, error) {
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 20
	}
	return uc.userRepo.ListUsers(ctx, offset, limit)
}
