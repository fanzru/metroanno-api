ALTER TABLE documents
DROP INDEX idx_created_by_user_id,
DROP COLUMN created_by_user_id;