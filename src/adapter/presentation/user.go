package presentation

import (
	"context"
	"fmt"

	"connectrpc.com/connect"

	userv1 "github.com/yuorei/yuorei-ads/gen/rpc/user/v1"
	"github.com/yuorei/yuorei-ads/src/usecase"
)

type UserServer struct {
	usecase *usecase.UseCase
}

func NewUserServer(repository *usecase.Repository) *UserServer {
	return &UserServer{
		usecase: usecase.NewUseCase(repository),
	}
}

func (s *UserServer) CreateUser(ctx context.Context, req *connect.Request[userv1.CreateUserRequest]) (*connect.Response[userv1.CreateUserResponse], error) {
	// TODO: 認証後のユーザIDを取得
	result, err := s.usecase.CreateUser(ctx, "id", req.Msg.Role)
	if err != nil {
		return nil, fmt.Errorf("failed to create campaign: %w", err)
	}

	res := connect.NewResponse(&userv1.CreateUserResponse{
		UserId: result.ID,
	})

	return res, nil
}
