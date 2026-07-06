package model

import "time"

type Transaction struct {
	ID                    int       `json:"id"`
	WalletID              int       `json:"wallet_id"`
	Type                  string    `json:"type"`
	Amount                int64     `json:"amount"`
	BalanceBefore         int64     `json:"balance_before"`
	BalanceAfter          int64     `json:"balance_after"`
	Reference             string    `json:"reference"`
	Description           string    `json:"description"`
	IdempotencyKey        string    `json:"-"`
	RefundedTransactionID *int      `json:"refunded_transaction_id"`
	CreatedAt             time.Time `json:"created_at"`
}
