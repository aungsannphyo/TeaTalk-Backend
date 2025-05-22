SELECT
    u.id,
    u.username,
    u.user_identity,
    u.email,
    pd.profile_image,
    COALESCE(pd.is_online, FALSE) AS is_online,
    COALESCE(
        pd.last_seen,
        CURRENT_TIMESTAMP
    ) AS last_seen
FROM (
        SELECT DISTINCT
            CASE
                WHEN f.user_id = ? THEN f.friend_id
                ELSE f.user_id
            END AS friend_id
        FROM friends f
        WHERE (
                f.user_id = ?
                OR f.friend_id = ?
            )
            AND f.user_id < f.friend_id
    ) AS friend_ids
    JOIN users u ON u.id = friend_ids.friend_id
    LEFT JOIN personal_details pd ON pd.user_id = u.id;