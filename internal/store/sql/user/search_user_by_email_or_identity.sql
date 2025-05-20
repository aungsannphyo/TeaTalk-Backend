SELECT
    id,
    email,
    username,
    user_identity,
    created_at
FROM users
WHERE
    email = ?
    OR user_identity = ?
LIMIT 1