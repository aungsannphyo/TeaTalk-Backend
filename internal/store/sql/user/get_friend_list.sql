SELECT u.id, u.username, u.user_identity, u.email, pd.profile_image, pd.is_online, pd.last_seen
FROM
    friends f
    JOIN users u ON u.id = CASE
        WHEN f.user_id = ? THEN f.friend_id
        ELSE f.user_id
    END
    LEFT JOIN personal_details pd ON pd.user_id = u.id
WHERE
    f.user_id = ?
    OR f.friend_id = ?;