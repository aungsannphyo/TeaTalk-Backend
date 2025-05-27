SELECT
    u.id,
    u.username,
    u.user_identity,
    u.email,
    pd.profile_image,
    COALESCE(pd.is_online, FALSE) AS is_online,
    pd.last_seen AS last_seen,
    (
        SELECT c.id
        FROM
            conversations c
            JOIN conversation_members m1 ON c.id = m1.conversation_id
            JOIN conversation_members m2 ON c.id = m2.conversation_id
        WHERE
            c.is_group = FALSE
            AND (
                (
                    m1.user_id = ?
                    AND m2.user_id = u.id
                )
                OR (
                    m1.user_id = u.id
                    AND m2.user_id = ?
                )
            )
            AND m1.user_id != m2.user_id
        GROUP BY
            c.id
        HAVING
            COUNT(DISTINCT m1.user_id) = 1
            AND COUNT(DISTINCT m2.user_id) = 1
        LIMIT 1
    ) AS conversation_id
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