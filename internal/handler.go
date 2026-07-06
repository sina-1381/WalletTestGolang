package internal

import (
	"WalletTest/proto"
	"context"
	"net"

	"google.golang.org/grpc"
)

type Wallet struct {
	Server   *grpc.Server
	Listener net.Listener
	proto.UnimplementedWalletServiceServer
}

func (s *Wallet) CreateWallet(ctx context.Context, req *proto.CreateWalletRequest) (*proto.CreateWalletResponse, error) {
	return &proto.CreateWalletResponse{
		WalletId: req.UserId * 100,
		Message:  s.Listener.Addr().Network(),
	}, nil
}

func (s *Wallet) GetBalance(ctx context.Context, req *proto.GetBalanceRequest) (*proto.GetBalanceResponse, error) {
	return &proto.GetBalanceResponse{
		Balance:  50000,
		WalletId: req.UserId * 100,
	}, nil
}
