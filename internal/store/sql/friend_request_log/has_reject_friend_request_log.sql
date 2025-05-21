SELECT action
FROM friend_request_logs
WHERE (
        sender_id = ?
        AND receiver_id = ?
    )
    OR (
        sender_id = ?
        AND receiver_id = ?
    )
ORDER BY created_at DESC
LIMIT 1;