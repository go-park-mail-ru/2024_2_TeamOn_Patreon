CREATE TABLE notification (
    notification_id uuid PRIMARY KEY,
    user_id uuid REFERENCES people ON DELETE CASCADE NOT NULL,
    about text NOT NULL,
    created_at timestamp DEFAULT now()
    is_viewed boolean DEFAULT false
);
