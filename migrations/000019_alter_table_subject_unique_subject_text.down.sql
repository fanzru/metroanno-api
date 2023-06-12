ALTER TABLE subjects
DROP CONSTRAINT subjects_subject_text_unique;


ALTER TABLE subjects
MODIFY COLUMN subject_text LONGTEXT;
