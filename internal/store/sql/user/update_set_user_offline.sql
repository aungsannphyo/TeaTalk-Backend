UPDATE personal_details
SET
    is_online = FALSE,
    last_seen = CURRENT_TIMESTAMP
WHERE
    user_id = 1