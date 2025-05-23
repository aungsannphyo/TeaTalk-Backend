package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/aungsannphyo/ywartalk/pkg/utils" // your password hash utils
	_ "github.com/go-sql-driver/mysql"
)

func seedDatabase(db *sql.DB) {
	log.Println("[SEED] Dummy Data Start SEEDING ...")
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		log.Fatal("failed to check existing users:", err)
	}
	if count > 0 {
		log.Println("✅ Seed skipped (users already exist)")
		return
	}

	users := []struct {
		ID, Email, Username, Identity string
	}{
		{"user-1", "batman@example.com", "Batman", "alice123"},
		{"user-2", "superman@example.com", "Superman", "bob123"},
		{"user-3", "flash@example.com", "flash", "carol123"},
		{"user-4", "wonderwoman@example.com", "Wonder Woman", "dave123"},
		{"user-5", "ironman@example.com", "Iron Man", "eve123"},
		{"user-6", "spiderman@example.com", "Spider Man", "frank123"},
		{"user-7", "antman@example.com", "Ant Man", "grace123"},
		{"user-8", "thanos@example.com", "Thanos", "heidi123"},
		{"user-9", "doctorstrange@example.com", "Doctor Strange", "ivan123"},
		{"user-10", "hulk@example.com", "Hulk", "judy123"},
	}

	for _, u := range users {
		hashedPassword, err := utils.HashPassword("123123")
		if err != nil {
			log.Fatalf("failed to hash password: %v", err)
		}

		_, err = db.Exec(`
			INSERT INTO users (id, email, username, user_identity, password) 
			VALUES (?, ?, ?, ?, ?)
		`, u.ID, u.Email, u.Username, u.Identity, hashedPassword)
		if err != nil {
			log.Fatalf("failed to insert user %s: %v", u.Username, err)
		}

		_, err = db.Exec(`
			INSERT INTO personal_details (user_id, gender, date_of_birth, bio, is_online)
			VALUES (?, 'OTHER', '2000-01-01', ?, TRUE)
		`, u.ID, fmt.Sprintf("This is %s's bio", u.Username))
		if err != nil {
			log.Fatalf("failed to insert personal_details for user %s: %v", u.Username, err)
		}
	}

	// Private conversation between user-1 and user-2
	_, err = db.Exec(`INSERT INTO conversations (id, is_group, name, created_by) VALUES ('conv-p1', FALSE, NULL, 'user-1')`)
	if err != nil {
		log.Fatalf("failed to insert private conversation: %v", err)
	}
	_, err = db.Exec(`INSERT INTO conversation_members (conversation_id, user_id) VALUES ('conv-p1', 'user-1'), ('conv-p1', 'user-2')`)
	if err != nil {
		log.Fatalf("failed to insert private conversation members: %v", err)
	}
	_, err = db.Exec(`INSERT INTO messages (id, conversation_id, sender_id, content) VALUES
		('msg-p1-1', 'conv-p1', 'user-1', 'Hello Bob'),
		('msg-p1-2', 'conv-p1', 'user-2', 'Hey Alice!'),
		('msg-p1-3', 'conv-p1', 'user-1', 'How are you?')
	`)
	if err != nil {
		log.Fatalf("failed to insert private messages: %v", err)
	}

	// Study Group conversation
	_, err = db.Exec(`INSERT INTO conversations (id, is_group, name, created_by) VALUES ('conv-g1', TRUE, 'Study Group', 'user-3')`)
	if err != nil {
		log.Fatalf("failed to insert Study Group conversation: %v", err)
	}
	_, err = db.Exec(`INSERT INTO conversation_members (conversation_id, user_id) VALUES 
		('conv-g1', 'user-3'),
		('conv-g1', 'user-4'),
		('conv-g1', 'user-5'),
		('conv-g1', 'user-6')
	`)
	if err != nil {
		log.Fatalf("failed to insert Study Group members: %v", err)
	}
	_, err = db.Exec(`INSERT INTO group_admins (conversation_id, user_id) VALUES ('conv-g1', 'user-3')`)
	if err != nil {
		log.Fatalf("failed to insert Study Group admin: %v", err)
	}
	_, err = db.Exec(`INSERT INTO messages (id, conversation_id, sender_id, content) VALUES
		('msg-g1-1', 'conv-g1', 'user-3', 'Welcome everyone!'),
		('msg-g1-2', 'conv-g1', 'user-4', 'Thanks!'),
		('msg-g1-3', 'conv-g1', 'user-5', 'Excited to start.')
	`)
	if err != nil {
		log.Fatalf("failed to insert Study Group messages: %v", err)
	}

	// DC Group conversation
	_, err = db.Exec(`INSERT INTO conversations (id, is_group, name, created_by) VALUES ('conv-dc', TRUE, 'DC Group', 'user-1')`)
	if err != nil {
		log.Fatalf("failed to insert DC Group conversation: %v", err)
	}
	_, err = db.Exec(`INSERT INTO conversation_members (conversation_id, user_id) VALUES 
		('conv-dc', 'user-1'),
		('conv-dc', 'user-2'),
		('conv-dc', 'user-3'),
		('conv-dc', 'user-4')
	`)
	if err != nil {
		log.Fatalf("failed to insert DC Group members: %v", err)
	}
	_, err = db.Exec(`INSERT INTO group_admins (conversation_id, user_id) VALUES ('conv-dc', 'user-1')`)
	if err != nil {
		log.Fatalf("failed to insert DC Group admin: %v", err)
	}
	_, err = db.Exec(`INSERT INTO messages (id, conversation_id, sender_id, content) VALUES
		('msg-dc-1', 'conv-dc', 'user-1', 'Welcome to the DC group!'),
		('msg-dc-2', 'conv-dc', 'user-2', 'Glad to be here!'),
		('msg-dc-3', 'conv-dc', 'user-3', 'Let’s get started.')
	`)
	if err != nil {
		log.Fatalf("failed to insert DC Group messages: %v", err)
	}

	// Marvel Group conversation
	_, err = db.Exec(`INSERT INTO conversations (id, is_group, name, created_by) VALUES ('conv-marvel', TRUE, 'Marvel Group', 'user-5')`)
	if err != nil {
		log.Fatalf("failed to insert Marvel Group conversation: %v", err)
	}
	_, err = db.Exec(`INSERT INTO conversation_members (conversation_id, user_id) VALUES 
		('conv-marvel', 'user-5'),
		('conv-marvel', 'user-6'),
		('conv-marvel', 'user-7'),
		('conv-marvel', 'user-8'),
		('conv-marvel', 'user-9'),
		('conv-marvel', 'user-10')
	`)
	if err != nil {
		log.Fatalf("failed to insert Marvel Group members: %v", err)
	}
	_, err = db.Exec(`INSERT INTO group_admins (conversation_id, user_id) VALUES ('conv-marvel', 'user-5')`)
	if err != nil {
		log.Fatalf("failed to insert Marvel Group admin: %v", err)
	}
	_, err = db.Exec(`INSERT INTO messages (id, conversation_id, sender_id, content) VALUES
		('msg-marvel-1', 'conv-marvel', 'user-5', 'Welcome Marvel heroes!'),
		('msg-marvel-2', 'conv-marvel', 'user-6', 'Let’s save the world!'),
		('msg-marvel-3', 'conv-marvel', 'user-7', 'Ready when you are!')
	`)
	if err != nil {
		log.Fatalf("failed to insert Marvel Group messages: %v", err)
	}

	log.Println("✅ Seed completed.")
}
