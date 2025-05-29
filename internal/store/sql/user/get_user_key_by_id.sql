SELECT id, salt, encrypted_user_key, user_key_nonce
FROM users
WHERE u.id = ?;