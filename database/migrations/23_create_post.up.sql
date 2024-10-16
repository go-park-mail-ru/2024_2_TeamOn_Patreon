CREATE TABLE post (
    post_id uuid PRIMARY KEY,
    user_id uuid REFERENCES people ON DELETE CASCADE,
    title text,
    about text,
    content_id uuid REFERENCES content ON DELETE CASCADE,
    subscription_layer_id uuid REFERENCES subscription_layer ON DELETE RESTRICT,
    created_date timestamp,
    updated_date timestamp
);
