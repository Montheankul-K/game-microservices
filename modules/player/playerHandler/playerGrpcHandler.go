package playerHandler

import (
	"context"
	playerPb "github.com/Montheankul-K/game-microservices/modules/player/playerPb"
	"github.com/Montheankul-K/game-microservices/modules/player/playerUsecase"
)

type (
	PlayerGrpcHandler struct {
		playerPb.UnimplementedPlayerGrpcServiceServer
		playerUsecase playerUsecase.PlayerUsecase
	}
)

func NewPlayerGrpcHandler(playerUsecase playerUsecase.PlayerUsecase) *PlayerGrpcHandler {
	return &PlayerGrpcHandler{
		playerUsecase: playerUsecase,
	}
}

func (g *PlayerGrpcHandler) CredentialSearch(ctx context.Context, req *playerPb.CredentialSearchReq) (*playerPb.PlayerProfile, error) {
	return nil, nil
}

func (g *PlayerGrpcHandler) FindOnePlayerProfileToRefresh(ctx context.Context, req *playerPb.FindOnePlayerProfileToRefreshReq) (*playerPb.PlayerProfile, error) {
	return nil, nil
}

func (g *PlayerGrpcHandler) GetPlayerSavingAccount(ctx context.Context, req *playerPb.GetPlayerSavingAccountReq) (*playerPb.GetPlayerSavingAccountRes, error) {
	return nil, nil
}
