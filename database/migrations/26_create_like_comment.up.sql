CREATE TABLE like_comment (
    like_comment_id uuid PRIMARY KEY,
    comment_id uuid REFERENCES comment ON DELETE CASCADE NOT NULL,
    user_id uuid REFERENCES people ON DELETE CASCADE NOT NULL,
    posted_date timestamp NOT NULL DEFAULT now()
);
