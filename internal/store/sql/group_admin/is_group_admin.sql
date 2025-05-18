SELECT COUNT(*)
FROM group_admins
WHERE
    conversation_id = ?
    AND user_id = ?