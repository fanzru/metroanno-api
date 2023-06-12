ALTER TABLE subjects
MODIFY COLUMN subject_text VARCHAR(255);

ALTER TABLE subjects
ADD CONSTRAINT subjects_subject_text_unique UNIQUE (subject_text);
