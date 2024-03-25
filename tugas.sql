
CREATE TABLE IF NOT EXISTS users (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255),
    occupation VARCHAR(255),
    email VARCHAR(255),
    password_hash VARCHAR(255),
    avatar_file_name VARCHAR(255),
    role VARCHAR(255),
    token VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME,
    PRIMARY KEY (id)
);

DESC users;

CREATE TABLE IF NOT EXISTS campaigns (
     id INT NOT NULL AUTO_INCREMENT,
     user_id INT NOT NULL,
     name VARCHAR(255) NOT NULL,
     short_description TEXT,
     description TEXT,
     perks TEXT,
     backer_count INT DEFAULT 0,
     goal_amount INT,
     current_amount INT DEFAULT 0,
     slug VARCHAR(255),
     created_at DATETIME NOT NULL,
     updated_at DATETIME NOT NULL,
     PRIMARY KEY (id),
     FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS campaign_images (
   id INT NOT NULL AUTO_INCREMENT,
   campaign_id INT NOT NULL,
   file_name VARCHAR(255),
   is_primary INT,
   created_at DATETIME NOT NULL,
   updated_at DATETIME NOT NULL,
   PRIMARY KEY (id),
   FOREIGN KEY (campaign_id) REFERENCES campaigns(id) ON DELETE CASCADE
);
