CREATE TABLE message_keys (
    conversation_id VARCHAR(255) PRIMARY KEY,
    aes_key_base64 TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (conversation_id) REFERENCES conversations(id) ON DELETE CASCADE
);
