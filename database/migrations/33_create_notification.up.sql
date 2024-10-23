CREATE TABLE notification (
    notification_id uuid PRIMARY KEY,
    user_id uuid REFERENCES people ON DELETE CASCADE NOT NULL,
    event_id uuid REFERENCES event ON DELETE CASCADE NOT NULL,
    sent_date timestamp NOT NULL DEFAULT now
);
