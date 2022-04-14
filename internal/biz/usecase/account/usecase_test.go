package account

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"usm/internal/biz/repo"
	"usm/internal/biz/repo/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	testNow               = time.Now()
	testErrNotFound       = errors.New("not found")
	testErrAlreadyExisted = errors.New("already existed")
	testUserStore         = map[int]*repo.User{
		1: {
			ID:         1,
			Username:   "liubo",
			Email:      "findliubo@163.com",
			Password:   "Admin@169+-", // default
			CreateTime: testNow,
			UpdateTime: testNow,
		},
	}
)

func TestUsecase_CreateUser(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	mockTran := mock.NewMockTransaction(ctrl)
	mockRepo := mock.NewMockUserRepo(ctrl)
	mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).AnyTimes().DoAndReturn(func(ctx context.Context, u *repo.User) (*repo.User, error) {
		if u.Username == "duplicated" {
			return nil, testErrAlreadyExisted
		}
		return &repo.User{
			ID:         1,
			Username:   u.Username,
			Email:      u.Email,
			Password:   "Admin@169+-", // default
			CreateTime: testNow,
			UpdateTime: testNow,
		}, nil
	})
	type args struct {
		user *repo.User
	}
	tests := []struct {
		name    string
		args    args
		want    *repo.User
		wantErr error
	}{
		{
			name: "should create user successfully",
			args: args{
				user: &repo.User{
					Username: "liubo",
					Email:    "findliubo@163.com",
				},
			},
			want: &repo.User{
				ID:         1,
				Username:   "liubo",
				Email:      "findliubo@163.com",
				Password:   "Admin@169+-",
				CreateTime: testNow,
				UpdateTime: testNow,
			},
			wantErr: nil,
		},
		{
			name: "should create user failed if record is duplicated",
			args: args{
				user: &repo.User{
					Username: "duplicated",
					Email:    "findliubo@163.com",
				},
			},
			want:    nil,
			wantErr: testErrAlreadyExisted,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &Usecase{
				tran:     mockTran,
				userRepo: mockRepo,
			}
			got, err := uc.CreateUser(ctx, tt.args.user)
			assert.Equal(t, tt.wantErr, err, "error=%v, wantErr=%v", err, tt.wantErr)
			assert.Equal(t, tt.want, got, "mismatch: got=%v, want=%v", got, tt.want)
		})
	}
}

func TestUsecase_UpdateUser(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	mockTran := mock.NewMockTransaction(ctrl)
	mockRepo := mock.NewMockUserRepo(ctrl)
	mockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).AnyTimes().DoAndReturn(func(ctx context.Context, u *repo.User) (*repo.User, error) {
		store, ok := testUserStore[u.ID]
		if !ok {
			return nil, testErrNotFound
		}
		store.Email = u.Email
		return store, nil
	})
	type args struct {
		user *repo.User
	}
	tests := []struct {
		name    string
		args    args
		want    *repo.User
		wantErr error
	}{
		{
			name: "should update user email successfully",
			args: args{
				user: &repo.User{
					ID:       1,
					Username: "liubo",
					Email:    "findliubo@361.com",
				},
			},
			want:    testUserStore[1],
			wantErr: nil,
		},
		{
			name: "should update user failed if not found",
			args: args{
				user: &repo.User{
					ID:    2,
					Email: "findliubo@361.com",
				},
			},
			want:    nil,
			wantErr: testErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &Usecase{
				tran:     mockTran,
				userRepo: mockRepo,
			}
			got, err := uc.UpdateUser(ctx, tt.args.user)
			assert.Equal(t, tt.wantErr, err, "error=%v, wantErr=%v", err, tt.wantErr)
			assert.Equal(t, tt.want, got, "mismatch: got=%v, want=%v", got, tt.want)
		})
	}
}

func TestUsecase_DeleteUser(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	mockTran := mock.NewMockTransaction(ctrl)
	mockRepo := mock.NewMockUserRepo(ctrl)
	mockRepo.EXPECT().Delete(gomock.Any(), gomock.Any()).AnyTimes().DoAndReturn(func(ctx context.Context, id int) error {
		if _, ok := testUserStore[id]; !ok {
			return testErrNotFound
		}
		return nil
	})
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "should delete user successfully",
			args: args{
				id: 1,
			},
			wantErr: nil,
		},
		{
			name: "should delete user failed if not found",
			args: args{
				id: 0,
			},
			wantErr: testErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &Usecase{
				tran:     mockTran,
				userRepo: mockRepo,
			}
			err := uc.DeleteUser(ctx, tt.args.id)
			assert.Equal(t, tt.wantErr, err, "error=%v, wantErr=%v", err, tt.wantErr)
		})
	}
}

func TestUsecase_GetUser(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	mockTran := mock.NewMockTransaction(ctrl)
	mockRepo := mock.NewMockUserRepo(ctrl)
	mockRepo.EXPECT().Get(gomock.Any(), gomock.Any()).AnyTimes().DoAndReturn(func(ctx context.Context, id int) (*repo.User, error) {
		store, ok := testUserStore[id]
		if !ok {
			return nil, testErrNotFound
		}
		return store, nil
	})
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    *repo.User
		wantErr error
	}{
		{
			name: "should get user successfully",
			args: args{
				id: 1,
			},
			want:    testUserStore[1],
			wantErr: nil,
		},
		{
			name: "should get user failed if not found",
			args: args{
				id: 0,
			},
			wantErr: testErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &Usecase{
				tran:     mockTran,
				userRepo: mockRepo,
			}
			got, err := uc.GetUser(ctx, tt.args.id)
			assert.Equal(t, tt.wantErr, err, "error=%v, wantErr=%v", err, tt.wantErr)
			assert.Equal(t, tt.want, got, "mismatch: got=%v, want=%v", got, tt.want)
		})
	}
}

func TestUsecase_ListUsers(t *testing.T) {
	type fields struct {
		tran     repo.Transaction
		userRepo repo.UserRepo
	}
	type args struct {
		ctx    context.Context
		offset int
		limit  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*repo.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &Usecase{
				tran:     tt.fields.tran,
				userRepo: tt.fields.userRepo,
			}
			got, err := uc.ListUsers(tt.args.ctx, tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.ListUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.ListUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_SetUserPassword(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	mockTran := mock.NewMockTransaction(ctrl)
	mockTran.EXPECT().WithTx(gomock.Any(), gomock.Any()).AnyTimes().DoAndReturn(func(ctx context.Context, fn func(ctx context.Context) error) error {
		return fn(ctx)
	})
	mockRepo := mock.NewMockUserRepo(ctrl)
	mockRepo.EXPECT().SetPassword(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().DoAndReturn(func(ctx context.Context, id int, pass string) error {
		store, ok := testUserStore[id]
		if !ok {
			return testErrNotFound
		}
		store.Password = pass
		return nil
	})
	type args struct {
		id   int
		pass string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "should set user password successfully",
			args: args{
				id:   1,
				pass: "Admin@169--",
			},
			wantErr: nil,
		},
		{
			name: "should set user password failed if not found",
			args: args{
				id:   0,
				pass: "Admin@169--",
			},
			wantErr: testErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &Usecase{
				tran:     mockTran,
				userRepo: mockRepo,
			}
			err := uc.SetUserPassword(ctx, tt.args.id, tt.args.pass)
			assert.Equal(t, tt.wantErr, err, "error=%v, wantErr=%v", err, tt.wantErr)
		})
	}
}

func TestUsecase_GetUserByUsername(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	mockTran := mock.NewMockTransaction(ctrl)
	mockRepo := mock.NewMockUserRepo(ctrl)
	mockRepo.EXPECT().GetByUsername(gomock.Any(), gomock.Any()).AnyTimes().DoAndReturn(func(ctx context.Context, username string) (*repo.User, error) {
		stores := map[string]*repo.User{
			"liubo": {
				ID:         1,
				Username:   "liubo",
				Email:      "findliubo@163.com",
				Password:   "Admin@169+-", // default
				CreateTime: testNow,
				UpdateTime: testNow,
			},
		}
		store, ok := stores[username]
		if !ok {
			return nil, testErrNotFound
		}
		return store, nil
	})
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    *repo.User
		wantErr error
	}{
		{
			name: "should get user successfully",
			args: args{
				username: "liubo",
			},
			want: &repo.User{
				ID:         1,
				Username:   "liubo",
				Email:      "findliubo@163.com",
				Password:   "Admin@169+-", // default
				CreateTime: testNow,
				UpdateTime: testNow,
			},
			wantErr: nil,
		},
		{
			name: "should get user failed if not found",
			args: args{
				username: "",
			},
			wantErr: testErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &Usecase{
				tran:     mockTran,
				userRepo: mockRepo,
			}
			got, err := uc.GetUserByUsername(ctx, tt.args.username)
			assert.Equal(t, tt.wantErr, err, "error=%v, wantErr=%v", err, tt.wantErr)
			assert.Equal(t, tt.want, got, "mismatch: got=%v, want=%v", got, tt.want)
		})
	}
}
