CREATE TABLE session (
    session_id uuid PRIMARY KEY,
    user_id uuid REFERENCES people ON DELETE CASCADE NOT NULL,
    created_date timestamp NOT NULL DEFAULT now(),
    finished_date timestamp NOT NULL
);
