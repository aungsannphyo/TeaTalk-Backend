CREATE UNIQUE INDEX idx_users_email ON users (email);

CREATE UNIQUE INDEX idx_users_user_identity ON users (user_identity);

CREATE INDEX idx_users_user_name ON users (username);

CREATE INDEX idx_users_id ON users (id);