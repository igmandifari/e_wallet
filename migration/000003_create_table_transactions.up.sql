BEGIN;
    CREATE TABLE transactions (
        id VARCHAR(36) PRIMARY KEY,
        sender_account_id VARCHAR(36) NOT NULL REFERENCES accounts(id),
        receiver_account_id VARCHAR(36) NOT NULL REFERENCES accounts(id),
        amount NUMERIC(15, 2) NOT NULL,
        status VARCHAR(50),
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    );

COMMIT;