UPDATE friend_requests
SET
    status = ?
WHERE
    receiver_id = ?
    AND id = ?