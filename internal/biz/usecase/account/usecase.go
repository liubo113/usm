package account

import (
	"context"

	"usm/internal/biz/repo"
)

type Usecase struct {
	tran     repo.Transaction
	userRepo repo.UserRepo
	// other repos...
}

func NewUsecase(tran repo.Transaction, userRepo repo.UserRepo) *Usecase {
	return &Usecase{
		tran:     tran,
		userRepo: userRepo,
	}
}

func (uc *Usecase) CreateUser(ctx context.Context, user *repo.User) (*repo.User, error) {
	return uc.userRepo.Create(ctx, user)
}

func (uc *Usecase) UpdateUser(ctx context.Context, user *repo.User) (*repo.User, error) {
	return uc.userRepo.Update(ctx, user)
}

func (uc *Usecase) DeleteUser(ctx context.Context, id int) error {
	return uc.userRepo.Delete(ctx, id)
}

func (uc *Usecase) GetUser(ctx context.Context, id int) (*repo.User, error) {
	return uc.userRepo.Get(ctx, id)
}

func (uc *Usecase) ListUsers(ctx context.Context, offset, limit int) ([]*repo.User, error) {
	return uc.userRepo.List(ctx, offset, limit)
}

func (uc *Usecase) SetUserPassword(ctx context.Context, id int, password string) error {
	// transaction example
	return uc.tran.WithTx(ctx, func(ctx context.Context) error {
		return uc.userRepo.SetPassword(ctx, id, password)
	})
}

func (uc *Usecase) GetUserByUsername(ctx context.Context, username string) (*repo.User, error) {
	return uc.userRepo.GetByUsername(ctx, username)
}
