INSERT INTO
    users (
        id,
        username,
        user_identity,
        email,
        password,
        salt,
        encrypted_user_key,
        user_key_nonce
    )
VALUES (?, ?, ?, ?, ?, ?, ?, ?)