CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(255) PRIMARY KEY DEFAULT(UUID()),
    email VARCHAR(255) UNIQUE NOT NULL,
    username VARCHAR(255) UNIQUE NOT NULL,
    user_identity VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    salt BLOB NOT NULL,                           
    encrypted_user_key BLOB NOT NULL,              
    user_key_nonce BLOB NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);