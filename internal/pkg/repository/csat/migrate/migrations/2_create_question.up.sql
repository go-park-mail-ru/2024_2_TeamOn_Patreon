CREATE TABLE question (
    question_id uuid PRIMARY KEY,
    question text NOT NULL,
    question_theme_id uuid REFERENCES question_theme ON DELETE RESTRICT NOT NULL
);
