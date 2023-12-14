BEGIN;
    CREATE TABLE activity_histories (
        id SERIAL PRIMARY KEY,
        user_id VARCHAR(36) NULL REFERENCES users(id),
        activity_type VARCHAR(50) NOT NULL,
        description TEXT,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    );

COMMIT;