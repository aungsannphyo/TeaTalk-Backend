CREATE TABLE group_join_requests (
    id VARCHAR(255) PRIMARY KEY  DEFAULT (UUID()),
    user_id VARCHAR(255) NOT NULL,
    conversation_id VARCHAR(255) NOT NULL,
    status ENUM('PENDING', 'ACCEPTED', 'REJECTED') DEFAULT 'PENDING',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, conversation_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (conversation_id) REFERENCES conversations(id) ON DELETE CASCADE
);
