CREATE TABLE event (
    event_id uuid PRIMARY KEY,
    event_type_id uuid REFERENCES event_type ON DELETE RESTRICT NOT NULL,
    about text NOT NULL,
    happened_date timestamp DEFAULT now()
);
