ALTER TABLE posts DROP FOREIGN KEY fk_user_id_posts;

-- rollback to previous data type
ALTER TABLE posts MODIFY COLUMN user_id INT;