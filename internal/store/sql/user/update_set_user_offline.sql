UPDATE personal_details
SET
    is_online = FALSE,
    last_seen = NOW()
WHERE
    user_id = ?