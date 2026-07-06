package service

import (
	"WalletTest/internal/model"
	"WalletTest/internal/repository"
	"context"
)

type Wallet struct {
	repository *repository.Repository
}

func NewWallet(repository *repository.Repository) *Wallet {
	return &Wallet{
		repository: repository,
	}
}

func (s *Wallet) Create(ctx context.Context, userID int) (model.Wallet, error) {
    wallet, err := s.repository.WalletCreate(ctx, userID)
    if err != nil {
        return model.Wallet{}, err
    }
    return wallet, nil
}

func (s *Wallet) GetBalance(ctx context.Context, userID int) (model.Wallet, error) {
    wallet, err := s.repository.WalletFindByUserID(ctx, userID)
    if err != nil {
        return model.Wallet{}, err
    }
    return wallet, nil
}