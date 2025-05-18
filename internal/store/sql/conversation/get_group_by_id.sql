SELECT c.id, c.is_group, c.name, c.created_by, c.created_at
FROM
    conversation_members cm
    JOIN conversations c ON cm.conversation_id = c.id
WHERE
    cm.user_id = ?
    AND c.is_group = TRUE