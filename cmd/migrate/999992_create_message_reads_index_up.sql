CREATE INDEX idx_reads_user_id ON message_reads (user_id);

CREATE INDEX idx_reads_message_id ON message_reads (message_id);

CREATE INDEX idx_reads_msg_user ON message_reads (message_id, user_id);

CREATE INDEX idx_message_reads_msg_user ON message_reads (message_id, user_id);