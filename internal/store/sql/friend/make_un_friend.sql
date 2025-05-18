DELETE FROM friends
WHERE (
        user_id = ?
        AND friend_id = ?
    )
    OR (
        friend_id = ?
        AND user_id = ?
    )