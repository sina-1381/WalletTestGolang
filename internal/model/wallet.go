package model

import (
	"WalletTest/proto"
	"time"
)

type Wallet struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ProtoWallet(w Wallet) *proto.Wallet {
	return &proto.Wallet{
		Id:        int64(w.ID),
		UserId:    int64(w.UserID),
		Balance:   w.Balance,
		CreatedAt: w.CreatedAt.Format(time.RFC3339),
		UpdatedAt: w.UpdatedAt.Format(time.RFC3339),
	}
}
