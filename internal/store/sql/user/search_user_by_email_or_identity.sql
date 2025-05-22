SELECT
    u.id,
    u.email,
    u.username,
    u.user_identity,
    IF(
        EXISTS (
            SELECT 1
            FROM friends f
            WHERE (
                    f.user_id = ?
                    AND f.friend_id = u.id
                )
                OR (
                    f.friend_id = ?
                    AND f.user_id = u.id
                )
        ),
        TRUE,
        FALSE
    ) AS is_friend,
    COALESCE(pd.profile_image, '') AS profile_image,
    COALESCE(pd.is_online, FALSE) AS is_online,
    COALESCE(
        pd.last_seen,
        CURRENT_TIMESTAMP
    ) AS last_seen
FROM users u
    LEFT JOIN personal_details pd ON u.id = pd.user_id
WHERE (
        u.email = ?
        OR u.user_identity = ?
        OR u.username LIKE CONCAT('%', ?, '%') -- ← partial match on username
    )
    AND u.id != ?;