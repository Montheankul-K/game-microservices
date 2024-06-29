package authHandler

import (
	"context"
	authPb "github.com/Montheankul-K/game-microservices/modules/auth/authPb"
	"github.com/Montheankul-K/game-microservices/modules/auth/authUsecase"
)

type (
	AuthGrpcHandler struct {
		authPb.UnimplementedAuthGrpcServiceServer
		authUsecase authUsecase.AuthUsecase
	}
)

func NewAuthGrpcHandler(authUsecase authUsecase.AuthUsecase) *AuthGrpcHandler {
	return &AuthGrpcHandler{
		authUsecase: authUsecase,
	}
}

func (g *AuthGrpcHandler) CredentialSearch(ctx context.Context, req *authPb.CredentialSearchReq) (*authPb.CredentialSearchRes, error) {
	return nil, nil
}

func (g *AuthGrpcHandler) RolesCount(ctx context.Context, req *authPb.RolesCountReq) (*authPb.RolesCountRes, error) {
	return nil, nil
}
