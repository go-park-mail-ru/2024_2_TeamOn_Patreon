CREATE TABLE event (
    event_id uuid PRIMARY KEY,
    event_type_id uuid REFERENCES event_type ON DELETE RESTRICT,
    about text,
    happened_date timestamp
);
