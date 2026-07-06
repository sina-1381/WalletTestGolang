CREATE TABLE
    IF NOT EXISTS transactions (
        id SERIAL PRIMARY KEY,
        wallet_id INTEGER NOT NULL REFERENCES wallets (id) ON DELETE CASCADE,
        type VARCHAR(10) NOT NULL CHECK (type IN ('credit', 'debit')),
        amount BIGINT NOT NULL CHECK (amount > 0),
        balance_before BIGINT NOT NULL,
        balance_after BIGINT NOT NULL,
        reference VARCHAR(255),
        description TEXT,
        idempotency_key VARCHAR(255) UNIQUE,
        created_at TIMESTAMPTZ DEFAULT NOW ()
    );

CREATE INDEX idx_transactions_wallet_id ON transactions (wallet_id);

CREATE INDEX idx_transactions_idempotency ON transactions (idempotency_key);

CREATE INDEX idx_transactions_created_at ON transactions (created_at);