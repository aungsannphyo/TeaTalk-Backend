SELECT
    m.id AS messageId,
    cm.user_id AS memberId,
    m.sender_id AS senderId,
    m.content,
    CASE WHEN mr.read_at IS NOT NULL THEN TRUE ELSE FALSE END AS isRead,
    u_reader.username AS seenByName,
    m.created_at AS messageCreatedAt
FROM
    messages m
    JOIN conversation_members cm ON cm.conversation_id = m.conversation_id
    LEFT JOIN message_reads mr ON mr.message_id = m.id AND mr.user_id = cm.user_id
    LEFT JOIN users u_reader ON u_reader.id = mr.user_id
WHERE
    m.conversation_id = ?
    AND m.created_at < IFNULL(?, NOW())
ORDER BY m.created_at DESC
LIMIT ?;
