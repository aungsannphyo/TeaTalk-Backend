CREATE TABLE IF NOT EXISTS group_details (
    conversation_id VARCHAR(255) PRIMARY KEY DEFAULT(UUID()),
    profile_image VARCHAR(255) NULL,
    bio TEXT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (conversation_id) REFERENCES users (id) ON DELETE CASCADE
);