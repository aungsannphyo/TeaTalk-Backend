SELECT u.id, u.email, u.username, u.user_identity, pd.profile_image, pd.gender, pd.date_of_birth, pd.bio, pd.is_online
FROM users u
    LEFT JOIN personal_details pd ON u.id = pd.user_id
WHERE
    u.id = ?;