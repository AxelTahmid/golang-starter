BEGIN;

CREATE TYPE role_type AS ENUM ('user', 'admin');

CREATE TABLE
    IF NOT EXISTS users (
        id BIGSERIAL PRIMARY KEY,
        name VARCHAR(255),
        email VARCHAR(255) NOT NULL UNIQUE,
        password VARCHAR(255) NOT NULL,
        verified BOOLEAN DEFAULT FALSE,
        role role_type DEFAULT 'user',
        created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    );

CREATE TRIGGER set_user_updated_at BEFORE
UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE set_updated_at ();

COMMIT;