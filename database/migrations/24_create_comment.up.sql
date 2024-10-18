CREATE TABLE comment (
    comment_id uuid PRIMARY KEY,
    post_id uuid REFERENCES post ON DELETE CASCADE NOT NULL,
    user_id uuid REFERENCES people ON DELETE CASCADE NOT NULL,
    about text NOT NULL,
    created_date timestamp NOT NULL DEFAULT now(),
    updated_date timestamp
);
