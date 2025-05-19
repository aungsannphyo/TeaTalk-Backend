SELECT
    id,
    username,
    email,
    password,
    created_at
FROM users
WHERE
    id = ?