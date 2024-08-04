package presentation

import (
	"context"
	"fmt"
	"time"

	"connectrpc.com/connect"

	userv1 "github.com/yuorei/yuorei-ads/gen/rpc/user/v1"
	"github.com/yuorei/yuorei-ads/src/adapter/infrastructure"
	"github.com/yuorei/yuorei-ads/src/domain"
	"github.com/yuorei/yuorei-ads/src/usecase"
)

type UserServer struct {
	usecase *usecase.UseCase
}

func NewUserServer(infra *infrastructure.Infrastructure) *UserServer {
	return &UserServer{
		usecase: usecase.NewUseCase(infra),
	}
}

func (s *UserServer) CreateUser(ctx context.Context, req *connect.Request[userv1.CreateUserRequest]) (*connect.Response[userv1.CreateUserResponse], error) {
	// TODO: 認証後のユーザIDを取得
	campaign := domain.NewUser("1", req.Msg.Username, req.Msg.Password, req.Msg.Email, []string{"general"}, time.Now(), time.Now(), nil)
	result, err := s.usecase.CreateUser(ctx, campaign)
	if err != nil {
		return nil, fmt.Errorf("failed to create campaign: %w", err)
	}

	res := connect.NewResponse(&userv1.CreateUserResponse{
		UserId: result.ID,
	})

	return res, nil
}
