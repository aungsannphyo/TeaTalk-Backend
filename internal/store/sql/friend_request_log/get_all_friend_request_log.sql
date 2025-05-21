SELECT frl.id AS request_id, frl.sender_id, u.username, u.email, pd.profile_image, frl.created_at
FROM
    friend_request_logs frl
    JOIN users u ON frl.sender_id = u.id
    LEFT JOIN personal_details pd ON u.id = pd.user_id
WHERE
    frl.receiver_id = ?
    AND frl.action = 'SENT';