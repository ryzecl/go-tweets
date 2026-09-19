-- migrate:up
CREATE TABLE IF NOT EXISTS comment_likes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    comment_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_id_comment_likes FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_comment_id_comment_likes FOREIGN KEY (comment_id) REFERENCES comments(id)
);

-- migrate:down
DROP TABLE IF EXISTS comment_likes;
