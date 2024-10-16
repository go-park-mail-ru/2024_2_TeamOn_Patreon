CREATE TABLE avatar (
    avatar_id uuid PRIMARY KEY,
    user_id uuid REFERENCES people ON DELETE CASCADE,
    avatar_url text
);
