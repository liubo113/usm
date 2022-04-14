package account

import (
	"context"

	pb "usm/api/account/v1"
	"usm/internal/biz"
	"usm/internal/biz/repo"
	acctuc "usm/internal/biz/usecase/account"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	pb.UnimplementedAccountServer

	log *log.Helper

	uc *acctuc.Usecase
}

func NewService(uc *acctuc.Usecase, logger log.Logger) *Service {
	return &Service{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *Service) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	log.Infof("create user %s, email=%s", req.Username, req.Email)
	u, err := s.uc.CreateUser(ctx, &repo.User{
		Username: req.Username,
		Email:    req.Email,
		Password: "Admin@169+-",
	})
	if err != nil {
		if err == biz.ErrResourceAlreadyExists {
			return nil, pb.ErrorUserAlreadyExisted("user %s already existed", req.Username)
		}
		return nil, err
	}
	return protoFromBizUser(u), nil
}

func (s *Service) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.User, error) {
	log.Infof("update user %d, email=%s", req.Id, req.Email)
	u, err := s.uc.UpdateUser(ctx, &repo.User{
		ID:    int(req.Id),
		Email: req.Email,
	})
	if err != nil {
		return nil, err
	}
	return protoFromBizUser(u), nil
}

func (s *Service) SetUserPassword(ctx context.Context, req *pb.SetUserPasswordRequest) (*pb.SetUserPasswordResponse, error) {
	log.Infof("user %d set password", req.Id)
	if err := s.uc.SetUserPassword(ctx, int(req.Id), req.Password); err != nil {
		return nil, err
	}
	return &pb.SetUserPasswordResponse{}, nil
}

func (s *Service) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	log.Infof("delete user %d", req.Id)
	if err := s.uc.DeleteUser(ctx, int(req.Id)); err != nil {
		return nil, err
	}
	return &pb.DeleteUserResponse{}, nil
}

func (s *Service) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	log.Infof("get user %d", req.Id)
	u, err := s.uc.GetUser(ctx, int(req.Id))
	if err != nil {
		if err == biz.ErrResourceNotFound {
			return nil, pb.ErrorUserNotFound("user %d not found", req.Id)
		}
		return nil, err
	}
	return protoFromBizUser(u), nil
}

func (s *Service) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	log.Infof("list users, offset=%d, limit=%d", req.Offset, req.Limit)
	us, err := s.uc.ListUsers(ctx, int(req.Offset), int(req.Limit))
	if err != nil {
		return nil, err
	}
	var resp pb.ListUsersResponse
	for _, u := range us {
		resp.Users = append(resp.Users, protoFromBizUser(u))
	}
	return &resp, nil
}

func (s *Service) Authenticate(ctx context.Context, req *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	switch method := req.AuthMethod.(type) {
	default:
		return nil, status.Errorf(codes.InvalidArgument, "invalid auth method")
	case *pb.AuthenticateRequest_BasicAuth_:
		auth := method.BasicAuth
		log.Infof("user %s authenticate, method: basic auth", auth.Username)
		u, err := s.uc.GetUserByUsername(ctx, auth.GetUsername())
		if err != nil {
			return nil, err
		}
		if u.Password != auth.GetPassword() {
			return nil, pb.ErrorMismatchUsernamePassword("mismatch password")
		}
	}
	return &pb.AuthenticateResponse{}, nil
}

func protoFromBizUser(u *repo.User) *pb.User {
	return &pb.User{
		Id:         int64(u.ID),
		Username:   u.Username,
		Email:      u.Email,
		CreateTime: u.CreateTime.Unix(),
		UpdateTime: u.UpdateTime.Unix(),
	}
}
