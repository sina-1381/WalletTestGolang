package handler

import (
	"WalletTest/internal/service"
	"WalletTest/proto"
	"context"
)

type Wallet struct {
	service *service.Wallet
	proto.UnimplementedWalletServiceServer
}

func NewWallet(service *service.Wallet) *Wallet {
	return &Wallet{
		service: service,
	}
}

func (h *Wallet) CreateWallet(ctx context.Context, req *proto.CreateWalletRequest) (*proto.CreateWalletResponse, error) {
	wallet, err := h.service.Create(ctx, int(req.UserId))
	if err != nil {
		return nil, err
	}
	return &proto.CreateWalletResponse{
		WalletId: int64(wallet.ID),
		Balance:  wallet.Balance,
	}, nil
}

func (h *Wallet) GetBalance(ctx context.Context, req *proto.GetBalanceRequest) (*proto.GetBalanceResponse, error) {
	wallet, err := h.service.GetBalance(ctx, int(req.UserId))
	if err != nil {
		return nil, err
	}
	return &proto.GetBalanceResponse{
		WalletId: int64(wallet.ID),
		Balance:  wallet.Balance,
	}, nil
}
