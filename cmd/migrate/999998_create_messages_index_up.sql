CREATE INDEX idx_messages_conversation_id ON messages (conversation_id);

CREATE INDEX idx_messages_sender_id ON messages (sender_id);

CREATE INDEX idx_messages_conv_sender ON messages (
    conversation_id,
    sender_id,
    created_at
);