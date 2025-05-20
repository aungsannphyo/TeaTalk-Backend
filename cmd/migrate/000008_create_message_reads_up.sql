CREATE TABLE message_reads (
  message_id VARCHAR(255),
  user_id VARCHAR(255),
  read_at TIMESTAMP,
  PRIMARY KEY (message_id, user_id),
  FOREIGN KEY (message_id) REFERENCES messages(id) ON DELETE CASCADE,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
