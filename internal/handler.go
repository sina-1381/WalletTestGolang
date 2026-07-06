package internal

import (
	"WalletTest/proto"
	"context"
)

type Wallet struct {
	proto.UnimplementedWalletServiceServer
}

func NewWallet() *Wallet {
	return &Wallet{}
}

func (w *Wallet) CreateWallet(ctx context.Context, req *proto.CreateWalletRequest) (*proto.CreateWalletResponse, error) {
	return &proto.CreateWalletResponse{
		WalletId: req.UserId * 100,
		Message:  "sssssssDdddddds",
	}, nil
}

func (w *Wallet) GetBalance(ctx context.Context, req *proto.GetBalanceRequest) (*proto.GetBalanceResponse, error) {
	return &proto.GetBalanceResponse{
		Balance:  50000,
		WalletId: req.UserId * 100,
	}, nil
}
