CREATE TABLE like_comment (
    like_comment_id uuid PRIMARY KEY,
    comment_id uuid REFERENCES comment ON DELETE CASCADE,
    user_id uuid REFERENCES people ON DELETE CASCADE,
    posted_date timestamp
);
