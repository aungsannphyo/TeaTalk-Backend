SELECT u.id, u.username, u.email, u.created_at
FROM
    conversation_members cm
    JOIN users u ON cm.user_id = u.id
    JOIN conversations c ON cm.conversation_id = c.id
WHERE
    c.id = ?
    AND c.is_group = TRUE