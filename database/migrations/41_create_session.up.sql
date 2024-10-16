CREATE TABLE session (
    session_id uuid PRIMARY KEY,
    user_id uuid REFERENCES people ON DELETE CASCADE,
    created_date timestamp,
    finished_date timestamp
);
