CREATE TABLE post_content (
    post_content_id uuid PRIMARY KEY,
    post_id uuid REFERENCES post ON DELETE CASCADE NOT NULL,
    content_id uuid REFERENCES content ON DELETE RESTRICT
);
