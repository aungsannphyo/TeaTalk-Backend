CREATE TABLE IF NOT EXISTS personal_details (
    user_id VARCHAR(255) PRIMARY KEY,
    profile_image VARCHAR(255) NULL,
    gender ENUM('MALE', 'FEMALE', 'OTHER') NULL,
    date_of_birth DATE NULL,
    bio TEXT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);