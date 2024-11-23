CREATE TABLE asked_question (
    asked_question_id uuid PRIMARY KEY,
    user_id text NOT NULL,
    question_id uuid REFERENCES question ON DELETE CASCADE NOT NULL,
    asked_date timestamp DEFAULT now(),
    result int -- может быть налл
);
