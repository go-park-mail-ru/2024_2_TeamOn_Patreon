CREATE TABLE subscription (
    subscription_id uuid PRIMARY KEY,
    user_id uuid REFERENCES people ON DELETE CASCADE NOT NULL,       -- без юзера нет подписки
    custom_subscription_id uuid REFERENCES custom_subscription ON DELETE CASCADE NOT NULL, -- без подписки нет подписки
    started_date timestamp NOT NULL DEFAULT now(),
    finished_date timestamp NOT NULL
);
