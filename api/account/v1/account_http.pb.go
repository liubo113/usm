// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.2.1

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type AccountHTTPServer interface {
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	GetUser(context.Context, *GetUserRequest) (*User, error)
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	SetUserPassword(context.Context, *SetUserPasswordRequest) (*SetUserPasswordResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*User, error)
}

func RegisterAccountHTTPServer(s *http.Server, srv AccountHTTPServer) {
	r := s.Route("/")
	r.POST("/account/v1/users", _Account_CreateUser0_HTTP_Handler(srv))
	r.PUT("/account/v1/users/{id}", _Account_UpdateUser0_HTTP_Handler(srv))
	r.POST("/account/v1/users:set-password", _Account_SetUserPassword0_HTTP_Handler(srv))
	r.DELETE("/account/v1/users/{id}", _Account_DeleteUser0_HTTP_Handler(srv))
	r.GET("/account/v1/users/{id}", _Account_GetUser0_HTTP_Handler(srv))
	r.GET("/account/v1/users", _Account_ListUsers0_HTTP_Handler(srv))
	r.POST("/account/v1/users:authenticate", _Account_Authenticate0_HTTP_Handler(srv))
}

func _Account_CreateUser0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.account.v1.Account/CreateUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateUser(ctx, req.(*CreateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*User)
		return ctx.Result(200, reply)
	}
}

func _Account_UpdateUser0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.account.v1.Account/UpdateUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUser(ctx, req.(*UpdateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*User)
		return ctx.Result(200, reply)
	}
}

func _Account_SetUserPassword0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetUserPasswordRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.account.v1.Account/SetUserPassword")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetUserPassword(ctx, req.(*SetUserPasswordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SetUserPasswordResponse)
		return ctx.Result(200, reply)
	}
}

func _Account_DeleteUser0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.account.v1.Account/DeleteUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteUser(ctx, req.(*DeleteUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteUserResponse)
		return ctx.Result(200, reply)
	}
}

func _Account_GetUser0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.account.v1.Account/GetUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*GetUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*User)
		return ctx.Result(200, reply)
	}
}

func _Account_ListUsers0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUsersRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.account.v1.Account/ListUsers")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUsers(ctx, req.(*ListUsersRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUsersResponse)
		return ctx.Result(200, reply)
	}
}

func _Account_Authenticate0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AuthenticateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.account.v1.Account/Authenticate")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Authenticate(ctx, req.(*AuthenticateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AuthenticateResponse)
		return ctx.Result(200, reply)
	}
}

type AccountHTTPClient interface {
	Authenticate(ctx context.Context, req *AuthenticateRequest, opts ...http.CallOption) (rsp *AuthenticateResponse, err error)
	CreateUser(ctx context.Context, req *CreateUserRequest, opts ...http.CallOption) (rsp *User, err error)
	DeleteUser(ctx context.Context, req *DeleteUserRequest, opts ...http.CallOption) (rsp *DeleteUserResponse, err error)
	GetUser(ctx context.Context, req *GetUserRequest, opts ...http.CallOption) (rsp *User, err error)
	ListUsers(ctx context.Context, req *ListUsersRequest, opts ...http.CallOption) (rsp *ListUsersResponse, err error)
	SetUserPassword(ctx context.Context, req *SetUserPasswordRequest, opts ...http.CallOption) (rsp *SetUserPasswordResponse, err error)
	UpdateUser(ctx context.Context, req *UpdateUserRequest, opts ...http.CallOption) (rsp *User, err error)
}

type AccountHTTPClientImpl struct {
	cc *http.Client
}

func NewAccountHTTPClient(client *http.Client) AccountHTTPClient {
	return &AccountHTTPClientImpl{client}
}

func (c *AccountHTTPClientImpl) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...http.CallOption) (*AuthenticateResponse, error) {
	var out AuthenticateResponse
	pattern := "/account/v1/users:authenticate"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.account.v1.Account/Authenticate"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AccountHTTPClientImpl) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...http.CallOption) (*User, error) {
	var out User
	pattern := "/account/v1/users"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.account.v1.Account/CreateUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AccountHTTPClientImpl) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...http.CallOption) (*DeleteUserResponse, error) {
	var out DeleteUserResponse
	pattern := "/account/v1/users/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.account.v1.Account/DeleteUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AccountHTTPClientImpl) GetUser(ctx context.Context, in *GetUserRequest, opts ...http.CallOption) (*User, error) {
	var out User
	pattern := "/account/v1/users/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.account.v1.Account/GetUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AccountHTTPClientImpl) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...http.CallOption) (*ListUsersResponse, error) {
	var out ListUsersResponse
	pattern := "/account/v1/users"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.account.v1.Account/ListUsers"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AccountHTTPClientImpl) SetUserPassword(ctx context.Context, in *SetUserPasswordRequest, opts ...http.CallOption) (*SetUserPasswordResponse, error) {
	var out SetUserPasswordResponse
	pattern := "/account/v1/users:set-password"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.account.v1.Account/SetUserPassword"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AccountHTTPClientImpl) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...http.CallOption) (*User, error) {
	var out User
	pattern := "/account/v1/users/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.account.v1.Account/UpdateUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}