CREATE TABLE tip (
    tip_id uuid PRIMARY KEY,
    user_id uuid REFERENCES people ON DELETE CASCADE,
    author_id uuid REFERENCES people (user_id) ON DELETE CASCADE,
    cost integer,
    message text,
    payed_date timestamp
);
