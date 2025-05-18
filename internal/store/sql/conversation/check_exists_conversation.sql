SELECT c.id, c.is_group, c.name, c.created_by, c.created_at
	FROM conversations c
	JOIN conversation_members m1 ON c.id = m1.conversation_id 
	JOIN conversation_members m2 ON c.id = m2.conversation_id 
	WHERE c.is_group = FALSE
	  AND m1.user_id = ?
	  AND m2.user_id = ?
	  AND m1.user_id != m2.user_id
	GROUP BY c.id
	HAVING COUNT(DISTINCT m1.user_id) = 1 AND COUNT(DISTINCT m2.user_id) = 1;