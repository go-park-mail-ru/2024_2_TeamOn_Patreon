CREATE TABLE comment (
    comment_id uuid PRIMARY KEY,
    post_id uuid REFERENCES post ON DELETE CASCADE,
    user_id uuid REFERENCES people ON DELETE CASCADE,
    about text,
    created_date timestamp,
    updated_date timestamp
);
