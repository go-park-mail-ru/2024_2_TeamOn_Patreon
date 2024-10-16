CREATE TABLE like_post (
    like_post_id uuid PRIMARY KEY,
    post_id uuid REFERENCES post ON DELETE CASCADE,
    user_id uuid REFERENCES people ON DELETE CASCADE,
    posted_date timestamp
);
