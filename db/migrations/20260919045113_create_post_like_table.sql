-- migrate:up
CREATE TABLE IF NOT EXISTS post_likes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    post_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_user_id_post_likes FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_post_id_post_likes FOREIGN KEY (post_id) REFERENCES posts(id)
);

-- migrate:down
DROP TABLE IF EXISTS post_likes;
