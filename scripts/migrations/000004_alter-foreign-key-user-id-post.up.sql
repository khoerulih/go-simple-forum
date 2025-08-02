-- make mistake while defining user_id data type in posts table so we must alter it to the right data type
ALTER TABLE posts MODIFY COLUMN user_id BIGINT;

ALTER TABLE posts ADD CONSTRAINT fk_user_id_posts FOREIGN KEY (user_id) REFERENCES users(id);