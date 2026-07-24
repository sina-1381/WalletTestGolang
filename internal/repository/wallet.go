package repository

import (
	"WalletTest/internal/model"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ============ PUBLIC METHODS ============

func (r *Repository) WalletCreate(ctx context.Context, userID int) (model.Wallet, error) {
	return r.walletCreate(ctx, r.postgresql.Connection, userID)
}

func (r *Repository) WalletFindByUserID(ctx context.Context, userID int) (model.Wallet, error) {
	return r.walletFindByUserID(ctx, r.postgresql.Connection, userID)
}

// ============ INTERNAL METHODS ============

func (r *Repository) walletCreate(ctx context.Context, exec Executor, userID int) (model.Wallet, error) {
	var wallet model.Wallet
	query := `INSERT INTO wallets (user_id) VALUES ($1) ON CONFLICT (user_id) DO NOTHING
              RETURNING id, user_id, balance, created_at, updated_at`
	err := exec.QueryRow(ctx, query, userID).
		Scan(&wallet.ID, &wallet.UserID, &wallet.Balance, &wallet.CreatedAt, &wallet.UpdatedAt)
	if err != nil {
		return model.Wallet{}, status.Errorf(codes.AlreadyExists, "Already Exists")
	}
	return wallet, nil
}

func (r *Repository) walletFindByUserID(ctx context.Context, exec Executor, userID int) (model.Wallet, error) {
	var wallet model.Wallet
	query := `SELECT id, user_id, balance, created_at, updated_at FROM wallets WHERE user_id = $1`
	err := exec.QueryRow(ctx, query, userID).
		Scan(&wallet.ID, &wallet.UserID, &wallet.Balance, &wallet.CreatedAt, &wallet.UpdatedAt)
	if err != nil {
		return model.Wallet{}, status.Errorf(codes.NotFound, "Not Found")
	}
	return wallet, nil
}
