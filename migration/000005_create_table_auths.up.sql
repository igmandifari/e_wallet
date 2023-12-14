BEGIN;
    CREATE TABLE auths (
        id SERIAL PRIMARY KEY,
        user_id VARCHAR(36) NOT NULL REFERENCES users(id),
        uid VARCHAR(36) NOT NULL,
        token text NOT NULL,
        expired_at TIMESTAMP WITH TIME ZONE NOT NULL,
        created_at TIMESTAMP  WITH TIME ZONE NOT NULL DEFAULT NOW()
    );
COMMIT;