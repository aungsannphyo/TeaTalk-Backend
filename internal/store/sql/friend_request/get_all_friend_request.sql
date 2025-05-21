SELECT fr.id AS request_id, fr.sender_id, u.username, u.email, pd.profile_image, fr.created_at
FROM
    friend_requests fr
    JOIN users u ON fr.sender_id = u.id
    LEFT JOIN personal_details pd ON u.id = pd.user_id
WHERE
    fr.receiver_id = ?
    AND fr.status = 'PENDING';