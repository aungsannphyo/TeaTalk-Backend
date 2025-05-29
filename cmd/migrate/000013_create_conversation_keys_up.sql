CREATE TABLE conversation_keys (
    conversation_id VARCHAR(255),
    user_id VARCHAR(255),
    encrypted_key BLOB NOT NULL,
    nonce BLOB NOT NULL,
    PRIMARY KEY (conversation_id, user_id),
    FOREIGN KEY (conversation_id) REFERENCES conversations(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
