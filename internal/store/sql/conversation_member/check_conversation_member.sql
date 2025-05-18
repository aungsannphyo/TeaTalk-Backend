SELECT COUNT(*)
FROM conversation_members
WHERE
    conversation_id = ?
    AND user_id = ?