ALTER TABLE student
    ADD COLUMN username VARCHAR(255) NOT NULL UNIQUE DEFAULT '',
    ADD COLUMN password VARCHAR(255) NOT NULL        DEFAULT '',
    ADD COLUMN role     VARCHAR(255) NOT NULL        DEFAULT 'student';
