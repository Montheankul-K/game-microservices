package grpccon

import (
	"errors"
	"github.com/Montheankul-K/game-microservices/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"

	authPb "github.com/Montheankul-K/game-microservices/modules/auth/authPb"
	inventoryPb "github.com/Montheankul-K/game-microservices/modules/inventory/inventoryPb"
	itemPb "github.com/Montheankul-K/game-microservices/modules/item/itemPb"
	playerPb "github.com/Montheankul-K/game-microservices/modules/player/playerPb"
)

type (
	GrpcClientFactory interface {
		Auth() authPb.AuthGrpcServiceClient
		Player() playerPb.PlayerGrpcServiceClient
		Item() itemPb.ItemGrpcServiceClient
		Inventory() inventoryPb.InventoryGrpcServiceClient
	}

	grpcClientFactory struct {
		client *grpc.ClientConn
	}

	grpcAuth struct {
	}
)

func (g *grpcClientFactory) Auth() authPb.AuthGrpcServiceClient {
	return authPb.NewAuthGrpcServiceClient(g.client)
}

func (g *grpcClientFactory) Player() playerPb.PlayerGrpcServiceClient {
	return playerPb.NewPlayerGrpcServiceClient(g.client)
}

func (g *grpcClientFactory) Item() itemPb.ItemGrpcServiceClient {
	return itemPb.NewItemGrpcServiceClient(g.client)
}

func (g *grpcClientFactory) Inventory() inventoryPb.InventoryGrpcServiceClient {
	return inventoryPb.NewInventoryGrpcServiceClient(g.client)
}

func NewGrpcClient(host string) (GrpcClientFactory, error) {
	opts := make([]grpc.DialOption, 0)
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	clientConn, err := grpc.NewClient(host, opts...)
	if err != nil {
		log.Printf("failed to connect to grpc server: %s", err.Error())
		return nil, errors.New("failed to connect to grpc server")
	}

	return &grpcClientFactory{
		client: clientConn,
	}, nil
}

func NewGrpcServer(cfg *config.Jwt, host string) (*grpc.Server, net.Listener) {
	opts := make([]grpc.ServerOption, 0)
	grpcServer := grpc.NewServer(opts...)

	listen, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	return grpcServer, listen
}
