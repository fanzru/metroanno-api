ALTER TABLE documents
ADD created_by_user_id INT(11),
ADD INDEX idx_created_by_user_id (created_by_user_id);