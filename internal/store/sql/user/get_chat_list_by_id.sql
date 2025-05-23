SELECT *
FROM (
        -- Private (Friend) Chats
        (
            SELECT
                c.id AS conversation_id, u.id AS receiver_id, FALSE AS is_group, u.username AS name, pd.profile_image, IF(pd.is_online, 1, 0) AS is_online, COALESCE(
                    pd.last_seen, CURRENT_TIMESTAMP
                ) AS last_seen, m.id AS last_message_id, m.content AS last_message_content, m.sender_id AS last_message_sender, m.created_at AS last_message_created_at, (
                    SELECT COUNT(*)
                    FROM
                        messages m2
                        LEFT JOIN message_reads mr ON m2.id = mr.message_id
                        AND mr.user_id = ?
                    WHERE
                        m2.conversation_id = c.id
                        AND m2.sender_id != ?
                        AND mr.message_id IS NULL
                ) AS unread_count
            FROM
                conversations c
                JOIN conversation_members cm1 ON cm1.conversation_id = c.id
                JOIN conversation_members cm2 ON cm2.conversation_id = c.id
                AND cm1.user_id != cm2.user_id
                JOIN users u ON u.id = cm2.user_id
                LEFT JOIN personal_details pd ON pd.user_id = u.id
                LEFT JOIN messages m ON m.id = (
                    SELECT id
                    FROM messages
                    WHERE
                        conversation_id = c.id
                    ORDER BY created_at DESC
                    LIMIT 1
                )
            WHERE
                cm1.user_id = ?
                AND c.is_group = FALSE
                AND EXISTS (
                    SELECT 1
                    FROM friends f
                    WHERE (
                            f.user_id = cm1.user_id
                            AND f.friend_id = cm2.user_id
                        )
                        OR (
                            f.user_id = cm2.user_id
                            AND f.friend_id = cm1.user_id
                        )
                )
        )
        UNION ALL
        -- Group Chats
        (
            SELECT
                c.id AS conversation_id, NULL AS receiver_id, TRUE AS is_group, c.name AS name, NULL AS profile_image, (
                    SELECT COUNT(*)
                    FROM
                        conversation_members cmg
                        JOIN personal_details pdg ON pdg.user_id = cmg.user_id
                    WHERE
                        cmg.conversation_id = c.id
                        AND pdg.is_online = TRUE
                        AND cmg.user_id != ?
                ) AS is_online, NULL AS last_seen, m.id AS last_message_id, m.content AS last_message_content, m.sender_id AS last_message_sender, m.created_at AS last_message_created_at, (
                    SELECT COUNT(*)
                    FROM
                        messages m2
                        LEFT JOIN message_reads mr ON m2.id = mr.message_id
                        AND mr.user_id = ?
                    WHERE
                        m2.conversation_id = c.id
                        AND m2.sender_id != ?
                        AND mr.message_id IS NULL
                ) AS unread_count
            FROM
                conversations c
                JOIN conversation_members cm ON cm.conversation_id = c.id
                LEFT JOIN messages m ON m.id = (
                    SELECT id
                    FROM messages
                    WHERE
                        conversation_id = c.id
                    ORDER BY created_at DESC
                    LIMIT 1
                )
            WHERE
                cm.user_id = ?
                AND c.is_group = TRUE
        )
    ) AS combined_chats
ORDER BY last_message_created_at DESC;