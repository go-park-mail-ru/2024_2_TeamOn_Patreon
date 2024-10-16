CREATE TABLE subscription (
    subscription_id uuid PRIMARY KEY,
    user_id uuid REFERENCES people ON DELETE CASCADE,       -- без юзера нет подписки
    custom_subscription_id uuid REFERENCES custom_subscription ON DELETE CASCADE, -- без подписки нет подписки
    started_date timestamp,
    finished_date timestamp
);
