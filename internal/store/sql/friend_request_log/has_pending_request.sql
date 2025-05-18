SELECT COUNT(*)
FROM friend_requests
WHERE (
        (
            sender_id = ?
            AND receiver_id = ?
        )
        OR (
            receiver_id = ?
            AND sender_id = ?
        )
    )
    AND status = "PENDING"