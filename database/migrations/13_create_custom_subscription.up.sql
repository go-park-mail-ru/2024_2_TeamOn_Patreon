CREATE TABLE custom_subscription (
    custom_subscription_id uuid PRIMARY KEY,
    author_id uuid REFERENCES people (user_id) ON DELETE CASCADE, -- без автора нет подписки
    custom_name text,
    cost integer,
    subscription_layer_id uuid REFERENCES subscription_layer ON DELETE RESTRICT,
    created_date timestamp
);
