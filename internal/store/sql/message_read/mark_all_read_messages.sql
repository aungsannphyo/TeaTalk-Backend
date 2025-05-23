INSERT IGNORE INTO
    message_reads (message_id, user_id, read_at)
SELECT m.id, ?, NOW()
FROM
    messages m
    LEFT JOIN message_reads mr ON mr.message_id = m.id
    AND mr.user_id = ?
WHERE
    m.conversation_id = ?
    AND mr.message_id IS NULL