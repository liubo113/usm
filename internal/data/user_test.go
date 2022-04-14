package data

import (
	"context"
	"testing"
	"time"

	"usm/internal/biz"
	"usm/internal/biz/repo"
	"usm/internal/data/ent"

	"github.com/stretchr/testify/assert"
)

func Test_userRepo_Create(t *testing.T) {
	ctx := context.Background()
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
					Password: "Admin@169+-",
				},
			},
			want: &repo.User{
				ID:       1,
				Username: "liubo",
				Email:    "findliubo@163.com",
				Password: "Admin@169+-",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, teardown := NewTestData(t)
			defer teardown()
			r := NewUserRepo(data)
			got, err := r.Create(ctx, tt.args.user)
			assert.Equal(t, tt.wantErr, err, "error=%v, wantErr=%v", err, tt.wantErr)
			got.CreateTime = time.Time{}
			got.UpdateTime = time.Time{}
			assert.Equal(t, tt.want, got, "mismatch: got=%v, want=%v", got, tt.want)
		})
	}
}

func Test_userRepo_Update(t *testing.T) {
	ctx := context.Background()
	defaultUser := &repo.User{
		ID:       1,
		Username: "liubo",
		Email:    "findliubo@163.com",
		Password: "Admin@169+-",
	}
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
			name: "should update user successfully",
			args: args{
				user: &repo.User{
					ID:    1,
					Email: "findliubo@361.com",
				},
			},
			want: &repo.User{
				ID:       1,
				Username: "liubo",
				Email:    "findliubo@361.com",
				Password: "Admin@169+-",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, teardown := NewTestData(t)
			defer teardown()
			r := NewUserRepo(data)
			r.Create(ctx, defaultUser)
			got, err := r.Update(ctx, tt.args.user)
			assert.Equal(t, tt.wantErr, err, "error=%v, wantErr=%v", err, tt.wantErr)
			got.CreateTime = time.Time{}
			got.UpdateTime = time.Time{}
			assert.Equal(t, tt.want, got, "mismatch: got=%v, want=%v", got, tt.want)
		})
	}
}

func Test_userRepo_Delete(t *testing.T) {
	ctx := context.Background()
	defaultUser := &repo.User{
		ID:       1,
		Username: "liubo",
		Email:    "findliubo@163.com",
		Password: "Admin@169+-",
	}
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, teardown := NewTestData(t)
			defer teardown()
			r := NewUserRepo(data)
			r.Create(ctx, defaultUser)
			err := r.Delete(ctx, tt.args.id)
			assert.Equal(t, tt.wantErr, err, "error=%v, wantErr=%v", err, tt.wantErr)
			_, err = r.Get(ctx, tt.args.id)
			assert.Equal(t, err, biz.ErrResourceNotFound, "mismatch get error: got=%v, want=%v", err, biz.ErrResourceNotFound)
		})
	}
}

func Test_userRepo_Get(t *testing.T) {
	ctx := context.Background()
	defaultUser := &repo.User{
		ID:       1,
		Username: "liubo",
		Email:    "findliubo@163.com",
		Password: "Admin@169+-",
	}
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
			want:    defaultUser,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, teardown := NewTestData(t)
			defer teardown()
			r := NewUserRepo(data)
			r.Create(ctx, defaultUser)
			got, err := r.Get(ctx, tt.args.id)
			assert.Equal(t, tt.wantErr, err, "error=%v, wantErr=%v", err, tt.wantErr)
			got.CreateTime = time.Time{}
			got.UpdateTime = time.Time{}
			assert.Equal(t, tt.want, got, "mismatch: got=%v, want=%v", got, tt.want)
		})
	}
}

func Test_userRepo_SetPassword(t *testing.T) {
	ctx := context.Background()
	defaultUser := &repo.User{
		ID:       1,
		Username: "liubo",
		Email:    "findliubo@163.com",
		Password: "Admin@169+-",
	}
	type args struct {
		id       int
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    *repo.User
		wantErr error
	}{
		{
			name: "should set user passwor successfully",
			args: args{
				id:       1,
				password: "Admin@169--",
			},
			want: &repo.User{
				ID:       1,
				Username: "liubo",
				Email:    "findliubo@163.com",
				Password: "Admin@169--",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, teardown := NewTestData(t)
			defer teardown()
			r := NewUserRepo(data)
			r.Create(ctx, defaultUser)
			err := r.SetPassword(ctx, tt.args.id, tt.args.password)
			assert.Equal(t, tt.wantErr, err, "error=%v, wantErr=%v", err, tt.wantErr)
			got, _ := r.Get(ctx, tt.args.id)
			got.CreateTime = time.Time{}
			got.UpdateTime = time.Time{}
			assert.Equal(t, tt.want, got, "mismatch: got=%v, want=%v", got, tt.want)
		})
	}
}

func Test_userRepo_GetByUsername(t *testing.T) {
	ctx := context.Background()
	defaultUser := &repo.User{
		ID:       1,
		Username: "liubo",
		Email:    "findliubo@163.com",
		Password: "Admin@169+-",
	}
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
			want:    defaultUser,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, teardown := NewTestData(t)
			defer teardown()
			r := NewUserRepo(data)
			r.Create(ctx, defaultUser)
			got, err := r.GetByUsername(ctx, tt.args.username)
			assert.Equal(t, tt.wantErr, err, "error=%v, wantErr=%v", err, tt.wantErr)
			got.CreateTime = time.Time{}
			got.UpdateTime = time.Time{}
			assert.Equal(t, tt.want, got, "mismatch: got=%v, want=%v", got, tt.want)
		})
	}
}

func Test_userRepo_userFromEntity(t *testing.T) {
	now := time.Now()
	type args struct {
		u *ent.User
	}
	tests := []struct {
		name string
		args args
		want *repo.User
	}{
		{
			name: "convert successfully",
			args: args{
				u: &ent.User{
					ID:         1,
					CreateTime: now,
					UpdateTime: now,
					Username:   "liubo",
					Email:      "findliubo@163.com",
					Password:   "Admin@169+-",
					Disabled:   true,
				},
			},
			want: &repo.User{
				ID:         1,
				CreateTime: now,
				UpdateTime: now,
				Username:   "liubo",
				Email:      "findliubo@163.com",
				Password:   "Admin@169+-",
				Disabled:   true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &userRepo{}
			got := r.userFromEntity(tt.args.u)
			assert.Equal(t, tt.want, got, "mismatch: got=%v, want=%v", got, tt.want)
		})
	}
}
