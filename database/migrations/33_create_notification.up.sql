CREATE TABLE notification (
    notification_id uuid PRIMARY KEY,
    user_id uuid REFERENCES people ON DELETE CASCADE,
    event_id uuid REFERENCES event ON DELETE CASCADE,
    sent_date timestamp
);
