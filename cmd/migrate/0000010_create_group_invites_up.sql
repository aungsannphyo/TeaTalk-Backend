CREATE TABLE group_invites (
    id VARCHAR(255) PRIMARY KEY DEFAULT (UUID()),
    conversation_id VARCHAR(255) NOT NULL,
    invited_by VARCHAR(255) NOT NULL,
    invited_user_id VARCHAR(255) NOT NULL,
    status VARCHAR(20) DEFAULT 'PENDING', -- pending, approved, rejected
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (conversation_id) REFERENCES conversations(id) ON DELETE CASCADE,
    FOREIGN KEY (invited_by) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (invited_user_id) REFERENCES users(id) ON DELETE CASCADE
);
