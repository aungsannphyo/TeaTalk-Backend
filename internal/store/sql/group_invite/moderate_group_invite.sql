UPDATE group_invites
SET
    status = ?
WHERE
    conversation_id = ?
    AND invited_user_id = ?