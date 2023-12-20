CREATE TABLE teacher
(
    id         SERIAL              NOT NULL PRIMARY KEY,
    name       VARCHAR(255)        NOT NULL,
    username   VARCHAR(255) UNIQUE NOT NULL,
    password   VARCHAR(255)        NOT NULL,
    role       VARCHAR(255)        NOT NULL DEFAULT 'teacher',
    created_at TIMESTAMP           NULL     DEFAULT now()
);

CREATE TABLE class
(
    id         SERIAL              NOT NULL PRIMARY KEY,
    name       VARCHAR(255) UNIQUE NOT NULL DEFAULT '',
    managed_by INT,
    start_date TIMESTAMP           NOT NULL,
    end_date   TIMESTAMP           NOT NULL,
    created_at TIMESTAMP           NULL     DEFAULT now(),
    CONSTRAINT class_teacher
        FOREIGN KEY (managed_by) REFERENCES teacher (id)
);

CREATE TABLE student_class
(
    class_id   INT NOT NULL,
    student_id INT NOT NULL,
    PRIMARY KEY (class_id, student_id),
    CONSTRAINT class_fk
        FOREIGN KEY (class_id) REFERENCES class (id),
    CONSTRAINT student_fk
        FOREIGN KEY (student_id) REFERENCES student (id)
)